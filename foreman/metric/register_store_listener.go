// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// StoreMetricAdded is a listener for Store
func (rs *Register) StoreMetricAdded(m *Metric) error {
	err := rs.UpdateMetric(m)
	if err != nil {
		return err
	}

	return nil
}
