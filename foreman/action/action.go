// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// ActionObject represents an abstract interface for the action.
type ActionObject interface {
	Object
	ProcessEvent(e *Event) (ResultSet, error)
}
