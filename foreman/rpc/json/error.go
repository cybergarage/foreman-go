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
	errorDetailTag        = "detail"
	errorDetailCodeTag    = "code"
	errorDetailMessageTag = "message"
)

// NewError returns a new error of the specified parameters.
func NewError(code int, msg string) interface{} {
	errObj := map[string]interface{}{}
	if 0 < code {
		errObj[errorCodeTag] = code
	}
	if 0 < len(msg) {
		errObj[errorMessageTag] = code
	}

	err := map[string]interface{}{
		errorContainerTag: errObj,
	}

	return err
}

// NewErrorWithError returns a new error of the specified error.
func NewErrorWithError(err *errors.Error) interface{} {
	errObj := map[string]interface{}{}
	if 0 < err.Code {
		errObj[errorCodeTag] = err.Code
	}
	if 0 < len(err.Message) {
		errObj[errorMessageTag] = err.Message
	}

	// Detail Errors
	if 0 < err.DetailCode || 0 < len(err.DetailMessage) {
		errDetailObj := map[string]interface{}{}
		if 0 < err.DetailCode {
			errDetailObj[errorDetailCodeTag] = err.DetailCode
		}
		if 0 < len(err.DetailMessage) {
			errDetailObj[errorDetailMessageTag] = err.DetailMessage
		}
		errObj[errorDetailTag] = errDetailObj
	}

	jsonError := map[string]interface{}{
		errorContainerTag: errObj,
	}

	return jsonError
}
