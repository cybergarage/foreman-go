// Copyright (C) 2017 Satoshi Konno. All rights reserved.
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
	graphite    *graphite.Server
	metricStore *MetricStore
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{}

	server.graphite = graphite.NewServer()
	server.graphite.CarbonListener = server
	server.graphite.RenderListener = server
	server.graphite.SetHTTPRequestListener(serverFQLPath, server)

	server.metricStore = NewMetricStore()

	return server
}

// LoadConfig loads a specified configutation file.
func (self *Server) LoadConfig(filename string) error {
	config, err := NewConfigFromFile(filename)
	if err != nil {
		return err
	}

	return self.setConfig(config)
}

// setConfig sets the specified configutations to the server.
func (self *Server) setConfig(config *Config) error {
	// FIXE : Not Implemented yet
	return nil
}

// Start starts the server.
func (self *Server) Start() error {
	err := self.graphite.Start()
	if err != nil {
		self.Stop()
		return err
	}

	err = self.metricStore.Open()
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

	err = self.metricStore.Close()
	if err != nil {
		self.Stop()
		return err
	}

	return nil
}

// Restart restats the server.
func (self *Server) Restart() error {
	err := self.Stop()
	if err != nil {
		return err
	}

	err = self.Start()
	if err != nil {
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

	for _, dp := range gm.DataPoints {
		// graphite.Metric to foreman.Metric
		fm := NewMetric()
		fm.Name = gm.Name
		fm.Timestamp = dp.Timestamp
		fm.Value = dp.Value

		err = self.metricStore.AddMetric(fm)
		if err != nil {
			// TODO : Handle the error
		}
	}
}

// QueryRequestReceived is a listener for Graphite Render
func (self *Server) QueryRequestReceived(gq *graphite.Query, err error) ([]*graphite.Metric, error) {
	// Ignore error requests
	if err != nil {
		return nil, nil
	}

	// graphite.Query to foreman.Query
	fq := NewMetricQuery()
	fq.Target = gq.Target
	fq.From = gq.From
	fq.Until = gq.Until

	rs, err := self.metricStore.Query(fq)
	if err != nil {
		return nil, err
	}

	mCount := rs.GetDataPointCount()
	m := make([]*graphite.Metric, mCount)

	dps := rs.GetFirstDataPoints()
	for n := 0; n < mCount; n++ {
		m[n] = graphite.NewMetric()
		if dps == nil {
			break
		}
		m[n].Name = dps.Name
		dpCount := len(dps.Values)
		m[n].DataPoints = graphite.NewDataPoints(dpCount)
		for i := 0; i < dpCount; i++ {
			dp := graphite.NewDataPoint()
			dp.Timestamp = dps.Values[i].Timestamp
			dp.Value = dps.Values[i].Value
			m[n].DataPoints[i] = dp
		}

		dps = rs.GetNextDataPoints()
	}

	return m, nil
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
