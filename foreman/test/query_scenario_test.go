// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"testing"
)

const (
	scenarioDirectory = "../../test/"
)

func TestQueryScenario(t *testing.T) {
	scenarioFiles := []string{
		"scenario_config_01.csv",
		//"scenario_metrics_01.csv",
		"scenario_register_01.csv",
	}

	s := NewQueryScenario()

	for _, file := range scenarioFiles {
		err := s.ExecuteFile(scenarioDirectory + file)
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
