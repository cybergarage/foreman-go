// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
	"github.com/cybergarage/foreman-go/foreman/rpc/graphite"
	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

// Client represents a client for the server.
type Client struct {
	Scheme     string
	Host       string
	HTTPPort   int
	CarbonPort int
	RenderPort int
	*Controller
}

// NewClient returns a new client.
func NewClient() *Client {
	client := &Client{
		Scheme:     DefaultRpcProtocol,
		Host:       DefaultServerHost,
		HTTPPort:   DefaultHttpPort,
		CarbonPort: go_graphite.DefaultCarbonPort,
		RenderPort: go_graphite.DefaultRenderPort,
		Controller: NewController(),
	}
	return client
}

// NewClientWithNode returns a new client for a specified node.
func NewClientWithNode(node Node) *Client {
	client := &Client{
		Scheme:     DefaultRpcProtocol,
		Host:       node.GetName(),
		HTTPPort:   node.GetRPCPort(),
		CarbonPort: node.GetCarbonPort(),
		RenderPort: node.GetRenderPort(),
		Controller: NewController(),
	}
	return client
}

// SetHost sets the specified host into the client.
func (client *Client) SetHost(host string) error {
	client.Host = host
	return nil
}

// SetHTTPPort sets the specified port.
func (client *Client) SetHTTPPort(port int) error {
	client.HTTPPort = port
	return nil
}

// SetCarbonPort sets the specified port.
func (client *Client) SetCarbonPort(port int) error {
	client.CarbonPort = port
	return nil
}

// SetRenderPort sets the specified port.
func (client *Client) SetRenderPort(port int) error {
	client.RenderPort = port
	return nil
}

// SetNode sets the target node to the client.
func (client *Client) SetNode(node Node) error {
	client.SetHost(node.GetAddress())
	client.SetHTTPPort(node.GetRPCPort())
	client.SetCarbonPort(node.GetCarbonPort())
	client.SetRenderPort(node.GetRenderPort())
	return nil
}

// PostMetric posts a metric over Graphite interface
func (client *Client) PostMetric(m *metric.Metric) error {
	gm, err := graphite.NewGraphiteMetricsWithMetric(m)
	if err != nil {
		return err
	}

	graphite := go_graphite.NewClient()
	graphite.CarbonPort = client.CarbonPort
	err = graphite.FeedMetrics(gm)
	if err != nil {
		return err
	}

	return nil
}

// GetMetrics gets the specified metrics over Graphite interface
func (client *Client) GetMetrics(q *metric.Query) (metric.ResultSet, error) {
	gq := graphite.NewGraphiteQueryWithMetricQuery(q)
	gc := go_graphite.NewClient()
	gc.RenderPort = client.RenderPort
	gms, err := gc.QueryRender(gq)
	if err != nil {
		return nil, err
	}

	rs, err := graphite.NewResultSetWithGraphiteMetrics(gms)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

// PostQuery posts a query string
func (client *Client) PostQuery(query string) (interface{}, error) {
	node := NewRemoteNode()
	node.Address = client.Host
	node.RPCPort = client.HTTPPort
	return node.PostQuery(query)
}

// PostQueryOverHTTP posts a query string over HTTP
func (client *Client) PostQueryOverHTTP(query string) (interface{}, int, error) {
	node := NewRemoteNode()
	node.Address = client.Host
	node.RPCPort = client.HTTPPort
	return node.PostQueryOverHTTP(query)
}
