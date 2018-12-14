// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import "reflect"

const (
	boostrapConfigErrorEmpty = "Empty configuration"
	boostrapConfigErrorKey   = "Invalid  configuration (%s)"
)

// boostrapConfig represents a boostrap configuration
type boostrapConfig struct {
	Objects map[string]interface{}
}

// NewBootstrapConfigWithObject returns a new cofiguration of the specified object
func NewBootstrapConfigWithObject(obj map[string]interface{}) *boostrapConfig {
	config := &boostrapConfig{
		Objects: obj,
	}
	return config
}

// Equals returns true when the specified quries is equals to this quries, otherwise false.
func (config *boostrapConfig) Equals(otherConfig *boostrapConfig) bool {
	return reflect.DeepEqual(config, otherConfig)
}
