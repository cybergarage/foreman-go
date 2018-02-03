// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

const (
	ErrorCodeUnknown                   = 0
	ErrorCodeParserError               = -32700
	ErrorCodeInvalidRequest            = -32600
	ErrorCodeMethodNotFound            = -32601
	ErrorCodeInvalidParams             = -32602
	ErrorCodeInternalError             = -32603
	ErrorCodeBadDestination            = -32000
	ErrorCodeMovedPermanently          = -32001
	ErrorCodeConditionFailed           = -32010
	ErrorCodeScriptEngineInternalError = -32020
	ErrorCodeScriptEngineNotFound      = -32021
	ErrorCodeScriptCompileError        = -32022
	ErrorCodeScriptRuntimeError        = -32023
)

func ErrorCodeToString(code int) string {
	errorMap := map[int]string{
		ErrorCodeUnknown:                   "Unknown Error",
		ErrorCodeParserError:               "Parse error",
		ErrorCodeInvalidRequest:            "Invalid Request",
		ErrorCodeMethodNotFound:            "Method not found",
		ErrorCodeInvalidParams:             "Invalid params",
		ErrorCodeInternalError:             "Internal error",
		ErrorCodeBadDestination:            "Bad Destination",
		ErrorCodeMovedPermanently:          "Moved Permanently",
		ErrorCodeConditionFailed:           "Condition Failed",
		ErrorCodeScriptEngineInternalError: "Script Engine Internal Error",
		ErrorCodeScriptEngineNotFound:      "Script Engine Not Found",
		ErrorCodeScriptCompileError:        "Script Compile Error",
		ErrorCodeScriptRuntimeError:        "Script Runtime Error",
	}
	errMsg, ok := errorMap[code]
	if !ok {
		return ""
	}
	return errMsg
}
