// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"os"
	"runtime"

	"github.com/cybergarage/foreman-go/foreman/node"

	"github.com/cybergarage/go-graphite/net/graphite"

	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/fd"
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
	*baseNode

	cluster string
	name    string

	metric.RegisterListener
	kb.KnowledgeBaseListener

	*Controller
	fd fd.Detector

	graphite    *graphite.Server
	registerMgr *register.Manager
	registryMgr *registry.Manager
	qosMgr      *qos.Manager
	metricMgr   *metric.Manager
	actionMgr   *action.Manager

	config     *Config
	configFile string

	status *Status
}

// NewServerWithConfig returns a new server with a specified configuration.
func NewServerWithConfig(conf *Config) (*Server, error) {

	server := &Server{
		cluster:     "",
		name:        "",
		Controller:  NewController(),
		graphite:    graphite.NewServer(),
		registryMgr: registry.NewManager(),
		registerMgr: register.NewManager(),
		qosMgr:      qos.NewManager(),
		metricMgr:   nil,
		actionMgr:   action.NewManager(),
		config:      conf,
		configFile:  "",
		status:      NewStatus(),
		fd:          fd.NewGossipDetector(),
	}

	server.baseNode = newBaseNodeWithNode(server)

	server.initialize()
	runtime.SetFinalizer(server, serverFinalizer)

	hostname, err := os.Hostname()
	if err == nil {
		server.name = hostname
	}

	// Metric Store

	metricStore, err := metric.NewStoreWithName(conf.Server.MetricsStore)
	if err != nil {
		return nil, err
	}
	server.metricMgr = metric.NewManagerWithStore(metricStore)

	// Registry Store

	server.actionMgr.SetRegistryStore(server.registryMgr.GetStore())
	server.actionMgr.SetRegisterStore(server.registerMgr.GetStore())

	// Register Store

	server.metricMgr.SetRegisterStore(server.registerMgr.GetStore())
	server.metricMgr.SetRegisterListener(server)

	// Graphite

	server.graphite.SetCarbonListener(server)
	server.graphite.SetRenderListener(server)

	// QoS

	server.qosMgr.AddListener(server)

	// Failure Detector

	server.fd.SetListener(server)

	// RPC

	fqlPath := server.config.FQL.Path
	server.graphite.SetHTTPRequestListener(fqlPath, server)

	return server, nil
}

// NewServerWithConfigFile returns a new server with a specified configuration file.
func NewServerWithConfigFile(confFile string) (*Server, error) {
	conf := NewDefaultConfig()

	err := conf.LoadFile(confFile)
	if err != nil {
		return nil, err
	}

	server, err := NewServerWithConfig(conf)
	if err != nil {
		return nil, err
	}

	return server, nil
}

// NewServer creates a new server
func NewServer() *Server {
	conf := NewDefaultConfig()
	conf.Server.MetricsStore = metric.MetricStoreSqlite
	server, _ := NewServerWithConfig(conf)
	return server
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

// SetCluster sets a cluster name
func (server *Server) SetCluster(name string) {
	server.cluster = name
}

// GetCluster returns the cluster name
func (server *Server) GetCluster() string {
	return server.cluster
}

// SetName sets a host name
func (server *Server) SetName(name string) {
	server.name = name
}

// GetName returns the host name
func (server *Server) GetName() string {
	return server.name
}

// GetHTTPPort returns the graphite HTTP port.
func (server *Server) GetHTTPPort() int {
	return server.graphite.Render.Port
}

// GetCarbonPort returns the graphite carbon port.
func (server *Server) GetCarbonPort() int {
	return server.graphite.Carbon.Port
}

// GetRenderPort returns the graphite render port.
func (server *Server) GetRenderPort() int {
	return server.graphite.Render.Port
}

// GetAddress returns the interface address
func (server *Server) GetAddress() string {
	return server.graphite.GetAddress()
}

// GetRPCPort returns the RPC port
func (server *Server) GetRPCPort() int {
	return server.GetHTTPPort()
}

// GetUniqueID returns a unique ID of the node
func (server *Server) GetUniqueID() string {
	return node.GetUniqueID(server)
}

// updateConfig sets latest configurations.
func (server *Server) updateConfig() error {
	// Set latest network configurations
	server.graphite.Carbon.Port = server.config.Server.CarbonPort
	server.graphite.Render.Port = server.config.Server.HTTPPort

	return nil
}

// LoadConfig sets the configurations in the specified file.
func (server *Server) LoadConfig(filename string) error {
	logging.Trace("Server loading config file from %s.", filename)
	err := server.config.LoadFile(filename)
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}
	return server.updateConfig()
}

// GetAllClusterNodes returns all nodes in the same cluster
func (server *Server) GetAllClusterNodes() []Node {
	// Return only the server if the server has no finder

	if !server.Controller.HasFinders() {
		return []Node{server}
	}

	// The server is included if the server has any finders

	allNodes, err := server.Controller.GetAllNodes()
	if err != nil {
		return make([]Node, 0)
	}

	clusterName := server.GetCluster()
	clusterNodes := make([]Node, 0)
	for _, node := range allNodes {
		if node.GetCluster() != clusterName {
			continue
		}
		clusterNodes = append(clusterNodes, node)
	}
	return clusterNodes
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
		if queryErr != nil {
			return nil, queryErr.Error()
		}
		if queryCnt == 1 {
			return resObject, nil
		}
		resObjects[n] = resObject
	}

	return resObjects, nil
}

// getStartupManagers returns all managers which should be started
func (server *Server) getStartupManagers() []Manager {
	managers := []Manager{
		server.registerMgr,
		server.metricMgr,
		server.qosMgr,
	}
	return managers
}

// Start starts the server.
func (server *Server) Start() error {
	err := server.updateConfig()
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}

	// Start all managers without graphite

	for _, mgr := range server.getStartupManagers() {
		err = mgr.Start()
		if err != nil {
			return err
		}
	}

	// Start Graphite manager

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

	// Boostrap

	boostrap := server.config.Server.Boostrap
	if boostrap != 0 {
		err = server.executeBoostrap()
		server.Stop()
		logging.Error("%s\n", err)
		return err
	}

	return nil
}

// getRunningMangers returns all managers which should be stopped.
func (server *Server) getRunningMangers() []Manager {
	managers := server.getStartupManagers()
	managers = append(managers, server.graphite)
	return managers
}

// Stop stops the server.
func (server *Server) Stop() error {

	// Start all managers

	var lastError error
	for _, mgr := range server.getRunningMangers() {
		err := mgr.Stop()
		if err != nil {
			lastError = err
		}
	}

	return lastError
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
