// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	rpc "github.com/cybergarage/foreman-go/foreman/rpc/graphite"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// Client represents a client for the server.
type Client struct {
	Scheme   string
	Host     string
	HTTPPort int
	graphite *graphite.Client
	*Controller
}

// NewClient returns a new client.
func NewClient() *Client {
	client := &Client{
		Scheme:     DefaultRpcProtocol,
		Host:       DefaultServerHost,
		HTTPPort:   DefaultHttpPort,
		graphite:   graphite.NewClient(),
		Controller: NewController(),
	}
	return client
}

// SetHost sets the specified host into the client.
func (client *Client) SetHost(host string) error {
	client.Host = host
	return nil
}

// SetHTTPPort sets the specified point into the client.
func (client *Client) SetHTTPPort(port int) error {
	client.HTTPPort = port
	return nil
}

// PostMetric posts a metric over Graphite interface
func (client *Client) PostMetric(m *metric.Metric) error {
	gm, err := rpc.NewGraphiteMetricsWithMetric(m)
	if err != nil {
		return err
	}

	err = client.graphite.PostMetrics(gm)
	if err != nil {
		return err
	}

	return nil
}

// PostQuery posts a query string
func (client *Client) PostQuery(query string) (interface{}, int, error) {
	node := NewRemoteNode()
	node.Address = client.Host
	node.RPCPort = client.HTTPPort
	return node.PostQuery(query)
}
