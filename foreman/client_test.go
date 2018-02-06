// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

func testClientQueries(t *testing.T, server *Server, client *Client) {
	testQueries := []string{
		"EXPORT%20CONFIG",
	}

	for _, query := range testQueries {
		_, err := client.PostQuery(query)
		if err != nil {
			t.Error(err)
		}
	}
}

func testClientConfigQuery(t *testing.T, server *Server, client *Client) {
	configObj, err := client.PostQuery("EXPORT%20CONFIG")
	if err != nil {
		t.Error(err)
	}

	configs := map[string]string{
		ConfigProductKey: ProductName,
		ConfigVersionKey: Version,
	}

	path := json.NewPathWithObject(configObj)
	for key, value := range configs {
		configValue, err := path.GetPathString(key)
		if err != nil {
			t.Error(err)
			continue
		}
		if configValue != value {
			t.Errorf("[%s] %s != %s", key, configValue, value)
		}
	}

}

func TestClientQueries(t *testing.T) {
	server := NewServer()

	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient()
	testClientQueries(t, server, client)
	testClientConfigQuery(t, server, client)

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}
