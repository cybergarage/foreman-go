// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

func TestNewStatement(t *testing.T) {
	NewStatementWithName("", StatementUnknownType)
}
