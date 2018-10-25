// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"crypto/sha256"
	"fmt"
)

// Queries represents a query list.
type Queries []Query

// NewQueries returns a new query list.
func NewQueries() Queries {
	queries := make([]Query, 0)
	return queries
}

// NewQueriesWithArray returns a new query list of the specified queries.
func NewQueriesWithArray(queries []Query) Queries {
	return queries
}

// Hash returns a hash string of the all queries.
func (queries Queries) Hash() (string, error) {
	if queries.Size() <= 0 {
		return "", fmt.Errorf("Empty queries")
	}

	queryStrs := ""
	for _, query := range queries {
		queryStrs += query.String()
	}

	sumHash := sha256.Sum256([]byte(queryStrs))
	return fmt.Sprintf("%X", sumHash), nil
}

// Equals returns true when the specified quries is equals to this quries, otherwise false.
func (queries Queries) Equals(otherQueries Queries) bool {
	if queries.Size() == 0 && otherQueries.Size() == 0 {
		return true
	}

	queriesHash, err := queries.Hash()
	if err != nil {
		return false
	}

	otherQueriesHash, err := otherQueries.Hash()
	if err != nil {
		return false
	}

	if queriesHash != otherQueriesHash {
		return false
	}

	return true
}

// Size returns a number of the queries.
func (queries Queries) Size() int {
	return len(queries)
}
