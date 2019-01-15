// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/test"
)

const (
	testScenarioDirectory = "../test/scenarios/"
)

func testQueryScenarioFilesWithConfig(t *testing.T, scenarioFiles []string, testConf *test.Config, scenarioOpt *test.ScenarioOption) {

	for _, file := range scenarioFiles {
		s := NewQueryScenario()
		filePath := filepath.Join(testScenarioDirectory, file)
		err := s.ExecuteFileWithOption(filePath, scenarioOpt)
		if err != nil {
			lastEvent := s.GetLastEvent()
			if testConf.EnableSkipError {
				t.Logf("%s (%d) : %s", file, lastEvent.GetNo(), err.Error())
			} else {
				t.Errorf("%s (%d) : %s", file, lastEvent.GetNo(), err.Error())
			}
			lastRes := s.GetLastResponse()
			if lastRes != nil {
				t.Logf("%d : %s\n", lastRes.GetStatusCode(), lastRes.GetQuery())
				t.Logf("%s", lastRes.GetContent())
			}
			break
		}
	}
}

func TestQueryBasicScenarios(t *testing.T) {

	scenarioFiles := []string{
		"scenario_action_01.csv",
		"scenario_config_01.csv",
		"scenario_finder_01.csv",
		"scenario_metrics_01.csv",
		"scenario_metrics_02.csv",
		"scenario_qos_01.csv",
		"scenario_qos_02.csv",
		"scenario_register_01.csv",
	}

	conf := test.NewDefaultConfig()
	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryBasicScenariosWithSleep(t *testing.T) {

	scenarioFiles := []string{
		"scenario_metrics_02.csv",
	}

	conf := test.NewDefaultConfig()

	opt := test.NewDefaultScenarioOption()
	opt.SetStepDuration(time.Second)

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryPythonActionScenarios(t *testing.T) {
	scenarioFiles := []string{
		"scenario_action_py_01.csv",
		"scenario_action_py_02.csv",
		"scenario_action_py_03.csv",
		"scenario_action_py_04.csv",
	}

	conf := test.NewDefaultConfig()
	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryDebugScenarios(t *testing.T) {

	scenarioFiles := []string{
		"scenario_action_py_05.csv",
	}

	conf := test.NewDefaultConfig()
	conf.EnableSkipError = true

	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}
