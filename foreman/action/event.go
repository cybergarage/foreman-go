// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import "time"

// EventObject represents an interface of the event object.
type EventObject interface {
	RouteSource
}

// Event represents an event object.
type Event struct {
	source    EventObject
	timestamp time.Time
	params    Parameters
}

// NewEventWithSource returns a new event object with the specified object.
func NewEventWithSource(sourceObj EventObject) *Event {
	e := &Event{
		source:    sourceObj,
		timestamp: time.Now(),
		params:    NewParameters(),
	}
	return e
}

// GetSource returns the source object which creates this event.
func (e *Event) GetSource() EventObject {
	return e.source
}

// GetTimestamp returns a timestamp of the event.
func (e *Event) GetTimestamp() time.Time {
	return e.timestamp
}

// GetParameters returns the parameters of the event.
func (e *Event) GetParameters() Parameters {
	return e.params
}

// AddParameter adds a new parameter.
func (e *Event) AddParameter(param *Parameter) error {
	return e.AddParameter(param)
}
