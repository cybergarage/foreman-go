// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

const (
	errorEncorderUnknownObjectType    = "Unknown object type : %s"
	errorDecorderEmptyKey             = "Empty path"
	errorPathEmptyRootObject          = "Empty root object"
	errorPathInvalidArrayIndexKey     = "Invalid array index key : %s"
	errorPathInvalidArrayIndexRange   = "Invalid array index range : %d > %d"
	errorDecorderNotFoundKey          = "Key (%s) is not found"
	errorDecorderInvalidKeyObjectType = "Key (%s) object type is invalid"
)
