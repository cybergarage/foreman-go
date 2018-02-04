// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"strconv"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/registry"
)

func TestConfigInitialKeys(t *testing.T) {
	reg := registry.NewManager()

	err := reg.Start()
	if err != nil {
		t.Error(err)
		return
	}

	config := NewConfigWithRegistry(reg)

	err = config.initialize()
	if err != nil {
		t.Error(err)
		return
	}

	initialKeys := []string{
		ConfigProductKey,
		ConfigVersionKey,
	}

	for _, key := range initialKeys {
		_, err := config.GetString(key)
		if err != nil {
			t.Error(err)
		}
	}

	err = reg.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestConfigIntKeys(t *testing.T) {
	reg := registry.NewManager()

	err := reg.Start()
	if err != nil {
		t.Error(err)
		return
	}

	config := NewConfigWithRegistry(reg)

	intKeys := map[string]int{
		"test_key01": 1,
		"test_key02": 2,
		"test_key03": 3,
	}

	for key, value := range intKeys {
		err := config.SetKey(key, strconv.Itoa(value))
		if err != nil {
			t.Error(err)
		}
		regValue, err := config.GetInt(key)
		if err != nil {
			t.Error(err)
			continue
		}
		if regValue != value {
			t.Errorf("%d != %d", regValue, value)
		}
	}

	err = reg.Stop()
	if err != nil {
		t.Error(err)
	}
}
