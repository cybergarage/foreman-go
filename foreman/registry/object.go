// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package registry provides registry interfaces
package registry

import (
	"bytes"
	"fmt"
	"strings"
)

// Object represents a meta object in the registry store.
type Object struct {
	ID           string // ID is guaranteed the uniqueness in the registry store basically. "0" is a special ID, and it is reserved for the root object.
	ParentID     string
	Name         string
	Data         string
	propertyData string
}

// NewObject returns a new object.
func NewObject() *Object {
	m := &Object{}
	return m
}

// isID returns whether the object has a specified ID.
func (obj *Object) isID(id string) bool {
	if obj.ID != id {
		return false
	}
	return true
}

// IsParentID returns whether the object has a specified ID.
func (obj *Object) IsParentID(id string) bool {
	if obj.ParentID != id {
		return false
	}
	return true
}

// IsRootParentID returns whether the object has a root parent ID.
func (obj *Object) IsRootParentID() bool {
	return obj.IsParentID(RootObjectID)
}

// IsName returns whether the object has a specified name.
func (obj *Object) IsName(name string) bool {
	if obj.Name != name {
		return false
	}
	return true
}

// IsData returns whether the object has a specified data.
func (obj *Object) IsData(data string) bool {
	if obj.Data != data {
		return false
	}
	return true
}

// SetProperty set a property to the object.
func (obj *Object) SetProperty(prop *Property) error {
	props, err := obj.GetProperties()
	if err != nil {
		return err
	}

	err = props.SetProperty(prop)
	if err != nil {
		return err
	}

	err = obj.updatePropertyData(props)
	if err != nil {
		return err
	}

	return nil
}

// GetProperties returns all properties in the object.
func (obj *Object) GetProperties() (Properties, error) {
	props := NewProperties()
	if len(obj.propertyData) <= 0 {
		return props, nil
	}

	propLines := strings.Split(obj.propertyData, (string)((byte)(objectPropertiesDelim)))
	for _, propLine := range propLines {
		propData := strings.Split(propLine, (string)((byte)(objectPropertyDelim)))
		if len(propData) != 2 {
			return props, fmt.Errorf(errorInvalidPropertyData, obj.propertyData)
		}
		prop := NewProperty()
		prop.Name = propData[0]
		prop.Data = propData[1]
		props.SetProperty(prop)
	}

	return props, nil
}

// updatePropertyData updates the internal property data by the specified properties.
func (obj *Object) updatePropertyData(props Properties) error {
	propBytes := bytes.NewBuffer(make([]byte, 0))

	firstProperty := true
	for _, prop := range props {
		if !firstProperty {
			propBytes.WriteByte((byte)(objectPropertiesDelim))
		} else {
			firstProperty = false
		}
		propBytes.WriteString(prop.Name)
		propBytes.WriteByte((byte)(objectPropertyDelim))
		propBytes.WriteString(prop.Data)
	}

	obj.propertyData = string(propBytes.Bytes())

	return nil
}

// String returns a string description of the instance
func (obj *Object) String() string {
	return fmt.Sprintf("[%s] : %s, %s, %s", obj.ID, obj.Name, obj.Data, obj.propertyData)
}
