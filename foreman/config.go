// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// #include <foreman/foreman-c.h>
// #cgo LDFLAGS: -lforeman++ -lstdc++
import "C"

import (
	"github.com/cybergarage/go-config/config"
)

// Config represents a configuration for Foreman.
type Config struct {
	*config.Config
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	conf := &Config{}
	conf.Config, _ = config.NewConfig()
	return conf
}

// NewConfigFromFile returns a new Config based on the specified file.
func NewConfigFromFile(filename string) (*Config, error) {
	var err error
	conf := &Config{}
	conf.Config, err = config.NewConfigFromFile(filename)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
