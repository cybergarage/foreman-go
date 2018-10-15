// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/uecho-go/net/echonet"
)

// NewController returns a default controller for Echonet.
func NewController() *echonet.Controller {
	ctrl := echonet.NewController()
	ctrl.SetConfig(NewDefaultConfig())
	return ctrl
}
