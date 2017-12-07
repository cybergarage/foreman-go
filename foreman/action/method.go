// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

// EncodingType defines encoding types for the source code
type EncodingType int

const (
	EncodingNone   = 0
	EncodingBase64 = 1
)

// Method represents an action method.
type Method struct {
	Language string
	Name     string
	Code     []byte
	Encoding int
}

// NewMethod returns a new method.
func NewMethod() *Method {
	m := &Method{
		Encoding: EncodingNone,
	}
	return m
}
