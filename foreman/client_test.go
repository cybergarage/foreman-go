// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"strings"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

func testClientExportConfigQuery(t *testing.T, server *Server, client *Client) {
	configObj, err := client.PostQuery("EXPORT FROM CONFIG")
	if err != nil {
		t.Error(err)
		return
	}

	configs := map[string]string{
		ConfigProductKey: ProductName,
		ConfigVersionKey: Version,
	}

	path := json.NewPathWithObject(configObj)
	for key, value := range configs {
		keyPath := json.NewPathStringWithStrings([]string{strings.ToLower(fql.QueryTargetConfig), key})
		configValue, err := path.GetPathString(keyPath)
		if err != nil {
			t.Error(err)
		}
		if configValue != value {
			t.Errorf("[%s] %s != %s", keyPath, configValue, value)
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
	client.SetNode(server)
	testClientExportConfigQuery(t, server, client)

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}
