// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"testing"
)

func TestQueryScenario(t *testing.T) {
	scenarioFiles := []string{
		"data/scenario_config_01.csv",
		"data/scenario_metrics_01.csv",
	}

	s := NewQueryScenario()

	for _, file := range scenarioFiles {
		err := s.ExecuteFile(file)
		if err != nil {
			t.Error(err)
			return
		}
	}
}