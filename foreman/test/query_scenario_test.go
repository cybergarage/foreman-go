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

func testQueryScenarioFilesWithConfig(t *testing.T, scenarioFiles []string, conf *Config) {

	s := NewQueryScenario()

	for _, file := range scenarioFiles {
		err := s.ExecuteFile(testScenarioDirectory + file)
		if err != nil {
			if conf.EnableSkipError {
				t.Skip(err)
			} else {
				t.Error(err)
			}
			res := s.GetLastResponse()
			if res != nil {
				t.Logf("%d : %s\n", res.GetStatusCode(), res.GetQuery())
				t.Logf("%s", res.GetContent())
			}
			break
		}
	}
}

func TestQueryBasicScenarios(t *testing.T) {

	scenarioFiles := []string{
		"scenario_action_01.csv",
		"scenario_metrics_01.csv",
		"scenario_qos_01.csv",
		"scenario_register_01.csv",
	}

	conf := NewDefaultConfig()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf)
}

func TestQueryExtraScenarios(t *testing.T) {
	scenarioFiles := []string{
		"scenario_action_02.csv",
	}

	conf := NewDefaultConfig()
	conf.EnableSkipError = true

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf)
}
