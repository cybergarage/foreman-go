// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
foremantest a testing utility for Forman.

	NAME
	foremantest

	SYNOPSIS
	foremantest <scenario_file>

	DESCRIPTION
	foremantest is a a testing utility for Forman.

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
*/
package main

import (
	"fmt"
	"os"

	"github.com/cybergarage/foreman-go/foreman/test"
)

const (
	ProgramName = "foremantest"
)

func usage() {
	fmt.Printf("Usage : %s <scenario_file>\n", ProgramName)
}

func errorMessage(e error) {
	fmt.Printf("%s\n", e.Error())
}

func main() {

	s := test.NewQueryScenario()

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	filename := os.Args[1]
	err := s.LoadFile(filename)
	if err != nil {
		errorMessage(err)
		os.Exit(1)
	}

	err = s.Setup()
	if err != nil {
		errorMessage(err)
		os.Exit(1)
	}

	for n, e := range s.GetEvents() {
		err := s.Execute(e)
		if err != nil {
			fmt.Printf("[%d] ERROR : %s\n", n, err.Error())
			os.Exit(1)
		}
		fmt.Printf("[%d] %s : OK\n", n, e.GetData())
	}

	err = s.Cleanup()
	if err != nil {
		errorMessage(err)
		os.Exit(1)
	}

	os.Exit(0)
}
