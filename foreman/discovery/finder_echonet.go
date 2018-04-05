// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import "regexp"

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	*baseFinder
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	finder := &EchonetFinder{
		baseFinder: newBaseFinder(),
	}
	return finder
}

// Search searches only the specified nodes.
func (finder *baseFinder) Search(regexp.Regexp) error {
	return nil
}

// SearchAll searches all nodes.
func (finder *baseFinder) SearchAll() error {
	return nil
}
