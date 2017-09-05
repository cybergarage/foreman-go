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

	"github.com/cybergarage/foreman-go/foreman"
	"github.com/cybergarage/foreman-go/foreman/log"
	"github.com/cybergarage/go-config/config"
)

const (
	ProgramName   = "foremand"
	ConfigFile    = "/etc/dkeykey/foremand.conf"
	ConfigRoot    = ProgramName
	ConfigLogFile = "log_file"
)

func main() {

	// Command Line Option

	foreground := flag.Bool("f", false, "Foreground mode.")
	verbose := flag.Int("v", 0, "Output log level.")
	configFile := flag.String("c", ConfigFile, "Path to an configuration file")
	flag.Parse()

	// Load configuration

	if len(*configFile) <= 0 {
		fmt.Fprintf(os.Stderr, "Configuration file is specified\n")
		os.Exit(1)
	}

	config, err := config.NewConfigFromFile(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration file (%s) is not found\n", *configFile)
		os.Exit(1)
	}

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

	// Demonaize

	/*　FIXME : Disabled the demonaize to use upstart temporarily.
	if !*foreground {
		_, ret, _ := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
		if ret != 1 {
			os.Exit(0)
		}
	}
	*/

	// Startup Server

	server := foreman.NewServer()

	// Output log message

	sharedLogger := log.GetSharedLogger()

	log.Info(fmt.Sprintf("%s is start ...", ProgramName))
	log.Info(fmt.Sprintf("%s = %s", ConfigLogFile, sharedLogger.File))
	log.Info(fmt.Sprintf("log_level = %s", sharedLogger.GetLevelString()))

	// Start Server

	log.Info(fmt.Sprintf("%s is started", ProgramName))

	err = server.Start()
	if err != nil {
		log.Error(fmt.Sprintf("%s couldn't be started (%s)", ProgramName, err.Error()))
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("%s is stop", ProgramName))

	os.Exit(0)
}
