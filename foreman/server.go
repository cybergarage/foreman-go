// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"net/http"

	"github.com/cybergarage/go-graphite/net/graphite"
)

const (
	serverFQLPath       = "/fql"
	serverFQLQueryParam = "q"
)

// Server represents a Foreman Server.
type Server struct {
	graphite *graphite.Server
	store    *Store
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{}

	server.graphite = graphite.NewServer()
	server.graphite.CarbonListener = server
	server.graphite.RenderListener = server
	server.graphite.SetHTTPRequestListener(serverFQLPath, server)

	server.store = NewStore()

	return server
}

// Start starts the server.
func (self *Server) Start() error {
	err := self.graphite.Start()
	if err != nil {
		self.Stop()
		return err
	}

	err = self.store.Open()
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

	err = self.store.Close()
	if err != nil {
		self.Stop()
		return err
	}

	return nil
}

// MetricRequestReceived is a listener for Graphite Carbon
func (self *Server) MetricRequestReceived(gm *graphite.Metric, err error) {
	// Ignore error requests
	if err != nil {
		return
	}

	// graphite.Metric to foreman.Metric
	fm := NewMetric()
	fm.Name = gm.Path
	fm.Timestamp = gm.Timestamp
	fm.Value = gm.Value

	err = self.store.AddMetric(fm)
	if err != nil {
		// TODO : Handle the error
	}
}

// QueryRequestReceived is a listener for Graphite Render
func (self *Server) QueryRequestReceived(gq *graphite.Query, err error) ([]*graphite.Metric, error) {
	// Ignore error requests
	if err != nil {
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := NewQuery()
	fq.Target = gq.Target
	fq.From = gq.From
	fq.Until = gq.Until

	_, err = self.store.Query(fq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// HTTPRequestReceived is a listener for FQL
func (self *Server) HTTPRequestReceived(r *http.Request, w http.ResponseWriter) {
	switch r.URL.Path {
	case serverFQLPath:
		self.fqlRequestReceived(r, w)
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// fqlRequestReceived handles FQL requests
func (self *Server) fqlRequestReceived(r *http.Request, w http.ResponseWriter) {
	self.badRequestReceived(r, w)
}

// badRequestReceived returns http.StatusBadRequest
func (self *Server) badRequestReceived(r *http.Request, w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
