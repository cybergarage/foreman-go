// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FIXME : Disable TestGraphiteFeedAPI*() for Linux because of these tests timeout On CentOS
// +build !linux

package foreman

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

const (
	testGraphiteFeedDataDirectory = "../test/data/graphite/"
)

type testFeedGraphiteConf struct {
	feedFilenames     []string
	retentionInterval time.Duration
	feedDuration      time.Duration
}

func newTestFeedGraphiteDefaultConf() *testFeedGraphiteConf {
	return newTestCassandraFeedDataConf()
}

func newTestCassandraFeedDataConf() *testFeedGraphiteConf {
	return &testFeedGraphiteConf{
		feedFilenames: []string{
			"server_graphite_feed_test_01_01.dat",
			"server_graphite_feed_test_01_02.dat",
			"server_graphite_feed_test_01_03.dat",
			"server_graphite_feed_test_01_04.dat",
			"server_graphite_feed_test_01_05.dat",
		},
		retentionInterval: 60 * time.Second,
		feedDuration:      60 * time.Microsecond * 100,
	}
}

func newTestCollectdFeedDataConf() *testFeedGraphiteConf {
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
		retentionInterval: 10 * time.Second,
		feedDuration:      10 * time.Microsecond * 100,
	}
}

func testFeedGraphiteDataToServer(t *testing.T, server *Server, feedDataFilename string) {

	// Feed metrics (Carbon API)

	feedDataFilePath := filepath.Join(testGraphiteFeedDataDirectory, feedDataFilename)

	feedBytes, err := ioutil.ReadFile(feedDataFilePath)
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
			t.Errorf("%s : %s (%d != %d)", feedDataFilename, q.Target, len(ms), 1)
		}
	}

	// Get each inserted metrics data (Render API)

	storeInterval := server.GetMetricsStoreInterval()
	storeIntervalSec := storeInterval.Seconds()
	if storeIntervalSec <= 60 { /* Render API does not specify the seconds format */
		storeInterval += time.Second * 60
	}

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
			t.Errorf("%s : %s (%d <= %d)", feedDataFilename, q.Target, len(ms), 1)
		}
	}
}

func testGraphiteFeedWithConfig(t *testing.T, serverConf *Config, testConf *testFeedGraphiteConf) {
	//logging.SetVerbose(true)

	// Setup a target server

	server, err := NewServerWithConfig(serverConf)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.SetMetricsStoreInterval(testConf.retentionInterval)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.SetMetricsStorePeriod(0)
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

	feedDataConfs := []*testFeedGraphiteConf{
		newTestCassandraFeedDataConf(),
		newTestCollectdFeedDataConf(),
	}

	for _, feedDataConf := range feedDataConfs {
		testGraphiteFeedWithConfig(t, serverConf, feedDataConf)
	}
}

func TestMultiGraphiteFeedAPIWithDefaultConfigFile(t *testing.T) {
	//logging.SetVerbose(true)

	serverConf, err := NewConfigWithFile(configTestFilename)
	if err != nil {
		t.Error(err)
		return
	}

	// Feed data configurations

	feedData60Sec := newTestCassandraFeedDataConf()
	feedData10Sec := newTestCollectdFeedDataConf()

	// Setup a target server

	server, err := NewServerWithConfig(serverConf)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.SetMetricsStoreInterval(feedData10Sec.retentionInterval)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.SetMetricsStorePeriod(0)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	// Feed catpured metrics (Carbon API)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		conf := feedData60Sec
		for _, feelFilename := range conf.feedFilenames {
			testFeedGraphiteDataToServer(t, server, feelFilename)
			time.Sleep(conf.feedDuration)
		}
	}()

	go func() {
		defer wg.Done()
		conf := feedData10Sec
		for _, feelFilename := range conf.feedFilenames {
			testFeedGraphiteDataToServer(t, server, feelFilename)
			time.Sleep(conf.feedDuration)
		}
	}()

	wg.Wait()

	// Stop the target server

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}

}
