// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// Event represents a scenario event.
type Event struct {
	Data string
}

// NewEvent create an scenario event with the specified string.
func NewEvent(data string) *Event {
	e := &Event{
		Data: data,
	}
	return e
}
