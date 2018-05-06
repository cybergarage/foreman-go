// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

// Remote represents a remote node.
type RemoteNode struct {
	Node
	Cluster string
	Name    string
	Address string
	RPCPort int
}

// NewNode returns a new node.
func NewRemoteNode() *RemoteNode {
	node := &RemoteNode{}
	return node
}

// GetCuster returns the cluster name
func (node *RemoteNode) GetCuster() string {
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

// PostQuery posts a query string
func (node *RemoteNode) PostQuery(queryString string) (interface{}, int, error) {
	url := url.URL{
		Scheme: DefaultRpcProtocol,
		Host: fmt.Sprintf(
			"%s:%d",
			node.GetAddress(),
			node.GetRPCPort()),
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
