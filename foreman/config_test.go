// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	configTestFilename = "config_test.conf"
)

func TestDefaultConfig(t *testing.T) {
	conf := NewDefaultConfig()

	defaultTimeout := time.Duration(DefaultConnectionTimeout) * time.Second
	if conf.GetServerConnectionTimeout() <= 0 {
		t.Errorf("%v <= 0", conf.GetServerConnectionTimeout())
	}
	if conf.GetServerConnectionTimeout() != defaultTimeout {
		t.Errorf("%v != %v", conf.GetServerConnectionTimeout(), defaultTimeout)
	}

	if conf.GetMetricsStorePeriod() <= 0 {
		t.Errorf("%v <= 0", conf.GetMetricsStorePeriod())
	}

	if conf.GetMetricsStoreInterval() <= 0 {
		t.Errorf("%v <= 0", conf.GetMetricsStoreInterval())
	}
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
