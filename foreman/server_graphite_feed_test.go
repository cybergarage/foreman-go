// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/logging"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

func testFeedGraphiteDataToServer(t *testing.T, server *Server, feedDataFilename string) {

	// Setup a client for the target server

	client := go_graphite.NewClient()
	client.SetHost(server.GetAddress())
	client.SetCarbonPort(server.GetCarbonPort())
	client.SetRenderPort(server.GetRenderPort())

	// Feed metrics (Carbon API)

	feedData, err := ioutil.ReadFile(feedDataFilename)
	if err != nil {
		t.Error(err)
		return
	}

	err = client.FeedString(string(feedData))
	if err != nil {
		t.Error(err)
		return
	}

}

func testGraphiteFeedWithConfig(t *testing.T, serverConf *Config) {

	feedDataFilenames := []string{
		"server_graphite_feed_test_01.dat",
		"server_graphite_feed_test_02.dat",
		"server_graphite_feed_test_03.dat",
		"server_graphite_feed_test_04.dat",
		"server_graphite_feed_test_05.dat",
	}

	// Setup a target server

	server, err := NewServerWithConfig(serverConf)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	// Feed metrics (Carbon API)

	for _, feedDataFilename := range feedDataFilenames {
		testFeedGraphiteDataToServer(t, server, feedDataFilename)
	}

	// Stop the target server

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestGraphiteFeedWithLocalhost(t *testing.T) {
	logging.SetVerbose(true)

	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGrahiteHost

	testGraphiteFeedWithConfig(t, serverConf)
}

func TestGraphiteFeedWithHostName(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
		return
	}

	serverConf := NewDefaultConfig()
	serverConf.Server.Host = hostname

	testGraphiteFeedWithConfig(t, serverConf)
}

func TestGraphiteFeedWithDefaultConfigFile(t *testing.T) {
	serverConf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		t.Error(err)
		return
	}

	testGraphiteFeedWithConfig(t, serverConf)
}
