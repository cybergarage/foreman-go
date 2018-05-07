// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"os"
	"runtime"

	"github.com/cybergarage/go-graphite/net/graphite"

	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

const (
	serverBindRetryCount = 100
)

// Server represents a Foreman Server.
type Server struct {
	Node

	metric.RegisterListener
	kb.KnowledgeBaseListener

	graphite    *graphite.Server
	registerMgr *register.Manager
	registryMgr *registry.Manager
	qosMgr      *qos.Manager
	metricMgr   *metric.Manager
	actionMgr   *action.Manager

	config *Config

	configFile string
}

// NewServerWithConfigFile returns a new server with a specified configuration file.
func NewServerWithConfigFile(configFile string) *Server {

	server := &Server{
		graphite:    graphite.NewServer(),
		registryMgr: registry.NewManager(),
		registerMgr: register.NewManager(),
		qosMgr:      qos.NewManager(),
		metricMgr:   metric.NewManager(),
		actionMgr:   action.NewManager(),
		config:      nil,
		configFile:  configFile,
	}

	server.initialize()
	runtime.SetFinalizer(server, serverFinalizer)

	var err error
	server.config, err = NewConfigWithRegistry(server.registryMgr)
	if err != nil {
		logging.Fatal("Could not create config registry!")
		logging.Fatal("%s", err)
		return nil
	}

	err = server.LoadConfig(server.configFile)
	if err != nil {
		return nil
	}

	server.actionMgr.SetRegistryStore(server.registryMgr.GetStore())
	server.actionMgr.SetRegisterStore(server.registerMgr.GetStore())

	server.graphite.SetCarbonListener(server)
	server.graphite.SetRenderListener(server)
	FqlPath, err := server.config.GetString(ConfigFqlPathKey)
	if err != nil {
		FqlPath = HttpServerFqlPath
	}
	server.graphite.SetHTTPRequestListener(FqlPath, server)

	server.metricMgr.SetRegisterStore(server.registerMgr.GetStore())
	server.metricMgr.SetRegisterListener(server)

	server.qosMgr.AddListener(server)

	return server
}

func NewServer() *Server {
	return NewServerWithConfigFile("")
}

// initialize initialize the server.
func (server *Server) initialize() error {
	err := server.registryMgr.Start()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	return nil
}

// serverFinalizer closes all managers.
func serverFinalizer(server *Server) {
	server.registryMgr.Stop()
}

// updateConfig sets latest configurations.
func (server *Server) updateConfig() error {
	// Set latest network configurations

	port, err := server.config.GetInt(ConfigCarbonPortKey)
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}
	server.graphite.Carbon.Port = port

	port, err = server.config.GetInt(ConfigHttpPortKey)
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}
	server.graphite.Render.Port = port

	return err
}

func (server *Server) LoadConfig(filename string) error {
	logging.Trace("Server loading config file from %s.", filename)
	err := server.config.LoadFile(filename)
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}
	return server.updateConfig()
}

// GetHostname returns the hostname.
func (server *Server) GetHostname() (string, error) {
	return os.Hostname()
}

// LoadConfig loads a specified configuration file.
// GetGraphitePort returns the graphite carbon port.
func (server *Server) GetGraphitePort() int {
	return server.graphite.Carbon.Port
}

// GetHTTPPort returns the graphite HTTP port.
func (server *Server) GetHTTPPort() int {
	return server.graphite.Render.Port
}

// GetCuster returns the cluster name
func (server *Server) GetCuster() string {
	// TODO : Support cluster
	return ""
}

// GetName returns the host name
func (server *Server) GetName() string {
	name, err := server.GetHostname()
	if err != nil {
		return ""
	}
	return name
}

// GetAddress returns the interface address
func (server *Server) GetAddress() string {
	return server.graphite.GetAddress()
}

// GetRPCPort returns the RPC port
func (server *Server) GetRPCPort() int {
	return server.GetHTTPPort()
}

// PostQuery posts a query string
func (server *Server) PostQuery(query string) (interface{}, error) {
	parser := fql.NewParser()
	queries, err := parser.ParseString(query)
	if err != nil {
		return nil, err
	}

	queryCnt := len(queries)
	if queryCnt <= 0 {
		return nil, nil
	}

	var resObjects []interface{}
	if 1 < queryCnt {
		resObjects = make([]interface{}, queryCnt)
	}

	for n, query := range queries {
		resObject, queryErr := server.ExecuteQuery(query)
		if queryErr == nil {
			return nil, queryErr.Error()
		}
		if queryCnt == 1 {
			return resObject, nil
		}
		resObjects[n] = resObject
	}

	return resObjects, nil
}

// PostMetric posts a metric
func (server *Server) PostMetric(m *metric.Metric) error {
	return nodePostMetric(server, m)
}

// Start starts the server.
func (server *Server) Start() error {
	if server == nil {
		return nil
	}

	err := server.updateConfig()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	// Start all managers without the register manager

	graphiteRetryCount := 0
	err = server.graphite.Start()
	for (err != nil) || (graphiteRetryCount < serverBindRetryCount) {
		server.graphite.SetCarbonPort(server.graphite.GetCarbonPort() + 1)
		server.graphite.SetRenderPort(server.graphite.GetRenderPort() + 1)
		graphiteRetryCount++
		err = server.graphite.Start()
	}
	if err != nil {
		server.Stop()
		logging.Error("%s\n", err)
		return err
	}

	err = server.registerMgr.Start()
	if err != nil {
		server.Stop()
		logging.Error("%s\n", err)
		return err
	}

	err = server.metricMgr.Start()
	if err != nil {
		server.Stop()
		logging.Error("%s\n", err)
		return err
	}

	err = server.qosMgr.Start()
	if err != nil {
		server.Stop()
		logging.Error("%s\n", err)
		return err
	}

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	if server == nil {
		return nil
	}

	// Stop all managers without the register manager

	err := server.graphite.Stop()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	err = server.registerMgr.Stop()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	err = server.metricMgr.Stop()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	err = server.qosMgr.Stop()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	return nil
}

// Restart restats the server.
func (server *Server) Restart() error {
	logging.Info("Stopping server for restart...\n")
	err := server.Stop()
	if err != nil {
		logging.Error("Could not stop server: %s\n", err)
		return err
	}

	logging.Info("Reloading config...")
	server.LoadConfig(server.configFile)

	logging.Info("Restarting server...")
	err = server.Start()
	if err != nil {
		logging.Error("Could not start server: %s\n", err)
		return err
	}
	logging.Info("Restarted.")

	return nil
}
