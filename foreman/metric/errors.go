// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metric provides interfaces for MetricStore of Foreman C++.
package metric

const (
	errorClangObjectNotInitialized = "Clang object isn't initialized"
	errorStoreCouldNotOpen         = "Couldn't open : %s"
	errorStoreCouldNotClose        = "Couldn't close : %s"
	errorStoreCouldNotAddMetric    = "Couldn't add a metric : %s"
	errorStoreCouldNotQuery        = "Couldn't get the specified metric : %s"
	errorResultSetCouldGetValue    = "Couldn't get the specified index value : %d"
)
