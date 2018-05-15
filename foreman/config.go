// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// #include <foreman/foreman-c.h>
import "C"
import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/registry"

	"github.com/BurntSushi/toml"
)

type LogConfig struct {
	File  string
	Level string
}

type ServerConfig struct {
	Host       string
	CarbonPort int
	HTTPPort   int
}

type FQLConfig struct {
	Path  string
	Query string
}

func DefaultTOMLConfig() TOMLConfig {
	t := TOMLConfig{}
	t.Log.File = DefaultLogFile
	t.Log.Level = DefaultLogLevel

	t.Server.Host = DefaultServerHost
	t.Server.HTTPPort = DefaultHttpPort
	t.Server.CarbonPort = DefaultCarbonPort

	t.FQL.Path = HttpRequestFqlPath
	t.FQL.Query = HttpRequestFqlQueryParam
	return t
}

// type TOMLConfig represents a configuration read from or to be written to a TOML file
type TOMLConfig struct {
	Log    LogConfig
	Server ServerConfig
	FQL    FQLConfig
}

// Config represents a Config for Foreman.
type Config struct {
	*registry.Manager
	fql.QueryExecutor
}

// NewConfigWithRegistry returns a new Config with the specified registry.
func NewConfigWithRegistry(mgr *registry.Manager) (*Config, error) {
	config := &Config{
		Manager: mgr,
	}
	err := config.initialize()
	return config, err
}

func (config *Config) initialize() error {
	err := config.CreateCategoryObject(ConfigCategoryKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
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
			logging.Trace("Error: %s\n", err)
			return err
		}
	}

	return nil
}

func (config *Config) loadTOMLConfig(t TOMLConfig) (err error) {
	err = config.SetKey(ConfigLogFileKey, t.Log.File)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	err = config.SetKey(ConfigLogLevelKey, t.Log.Level)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}

	err = config.SetKey(ConfigHostKey, t.Server.Host)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	err = config.SetKey(ConfigCarbonPortKey, strconv.Itoa(t.Server.CarbonPort))
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	err = config.SetKey(ConfigHttpPortKey, strconv.Itoa(t.Server.HTTPPort))
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}

	err = config.SetKey(ConfigFqlPathKey, t.FQL.Path)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	err = config.SetKey(ConfigFqlQueryKey, t.FQL.Query)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	return
}

func (config Config) toTOMLConfig() (t TOMLConfig, err error) {
	t = TOMLConfig{}
	t.Log.File, err = config.GetString(ConfigLogFileKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	t.Log.Level, err = config.GetString(ConfigLogLevelKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}

	t.Server.Host, err = config.GetString(ConfigHostKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	t.Server.CarbonPort, err = config.GetInt(ConfigCarbonPortKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	t.Server.HTTPPort, err = config.GetInt(ConfigHttpPortKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}

	t.FQL.Path, err = config.GetString(ConfigFqlPathKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	t.FQL.Query, err = config.GetString(ConfigFqlQueryKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
		return
	}
	return
}

// LoadFile loads a specified Config file.
func (config *Config) LoadFile(filename string) error {
	t := DefaultTOMLConfig()
	if filename != "" {
		logging.Trace("TOML Config file path: %s", filename)
		_, err := toml.DecodeFile(filename, &t)
		if err != nil {
			return err
		}
		logging.Trace("Got config: %s", filename, t)
	}
	logging.SetLogLevel(logging.LogLevelFromString(t.Log.Level))
	logging.SetLogFile(t.Log.File)
	config.loadTOMLConfig(t)
	return nil
}

func (config Config) ToFile(filename string) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		outfile.Close()
	}()
	out := toml.NewEncoder(bufio.NewWriter(outfile))
	t, err := config.toTOMLConfig()
	if err != nil {
		return err
	}
	return out.Encode(t)
}

// SetKey sets a key value.
func (config *Config) SetKey(key string, value string) error {
	parentObj, err := config.GetCategoryObject(ConfigCategoryKey)
	if err != nil {
		logging.Trace("Error: %s\n", err)
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
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
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

	configContainer := map[string]interface{}{}
	configContainer[strings.ToLower(fql.QueryTargetConfig)] = configMap

	return configContainer, nil
}
