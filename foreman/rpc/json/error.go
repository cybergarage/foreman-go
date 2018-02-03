// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

const (
	errorContainerTag = "error"
	errorCodeTag      = "code"
	errorMessageTag   = "message"
)

// NewError returns a new error of the specified parameters.
func NewError(code int, msg string) interface{} {
	errDetail := map[string]interface{}{
		errorCodeTag:    code,
		errorMessageTag: msg,
	}
	err := map[string]interface{}{
		errorContainerTag: errDetail,
	}
	return err
}
