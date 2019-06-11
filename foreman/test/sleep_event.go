// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

// SleepEvent represents a sleep scenario event.
type SleepEvent struct {
	Executor
	time.Duration
}

// NewSleepEvent create a sleep scenario event.
func NewSleepEvent() *SleepEvent {
	q := &SleepEvent{
		Duration: 0,
	}
	return q
}

// NewSleepEventWithEvent create a sleep scenario event with the specified event
func NewSleepEventWithEvent(e *Event) (*SleepEvent, error) {
	se := NewSleepEvent()
	err := se.ParseEvent(e)
	if err != nil {
		return nil, err
	}
	return se, nil
}

// ParseEvent parses the specified event.
func (q *SleepEvent) ParseEvent(e *Event) error {
	eventData := e.Data
	if len(eventData) <= 0 {
		return fmt.Errorf(errorEmptyEvent)
	}

	params := strings.SplitN(eventData, " ", 2)
	if len(params) != 2 {
		return fmt.Errorf(errorInvalidEventParameterCount, len(params), e.Data)
	}

	sleepSec, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}

	q.Duration = time.Duration(sleepSec) * time.Second

	return nil
}

// Execute executes the specified event.
func (q *SleepEvent) Execute() (*Response, *errors.Error) {
	time.Sleep(q.Duration)

	res := NewQueryResponse()
	res.Query = fmt.Sprintf("%d", q.Duration)
	res.StatusCode = 0

	return res.Response, nil
}
