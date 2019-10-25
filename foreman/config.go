// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"bufio"
	"os"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/util"
)

type LogConfig struct {
	File  string
	Level string
}

type ServerConfig struct {
	Cluster               string
	Host                  string
	CarbonPort            int
	HTTPPort              int
	ConnectionTimeout     int
	ConnectionWaitTimeout int
	Bootstrap             int
	Finder                string
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

type BootstrapConfig struct {
	Query string
}

type ProfileConfig struct {
	Port int
}

// Config represents a configuration.
type Config struct {
	Log       LogConfig
	Server    ServerConfig
	Finder    discovery.FinderConfig
	FQL       FQLConfig
	Metrics   MetricsConfig
	Bootstrap BootstrapConfig
	Profile   ProfileConfig
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
	conf.Server.Bootstrap = DefaultBootstrapMode
	conf.Server.ConnectionTimeout = DefaultConnectionTimeout
	conf.Server.ConnectionWaitTimeout = DefaultConnectionWaitTimeout
	conf.Server.Finder = DefaultFinder

	conf.Finder.Hosts = []string{}

	conf.FQL.Path = HttpRequestFqlPath
	conf.FQL.Query = HttpRequestFqlQueryParam

	conf.Metrics.Store = DefaultMetricsStore
	conf.Metrics.Period = DefaultMetricsPeriod
	conf.Metrics.Interval = DefaultMetricsInterval

	conf.Bootstrap.Query = ""

	conf.Profile.Port = DefaultProfilePort

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

// SetServerConnectionTimeout sets a timeout for all server connections.
func (conf *Config) SetServerConnectionTimeout(d time.Duration) error {
	conf.Server.ConnectionTimeout = int(d / time.Second)
	return nil
}

// GetServerConnectionTimeout returns a timeout for all server connections.
func (conf *Config) GetServerConnectionTimeout() time.Duration {
	return time.Second * time.Duration(conf.Server.ConnectionTimeout)
}

// SetServerConnectionWaitTimeout sets a timeout for all server connections.
func (conf *Config) SetServerConnectionWaitTimeout(d time.Duration) error {
	conf.Server.ConnectionWaitTimeout = int(d / time.Second)
	return nil
}

// GetServerConnectionWaitTimeout returns a timeout for all server connections.
func (conf *Config) GetServerConnectionWaitTimeout() time.Duration {
	return time.Second * time.Duration(conf.Server.ConnectionWaitTimeout)
}

// SetMetricsStorePeriod sets a metrics period duration.
func (conf *Config) SetMetricsStorePeriod(d time.Duration) error {
	if d != 0 {
		conf.Metrics.Period = int(d / time.Second)
	} else {
		conf.Metrics.Period = 0
	}
	return nil
}

// GetMetricsStorePeriod returns the metrics period duration.
func (conf *Config) GetMetricsStorePeriod() time.Duration {
	return time.Second * time.Duration(conf.Metrics.Period)
}

// SetMetricsStoreInterval sets a metrics intervale.
func (conf *Config) SetMetricsStoreInterval(d time.Duration) error {
	conf.Metrics.Interval = int(d / time.Second)
	return nil
}

// GetMetricsStoreInterval returns the metrics interval.
func (conf *Config) GetMetricsStoreInterval() time.Duration {
	return time.Second * time.Duration(conf.Metrics.Interval)
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

// SetBootstrapEnabled set the boostrap flag.
func (conf *Config) SetBootstrapEnabled(flag bool) {
	if flag {
		conf.Server.Bootstrap = 1
		return
	}
	conf.Server.Bootstrap = 0
}

// IsBootstrapEnabled returns true when the boostrap flag is true, otherwise false.
func (conf *Config) IsBootstrapEnabled() bool {
	if conf.Server.Bootstrap == 0 {
		return false
	}
	return true
}

// SetBootstrapQuery set a boostrap query file or directory.
func (conf *Config) SetBootstrapQuery(query string) bool {
	conf.Bootstrap.Query = query
	return true
}

// HasBootstrapQuery returns true when the query file is specified, otherwise false.
func (conf *Config) HasBootstrapQuery() bool {
	if len(conf.Bootstrap.Query) <= 0 {
		return false
	}
	return true
}

// GetBootstrapQueryFiles returns all boostrap file name.
func (conf *Config) GetBootstrapQueryFiles() ([]*util.File, error) {
	queryFiles := make([]*util.File, 0)

	if len(conf.Bootstrap.Query) <= 0 {
		return queryFiles, nil
	}

	queryFile := util.NewFileWithPath(conf.Bootstrap.Query)
	if !queryFile.IsDir() {
		return append(queryFiles, util.NewFileWithPath(queryFile.Path)), nil
	}

	queryFiles, err := queryFile.ListFilesWithExtention(action.ActionFileExtention)
	if err != nil {
		return nil, err
	}

	return queryFiles, nil
}

// ToFile stores the current configuration to a file.
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
