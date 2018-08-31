// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"testing"
)

func TestNewTargets(t *testing.T) {
	targetTypes := map[*Target]TargetType{
		NewTarget():                              QueryTargetTypeNone,
		NewTargetWithString(QueryTargetQos):      QueryTargetTypeShared,
		NewTargetWithString(QueryTargetConfig):   QueryTargetTypeStandalone,
		NewTargetWithString(QueryTargetMetrics):  QueryTargetTypeFederated,
		NewTargetWithString(QueryTargetRegister): QueryTargetTypeStandalone,
		NewTargetWithString(QueryTargetRegistry): QueryTargetTypeShared,
		NewTargetWithString(QueryTargetAction):   QueryTargetTypeShared,
		NewTargetWithString(QueryTargetRoute):    QueryTargetTypeShared,
	}

	for target, targetType := range targetTypes {
		if target.GetTargetType() != targetType {
			t.Errorf("%s : %d != %d", target, target.GetTargetType(), targetType)
		}
	}
}
