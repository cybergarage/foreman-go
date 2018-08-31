// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

import (
	"crypto/sha256"
	"fmt"
)

// Node represents an abstract node interface
type Node interface {
	Config
	Status
	// GetUniqueID returns a unique ID of the node
	GetUniqueID() string
}

// Equal returns true if the other node is same with this node
func Equal(this, other Node) bool {
	if !ConfigEqual(this, other) {
		return false
	}
	return true
}

// GetUniqueID returns a unique ID of the node
func GetUniqueID(node Node) string {
	seed := fmt.Sprintf("%s%s%s%d",
		node.GetCluster(),
		node.GetName(),
		node.GetAddress(),
		node.GetRPCPort())
	h := sha256.New()
	h.Write([]byte(seed))
	return fmt.Sprintf("%x", h.Sum(nil))
}
