// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"net"
)

// Node represents an abstract node interface
type Node interface {
	// GetCuster returns the cluster name
	GetCuster() string
	// GetAddress returns the interface address
	GetAddress() net.Addr
	// GetRPCPort returns the RPC port
	GetRPCPort() int
}
