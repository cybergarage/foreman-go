// Copyright 2014 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
foremand is a deamon command of Forman.

	NAME
	foremand

	SYNOPSIS
	foremand [-v] [-c config file]  [OPTIONS]

	DESCRIPTION
	foremand is a deamon process to Foreman.

	Logs to stdout by default, can be changed in the config file.

	OPTIONS
	-config : /path/to/foremand.conf Path to a configuration files.
	-v : Enable verbose output.

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
*/
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/cybergarage/foreman-go/foreman"
	"github.com/cybergarage/foreman-go/foreman/logging"
)

const (
	ProgramName   = "foremand"
	ConfigFile    = "/etc/foreman/foremand.conf"
	ConfigRoot    = ProgramName
	ConfigLogFile = "log_file"
)

func becomeVerbose(verbose bool) {
	if verbose {
		logging.SetLogLevel(logging.LevelTrace)
		logging.Trace("Enabled verbose output.")
	}
}

func main() {
	// Command Line Options

	//	foreground := flag.Bool("f", false, "Foreground mode.")
	verbose := flag.Bool("v", false, "Verbose logging")
	configFile := flag.String("config", ConfigFile, "Path to an configuration file")
	flag.Parse()

	// logging Level
	logging.SetVerbose(*verbose)

	// Load configuration
	server := foreman.NewServerWithConfigFile(*configFile)

	if server == nil {
		logging.Fatal("Could not start server. Terminating...")
		os.Exit(1)
	}

	// Start Server
	logging.Info("%s is starting ...", ProgramName)

	err := server.Start()
	if err != nil {
		logging.Fatal("%s couldn't be started (%s)", ProgramName, err.Error())
		os.Exit(1)
	}

	logging.Info("%s started", ProgramName)

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM)

	exitCh := make(chan int)

	go func() {
		for {
			s := <-sigCh
			switch s {
			case syscall.SIGHUP:
				logging.Info("Caught SIGHUP, restarting...")
				err = server.Restart()
				if err != nil {
					logging.Fatal("%s couldn't be restarted (%s)", ProgramName, err.Error())
					os.Exit(1)
				}
			case syscall.SIGINT, syscall.SIGTERM:
				logging.Info("Caught %s, stopping...", s.String())
				err = server.Stop()
				if err != nil {
					logging.Error("%s couldn't be stopped (%s)", ProgramName, err.Error())
					os.Exit(1)
				}
				exitCh <- 0
			}
		}
	}()

	code := <-exitCh

	logging.Info("%s stopped.", ProgramName)

	os.Exit(code)
}
