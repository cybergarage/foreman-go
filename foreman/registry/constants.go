// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

const (
	// RootObjectID is a special object ID, and it is reserved as zero.
	RootObjectID = "0"
	// RootPath is a special path which represents a root path.
	RootPath = "/"
	// PathDelim is a separator of object IDs.
	PathDelim = "/"

	objectPropertiesDelim = 7
	objectPropertyDelim   = 9
)
