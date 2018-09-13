// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
)

// NewError returns a new error of the specified parameters.
func NewError(code int, msg string) interface{} {
	errObj := map[string]interface{}{}
	if 0 < code {
		errObj[RpcErrorCode] = code
	}
	if 0 < len(msg) {
		errObj[RpcErrorMessage] = code
	}

	err := map[string]interface{}{
		RpcError: errObj,
	}

	return err
}

// NewErrorWithError returns a new error of the specified error.
func NewErrorWithError(err *errors.Error) interface{} {
	errObj := map[string]interface{}{}
	if 0 < err.Code {
		errObj[RpcErrorCode] = err.Code
	}
	if 0 < len(err.Message) {
		errObj[RpcErrorMessage] = err.Message
	}

	// Detail Errors
	if 0 < err.DetailCode || 0 < len(err.DetailMessage) {
		errDetailObj := map[string]interface{}{}
		if 0 < err.DetailCode {
			errDetailObj[RpcErrorDetailCode] = err.DetailCode
		}
		if 0 < len(err.DetailMessage) {
			errDetailObj[RpcErrorDetailMessage] = err.DetailMessage
		}
		errObj[RpcErrorDetail] = errDetailObj
	}

	jsonError := map[string]interface{}{
		RpcError: errObj,
	}

	return jsonError
}
