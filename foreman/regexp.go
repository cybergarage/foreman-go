// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/discovery"
)

// Regexp represents a regexp for the finder.
type Regexp struct {
	*discovery.Regexp
}

// NewRegexp returns a new regexp.
func NewRegexp() *Regexp {
	regexp := &Regexp{
		Regexp: discovery.NewRegexp(),
	}
	return regexp
}
