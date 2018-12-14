// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

const (
	boostrapQueryTestFile         = "../test/data/query/boostrap/boostrap_query_01_qos.fql"
	boostrapQueryTestFileScenario = "../test/scenarios/scenario_boostrap_query_01.csv"

	boostrapQueryTestDir         = "../test/data/query/boostrap"
	boostrapQueryTestDirScenario = "../test/scenarios/scenario_boostrap_query_02.csv"
)

func TestSeverBootstrapQueryFile(t *testing.T) {
	server := NewServer()
	server.SetBootstrapQuery(boostrapQueryTestFile)

	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestSeverBootstrapQueryFileWithScenario(t *testing.T) {
	server := NewServer()
	server.SetBootstrapQuery(boostrapQueryTestFile)

	s := NewQueryScenarioWithServer(server)
	err := s.ExecuteFile(boostrapQueryTestFileScenario)
	if err != nil {
		t.Error(err)
	}
}

func TestSeverBootstrapQueryDir(t *testing.T) {
	server := NewServer()
	server.SetBootstrapQuery(boostrapQueryTestDir)

	err := server.Start()
	if err != nil {
		t.Error(err)
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestSeverBootstrapQueryDirWithScenario(t *testing.T) {
	server := NewServer()
	server.SetBootstrapQuery(boostrapQueryTestDir)

	s := NewQueryScenarioWithServer(server)
	err := s.ExecuteFile(boostrapQueryTestDirScenario)
	if err != nil {
		t.Error(err)
	}
}
