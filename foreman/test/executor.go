// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
)

// Executor represents a scenario executor
type Executor interface {
	// Execute executes the scenario
	Execute() (*Response, *errors.Error)
}
