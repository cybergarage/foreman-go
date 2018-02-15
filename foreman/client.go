// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cybergarage/foreman-go/foreman/metric"
	rpc "github.com/cybergarage/foreman-go/foreman/rpc/graphite"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// Client represents a client for the server.
type Client struct {
	Host     string
	HTTPPort int
	graphite *graphite.Client
}

// NewClient returns a new client.
func NewClient() *Client {
	client := &Client{
		Host:     DefaultServerHost,
		HTTPPort: DefaultHttpPort,
		graphite: graphite.NewClient(),
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

	err = client.graphite.PostMetric(gm)
	if err != nil {
		return err
	}

	return nil
}

// PostQuery posts a query string over Graphite interface
func (client *Client) PostQuery(queryString string) (interface{}, int, error) {
	url := fmt.Sprintf("http://%s:%d%s?%s=%s",
		client.Host,
		client.HTTPPort,
		HttpServerFqlPath,
		HttpServerFqlQuery,
		queryString)

	res, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}

	jsonResBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, 0, err
	}

	resCode := res.StatusCode

	var resObj interface{}

	dec := json.NewDecorder()
	resObj, err = dec.Decode(string(jsonResBytes))
	if err != nil {
		return nil, 0, err
	}

	return resObj, resCode, nil
}
