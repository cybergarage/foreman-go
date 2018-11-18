// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// #include <foreman/foreman-c.h>
import "C"

import (
	"bufio"
	"os"

	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/BurntSushi/toml"
)

type LogConfig struct {
	File  string
	Level string
}

type ServerConfig struct {
	Cluster    string
	Host       string
	CarbonPort int
	HTTPPort   int
	Boostrap   int
	Finder     string
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

	conf.Server.Cluster = DefaultCluster
	conf.Server.Host = DefaultServerHost
	conf.Server.HTTPPort = DefaultHttpPort
	conf.Server.CarbonPort = DefaultCarbonPort
	conf.Server.Boostrap = DefaultBoostrapMode
	conf.Server.Finder = DefaultFinder

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

// SetBoostrapEnabled set the boostrap flag.
func (conf *Config) SetBoostrapEnabled(flag bool) {
	if flag {
		conf.Server.Boostrap = 1
		return
	}
	conf.Server.Boostrap = 0
}

// IsBoostrapEnabled returns true when the boostrap flag is true, otherwise false.
func (conf *Config) IsBoostrapEnabled() bool {
	if conf.Server.Boostrap == 0 {
		return false
	}
	return true
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
