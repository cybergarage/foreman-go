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
	testStoreMetricsCount    = 100
	testStoreMetricsPrefix   = "path"
	testStoreMetricsInterval = 5 * 60
)

func testStore(t *testing.T, store *Store) {
	err := store.Open()
	if err != nil {
		t.Error(t)
	}

	var m [testStoreMetricsCount]*Metric
	for n := 0; n < testStoreMetricsCount; n++ {
		m[n] = NewMetric()
		m[n].Name = fmt.Sprintf("%s%s", testStoreMetricsPrefix, n)
	}

	from := time.Now()
	interval := 5 * 60
	until := from.Add(300 * time.Minute)

	// Insert metric values

	for n := 0; n < testStoreMetricsCount; n++ {
		m[n].Timestamp = from
		m[n].Value = float64(n)
		err = store.AddMetric(m[n])
		if err != nil {
			t.Error(t)
		}
	}

	// Query metric values

	q := NewQuery()
	q.From = from
	q.Interval = interval
	q.Until = until
	for n := 0; n < testStoreMetricsCount; n++ {
		q.Target = m[n].Name
		rs, err := store.Query(q)
		if err != nil {
			t.Error(t)
		}
		rsCount, err := rs.GetCount()
		if err != nil {
			t.Error(t)
		}
		if rsCount != 1 {
			//t.Error(fmt.Errorf("ResultSet is invalid : %d", rsCount))
		}

		var i int64
		for i = 0; i < rsCount; i++ {
			value, err := rs.GetValue(int64(i))
			if err != nil {
				t.Error(t)
			}
			fmt.Printf("[%d] : %f\n", i, value)
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
