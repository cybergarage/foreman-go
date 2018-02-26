// Copyright 2014 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
foremand is a deamon command of Forman.

	NAME
	foremand

	SYNOPSIS
	foremand [-s service] [-c cert file]  [OPTIONS]

	DESCRIPTION
	foremand is a deamon process to Foreman.

	Logs are located at /var/logging/foremand.logging

	OPTIONS
	-c : * /path/to/foremand.conf * Path to an configuration files.
	-v : *level* Enable verbose output.
	-f : *true* Run as forground process.

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

func main() {

	// Command Line Option

	//	foreground := flag.Bool("f", false, "Foreground mode.")
	verbose := flag.Int("v", 0, "Output logging level.")
	configFile := flag.String("config", ConfigFile, "Path to an configuration file")
	flag.Parse()

	// Load configuration
	server := foreman.NewServerWithConfigFile(*configFile)

	if server == nil {
		logging.Fatal("Could not start server. Terminating...")
		os.Exit(1)
	}

	// logging Level

	logLevel := logging.LevelInfo
	if 0 < *verbose {
		logLevel = logging.LevelTrace
	}
	logging.SetLogLevel(logLevel)

	// Start Server
	logging.Info("%s is starting ...", ProgramName)

	err := server.Start()
	if err != nil {
		logging.Fatal("%s couldn't be started (%s)", ProgramName, err.Error())
		os.Exit(1)
	}

	logging.Info("%s is started", ProgramName)

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
				err = server.Restart()
				if err != nil {
					logging.Fatal("%s couldn't be restarted (%s)", ProgramName, err.Error())
					os.Exit(1)
				}
			case syscall.SIGINT, syscall.SIGTERM:
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

	logging.Info("Stopping %s...", ProgramName)

	os.Exit(code)
}
