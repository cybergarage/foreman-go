// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/uecho-go/net/echonet"
)

// EchonetController represents a base controller for Echonet.
type EchonetController struct {
	*echonet.Controller
	*EchonetDevice
}

// NewController returns a default controller for Echonet.
func NewController() *EchonetController {
	ctrl := echonet.NewController()

	ctrl.SetConfig(NewDefaultConfig())
	ctrl.SetManufacturerCode(ManufacturerCode)

	dev := NewDevice()
	ctrl.AddDevice(dev.Device)

	return &EchonetController{
		Controller:    ctrl,
		EchonetDevice: dev,
	}
}
