// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"github.com/cybergarage/foreman-go/foreman/metric"
)

// RegisterMetricAdded is a listener for metric.Register
func (qos *QoS) RegisterMetricAdded(rm *metric.RegisterMetric) {
	q := NewQuery()
	q.Target = rm.GetName()

	formulas, err := qos.FindRelatedFormulas(q)
	if err != nil {
		return
	}

	for _, formula := range formulas {
		formula.SetMetricEntity(rm.Metric)
	}
}

// RegisterMetricUpdated is a listener for metric.Register
func (qos *QoS) RegisterMetricUpdated(rm *metric.RegisterMetric) {
}
