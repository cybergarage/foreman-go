// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

const (
	boostrapConfigErrorEmpty = "Empty configuration"
	boostrapConfigErrorKey   = "Invalid  configuration (%s)"
)

// boostracConfig represents a boostrap configuration
type boostracConfig struct {
	Objects map[string]interface{}
}

// NewBoostrapConfigWithObject returns a new cofiguration of the specified object
func NewBoostrapConfigWithObject(obj map[string]interface{}) *boostracConfig {
	config := &boostracConfig{
		Objects: obj,
	}
	return config
}

// Hash returns a hash string of the all queries.
func (config *boostracConfig) Hash() (string, error) {
	if config.Objects == nil {
		return "", fmt.Errorf(boostrapConfigErrorEmpty)
	}

	configKeys := make([]string, 0, len(config.Objects))
	for configKey := range config.Objects {
		configKeys = append(configKeys, configKey)
	}

	sort.Strings(configKeys)

	configStr := ""
	for _, configKey := range configKeys {
		config, ok := config.Objects[configKey]
		if !ok {
			return "", fmt.Errorf(boostrapConfigErrorKey, configKey)
		}
		configStr += fmt.Sprintf("%v", config)
	}
	sumHash := sha256.Sum256([]byte(configStr))
	return fmt.Sprintf("%X", sumHash), nil
}

// Equals returns true when the specified quries is equals to this quries, otherwise false.
func (config *boostracConfig) Equals(otherConfig *boostracConfig) bool {
	configHash, err := config.Hash()
	if err != nil {
		return false
	}

	otherConfigHash, err := otherConfig.Hash()
	if err != nil {
		return false
	}

	if configHash != otherConfigHash {
		return false
	}

	return true
}
