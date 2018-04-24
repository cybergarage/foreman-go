// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cybergarage/foreman-go/foreman/metric"
	rpc "github.com/cybergarage/foreman-go/foreman/rpc/graphite"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
	"github.com/cybergarage/go-graphite/net/graphite"
)

// Client represents a client for the server.
type Client struct {
	Scheme   string
	Host     string
	HTTPPort int
	graphite *graphite.Client
}

// NewClient returns a new client.
func NewClient() *Client {
	client := &Client{
		Scheme:   DefaultRpcProtocol,
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

	err = client.graphite.PostMetrics(gm)
	if err != nil {
		return err
	}

	return nil
}

// PostQuery posts a query string over Graphite interface
func (client *Client) PostQuery(queryString string) (interface{}, int, error) {
	url := url.URL{
		Scheme: client.Scheme,
		Host: fmt.Sprintf(
			"%s:%d",
			client.Host,
			client.HTTPPort),
		Path: HttpServerFqlPath,
		RawQuery: fmt.Sprintf("%s=%s",
			HttpServerFqlQuery,
			url.QueryEscape(queryString)),
	}

	res, err := http.Get(url.String())
	if err != nil {
		return nil, 0, err
	}

	resCode := res.StatusCode

	jsonResBytes, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, 0, err
	}

	if len(jsonResBytes) <= 0 {
		return nil, resCode, nil
	}

	var resObj interface{}

	dec := json.NewDecorder()
	resObj, err = dec.Decode(string(jsonResBytes))
	if err != nil {
		return nil, 0, err
	}

	return resObj, resCode, nil
}
