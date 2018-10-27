// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

const (
	errorNodeExportObjectNotFound = "Export Object Not Found (%s) : %v"
)

// baseNode is a super class to define methods using Node interface.
type baseNode struct {
	Node Node
}

// newBaseNodeWithNode returns a remote new node.
func newBaseNodeWithNode(node Node) *baseNode {
	baseNode := &baseNode{
		Node: node,
	}
	return baseNode
}
