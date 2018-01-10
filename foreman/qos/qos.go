// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// QoS includes all QoS rules.
type QoS struct {
	Rules     []Rule
	variables map[string]Variable
}

// NewQoS returns a new null object.
func NewQoS() *QoS {
	qos := &QoS{
		Rules:     make([]Rule, 0),
		variables: make(map[string]Variable),
	}
	return qos
}
