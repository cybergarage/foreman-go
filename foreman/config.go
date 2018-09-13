// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// #include <foreman/foreman-c.h>
import "C"

import (
	"bufio"
	"os"
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/rpc"

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
	Boostrap   int
}

type FQLConfig struct {
	Path  string
	Query string
}

type MetricsConfig struct {
	Store    string
	Period   int
	Interval int
}

// Config represents a configuration.
type Config struct {
	Log     LogConfig
	Server  ServerConfig
	FQL     FQLConfig
	Metrics MetricsConfig
}

// NewDefaultConfig return a default configuration.
func NewDefaultConfig() *Config {
	conf := Config{}

	conf.Log.File = DefaultLogFile
	conf.Log.Level = DefaultLogLevel

	conf.Server.Host = DefaultServerHost
	conf.Server.HTTPPort = DefaultHttpPort
	conf.Server.CarbonPort = DefaultCarbonPort
	conf.Server.Boostrap = DefaultBoostrapMode

	conf.FQL.Path = HttpRequestFqlPath
	conf.FQL.Query = HttpRequestFqlQueryParam

	conf.Metrics.Store = DefaultMetricsStore
	conf.Metrics.Period = DefaultMetricsPeriod
	conf.Metrics.Interval = DefaultMetricsInterval

	return &conf
}

// NewConfigWithFile return a specified configuration.
func NewConfigWithFile(filename string) (*Config, error) {
	conf := NewDefaultConfig()
	err := conf.LoadFile(filename)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// GetLogLevel returns the log level type.
func (conf *Config) GetLogLevel() logging.LogLevel {
	return logging.LogLevelFromString(conf.Log.Level)
}

// LoadFile loads a specified Config file.
func (conf *Config) LoadFile(filename string) error {
	if filename != "" {
		logging.Trace("TOML Config file path: %s", filename)
		_, err := toml.DecodeFile(filename, conf)
		if err != nil {
			return err
		}
		logging.Trace("Got config: %s", filename)
	}
	return nil
}

func (conf *Config) ToFile(filename string) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		outfile.Close()
	}()
	out := toml.NewEncoder(bufio.NewWriter(outfile))
	return out.Encode(conf)
}

// ExecuteQuery must return the result as a standard array or map.
func (conf *Config) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	configMap := map[string]interface{}{}

	configMap[ConfigProductKey] = ProductName
	configMap[ConfigVersionKey] = Version

	configMap[rpc.ConfigHttpPortKey] = conf.Server.HTTPPort
	configMap[rpc.ConfigCarbonPortKey] = conf.Server.CarbonPort

	configMap[rpc.ConfigLogLevelKey] = conf.Log.Level

	configMap[rpc.ConfigBoostrapKey] = conf.Server.Boostrap

	configContainer := map[string]interface{}{}
	configContainer[strings.ToLower(fql.QueryTargetConfig)] = configMap

	return configContainer, nil
}
