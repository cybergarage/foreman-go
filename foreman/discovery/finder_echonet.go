// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

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

// SearchAll searches all nodes.
func (finder *EchonetFinder) SearchAll() error {
	return nil
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return nil
}
