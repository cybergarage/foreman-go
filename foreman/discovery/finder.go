// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

// FinderSearchListener a listener for Finder.
type FinderSearchListener interface {
	FinderSearchResponseReceived(*Node)
}

// Finder represents an abstract interface
type Finder interface {
	Search() error
	SetSearchListener(FinderSearchListener) error
	FindNodes(string) ([]*Node, error)
	GetNodes() ([]*Node, error)
}
