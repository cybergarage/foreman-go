// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import "time"

// Event represents an event of a paremeter for the source and destination objects.
type Event interface {
	GetName() string
	GetSource() RouteObject
	GetTimeStamp() time.Time
	GetParameters() Parameters
}
