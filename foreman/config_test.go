// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/metric"
)

const (
	errorConfigTestFilename = "config_test.conf"
)

func TestDefaultConfig(t *testing.T) {
	NewDefaultConfig()
}

func TestConfigLoadFile(t *testing.T) {
	testDir, _ := os.Getwd()
	filename := filepath.Join(testDir, errorConfigTestFilename)

	conf, err := NewConfigWithFile(filename)
	if err != nil {
		t.Error(err)
	}

	_, err = metric.NewStoreWithName(conf.Metrics.Store)
	if err != nil {
		t.Error(err)
	}
}
