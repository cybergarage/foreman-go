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
	loaderTestFilename = "loader_test.fql"
)

func TestLoadQueryFromFile(t *testing.T) {
	testDir, _ := os.Getwd()
	filename := filepath.Join(testDir, loaderTestFilename)

	loader := NewLoader()
	_, err := loader.LoadFromFile(filename)
	if err != nil {
		t.Error(err)
	}
}
