// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import "github.com/cybergarage/uecho-go/net/echonet"

const (
	FinderDeviceCode = 0x0F9101
)

const (
	FinderConditionCode  = 0x80 // (Same as operation status code for Echonet)
	FinderClusterCode    = 0xA0
	FinderNameCode       = 0xA1
	FinderAddressCode    = 0xA2
	FinderRPCPortCode    = 0xA3
	FinderRenderPortCode = 0xA4
	FinderCarbonPortCode = 0xA5

	FinderClockCode   = 0xB0
	FinderVersionCode = 0xB1
)

const (
	FinderRPCPortSize    = 4
	FinderRenderPortSize = 4
	FinderCarbonPortSize = 4

	FinderClockSize   = 8
	FinderVersionSize = 8
)

func getAllFinderDevicePropertyCodes() []echonet.PropertyCode {
	props := []echonet.PropertyCode{
		FinderConditionCode,
		FinderClusterCode,
		FinderNameCode,
		FinderAddressCode,
		FinderRPCPortCode,
		FinderRenderPortCode,
		FinderCarbonPortCode,
		FinderClockCode,
		FinderVersionCode,
	}
	return props
}

// NewDevice returns a finder device.
func NewDevice() (*echonet.Device, error) {
	dev := echonet.NewDevice()
	dev.SetCode(FinderDeviceCode)

	for _, propCode := range getAllFinderDevicePropertyCodes() {
		dev.CreateProperty(propCode, echonet.PropertyAttributeRead)
	}

	return dev, nil
}
