// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// Event represents a scenario event.
type Event struct {
	No   int
	Data string
}

// NewEvent create an scenario event with the specified string.
func NewEvent(n int, data string) *Event {
	e := &Event{
		No:   n,
		Data: data,
	}
	return e
}

// GetNo returns the scenario number.
func (e *Event) GetNo() int {
	return e.No
}

// GetData returns the scenario data.
func (e *Event) GetData() string {
	return e.Data
}
