// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

const (
	ErrorInternalServerError         = 500
	ErrorCodeQueryInvalid            = 1000
	ErrorCodeQueryMethodNotSupported = 1001
	ErrorCodeQueryEmptyTarget        = 1010
	ErrorCodeQueryTargetNotSupported = 1011
	ErrorCodeQueryInvalidValues      = 1012
	ErrorCodeQueryInvalidColumns     = 1013
	ErrorCodeQueryInvalidConditions  = 1014
)

func ErrorCodeToString(code int) string {
	errorMap := map[int]string{}
	errMsg, ok := errorMap[code]
	if !ok {
		return ""
	}
	return errMsg
}
