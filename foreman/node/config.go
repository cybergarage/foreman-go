// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

// Config represents an abstract node interface for the configuration
type Config interface {
	// GetCluster returns the cluster name
	GetCluster() string
	// GetName returns the host name
	GetName() string
	// GetAddress returns the interface address
	GetAddress() string
	// GetRPCPort returns the RPC port
	GetRPCPort() int
	// GetRenderPort returns the Graphite render port
	GetRenderPort() int
	// GetCarbonPort returns the Graphite carbon port
	GetCarbonPort() int
}

// ConfigEqual returns true if the other node is same with this node
func ConfigEqual(this, other Config) bool {
	if this.GetCluster() != other.GetCluster() {
		return false
	}

	if 0 < len(this.GetName()) && 0 < len(other.GetName()) {
		if this.GetName() != other.GetName() {
			return false
		}
	}

	if 0 < len(this.GetAddress()) && 0 < len(other.GetAddress()) {
		if this.GetAddress() != other.GetAddress() {
			return false
		}
	}

	if this.GetRPCPort() != other.GetRPCPort() {
		return false
	}

	if this.GetRenderPort() != other.GetRenderPort() {
		return false
	}

	if this.GetCarbonPort() != other.GetCarbonPort() {
		return false
	}

	return true
}
