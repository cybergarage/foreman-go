// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	configTestFilename       = "config_test.conf"
	finderConfigTestFilename = "finder_config_test.conf"
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

	if len(conf.Finder.Hosts) != 0 {
		t.Errorf("%v != 0", len(conf.Finder.Hosts))
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

func TestFinderConfigLoadFile(t *testing.T) {
	conf, err := NewConfigWithFile(finderConfigTestFilename)
	if err != nil {
		t.Error(err)
	}

	if len(conf.Finder.Hosts) != 3 {
		t.Errorf("%v != 0", len(conf.Finder.Hosts))
	}

	for i := 0; i < 3; i++ {
		expectedHost := fmt.Sprintf("org.cybergarage.foreman00%d", i+1)
		if conf.Finder.Hosts[i] != expectedHost {
			t.Errorf("%v != %v", conf.Finder.Hosts[i], expectedHost)
		}
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
