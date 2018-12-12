// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FIXME : Disable TestGraphiteFeedAPI*() for Linux because of these tests timeout On CentOS
// +build !linux

package foreman

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/logging"
	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

type testFeedGraphiteConf struct {
	feedFilenames []string
	feedDuration  time.Duration
}

func newTestFeedGraphiteDefaultConf() *testFeedGraphiteConf {
	return newTestFeedGraphite60SecConf()
}

func newTestFeedGraphite60SecConf() *testFeedGraphiteConf {
	return &testFeedGraphiteConf{
		feedFilenames: []string{
			"server_graphite_feed_test_01_01.dat",
			"server_graphite_feed_test_01_02.dat",
			"server_graphite_feed_test_01_03.dat",
			"server_graphite_feed_test_01_04.dat",
			"server_graphite_feed_test_01_05.dat",
		},
		feedDuration: 60 * time.Microsecond * 100,
	}
}

func newTestFeedGraphite10SecConf() *testFeedGraphiteConf {
	return &testFeedGraphiteConf{
		feedFilenames: []string{
			"server_graphite_feed_test_02_01_01.dat",
			"server_graphite_feed_test_02_01_02.dat",
			"server_graphite_feed_test_02_02_01.dat",
			"server_graphite_feed_test_02_02_02.dat",
			"server_graphite_feed_test_02_02_03.dat",
			"server_graphite_feed_test_02_03_01.dat",
			"server_graphite_feed_test_02_03_02.dat",
			"server_graphite_feed_test_02_04_01.dat",
			"server_graphite_feed_test_02_04_02.dat",
			"server_graphite_feed_test_02_05_01.dat",
			"server_graphite_feed_test_02_05_02.dat",
			"server_graphite_feed_test_02_05_03.dat",
			"server_graphite_feed_test_02_06_01.dat",
			"server_graphite_feed_test_02_06_02.dat",
			"server_graphite_feed_test_02_07_01.dat",
			"server_graphite_feed_test_02_07_02.dat",
			"server_graphite_feed_test_02_07_03.dat",
			"server_graphite_feed_test_02_08_01.dat",
			"server_graphite_feed_test_02_08_02.dat",
			"server_graphite_feed_test_02_09_01.dat",
			"server_graphite_feed_test_02_09_02.dat",
			"server_graphite_feed_test_02_09_03.dat",
			"server_graphite_feed_test_02_10_01.dat",
			"server_graphite_feed_test_02_10_02.dat",
		},
		feedDuration: 10 * time.Microsecond * 100,
	}
}

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

func testGraphiteFeedWithConfig(t *testing.T, serverConf *Config, testConf *testFeedGraphiteConf) {
	logging.SetVerbose(true)

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

	for _, feedDataFilename := range testConf.feedFilenames {
		testFeedGraphiteDataToServer(t, server, feedDataFilename)
		time.Sleep(testConf.feedDuration)
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

	testGraphiteFeedWithConfig(t, serverConf, newTestFeedGraphiteDefaultConf())
}

func TestGraphiteFeedAPIWithHostName(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
		return
	}

	serverConf := NewDefaultConfig()
	serverConf.Server.Host = hostname

	testGraphiteFeedWithConfig(t, serverConf, newTestFeedGraphiteDefaultConf())
}

func TestGraphiteFeedAPIWithDefaultConfigFile(t *testing.T) {
	serverConf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		t.Error(err)
		return
	}

	testGraphiteFeedWithConfig(t, serverConf, newTestFeedGraphiteDefaultConf())
}

/*
func TestMultiGraphiteFeedAPIWithDefaultConfigFile(t *testing.T) {
	serverConf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		t.Error(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		testGraphiteFeedWithConfig(t, serverConf, newTestFeedGraphite60SecConf())
	}()

	go func() {
		defer wg.Done()
		testGraphiteFeedWithConfig(t, serverConf, newTestFeedGraphite10SecConf())
	}()

	wg.Wait()
}
*/
