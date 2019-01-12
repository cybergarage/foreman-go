// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"os"
	"runtime"
	"time"

	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/fd"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/node"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
	"github.com/cybergarage/go-graphite/net/graphite"
)

const (
	serverBindRetryCount = 100
)

const (
	serverLoadingQueryFileMessageFormat = "Loading : %s"
	serverExecutingQueryMessageFormat   = "[%d] %s"
)

// Server represents a Foreman Server.
type Server struct {
	Node
	*baseNode

	metric.RegisterListener
	kb.KnowledgeBaseListener

	*Controller
	fd fd.Detector

	graphiteMgr *graphite.Manager
	registerMgr *register.Manager
	registryMgr *registry.Manager
	qosMgr      *qos.Manager
	metricMgr   *metric.Manager
	actionMgr   *action.Manager

	*Config
	configFile string
	queryFile  string

	status *Status
}

// NewServerWithConfig returns a new server with a specified configuration.
func NewServerWithConfig(conf *Config) (*Server, error) {

	server := &Server{
		Controller:  NewController(),
		graphiteMgr: graphite.NewManager(),
		registryMgr: registry.NewManager(),
		registerMgr: register.NewManager(),
		qosMgr:      qos.NewManager(),
		metricMgr:   nil,
		actionMgr:   action.NewManager(),
		Config:      conf,
		configFile:  "",
		queryFile:   "",
		status:      NewStatus(),
		fd:          fd.NewGossipDetector(),
	}

	server.baseNode = newBaseNodeWithNode(server)

	runtime.SetFinalizer(server, serverFinalizer)

	// Graphite

	server.graphiteMgr.SetAutoInterfaceBindingEnabled(true)
	server.graphiteMgr.SetEachInterfaceBindingEnabled(true)
	server.graphiteMgr.SetCarbonListener(server)
	server.graphiteMgr.SetRenderListener(server)

	// Registry Store

	server.actionMgr.SetRegistryStore(server.registryMgr.GetStore())
	server.actionMgr.SetRegisterStore(server.registerMgr.GetStore())

	// QoS

	server.qosMgr.AddListener(server)

	// Failure Detector

	server.fd.SetListener(server)

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
	server, _ := NewServerWithConfig(conf)
	return server
}

// serverFinalizer closes all managers.
func serverFinalizer(server *Server) {
	server.Stop()
}

// GetController returns a controller of the node.
func (server *Server) GetController() *Controller {
	return server.Controller
}

// SetCluster sets a cluster name
func (server *Server) SetCluster(name string) {
	server.Server.Cluster = name
}

// GetCluster returns the cluster name
func (server *Server) GetCluster() string {
	return server.Server.Cluster
}

// SetName sets a host name
func (server *Server) SetName(name string) {
	server.Server.Host = name
}

// GetName returns the host name
func (server *Server) GetName() string {
	host := server.Server.Host
	if 0 < len(host) {
		return host
	}

	hostname, err := os.Hostname()
	if err == nil {
		return hostname
	}

	return DefaultServerHost
}

// GetHTTPPort returns the graphite HTTP port.
func (server *Server) GetHTTPPort() int {
	return server.GetRenderPort()
}

// GetCarbonPort returns the graphite carbon port.
func (server *Server) GetCarbonPort() int {
	return server.graphiteMgr.GetCarbonPort()
}

// GetRenderPort returns the graphite render port.
func (server *Server) GetRenderPort() int {
	return server.graphiteMgr.GetRenderPort()
}

// GetAddress returns the interface address
func (server *Server) GetAddress() string {
	addr, err := server.graphiteMgr.GetBoundAddress()
	if err != nil {
		return ""
	}
	return addr
}

// GetRPCPort returns the RPC port
func (server *Server) GetRPCPort() int {
	return server.GetHTTPPort()
}

// GetUniqueID returns a unique ID of the node
func (server *Server) GetUniqueID() string {
	return node.GetUniqueID(server)
}

// SetConfigFile sets a initial configuration file for reloading.
func (server *Server) SetConfigFile(filename string) {
	server.configFile = filename
}

// GetConfigFile returns a configuration file for reloading.
func (server *Server) GetConfigFile() string {
	return server.configFile
}

// LoadConfig sets the configurations in the specified file.
func (server *Server) LoadConfig(filename string) error {
	logging.Trace("Server loading config file from %s.", filename)
	return server.LoadFile(filename)
}

