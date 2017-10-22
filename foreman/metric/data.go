// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"fmt"
	"time"
)

// Data represents a Foreman Data.
type Data struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// NewData returns a new Data.
func NewData() *Data {
	m := &Data{}

	return m
}

// String returns a string description of the instance
func (self *Data) String() string {
	return fmt.Sprintf("%s : %f (%d)", self.Name, self.Value, self.Timestamp.Unix())
}
