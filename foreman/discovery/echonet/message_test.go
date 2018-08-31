// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"testing"

	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

func TestNewRequestMessage(t *testing.T) {
	msg := NewRequestAllPropertiesMessage()

	_, err := protocol.NewMessageWithBytes(msg.Bytes())
	if err != nil {
		t.Error(err)
	}
}
