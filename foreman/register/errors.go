// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"errors"
)

var errorNilObject = errors.New("Nil object")
var errorQueryEmptyID = errors.New("Empty ID")

const (
	errorInvalidObject     = "Invalid object (%s)"
	errorInvalidObjectData = "Invalid object (%s) data (%s)"
	errorInvalidStore      = "Invalid store (%s)"
	errorQueryNotFoundID   = "Not found ID (%s)"
)
