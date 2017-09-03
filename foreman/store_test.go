// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"testing"
	"time"
)

const (
	testStoreMetricsCount       = 100
	testStoreMetricsPrefix      = "path"
	testStoreMetricsInterval    = DefaultRetiontionInterval
	testStoreMetricsPeriodCount = 10
)

func testStore(t *testing.T, store *Store) {
	store.SetRetentionInterval(testStoreMetricsInterval)

	err := store.Open()
	if err != nil {
		t.Error(t)
	}

	// Setup metrics

	var m [testStoreMetricsCount]*Metric
	for n := 0; n < testStoreMetricsCount; n++ {
		m[n] = NewMetric()
		m[n].Name = fmt.Sprintf("%s%d", testStoreMetricsPrefix, n)
	}

	// Insert metrics

	from := time.Now()
	until := from
	for i := 0; i < testStoreMetricsPeriodCount; i++ {
		for j := 0; j < testStoreMetricsCount; j++ {
			m[j].Timestamp = until
			m[j].Value = float64(i * j)
			err = store.AddMetric(m[j])
			if err != nil {
				t.Error(t)
			}
		}
		until = until.Add(testStoreMetricsInterval)
	}

	// Query metric values

	q := NewQuery()
	q.From = &from
	q.Interval = testStoreMetricsInterval
	q.Until = &until
	for j := 0; j < testStoreMetricsCount; j++ {
		q.Target = m[j].Name
		rs, err := store.Query(q)
		if err != nil {
			t.Error(t)
		}
		rsCount := rs.GetDataPointCount()
		if rsCount != 1 {
			t.Error(fmt.Errorf("ResultSet is invalid : %d", rsCount))
		}

		dps := rs.GetFirstDataPoints()
		if dps == nil {
			t.Error(fmt.Errorf("DataPoint is not found"))
			return
		}

		if dps.Name != m[j].Name {
			t.Error(fmt.Errorf("Invalid DataPoint Name %s != %s", dps.Name, m[j].Name))
		}

		dpsCount := len(dps.Values)
		for i := 0; i < dpsCount; i++ {
			dp := dps.Values[i]
			value := dp.Value
			//fmt.Printf("[%d] : %f\n", i, value)
			if int(value) != int(i*j) {
				t.Error(fmt.Errorf("ResultSet value is invalid : %d != %d", int(value), int(i*j)))
			}
		}
	}

	err = store.Close()
	if err != nil {
		t.Error(t)
	}
}

func TestNewSQLiteStore(t *testing.T) {
	testStore(t, NewSQLiteStore())
}
