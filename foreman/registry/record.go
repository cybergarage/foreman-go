// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"fmt"
	"time"
)

// Record represents a meta record.
type Record struct {
	Key       string
	Value     string
	Version   int
	Timestamp time.Time
}

// NewRecord returns a new record.
func NewRecord() *Record {
	m := &Record{}
	return m
}

// String returns a string description of the instance
func (self *Record) String() string {
	return fmt.Sprintf("%s : %f (%d)", self.Key, self.Value, self.Timestamp.Unix())
}
