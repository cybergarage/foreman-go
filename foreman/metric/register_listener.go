// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

// RegisterListener represents a listener for metric register.
type RegisterListener interface {
	// RegisterMetricAdded is called when a new metric is added
	RegisterMetricAdded(*RegisterMetric)
	// RegisterMetricUpdated is called when a metric which is already added is updated
	RegisterMetricUpdated(*RegisterMetric)
}
