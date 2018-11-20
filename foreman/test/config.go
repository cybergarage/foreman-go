// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// Config represents a configuration for a scenario test.
type Config struct {
	EnableSkipError bool
}

// NewDefaultConfig retura an new default configuration.
func NewDefaultConfig() *Config {
	conf := &Config{
		EnableSkipError: false,
	}
	return conf
}
