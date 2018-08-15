// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"regexp"
)

// FinderSearchListener a listener for Finder.
type FinderSearchListener interface {
	FinderSearchResponseReceived(*Node)
}

// FinderNotifyListener a listener for Finder.
type FinderNotifyListener interface {
	FinderNotifyReceived(*Node)
}

// Finder represents an abstract interface
type Finder interface {
	// SearchAll searches all nodes.
	SearchAll() error
	// SetSearchListener sets a specified listener.
	SetSearchListener(FinderSearchListener) error
	// SetNotifyListener sets a specified listener.
	SetNotifyListener(FinderNotifyListener) error
	// GetAllNodes returns all found nodes.
	GetAllNodes() ([]Node, error)
	// GetPrefixNodes returns only nodes matching with a specified start string
	GetPrefixNodes(string) ([]Node, error)
	// GetRegexpNodes returns only nodes matching with a specified regular expression
	GetRegexpNodes(*regexp.Regexp) ([]Node, error)
	// Start starts the finder.
	Start() error
	// Stop stops the finder.
	Stop() error
}
