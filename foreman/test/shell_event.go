// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// #include <stdlib.h>
import "C"

import (
	"errors"
	"fmt"
	"strings"
)

// ShellEvent represents a scenario shell event.
type ShellEvent struct {
	Command string
}

// NewShellEvent create an scenario shell event
func NewShellEvent() *ShellEvent {
	q := &ShellEvent{}
	return q
}

// ParseEvent parses the specified event.
func (q *ShellEvent) ParseEvent(e *Event) error {
	eventData := e.Data
	if len(eventData) <= 0 {
		return errors.New(errorEmptyEvent)
	}

	params := strings.SplitN(eventData, " ", 2)
	if len(params) != 2 {
		return fmt.Errorf(errorInvalidEventParameterCount, len(params), e.Data)
	}

	q.Command = params[1]

	return nil
}

const (
	errorExecCmd = "%s (%d)"
)

// Execute executes the specified event.
func (q *ShellEvent) Execute() error {
	cmd := q.Command
	cmd = "\"" + q.Command + "\""
	ret := C.system(C.CString(cmd))
	if ret != 0 {
		return fmt.Errorf(errorExecCmd, cmd, ret)
	}
	return nil
}
