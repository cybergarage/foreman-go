// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"time"
)

const (
	RetiontionIntervalFiveMinute = time.Duration(5) * time.Minute
	DefaultRetiontionInterval    = RetiontionIntervalFiveMinute
)
