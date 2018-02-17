// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package logger implements a logger package for DOJO-OS (Under development version).

The package is under development. Using the package, you will can output some level message as the following.

	logger.Trace(....)
	logger.Info(....)
	logger.Error(....)
	logger.Warn(....)
	logger.Fatal(....)
	logger.Trace(....)
*/
package logging

import (
	"os"

	"github.com/romana/rlog"
)

type LogLevel int

const (
	LevelNone LogLevel = 1 << iota
	LevelTrace
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelNone:
		return "NONE"
	case LevelTrace:
		return "TRACE"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

func LogLevelFromString(ls string) LogLevel {
	switch ls {
	case "NONE":
		return LevelNone
	case "TRACE":
		return LevelTrace
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	}
	Warn("Invalid log level: %s. Assuming TRACE", ls)
	return LevelTrace
}

func SetLogLevelFromString(ls string) {
	os.Setenv("RLOG_LOG_LEVEL", ls)
	rlog.UpdateEnv()
}

func SetLogLevel(l LogLevel) {
	ls := l.String()
	switch l { // for rlog
	case LevelTrace:
		ls = "DEBUG"
	case LevelFatal:
		ls = "CRITICAL"
	}
	SetLogLevelFromString(ls)
}

func SetLogFile(file string) {
	if file == "-" {
		SetLogtToStdout()
	} else {
		os.Setenv("RLOG_LOG_FILE", file)
		rlog.UpdateEnv()
	}
}

func SetLogtToStdout() {
	os.Setenv("RLOG_LOG_STREAM", "stdout")
	rlog.UpdateEnv()
}

func Trace(format string, a ...interface{}) int {
	rlog.Debugf(format, a...)
	return 0
}

func Info(format string, a ...interface{}) int {
	rlog.Infof(format, a...)
	return 0
}

func Warn(format string, a ...interface{}) int {
	rlog.Warnf(format, a...)
	return 0
}

func Error(format string, a ...interface{}) int {
	rlog.Errorf(format, a...)
	return 0
}

func Fatal(format string, a ...interface{}) int {
	rlog.Criticalf(format, a...)
	return 0
}
