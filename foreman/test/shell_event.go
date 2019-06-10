// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
)

const (
	errorExecCmd = "%s (%d)"
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

// NewShellEventWithEvent create an scenario shell event
func NewShellEventWithEvent(e *Event) (*ShellEvent, error) {
	se := NewShellEvent()
	err := se.ParseEvent(e)
	if err != nil {
		return nil, err
	}
	return se, nil
}

// ParseEvent parses the specified event.
func (q *ShellEvent) ParseEvent(e *Event) error {
	eventData := e.Data
	if len(eventData) <= 0 {
		return fmt.Errorf(errorEmptyEvent)
	}

	params := strings.SplitN(eventData, " ", 2)
	if len(params) != 2 {
		return fmt.Errorf(errorInvalidEventParameterCount, len(params), e.Data)
	}

	q.Command = params[1]

	return nil
}

// Execute executes the scenario
func (q *ShellEvent) Execute() (*Response, *errors.Error) {
	cmd := q.Command
	//cmd = "\"" + q.Command + "\""

	/* cgo
	ret := C.system(C.CString(cmd))
	if ret != 0 {
		return fmt.Errorf(errorExecCmd, cmd, ret)
	}
	*/

	err := exec.Command("sh", "-c", cmd).Run()

	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	res := NewQueryResponse()
	res.Query = q.Command
	res.StatusCode = 0

	return res.Response, nil
}
