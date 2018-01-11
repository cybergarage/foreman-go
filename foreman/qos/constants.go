// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

const (
	QosOperatorStringLT    = "<"
	QosOperatorStringLE    = "<="
	QosOperatorStringGT    = ">"
	QosOperatorStringGE    = ">="
	QosOperatorStringEQ    = "=="
	QosOperatorStringNotEQ = "!="
)

const (
	qosOperatorTypeUnknown = iota
	qosOperatorTypeLT
	qosOperatorTypeLE
	qosOperatorTypeGT
	qosOperatorTypeGE
	qosOperatorTypeEQ
	qosOperatorTypeNotEQ
)

var qosOperatorTypes = map[string]int{
	QosOperatorStringLT:    qosOperatorTypeLT,
	QosOperatorStringLE:    qosOperatorTypeLE,
	QosOperatorStringGT:    qosOperatorTypeGT,
	QosOperatorStringGE:    qosOperatorTypeGE,
	QosOperatorStringEQ:    qosOperatorTypeEQ,
	QosOperatorStringNotEQ: qosOperatorTypeNotEQ,
}
