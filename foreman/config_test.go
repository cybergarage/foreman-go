// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	configTestFilename = "config_test.conf"
)

func TestDefaultConfig(t *testing.T) {
	NewDefaultConfig()
}

func TestConfigLoadFile(t *testing.T) {
	conf, err := newDefaultTestServerConfig()
	if err != nil {
		t.Error(err)
	}

	_, err = metric.NewStoreWithName(conf.Metrics.Store)
	if err != nil {
		t.Error(err)
	}
}

func newDefaultTestServerConfig() (*Config, error) {
	conf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		return nil, err
	}

	// Disable : [Bootstrap] query = "/etc/foreman/bootstrap"
	conf.SetBootstrapQuery("")

	return conf, nil
}
