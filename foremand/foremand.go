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

	Logs are located at /var/log/foremand.log

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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cybergarage/foreman-go/foreman"
	"github.com/cybergarage/foreman-go/foreman/log"
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
	//	verbose := flag.Int("v", 0, "Output log level.")
	configFile := flag.String("c", ConfigFile, "Path to an configuration file")
	flag.Parse()

	server := foreman.NewServer()

	// Load configuration

	if 0 < len(*configFile) {
		err := server.LoadConfig(*configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			//os.Exit(1)
		}
	}

	/*
		// Log Level

		logLevel := log.LoggerLevelInfo
		if 0 < *verbose {
			logLevel = log.LoggerLevelTrace
		}

		// Setup logger

		logFile, err := config.GetKeyStringByPath(ConfigRoot + "/" + ConfigLogFile)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		if *foreground {
			log.SetSharedLogger(log.NewStdoutLogger(logLevel))
		} else {
			log.SetSharedLogger(log.NewFileLogger(logFile, logLevel))
		}
		defer log.SetSharedLogger(nil)

		// Output log message

		sharedLogger := log.GetSharedLogger()

		log.Info(fmt.Sprintf("%s is start ...", ProgramName))
		log.Info(fmt.Sprintf("%s = %s", ConfigLogFile, sharedLogger.File))
		log.Info(fmt.Sprintf("log_level = %s", sharedLogger.GetLevelString()))

	*/

	// Start Server

	err := server.Start()
	if err != nil {
		log.Error(fmt.Sprintf("%s couldn't be started (%s)", ProgramName, err.Error()))
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("%s is started", ProgramName))

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
					log.Error(fmt.Sprintf("%s couldn't be restarted (%s)", ProgramName, err.Error()))
					os.Exit(1)
				}
			case syscall.SIGINT, syscall.SIGTERM:
				err = server.Stop()
				if err != nil {
					log.Error(fmt.Sprintf("%s couldn't be stopped (%s)", ProgramName, err.Error()))
					os.Exit(1)
				}
				exitCh <- 0
			}
		}
	}()

	code := <-exitCh

	log.Info(fmt.Sprintf("%s is stop", ProgramName))

	os.Exit(code)
}
