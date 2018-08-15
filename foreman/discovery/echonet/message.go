// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

// NewRequestAllPropertiesMessage create a request message to get all properties.
func NewRequestAllPropertiesMessage() *protocol.Message {
	msg := protocol.NewMessage()
	return msg
}
