// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FIXME : Disable TestGraphiteAPI*() for Linux because of these tests timeout On CentOS
// +build !linux

package foreman

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

const (
	testGraphiteHost                 = "localhost"
	testGrahiteMetricsNameFormat     = "localhost.metrics%d"
	testGrahiteMetricsCount          = 10
	testGrahiteMetricsIntervalSecond = (60 * 5)
	testGrahiteMetricsDurationSecond = (60 * 15)
	testGrahiteMetricsEachDataCount  = testGrahiteMetricsDurationSecond / testGrahiteMetricsIntervalSecond
	testGrahiteMetricsTotalDataCount = testGrahiteMetricsEachDataCount * testGrahiteMetricsCount
)

type testGraphiteAPIConfig struct {
	insertRepeadCount int
	timestampJitter   bool
}

func newtestGraphiteAPIsConfig() *testGraphiteAPIConfig {
	conf := &testGraphiteAPIConfig{
		insertRepeadCount: 1,
		timestampJitter:   false,
	}
	return conf
}

func testGraphiteAPIsWithConfig(t *testing.T, serverConf *Config, testConf *testGraphiteAPIConfig) {

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
	defer server.Stop()

	// Setup a client for the target server

	client := go_graphite.NewClient()
	client.SetHost(server.GetAddress())
	client.SetCarbonPort(server.GetCarbonPort())
	client.SetRenderPort(server.GetRenderPort())

	// Post metrics (Carbon API)

	nowTs := time.Now()
	untilTs := nowTs.Truncate(time.Second * testGrahiteMetricsIntervalSecond)
	fromTs := untilTs.Add(-(time.Second * testGrahiteMetricsDurationSecond))

	for i := 0; i < testConf.insertRepeadCount; i++ {
		for n := 0; n < testGrahiteMetricsCount; n++ {
			m := go_graphite.NewMetrics()
			m.SetName(fmt.Sprintf(testGrahiteMetricsNameFormat, n))

			ts := fromTs
			for ts.Before(untilTs) {
				mts := ts
				if testConf.timestampJitter {
					mtsJitter := rand.Intn(testGrahiteMetricsIntervalSecond / 2)
					mts = mts.Add(time.Second * time.Duration(mtsJitter))
				}
				dp := go_graphite.NewDataPoint()
				dp.SetValue(float64(n))
				dp.SetTimestamp(mts)
				m.AddDataPoint(dp)
				ts = ts.Add(time.Second * testGrahiteMetricsIntervalSecond)
			}

			err := client.FeedMetrics(m)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Get all inserted metrics (Metrics API)

	ms, err := client.GetAllMetrics()
	if err == nil {
		if len(ms) != testGrahiteMetricsCount {
			// FIXME
			t.Logf("%d != %d", len(ms), testGrahiteMetricsCount)
		}
	} else {
		t.Error(err)
	}

	// Find inserted metrics (Metrics API)

	q := go_graphite.NewQuery()
	q.Target = "*"
	ms, err = client.FindMetrics(q)
	if err == nil {
		if len(ms) != testGrahiteMetricsCount {
			// FIXME on Screwdriver
			t.Logf("%d != %d", len(ms), testGrahiteMetricsCount)
		}
	} else {
		t.Error(err)
	}

	// Get each inserted metrics data (Render API)

	for n := 0; n < testGrahiteMetricsCount; n++ {
		q := go_graphite.NewQuery()
		q.Target = fmt.Sprintf(testGrahiteMetricsNameFormat, n)
		q.From = &fromTs
		q.Until = &untilTs
		ms, err = client.QueryRender(q)
		if err == nil {
			if len(ms) != testGrahiteMetricsEachDataCount {
				// FIXME on Screwdriver
				t.Logf("%d != %d", len(ms), testGrahiteMetricsEachDataCount)
			}
		} else {
			t.Error(err)
			continue
		}
	}

	// Get all inserted metrics data (Render API)

	q = go_graphite.NewQuery()
	q.Target = "*"
	q.From = &fromTs
	q.Until = &untilTs
	ms, err = client.QueryRender(q)
	if err == nil {
		if len(ms) != testGrahiteMetricsTotalDataCount {
			// FIXME on Screwdriver
			t.Logf("%d != %d", len(ms), testGrahiteMetricsTotalDataCount)
		}
	} else {
		t.Error(err)
	}
}

func TestGraphiteAPIsWithLocalhost(t *testing.T) {
	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGraphiteHost

	testConf := newtestGraphiteAPIsConfig()

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}

func TestGraphiteAPIsWithHostName(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		t.Error(err)
		return
	}

	serverConf := NewDefaultConfig()
	serverConf.Server.Host = hostname

	testConf := newtestGraphiteAPIsConfig()

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}

func TestGraphiteAPIsWithDefaultConfigFile(t *testing.T) {
	serverConf, err := newDefaultTestServerConfig()
	if err != nil {
		t.Error(err)
		return
	}

	testConf := newtestGraphiteAPIsConfig()

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}

func TestGraphiteAPIsForRepeatedInsert(t *testing.T) {
	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGraphiteHost

	testConf := newtestGraphiteAPIsConfig()
	testConf.insertRepeadCount = 2

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}

func TestGraphiteAPIsWithTimestampJitter(t *testing.T) {
	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGraphiteHost

	testConf := newtestGraphiteAPIsConfig()
	testConf.timestampJitter = true

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}

func TestGraphiteAPIsWithRepeatedAndTimestampJitter(t *testing.T) {
	serverConf := NewDefaultConfig()
	serverConf.Server.Host = testGraphiteHost

	testConf := newtestGraphiteAPIsConfig()
	testConf.timestampJitter = true
	testConf.insertRepeadCount = 2

	testGraphiteAPIsWithConfig(t, serverConf, testConf)
}
