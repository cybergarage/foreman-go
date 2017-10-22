// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"fmt"
	"testing"
)

func TestNewResultSet(t *testing.T) {
	rs := NewResultSet()

	dpsCount := rs.GetDataPointCount()
	if dpsCount != 0 {
		t.Error(fmt.Errorf("DataPoints is found : %d", dpsCount))
	}

	dps := rs.GetFirstDataPoints()
	if dps != nil {
		t.Error(fmt.Errorf("DataPoints is not nil"))
		return
	}
}
