// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/test"
)

const (
	testScenarioDirectory = "../test/scenarios/"
)

func testQueryScenarioFilesWithConfig(t *testing.T, scenarioFiles []string, testConf *test.Config, scenarioOpt *test.ScenarioOption) {
	t.Helper()

	for _, file := range scenarioFiles {
		t.Run(file, func(t *testing.T) {
			s := NewQueryScenario()
			filePath := filepath.Join(testScenarioDirectory, file)
			err := s.ExecuteFileWithOption(filePath, scenarioOpt)
			if err == nil {
				return
			}
			lastEvent := s.GetLastEvent()
			if testConf.EnableSkipError {
				logging.Error(err.Error().Error())
				t.Logf("%s (%d) : %s", file, lastEvent.GetNo(), err.Error())
			} else {
				t.Errorf("%s (%d) : %s", file, lastEvent.GetNo(), err.Error())
			}
			lastRes := s.GetLastResponse()
			if lastRes != nil {
				t.Logf("%d : %s\n", lastRes.GetStatusCode(), lastRes.GetQuery())
				t.Logf("%s", lastRes.GetContent())
			}
		})
	}
}

func TestQueryBasicScenarios(t *testing.T) {

	scenarioFiles := []string{
		"scenario_action_system_ls.test",
		"scenario_config_export.test",
		"scenario_finder_export.test",
		"scenario_metrics_export.test",
		"scenario_register_set.test",
		"scenario_register_export.test",
	}

	conf := test.NewDefaultConfig()
	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryBasicScenariosWithSleep(t *testing.T) {
	scenarioFiles := []string{}

	conf := test.NewDefaultConfig()

	opt := test.NewDefaultScenarioOption()
	opt.SetStepDuration(time.Second)

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryPythonActionScenarios(t *testing.T) {
	scenarioFiles := []string{
		"scenario_action_py_echo.test",
		"scenario_action_py_error_functions.test",
		// "scenario_action_py_qos_actions.test",
		"scenario_action_py_register.test",
		"scenario_action_py_query.test",
		"scenario_action_py_analyze_metrics.test",
		"scenario_action_py_export_metrics.test",
		"scenario_qos_unsatisfied_one_metrics.test",
		"scenario_qos_unsatisfied_two_metrics.test",
		"scenario_qos_unsatisfied_three_metrics.test",
		"scenario_qos_multi_routing.test",
		"scenario_qos_function_abs.test",
	}

	conf := test.NewDefaultConfig()
	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestPlatformDependentScenarios(t *testing.T) {
	scenarioFiles := []string{
		// FIXME On CentOS
		"scenario_qos_export_unsatisfied_metrics_graphite.test",
	}

	conf := test.NewDefaultConfig()
	conf.EnableSkipError = true

	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryNoRepeatabilityScenarios(t *testing.T) {
	scenarioFiles := []string{
		"scenario_qos_export_unsatisfied_metrics.test",
	}

	conf := test.NewDefaultConfig()
	conf.EnableSkipError = true

	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}

func TestQueryDebuggingScenarios(t *testing.T) {
	scenarioFiles := []string{}

	conf := test.NewDefaultConfig()

	opt := test.NewDefaultScenarioOption()

	testQueryScenarioFilesWithConfig(t, scenarioFiles, conf, opt)
}
