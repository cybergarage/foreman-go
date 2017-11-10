// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

import (
	"fmt"
	"strconv"
)

// Property represents a property in the registry store.
type Property struct {
	Name string
	Data string
}

// NewObject returns a new object.
func NewProperty() *Property {
	prop := &Property{}
	return prop
}

// String returns a string description of the instance
func (self *Property) String() string {
	return fmt.Sprintf("[%s] : %s", self.Name, self.Data)
}

// SetString set a string value to the data
func (self *Property) SetString(value string) error {
	self.Data = value
	return nil
}

// GetString returns a data by string
func (self *Property) GetString() (string, error) {
	return self.Data, nil
}

// SetInteger set a integer value to the data
func (self *Property) SetInteger(value int64) error {
	self.Data = fmt.Sprintf("%d", value)
	return nil
}

// GetInteger returns a data as interger number
func (self *Property) GetInteger() (int64, error) {
	return strconv.ParseInt(self.Data, 10, 64)
}

// SetReal set a real value to the data
func (self *Property) SetReal(value float64) error {
	self.Data = fmt.Sprintf("%f", value)
	return nil
}

// GetReal returns a data as real number
func (self *Property) GetReal() (float64, error) {
	return strconv.ParseFloat(self.Data, 64)
}
