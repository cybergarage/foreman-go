// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
)

const (
	errorContainerTag     = "error"
	errorCodeTag          = "code"
	errorMessageTag       = "message"
	errorDetailCodeTag    = "detailCode"
	errorDetailMessageTag = "detailMessage"
)

// NewError returns a new error of the specified parameters.
func NewError(code int, msg string) interface{} {
	errDetail := map[string]interface{}{}
	if 0 < code {
		errDetail[errorCodeTag] = code
	}
	if 0 < len(msg) {
		errDetail[errorMessageTag] = code
	}

	err := map[string]interface{}{
		errorContainerTag: errDetail,
	}

	return err
}

// NewErrorWithError returns a new error of the specified error.
func NewErrorWithError(err *errors.Error) interface{} {
	errDetail := map[string]interface{}{}
	if 0 < err.Code {
		errDetail[errorCodeTag] = err.Code
	}
	if 0 < len(err.Message) {
		errDetail[errorMessageTag] = err.Message
	}
	if 0 < err.DetailCode {
		errDetail[errorDetailCodeTag] = err.DetailCode
	}
	if 0 < len(err.DetailMessage) {
		errDetail[errorDetailMessageTag] = err.DetailMessage
	}

	jsonError := map[string]interface{}{
		errorContainerTag: errDetail,
	}

	return jsonError
}
