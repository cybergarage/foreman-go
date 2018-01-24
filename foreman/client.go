// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	rpc "github.com/cybergarage/foreman-go/foreman/rpc/graphite"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// Client represents a client for the server.
type Client struct {
	Host     string
	graphite *graphite.Client
}

// NewClient returns a new client.
func NewClient() *Client {
	client := &Client{
		Host:     DefaultServerHost,
		graphite: graphite.NewClient(),
	}
	return client
}

// Dial tries to connect to the specified host.
func (client *Client) Dial(host string) error {
	return nil
}

// SendMetric send a new metric.
func (client *Client) SendMetric(host string, m *metric.Metric) error {
	return nil
}

// PostMetric posts a metric over Graphite interface
func (client *Client) PostMetric(m *metric.Metric) error {
	gm, err := rpc.NewGraphiteMetricsWithMetric(m)
	if err != nil {
		return err
	}

	err = client.graphite.PostMetric(gm)
	if err != nil {
		return err
	}

	return nil
}

// QueryMetrics posts a query over Graphite interface
func (client *Client) QueryMetrics(q *metric.Query) (metric.ResultSet, error) {

	gq, err := rpc.NewGraphiteQueryWithMetricQuery(q)
	if err != nil {
		return nil, err
	}

	gms, err := client.graphite.PostQuery(gq)
	if err != nil {
		return nil, err
	}

	rs, err := rpc.NewResultSetWithGraphiteMetrics(gms)
	if err != nil {
		return nil, err
	}

	return rs, nil
}