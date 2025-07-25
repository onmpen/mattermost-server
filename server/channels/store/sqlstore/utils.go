// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/wiggin77/merror"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
)

var escapeLikeSearchChar = []string{
	"%",
	"_",
}

func sanitizeSearchTerm(term string, escapeChar string) string {
	term = strings.Replace(term, escapeChar, "", -1)

	for _, c := range escapeLikeSearchChar {
		term = strings.Replace(term, c, escapeChar+c, -1)
	}

	return term
}

// Converts a list of strings into a list of query parameters and a named parameter map that can
// be used as part of a SQL query.
func MapStringsToQueryParams(list []string, paramPrefix string) (string, map[string]any) {
	var keys strings.Builder
	params := make(map[string]any, len(list))
	for i, entry := range list {
		if keys.Len() > 0 {
			keys.WriteString(",")
		}

		key := paramPrefix + strconv.Itoa(i)
		keys.WriteString(":" + key)
		params[key] = entry
	}

	return "(" + keys.String() + ")", params
}

// finalizeTransactionX ensures a transaction is closed after use, rolling back if not already committed.
func finalizeTransactionX(transaction *sqlxTxWrapper, perr *error) {
	// Rollback returns sql.ErrTxDone if the transaction was already closed.
	if err := transaction.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		*perr = merror.Append(*perr, err)
	}
}

func deferClose(c io.Closer, perr *error) {
	err := c.Close()
	*perr = merror.Append(*perr, err)
}

// removeNonAlphaNumericUnquotedTerms removes all unquoted words that only contain
// non-alphanumeric chars from given line
func removeNonAlphaNumericUnquotedTerms(line, separator string) string {
	words := strings.Split(line, separator)
	filteredResult := make([]string, 0, len(words))

	for _, w := range words {
		if isQuotedWord(w) || containsAlphaNumericChar(w) {
			filteredResult = append(filteredResult, strings.TrimSpace(w))
		}
	}
	return strings.Join(filteredResult, separator)
}

// containsAlphaNumericChar returns true in case any letter or digit is present, false otherwise
func containsAlphaNumericChar(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// isQuotedWord return true if the input string is quoted, false otherwise. Ex :-
//
//	"quoted string"  -  will return true
//	unquoted string  -  will return false
func isQuotedWord(s string) bool {
	if len(s) < 2 {
		return false
	}

	return s[0] == '"' && s[len(s)-1] == '"'
}

func constructArrayArgs(ids []string) (string, []any) {
	var placeholder strings.Builder
	values := make([]any, 0, len(ids))
	for _, entry := range ids {
		if placeholder.Len() > 0 {
			placeholder.WriteString(",")
		}

		placeholder.WriteString("?")
		values = append(values, entry)
	}

	return "(" + placeholder.String() + ")", values
}

func wrapBinaryParamStringMap(ok bool, props model.StringMap) model.StringMap {
	if props == nil {
		props = make(model.StringMap)
	}
	props[model.BinaryParamKey] = strconv.FormatBool(ok)
	return props
}

// morphWriter is a target to pass to the logger instance of morph.
// For now, everything is just logged at a debug level. If we need to log
// errors/warnings from the library also, that needs to be seen later.
type morphWriter struct{}

func (l *morphWriter) Write(in []byte) (int, error) {
	mlog.Debug(strings.TrimSpace(string(in)))
	return len(in), nil
}

func DSNHasBinaryParam(dsn string) (bool, error) {
	url, err := url.Parse(dsn)
	if err != nil {
		return false, err
	}
	return url.Query().Get("binary_parameters") == "yes", nil
}

// AppendBinaryFlag updates the byte slice to work using binary_parameters=yes.
func AppendBinaryFlag(buf []byte) []byte {
	return append([]byte{0x01}, buf...)
}

const maxTokenSize = 50

// trimInput limits the string to a max size to prevent clogging up disk space
// while logging
func trimInput(input string) string {
	if len(input) > maxTokenSize {
		input = input[:maxTokenSize] + "..."
	}
	return input
}

// Returns the column name for PostgreSQL.
func quoteColumnName(driver string, columnName string) string {
	return columnName
}

// scanRowsIntoMap scans SQL rows into a map, using a provided scanner function to extract key-value pairs
func scanRowsIntoMap[K comparable, V any](rows *sql.Rows, scanner func(rows *sql.Rows) (K, V, error), defaults map[K]V) (map[K]V, error) {
	results := make(map[K]V, len(defaults))

	// Initialize with default values if provided
	maps.Copy(results, defaults)

	for rows.Next() {
		key, value, err := scanner(rows)
		if err != nil {
			return nil, err
		}
		results[key] = value
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating rows: %w", err)
	}

	return results, nil
}
