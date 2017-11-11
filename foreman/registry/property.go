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

// NewProperty returns a propery object.
func NewProperty() *Property {
	prop := &Property{}
	return prop
}

// SetString set a string value to the data
func (prop *Property) SetString(value string) error {
	prop.Data = value
	return nil
}

// GetString returns a data by string
func (prop *Property) GetString() (string, error) {
	return prop.Data, nil
}

// SetInteger set a integer value to the data
func (prop *Property) SetInteger(value int64) error {
	prop.Data = fmt.Sprintf("%d", value)
	return nil
}

// GetInteger returns a data as interger number
func (prop *Property) GetInteger() (int64, error) {
	return strconv.ParseInt(prop.Data, 10, 64)
}

// SetReal set a real value to the data
func (prop *Property) SetReal(value float64) error {
	prop.Data = fmt.Sprintf("%f", value)
	return nil
}

// GetReal returns a data as real number
func (prop *Property) GetReal() (float64, error) {
	return strconv.ParseFloat(prop.Data, 64)
}

// String returns a string description of the instance
func (prop *Property) String() string {
	return fmt.Sprintf("[%s] : %s", prop.Name, prop.Data)
}
