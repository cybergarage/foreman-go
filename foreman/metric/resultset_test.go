// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides query interfaces for metric store.
package metric

import (
	"fmt"
	"testing"
)

func TestNewResultSet(t *testing.T) {
	rs := NewResultSet()

	dpsCount := rs.GetMetricsCount()
	if dpsCount != 0 {
		t.Error(fmt.Errorf("DataPoints is found : %d", dpsCount))
	}

	ms := rs.GetFirstMetrics()
	if ms != nil {
		t.Error(fmt.Errorf("DataPoints is not nil"))
		return
	}
}
