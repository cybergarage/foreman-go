// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cybergarage/foreman-go/foreman/discovery"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// RemoteNode represents a remote node.
type RemoteNode struct {
	Node

	Cluster    string
	Name       string
	Address    string
	RPCPort    int
	RenderPort int
	CarbonPort int
}

// NewRemoteNode returns a remote new node.
func NewRemoteNode() *RemoteNode {
	node := &RemoteNode{}
	return node
}

// NewRemoteNodeWithDiscoveryNode returns a remote new node.
func NewRemoteNodeWithDiscoveryNode(discoveryNode discovery.Node) *RemoteNode {
	node := &RemoteNode{
		Cluster:    discoveryNode.GetCluster(),
		Name:       discoveryNode.GetName(),
		Address:    discoveryNode.GetAddress(),
		RPCPort:    discoveryNode.GetRPCPort(),
		RenderPort: discoveryNode.GetRenderPort(),
		CarbonPort: discoveryNode.GetCarbonPort(),
	}
	return node
}

// GetCluster returns the cluster name
func (node *RemoteNode) GetCluster() string {
	return node.Cluster
}

// GetName returns the host name
func (node *RemoteNode) GetName() string {
	return node.Name
}

// GetAddress returns the interface address
func (node *RemoteNode) GetAddress() string {
	return node.Address
}

// GetRPCPort returns the RPC port
func (node *RemoteNode) GetRPCPort() int {
	return node.RPCPort
}

// GetRenderPort returns the Graphite render port
func (node *RemoteNode) GetRenderPort() int {
	return node.RenderPort
}

// GetCarbonPort returns the Graphite carbon port
func (node *RemoteNode) GetCarbonPort() int {
	return node.CarbonPort
}

// PostQueryOverHTTP posts a query string over HTTP
func (node *RemoteNode) PostQueryOverHTTP(query string) (interface{}, int, error) {
	return node.postQueryOverHTTPWithRetransmissionFlag(query, false)
}

// postQueryOverHTTPWithRetransmissionFlag posts a query string over HTTP
func (node *RemoteNode) postQueryOverHTTPWithRetransmissionFlag(query string, retransmissionFlag bool) (interface{}, int, error) {
	rawQuery := fmt.Sprintf("%s=%s",
		HttpRequestFqlQueryParam,
		url.QueryEscape(query))

	if retransmissionFlag {
		rawQuery += fmt.Sprintf("&%s=%t",
			HttpRequestFqlRetransmissionParam,
			retransmissionFlag)
	}

	url := url.URL{
		Scheme: DefaultRpcProtocol,
		Host: fmt.Sprintf(
			"%s:%d",
			node.GetAddress(),
			node.GetRPCPort()),
		Path:     HttpRequestFqlPath,
		RawQuery: rawQuery,
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

// PostQuery posts a query string
func (node *RemoteNode) PostQuery(query string) (interface{}, error) {
	resObj, _, err := node.postQueryOverHTTPWithRetransmissionFlag(query, false)
	return resObj, err
}

// PostRetransmissionQuery posts a query string
func (node *RemoteNode) PostRetransmissionQuery(query string) (interface{}, error) {
	resObj, _, err := node.postQueryOverHTTPWithRetransmissionFlag(query, true)
	return resObj, err
}
