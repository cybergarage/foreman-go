// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// Error represents an error.
type Error struct {
	Code          int
	Message       string
	DetailCode    int
	DetailMessage string
}

// NewErrorWithCodeAndMessage returns a new error.
func NewErrorWithCodeAndMessage(code int, msg string) *Error {
	err := &Error{
		Code:          code,
		Message:       msg,
		DetailCode:    0,
		DetailMessage: "",
	}
	return err
}

// NewError returns a new error.
func NewError() *Error {
	err := &Error{
		Code:          0,
		Message:       "",
		DetailCode:    0,
		DetailMessage: "",
	}
	return err
}

// String returns a string description of the instance
func (err *Error) String() string {
	errMsg := fmt.Sprintf("[%d] %s", err.Code, err.Message)
	if (err.DetailCode <= 0) || (len(err.DetailMessage) <= 0) {
		return errMsg
	}
	errMsg += fmt.Sprintf(" [%d] %s", err.DetailCode, err.DetailMessage)
	return errMsg
}

// Error returns a error
func (err *Error) Error() error {
	return errors.New(err.String())
}

// JSONError returns a JSON error
func (err *Error) JSONError() interface{} {
	return json.NewError(err.Code, err.Message)
}
