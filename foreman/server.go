// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"net/http"
	"os"

	"github.com/cybergarage/go-graphite/net/graphite"

	"github.com/cybergarage/foreman-go/foreman/action"
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/qos"
	"github.com/cybergarage/foreman-go/foreman/registry"
)

const (
	serverFQLPath       = "/fql"
	serverFQLQueryParam = "q"
)

// Server represents a Foreman Server.
type Server struct {
	graphite      *graphite.Server
	registryStore *registry.Store
	qosMgr        *qos.Manager
	metricMgr     *metric.Manager
	actionMgr     *action.Manager
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{
		graphite:      graphite.NewServer(),
		registryStore: registry.NewStore(),
		qosMgr:        qos.NewManager(),
		metricMgr:     metric.NewManager(),
		actionMgr:     action.NewManager(),
	}

	server.graphite.CarbonListener = server
	server.graphite.RenderListener = server
	server.graphite.SetHTTPRequestListener(serverFQLPath, server)

	server.metricMgr.SetRegisterListener(server.qosMgr)

	return server
}

// GetHostname returns the hostname.
func (server *Server) GetHostname() (string, error) {
	return os.Hostname()
}

// LoadConfig loads a specified configuration file.
func (server *Server) LoadConfig(filename string) error {
	config, err := NewConfigFromFile(filename)
	if err != nil {
		return err
	}

	return server.setConfig(config)
}

// setConfig sets the specified configuration to the server.
func (server *Server) setConfig(config *Config) error {
	// FIXE : Not Implemented yet
	return nil
}

// Start starts the server.
func (server *Server) Start() error {
	err := server.graphite.Start()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.registryStore.Open()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.qosMgr.Start()
	if err != nil {
		server.Stop()
		return err
	}

	err = server.metricMgr.Start()
	if err != nil {
		server.Stop()
		return err
	}

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	err := server.graphite.Stop()
	if err != nil {
		return err
	}

	err = server.registryStore.Close()
	if err != nil {
		return err
	}

	err = server.qosMgr.Stop()
	if err != nil {
		return err
	}

	err = server.metricMgr.Stop()
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

// HTTPRequestReceived is a listener for FQL
func (server *Server) HTTPRequestReceived(r *http.Request, w http.ResponseWriter) {
	switch r.URL.Path {
	case serverFQLPath:
		server.fqlRequestReceived(r, w)
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// fqlRequestReceived handles FQL requests
func (server *Server) fqlRequestReceived(r *http.Request, w http.ResponseWriter) {
	server.badRequestReceived(r, w)
}

// badRequestReceived returns http.StatusBadRequest
func (server *Server) badRequestReceived(r *http.Request, w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
