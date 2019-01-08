// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

const (
	QueryInsertString         = "INSERT"
	QuerySelectString         = "SELECT"
	QueryUpdateString         = "UPDATE"
	QueryDeleteString         = "DELETE"
	QueryAnalyzeString        = "ANALYZE"
	QueryExecuteString        = "EXECUTE"
	QueryTargetQos            = "QOS"
	QueryTargetConfig         = "CONFIG"
	QueryTargetMetrics        = "METRICS"
	QueryTargetRegister       = "REGISTER"
	QueryTargetRegistry       = "REGISTRY"
	QueryTargetAction         = "ACTION"
	QueryTargetRoute          = "ROUTE"
	QueryTargetFinder         = "FINDER"
	QueryTargetTypeNone       = 0
	QueryTargetTypeShared     = 1
	QueryTargetTypeFederated  = 2
	QueryTargetTypeStandalone = 3
	QueryConditionOperatorLt  = "<"
	QueryConditionOperatorLe  = "<="
	QueryConditionOperatorGt  = ">"
	QueryConditionOperatorGe  = ">="
	QueryConditionOperatorEq  = "=="
	QueryConditionOperatorNeq = "!="
	QueryColumnAll            = "*"
	QueryColumnName           = "name"
	QueryColumnTimestamp      = "ts"
	QueryColumnValue          = "value"
	QueryColumnVersion        = "ver"
)
