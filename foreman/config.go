// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"strconv"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Config represents a Config for Foreman.
type Config struct {
	*registry.Manager
	fql.QueryExecutor
}

// NewConfigWithRegistry returns a new Config with the specified registry.
func NewConfigWithRegistry(mgr *registry.Manager) *Config {
	config := &Config{
		Manager: mgr,
	}
	config.initialize()
	return config
}

func (config *Config) initialize() error {
	err := config.CreateCategoryObject(ConfigCategoryKey)
	if err != nil {
		return err
	}

	initialKeys := map[string]string{
		ConfigProductKey:    ProductName,
		ConfigVersionKey:    Version,
		ConfigHostKey:       DefaultServerHost,
		ConfigCarbonPortKey: strconv.Itoa(DefaultCarbonPort),
		ConfigHttpPortKey:   strconv.Itoa(DefaultHttpPort),
	}

	for key, value := range initialKeys {
		err = config.SetKey(key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadFile loads a specified Config file.
func (config *Config) LoadFile(filename string) error {
	// TODO: Not implemented yet
	return nil
}

// SetKey sets a key value.
func (config *Config) SetKey(key string, value string) error {
	parentObj, err := config.GetCategoryObject(ConfigCategoryKey)
	if err != nil {
		return err
	}

	obj := registry.NewObject()
	obj.ParentID = parentObj.ID
	obj.Name = key
	obj.Data = value

	return config.CreateObject(obj)
}

// GetString returns the specified key value.
func (config *Config) GetString(key string) (string, error) {
	keyObj, err := config.GetCategoryKeyObject(ConfigCategoryKey, key)
	if err != nil {
		return "", err
	}
	return keyObj.Data, nil
}

// GetInt returns the specified key value.
func (config *Config) GetInt(key string) (int, error) {
	keyStr, err := config.GetString(key)
	if err != nil {
		return 0, err
	}
	value64, err := strconv.ParseInt(keyStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(value64), nil
}

// ExecuteQuery must return the result as a standard array or map.
func (config *Config) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotFound)
	}

	configMap := map[string]string{}

	configObj, err := config.GetCategoryObject(ConfigCategoryKey)
	if err == nil {
		objs, err := config.GetChildObjects(configObj.ID)
		if err == nil {
			for _, obj := range objs {
				configMap[obj.Name] = obj.Data
			}
		}
	}

	return configMap, nil
}
