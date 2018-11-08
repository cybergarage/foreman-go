// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

const (
	serverRestartTestConfig01Filename = "server_restart_test_01.conf"
	serverRestartTestConfig02Filename = "server_restart_test_02.conf"
	serverRestartTestQuery01Filename  = "server_restart_test_01.fql"
	serverRestartTestQuery02Filename  = "server_restart_test_02.fql"
)

func TestSeverRestart(t *testing.T) {

	clusterName := "foreman01"
	actionName := "action01"

	conf, err := NewConfigWithFile(serverRestartTestConfig01Filename)
	if err != nil {
		t.Error(err)
		return
	}

	server, err := NewServerWithConfig(conf)
	if err != nil {
		t.Error(err)
		return
	}

	err = server.LoadQuery(serverRestartTestQuery01Filename)
	if err != nil {
		t.Error(err)
		return
	}

	// Start

	err = server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	if server.GetCluster() != clusterName {
		t.Errorf("%s != %s", server.GetCluster(), clusterName)
	}

	if !server.actionMgr.HasMethod(actionName) {
		t.Errorf("%s", actionName)
	}

	// Restart (Emulate how the initial configuration and query files are modified by changing the loading file names)

	oldActionName := actionName
	newClusterName := "foreman02"
	newActionName := "action02"

	server.SetConfigFile(serverRestartTestConfig02Filename)
	server.SetQueryFile(serverRestartTestQuery02Filename)

	err = server.Restart()
	if err != nil {
		t.Error(err)
		return
	}

	if server.GetCluster() != newClusterName {
		t.Errorf("%s != %s", server.GetCluster(), newClusterName)
	}

	if server.actionMgr.HasMethod(oldActionName) {
		t.Errorf("%s", oldActionName)
	}

	if !server.actionMgr.HasMethod(newActionName) {
		t.Errorf("%s", newActionName)
	}

	// Stop

	err = server.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}
