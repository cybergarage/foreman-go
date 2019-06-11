// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import "strings"

const (
	shellEventPrefix = "SHELL"
	sleepEventPrefix = "SLEEP"
)

const (
	errorEmptyEvent                 = "Empty event"
	errorInvalidEventParameterCount = "Invalid event parameter count (%d) : %s"
	errorInvalidEventCode           = "Invalid event code : %s"
)

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

// HasDataPrefix returns true when the event databegins with the specified prefix, otherwise false.
func (e *Event) HasDataPrefix(prefix string) bool {
	return strings.HasPrefix(e.Data, prefix)
}

// IsShellEvent returns true when this event has the prefix, otherwise false.
func (e *Event) IsShellEvent() bool {
	return e.HasDataPrefix(shellEventPrefix)
}

// IsSleepEvent returns true when this event has the prefix, otherwise false.
func (e *Event) IsSleepEvent() bool {
	return e.HasDataPrefix(sleepEventPrefix)
}
