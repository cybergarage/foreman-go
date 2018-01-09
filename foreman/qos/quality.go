// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Quality represents a quality formula in a QoS clause.
type Quality interface {
	GetVariable() Variable
}