// SetQueryFile sets a initial query file for reloading.
func (server *Server) SetQueryFile(filename string) {
	server.queryFile = filename
}

// GetQueryFile return a initial query file for reloading.
func (server *Server) GetQueryFile() string {
	return server.queryFile
}

// LoadQuery executes the queries in the specified file.
func (server *Server) LoadQuery(filename string) error {
	logging.Info(serverLoadingQueryFileMessageFormat, filename)

	loader := fql.NewLoader()
	queries, err := loader.LoadFromFile(filename)
	if err != nil {
		return err
	}

	for n, query := range queries {
		logging.Info(serverExecutingQueryMessageFormat, n, query.String())
		_, err := server.ExecuteQuery(query)
		if err != nil {
			logging.Error(err.Error().Error())
			return err.Error()
		}
	}

	return nil
}

// GetAllClusterNodes returns all nodes in the same cluster
func (server *Server) GetAllClusterNodes() []Node {
	// Return only the server if the server has no finder

	if !server.Controller.HasFinder() {
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

// applyConfig sets latest configurations to the server
func (server *Server) applyConfig() error {
	// Controller

	switch server.Server.Finder {
	case FinderEchonet:
		server.Controller.SetFinder(discovery.NewEchonetFinderWithLocalNode(server))
	}

	// Metric Store

	metricStore, err := metric.NewStoreWithName(server.Metrics.Store)
	if err != nil {
		return err
	}
	metricStore.SetRetentionPeriod(time.Duration(server.Metrics.Period) * time.Second)
	metricStore.SetRetentionInterval(time.Duration(server.Metrics.Interval) * time.Second)

	server.metricMgr = metric.NewManagerWithStore(metricStore)
	server.metricMgr.SetRegisterStore(server.registerMgr.GetStore())
	server.metricMgr.SetRegisterListener(server)

	// Graphite Ports

	server.graphiteMgr.SetCarbonPort(server.Server.CarbonPort)
	server.graphiteMgr.SetRenderPort(server.Server.HTTPPort)

	// RPC

	server.graphiteMgr.SetHTTPRequestListener(server.FQL.Path, server)

	return nil
}

// getStartupManagers returns all managers which should be started
func (server *Server) getStartupManagers() []Manager {
	managers := []Manager{
		server.registryMgr,
		server.registerMgr,
		server.metricMgr,
		server.qosMgr,
		server.actionMgr,
	}
	return managers
}

// Start starts the server.
func (server *Server) Start() error {
	// Apply configuration

	err := server.applyConfig()
	if err != nil {
		return err
	}

	// Start all managers without graphite

	for _, mgr := range server.getStartupManagers() {
		err := mgr.Start()
		if err != nil {
			return err
		}
	}

	// Start Graphite manager

	graphiteRetryCount := 0
	err = server.graphiteMgr.Start()
	for (err != nil) && (graphiteRetryCount < serverBindRetryCount) {
		server.graphiteMgr.SetCarbonPort(server.graphiteMgr.GetCarbonPort() + 1)
		server.graphiteMgr.SetRenderPort(server.graphiteMgr.GetRenderPort() + 1)
		graphiteRetryCount++
		err = server.graphiteMgr.Start()
	}

	if err != nil {
		server.Stop()
		logging.Error("%s", err)
		return err
	}

	// Controller

	err = server.Controller.Start()
	if err != nil {
		server.Stop()
		logging.Error("%s", err)
		return err
	}

	err = server.Controller.Search()
	if err != nil {
		server.Stop()
		logging.Error("%s", err)
		return err
	}

	// Bootstrap

	err = server.executeBootstrap()
	if err != nil {
		server.Stop()
		logging.Error("%s", err)
		return err
	}

	return nil
}

// getRunningMangers returns all managers which should be stopped.
func (server *Server) getRunningMangers() []Manager {
	managers := server.getStartupManagers()
	managers = append(managers, server.graphiteMgr)
	return managers
}

// Stop stops the server.
func (server *Server) Stop() error {

	var lastError error

	// Controller

	err := server.Controller.Stop()
	if err != nil {
		logging.Error("%s", err)
		lastError = err
	}

	// Start all managers

	for _, mgr := range server.getRunningMangers() {
		err := mgr.Stop()
		if err != nil {
			logging.Error("%s", err)
			lastError = err
		}
	}

	return lastError
}
