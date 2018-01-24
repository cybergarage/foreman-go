// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

func testClientDial(t *testing.T, client *Client) {
	server := NewServer()

	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	_, err = server.GetHostname()
	if err != nil {
		t.Error(err)
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	testClientDial(t, client)
}