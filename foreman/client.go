// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
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

// PostMetric posts a specified metric into all metric datapoints to Carbon.
func (client *Client) PostMetric(*metric.Metric) error {
	//err := client.graphite.PostMetric(m)
	//return err
	return nil
}
