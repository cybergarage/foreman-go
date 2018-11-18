// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Loader represents a loader to read from the specefied data source.
type Loader struct {
}

// NewLoader returns a new loader.
func NewLoader() *Loader {
	loader := &Loader{}
	return loader
}

// LoadFromFile loads queries from the specified file.
func (loader *Loader) LoadFromFile(filename string) (Queries, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	queries := make([]Query, 0)

	for n, line := range strings.Split(string(content), "\n") {

		if len(line) <= 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}

		parser := NewParser()
		newQueries, err := parser.ParseString(line)
		if err != nil {
			return nil, fmt.Errorf("[%d] : %s (%s)", (n + 1), line, err.Error())
		}

		queries = append(queries, newQueries...)
	}

	return queries, nil
}
