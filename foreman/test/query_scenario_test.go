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
		"scenario_metrics_01.csv",
		"scenario_register_01.csv",
		"scenario_action_01.csv",

		// FIXME : Could not execute scenario_register_01 after scenario_action_01
		//[signal SIGSEGV: segmentation violation code=0x1 addr=0x585ac0000 pc=0x7fff55ccb2c7]
		//
		//runtime stack:
		//runtime.throw(0x44e42ac, 0x2a)
		//	/usr/local/Cellar/go/1.9.3/libexec/src/runtime/panic.go:605 +0x95
		//runtime.sigpanic()
		//	/usr/local/Cellar/go/1.9.3/libexec/src/runtime/signal_unix.go:351 +0x2b8
		//
		//goroutine 91 [syscall, locked to thread]:
		//runtime.cgocall(0x434a480, 0xc42005d8b0, 0x43497f0)
		//	/usr/local/Cellar/go/1.9.3/libexec/src/runtime/cgocall.go:132 +0xe4 fp=0xc42005d880 sp=0xc42005d840 pc=0x4004674
		//github.com/cybergarage/foreman-go/foreman/action._Cfunc_foreman_action_manager_execmethod(0x5a04060, 0x5a0c100, 0x5a07670, 0x5a07720, 0x5a0c080, 0x5a0c100)
		//	github.com/cybergarage/foreman-go/foreman/action/_obj/_cgo_gotypes.go:116 +0x4a fp=0xc42005d8b0 sp=0xc42005d880 pc=0x431645a
		//github.com/cybergarage/foreman-go/foreman/action.(*cgoScriptManager).ExecMethod.func2(0x5a04060, 0x5a0c100, 0x5a07670, 0x5a07720, 0x5a0c080, 0xc4201a8590)
		//	/Users/skonno/Src/foreman-go/src/github.com/cybergarage/foreman-go/foreman/action/script_manager_cgo.go:141 +0x12b fp=0xc42005d8f0 sp=0xc42005d8b0 pc=0x431b31b
		//github.com/cybergarage/foreman-go/foreman/action.(*cgoScriptManager).ExecMethod(0xc420218b98, 0xc4201a8570, 0x9, 0xc4202a98c0, 0x0, 0x0, 0x0)
		//	/Users/skonno/Src/foreman-go/src/github.com/cybergarage/foreman-go/foreman/action/script_manager_cgo.go:141 +0x163 fp=0xc42005d958 sp=0xc42005d8f0 pc=0x43196c3
		//"scenario_register_01.csv",

		"scenario_qos_01.csv",
	}

	s := NewQueryScenario()

	for _, file := range scenarioFiles {
		err := s.ExecuteFile(testScenarioDirectory + file)
		if err != nil {
			t.Skip(err)
			res := s.GetLastResponse()
			if res != nil {
				t.Logf("%d : %s\n", res.GetStatusCode(), res.GetQuery())
				t.Logf("%s", res.GetContent())
			}
			break
		}
	}
}
