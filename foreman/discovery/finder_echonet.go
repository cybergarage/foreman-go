// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"time"

	foreman_echonet "github.com/cybergarage/foreman-go/foreman/discovery/echonet"
	"github.com/cybergarage/uecho-go/net/echonet"
)

const (
	echonetFinderSearchSleepSecond = 1
)

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	echonet.ControllerListener
	*baseFinder
	*echonet.Controller
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	finder := &EchonetFinder{
		baseFinder: newBaseFinder(),
		Controller: foreman_echonet.NewController(),
	}
	finder.Controller.SetListener(finder)
	return finder
}

// Search searches all nodes.
func (finder *EchonetFinder) Search() error {
	err := finder.Controller.SearchAllObjects()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * echonetFinderSearchSleepSecond)
	return nil
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return finder.Controller.Start()
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return finder.Controller.Stop()
}
