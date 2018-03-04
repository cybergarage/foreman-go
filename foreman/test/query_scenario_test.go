// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"testing"
)

const (
	testScenarioDirectory = "../../test/scenarios/"
)

func TestQueryScenario(t *testing.T) {
	scenarioFiles := []string{
		"scenario_config_01.csv",
		"scenario_register_01.csv",
		"scenario_metrics_01.csv",
		"scenario_action_01.csv",
		// FIXME
		// fatal error: unexpected signal during runtime execution
		// [signal SIGSEGV: segmentation violation code=0x1 addr=0x5859c3000 pc=0x7fff4be8a2c7]
		//"scenario_register_01.csv",
	}

	s := NewQueryScenario()

	for _, file := range scenarioFiles {
		err := s.ExecuteFile(testScenarioDirectory + file)
		if err != nil {
			t.Error(err)
			res := s.GetLastResponse()
			if res != nil {
				t.Logf("%d : %s\n", res.GetStatusCode(), res.GetQuery())
				t.Logf("%s", res.GetContent())
			}
			break
		}
	}
}
