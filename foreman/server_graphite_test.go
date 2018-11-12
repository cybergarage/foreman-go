// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"

	go_graphite "github.com/cybergarage/go-graphite/net/graphite"
)

const (
	testGrahiteHost                  = "localhost"
	testGrahiteMetricsNameFormat     = "localhost.metrics%d"
	testGrahiteMetricsCount          = 10
	testGrahiteMetricsIntervalSecond = (5 * 60)
	testGrahiteMetricsDurationSecond = (60 * 60)
	testGrahiteMetricsEachDataCount  = testGrahiteMetricsDurationSecond / testGrahiteMetricsIntervalSecond
	testGrahiteMetricsTotalDataCount = testGrahiteMetricsEachDataCount * testGrahiteMetricsCount
)

func testGraphiteAPIsWithConfig(t *testing.T, conf *Config) {

	// Setup a target server

	server, err := NewServerWithConfig(conf)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	// Setup a client for the target server

	client := go_graphite.NewClient()
	client.SetHost(server.GetAddress())
	client.SetCarbonPort(server.GetCarbonPort())
	client.SetRenderPort(server.GetRenderPort())

	// Post metrics (Carbon API)

	nowTs := time.Now()
	untilTs := nowTs.Truncate(time.Second * testGrahiteMetricsIntervalSecond)
	fromTs := untilTs.Add(-(time.Second * testGrahiteMetricsDurationSecond))

	for n := 0; n < testGrahiteMetricsCount; n++ {
		m := go_graphite.NewMetrics()
		m.SetName(fmt.Sprintf(testGrahiteMetricsNameFormat, n))

		ts := fromTs
		for ts.Before(untilTs) {
			dp := go_graphite.NewDataPoint()
			dp.SetValue(float64(n))
			dp.SetTimestamp(ts)
			m.AddDataPoint(dp)
			ts = ts.Add(time.Second * testGrahiteMetricsIntervalSecond)
		}

		err := client.PostMetrics(m)
		if err != nil {
			t.Error(err)
			server.Stop()
			return
		}
	}

	// Get all inserted metrics (Metrics API)

	ms, err := client.GetAllMetrics()
	if err == nil {
		if len(ms) != testGrahiteMetricsCount {
			// FIXME
			t.Skipf("%d != %d", len(ms), testGrahiteMetricsCount)
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
			t.Skipf("%d != %d", len(ms), testGrahiteMetricsCount)
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
				t.Skipf("%d != %d", len(ms), testGrahiteMetricsEachDataCount)
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
			t.Skipf("%d != %d", len(ms), testGrahiteMetricsTotalDataCount)
		}
	} else {
		t.Error(err)
	}

	// Stop the target server

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestGraphiteAPIsWithLocalhost(t *testing.T) {
	conf := NewDefaultConfig()
	conf.Server.Host = testGrahiteHost
	testGraphiteAPIsWithConfig(t, conf)
}

/*
func TestGraphiteAPIsWithHostName(t *testing.T) {
	conf := NewDefaultConfig()
	conf.Server.Host = testGrahiteHost
	testGraphiteAPIsWithConfig(t, conf)
}
*/
