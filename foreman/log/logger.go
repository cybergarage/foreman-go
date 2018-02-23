// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package log implements a logger package for DOJO-OS (Under development version).

The package is under development. Using the package, you will can output some level message as the following.

	log.SetSharedLogger(log.NewStdoutLogger(log.LoggerLevelTrace))
	....
	log.Trace(....)
	log.Info(....)
	log.Error(....)
	log.Warn(....)
	log.Fatal(....)
	log.Trace(....)
*/
package log

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type LoggerOutpter func(file string, level int, msg string) (int, error)

type Logger struct {
	File     string
	Level    int
	outputer LoggerOutpter
}

const (
	Format           = "%s %s %s"
	LF               = "\n"
	FilePerm         = 0644
	LoggerLevelTrace = (1 << 5)
	LoggerLevelInfo  = (1 << 4)
	LoggerLevelWarn  = (1 << 3)
	LoggerLevelError = (1 << 2)
	LoggerLevelFatal = (1 << 1)
	LoggerLevelNone  = 0

	LoggerLevelUnknownString = "UNKNOWN"
	LoggerLStdout            = "stdout"
)

var sharedLogger *Logger

func SetSharedLogger(logger *Logger) {
	sharedLogger = logger
}

func GetSharedLogger() *Logger {
	return sharedLogger
}

var sharedLogStrings = map[int]string{
	LoggerLevelTrace: "TRACE",
	LoggerLevelInfo:  "INFO",
	LoggerLevelWarn:  "WARN",
	LoggerLevelError: "ERROR",
	LoggerLevelFatal: "FATAL",
}

func (logger *Logger) GetLevel() int {
	return logger.Level
}

func (logger *Logger) GetLevelString() string {
	logString, hasString := sharedLogStrings[logger.Level]
	if !hasString {
		return LoggerLevelUnknownString
	}
	return logString
}

func NewStdoutLogger(level int) *Logger {
	logger := &Logger{
		File:     LoggerLStdout,
		Level:    level,
		outputer: outputToStdout}
	return logger
}

func outputToStdout(file string, level int, msg string) (int, error) {
	fmt.Fprintln(os.Stdout, msg)
	return len(msg), nil
}

func NewFileLogger(file string, level int) *Logger {
	logger := &Logger{
		File:     file,
		Level:    level,
		outputer: outputToFile}
	return logger
}

func outputToFile(file string, level int, msg string) (int, error) {
	msgBytes := []byte(msg + LF)
	fd, err := os.OpenFile(file, (os.O_WRONLY | os.O_CREATE | os.O_APPEND), FilePerm)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(fd)
	writer.Write(msgBytes)
	writer.Flush()

	fd.Close()

	return len(msgBytes), nil
}

func output(outputLevel int, userFormat string, a ...interface{}) int {
	if sharedLogger == nil {
		return 0
	}

	logLevel := sharedLogger.GetLevel()
	if (logLevel < outputLevel) || (logLevel <= LoggerLevelFatal) || (LoggerLevelTrace < logLevel) {
		return 0
	}

	t := time.Now()
	logDate := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// TODO : Support parameters
	//　userMsg := fmt.Sprintf(userFormat, a...)
	userMsg := userFormat

	headerString := fmt.Sprintf("[%s]", sharedLogger.GetLevelString())
	logMsg := fmt.Sprintf(Format, logDate, headerString, userMsg)
	logMsgLen := len(logMsg)

	if 0 < logMsgLen {
		logMsgLen, _ = sharedLogger.outputer(sharedLogger.File, logLevel, logMsg)
	}

	return logMsgLen
}

func Trace(format string, a ...interface{}) int {
	return output(LoggerLevelTrace, format, a)
}

func Info(format string, a ...interface{}) int {
	return output(LoggerLevelInfo, format, a)
}

func Warn(format string, a ...interface{}) int {
	return output(LoggerLevelWarn, format, a)
}

func Error(format string, a ...interface{}) int {
	return output(LoggerLevelError, format, a)
}

func Fatal(format string, a ...interface{}) int {
	return output(LoggerLevelFatal, format, a)
}
