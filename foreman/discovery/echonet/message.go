// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	"github.com/cybergarage/uecho-go/net/echonet/protocol"
)

// NewRequestAllPropertiesMessage create a request message to get all properties.
func NewRequestAllPropertiesMessage() *uecho_echonet.Message {
	msg := uecho_echonet.NewMessage()
	msg.SetESV(protocol.ESVReadRequest)
	msg.SetDEOJ(FinderDeviceCode)
	msg.AddProperties(uecho_echonet.NewPropertiesWithCodes(FinderDeviceAllPropertyCodes()))
	return msg
}
