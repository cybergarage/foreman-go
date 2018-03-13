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
	ErrorCodeQueryNotFoundData       = 1015
	ErrorCodeQueryInvalidTimeFormat  = 1100
)

func ErrorCodeToString(code int) string {
	errorMap := map[int]string{
		ErrorInternalServerError:         "Internal Server Error",
		ErrorCodeQueryInvalid:            "Invalid Query",
		ErrorCodeQueryMethodNotSupported: "Query method not supported",
		ErrorCodeQueryEmptyTarget:        "Empty query target",
		ErrorCodeQueryTargetNotSupported: "Query target not supported",
		ErrorCodeQueryInvalidValues:      "Invalid query values",
		ErrorCodeQueryInvalidColumns:     "Invalid query columns",
		ErrorCodeQueryInvalidConditions:  "Invalid query conditions",
		ErrorCodeQueryNotFoundData:       "Not found data",
		ErrorCodeQueryInvalidTimeFormat:  "Invalid time format",
	}
	errMsg, ok := errorMap[code]
	if !ok {
		return ""
	}
	return errMsg
}
