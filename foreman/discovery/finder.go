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
	Search(regexp.Regexp) error
	SearchAll() error
	SetSearchListener(FinderSearchListener) error
	SetNotifyListener(FinderNotifyListener) error
	GetNodes(regexp.Regexp) ([]*Node, error)
	GetAllNodes() ([]*Node, error)
}
