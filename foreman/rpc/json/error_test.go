// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

func TestNewError(t *testing.T) {
	err := errors.NewError()
	NewErrorWithError(err)
}
