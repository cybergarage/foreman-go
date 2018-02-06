// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"strings"
)

// NewPathStringWithStrings a joined string of the specified strings.
func NewPathStringWithStrings(paths []string) string {
	pathString := PathSep
	for n, path := range paths {
		if 0 < n {
			pathString += PathSep
		}
		pathString += path
	}
	return pathString
}

// trimPathStrings trims PathSep from the specified path string.
func trimPathStrings(pathString string) string {
	return strings.Trim(pathString, PathSep)
}
