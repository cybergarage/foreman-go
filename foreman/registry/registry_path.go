// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"strings"
)

// registryGetPathStrings splits the specified path string into an string array.
func registryGetPathStrings(path string) ([]string, error) {
	allNames := strings.Split(path, PathDelim)
	names := make([]string, 0)
	for _, name := range allNames {
		if len(name) <= 0 {
			continue
		}
		names = append(names, name)
	}
	return names, nil
}
