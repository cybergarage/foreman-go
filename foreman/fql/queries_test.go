// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	queryHashTestFilename = loaderTestFilename
)

func TestNewQueries(t *testing.T) {
	NewQueries()
}

func TestNewQueriesHash(t *testing.T) {
	testDir, _ := os.Getwd()
	filename := filepath.Join(testDir, queryHashTestFilename)

	loader := NewLoader()

	// Q1

	q1, err := loader.LoadFromFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	q1Hash, err := q1.Hash()
	if err != nil {
		t.Error(err)
		return
	}

	// Q2

	q2, err := loader.LoadFromFile(filename)
	if err != nil {
		t.Error(err)
		return
	}
	q2Hash, err := q2.Hash()
	if err != nil {
		t.Error(err)
		return
	}

	// Compare

	if !q1.Equals(q2) {
		t.Errorf("%s == %s", q1Hash, q2Hash)
	}
}
