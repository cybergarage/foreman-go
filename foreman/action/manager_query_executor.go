// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/fql"
)

// ExecuteQuery must return the result as a standard array or map.
func (mgr *Manager) ExecuteQuery(q fql.Query) (interface{}, error) {
	return nil, fmt.Errorf("")
}
