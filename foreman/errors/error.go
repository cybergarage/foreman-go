// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
)

// Error represents an error.
type Error struct {
	Code          int
	Message       string
	DetailCode    int
	DetailMessage string
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

// NewErrorWithCode returns a new error with the specified code.
func NewErrorWithCode(code int) *Error {
	err := &Error{
		Code:          code,
		Message:       ErrorCodeToString(code),
		DetailCode:    0,
		DetailMessage: "",
	}
	return err
}

// NewErrorWithCodeAndMessage returns a new error with the specified code and message.
func NewErrorWithCodeAndMessage(code int, msg string) *Error {
	err := &Error{
		Code:          code,
		Message:       msg,
		DetailCode:    0,
		DetailMessage: "",
	}
	return err
}

// String returns a string description of the instance
func (err *Error) String() string {
	var errMsg string
	if 0 < err.Code {
		errMsg += fmt.Sprintf("[%d] ", err.Code)
	}
	if 0 < len(err.Message) {
		errMsg += fmt.Sprintf("[%s] ", err.Message)
	}
	if 0 < err.DetailCode {
		errMsg += fmt.Sprintf("[%d] ", err.DetailCode)
	}
	if 0 < len(err.DetailMessage) {
		errMsg += fmt.Sprintf("[%s] ", err.DetailMessage)
	}
	return errMsg
}

// Error returns a error
func (err *Error) Error() error {
	return errors.New(err.String())
}
