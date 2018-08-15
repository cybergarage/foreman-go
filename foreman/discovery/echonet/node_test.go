// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"testing"
)

func TestNode(t *testing.T) {
	node, err := NewEchonetNodeWithNode(nil)
	if err != nil {
		t.Error(err)
	}

	_, err = node.GetNodeDevice()
	if err != nil {
		t.Error(err)
	}
}
