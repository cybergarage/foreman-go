// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"testing"
)

const (
	boostrapQueryTestFile = "../test/data/query/boostrap/boostrap_query_01_qos.fql"
	boostrapQueryTestDir  = "../test/data/query/boostrap"
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
