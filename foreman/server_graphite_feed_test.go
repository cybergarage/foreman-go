// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"io/ioutil"
	"os"
	"testing"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

func testFeedGraphiteDataToServer(t *testing.T, server *Server, feedDataFilename string) {

	// Feed metrics (Carbon API)

	feedBytes, err := ioutil.ReadFile(feedDataFilename)
	if err != nil {
		t.Error(err)
		return
	}

	feedData := string(feedBytes)

	feedMetrics, err := go_graphite.NewMetricsWithPlainText(feedData)
	if err != nil {
		t.Error(err)
		return
	}

	// Setup a client for the target server

	client := go_graphite.NewClient()
	client.SetHost(server.GetAddress())
	client.SetCarbonPort(server.GetCarbonPort())
	client.SetRenderPort(server.GetRenderPort())

	// Feed metrics (Carbon API)

	err = client.FeedString(feedData)
	if err != nil {
		t.Error(err)
		return
	}

	// Find inserted metrics (Metrics API)

	for _, m := range feedMetrics {
		q := go_graphite.NewQuery()
		q.Target = m.GetName()
		ms, err := client.FindMetrics(q)
		if err != nil {
			t.Error(err)
			continue
		}
		if len(ms) != 1 {
			t.Errorf("%d != %d", len(ms), 1)
		}
	}

	// Get each inserted metrics data (Render API)

	storeInterval := server.GetMetricsStoreInterval()

	for _, m := range feedMetrics {
		dp, err := m.GetDataPoint(0)
		if err != nil {
			t.Error(err)
			continue
		}

		ts := dp.GetTimestamp()
		from := ts.Add(-storeInterval)
		until := ts.Add(storeInterval)

		q := go_graphite.NewQuery()
		q.Target = m.GetName()
		q.From = &from
		q.Until = &until

		ms, err := client.QueryRender(q)
		if err != nil {
			t.Error(err)
		}

		if len(ms) <= 1 {
			t.Errorf("%d <= %d", len(ms), 1)
		}
	}
}

func testGraphiteFeedWithConfig(t *testing.T, serverConf *Config) {
	//logging.SetVerbose(true)

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

func TestGraphiteFeedAPIWithLocalhost(t *testing.T) {
	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGrahiteHost

	testGraphiteFeedWithConfig(t, serverConf)
}

func TestGraphiteFeedAPIWithHostName(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
		return
	}

	serverConf := NewDefaultConfig()
	serverConf.Server.Host = hostname

	testGraphiteFeedWithConfig(t, serverConf)
}

func TestGraphiteFeedAPIWithDefaultConfigFile(t *testing.T) {
	serverConf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		t.Error(err)
		return
	}
	testGraphiteFeedWithConfig(t, serverConf)
}
