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
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/register"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

// Server represents a Foreman Server.
type Server struct {
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

// NewServer returns a new Server.
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

	server.graphite.CarbonListener = server
	server.graphite.RenderListener = server
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

// GetHostname returns the hostname.
func (server *Server) GetHostname() (string, error) {
	return os.Hostname()
}

// LoadConfig loads a specified configuration file.
func (server *Server) LoadConfig(filename string) error {
	logging.Trace("Server loading config file from %s.", filename)
	err := server.config.LoadFile(filename)
	if err != nil {
		logging.Error("%s\n", err)
		return err
	}
	return server.updateConfig()
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
		logging.Error("%s\n", err)
		return err
	}

	// Start all managers without the register manager

	err = server.graphite.Start()
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
