// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package registry

// Properties represents a property map.
type Properties map[string]*Property

// NewProperties returns a property map.
func NewProperties() Properties {
	return make(Properties)
}

// Clear removes all properties
func (props Properties) Clear() {
	for name := range props {
		delete(props, name)
	}
}

// SetProperty sets a property into the property map
func (props Properties) SetProperty(prop *Property) error {
	props[prop.Name] = prop
	return nil
}

// GetPropertyByName returns a property which has specified name
func (props Properties) GetPropertyByName(name string) (*Property, bool) {
	prop, ok := props[name]
	return prop, ok
}
