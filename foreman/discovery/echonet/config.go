// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echonet

import (
	"github.com/cybergarage/uecho-go/net/echonet"
)

// NewDefaultConfig returns a default configuration for Echonet node and controller
func NewDefaultConfig() *echonet.Config {
	conf := echonet.NewDefaultConfig()
	conf.SetTCPEnabled(true)
	return conf
}
