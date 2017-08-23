// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/go-graphite/net/graphite"
)

// Server represents a Foreman Server.
type Server struct {
	graphite *graphite.Server
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{}

	server.graphite = graphite.NewServer()
	server.graphite.CarbonListener = server
	server.graphite.RenderListener = server

	return server
}

// Start starts the server.
func (self *Server) Start() error {
	err := self.graphite.Start()
	if err != nil {
		self.Stop()
		return err
	}
	return nil
}

// Stop stops the server.
func (self *Server) Stop() error {
	err := self.graphite.Stop()
	if err != nil {
		return err
	}

	return nil
}

// MetricRequestReceived is a listener for Graphite Carbon
func (self *Server) MetricRequestReceived(m *graphite.Metric, err error) {
	if err != nil {
		return
	}
}

// QueryRequestReceived is a listener for Graphite Render
func (self *Server) QueryRequestReceived(query *graphite.Query, err error) ([]*graphite.Metric, error) {
	if err != nil {
		return nil, nil
	}
	return nil, nil
}
