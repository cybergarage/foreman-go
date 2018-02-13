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
	"github.com/cybergarage/foreman-go/foreman/kb"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Server represents a Foreman Server.
type Server struct {
	metric.RegisterListener
	kb.RuleListener

	graphite    *graphite.Server
	registerMgr *register.Manager
	registryMgr *registry.Manager
	qosMgr      *qos.Manager
	metricMgr   *metric.Manager
	actionMgr   *action.Manager

	config *Config
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{
		graphite:    graphite.NewServer(),
		registryMgr: registry.NewManager(),
		registerMgr: register.NewManager(),
		qosMgr:      qos.NewManager(),
		metricMgr:   metric.NewManager(),
		actionMgr:   action.NewManager(),
		config:      nil,
	}

	server.initialize()
	runtime.SetFinalizer(server, serverFinalizer)

	server.config = NewConfigWithRegistry(server.registryMgr)

	server.graphite.CarbonListener = server

	server.graphite.RenderListener = server
	server.graphite.SetHTTPRequestListener(HttpServerFqlPath, server)

	server.metricMgr.SetRegisterStore(server.registerMgr.GetStore())
	server.metricMgr.SetRegisterListener(server)

	server.actionMgr.AddRouteContainer(server.qosMgr)

	return server
}

// GetHostname returns the hostname.
func (server *Server) GetHostname() (string, error) {
	return os.Hostname()
}

// LoadConfig loads a specified configuration file.
func (server *Server) LoadConfig(filename string) error {
	err := server.config.LoadFile(filename)
	if err != nil {
		return err
	}
	return nil
}

// initialize initialize the server.
func (server *Server) initialize() error {
	err := server.registryMgr.Start()
	if err != nil {
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
		return err
	}
	server.graphite.Carbon.Port = port

	port, err = server.config.GetInt(ConfigHttpPortKey)
	if err != nil {
		return err
	}
	server.graphite.Render.Port = port

	return err
}

// GetGraphitePort returns the graphite carbon port.
func (server *Server) GetGraphitePort() int {
	return server.graphite.Carbon.Port
}

// GetHTTPPort returns the graphite HTTP port.
func (server *Server) GetHTTPPort() int {
	return server.graphite.Render.Port
}

// Start starts the server.
func (server *Server) Start() error {
	err := server.updateConfig()
	if err != nil {
		return err
	}

	// Start all managers without the register manager

	err = server.graphite.Start()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.registerMgr.Start()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.metricMgr.Start()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.qosMgr.Start()
	if err != nil {
		server.Stop()
		return err
	}

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	// Stop all managers without the register manager

	err := server.graphite.Stop()
	if err != nil {
		return err
	}

	err = server.registerMgr.Stop()
	if err != nil {
		return err
	}

	err = server.metricMgr.Stop()
	if err != nil {
		return err
	}

	err = server.qosMgr.Stop()
	if err != nil {
		return err
	}

	return nil
}

// Restart restats the server.
func (server *Server) Restart() error {
	err := server.Stop()
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	return nil
}
