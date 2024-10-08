// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSDstyle
// license that can be found in the LICENSE file.

/*
Package logger implements a logger package for DOJOOS (Under development version).

The package is under development. Using the package, you will can output some level message as the following.

	logger.Trace(....)
	logger.Info(....)
	logger.Error(....)
	logger.Warn(....)
	logger.Fatal(....)
*/
package logging

// #include <foreman/foreman-c.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// SetLogLevel sets the specified logging level to the all outputters.
func SetLogLevel(level int) {
	C.foreman_logger_setlevel(C.foreman_logger_getsharedinstance(), C.int(level))
}

// SetLogLevelString sets the specified logging level string to the all outputters.
func SetLogLevelString(level string) {
	clevel := C.CString(level)
	defer C.free(unsafe.Pointer(clevel))
	ret := bool(C.foreman_logger_setlevelstring(C.foreman_logger_getsharedinstance(), clevel))
	if ret {
		return
	}
	SetLogLevel(LoggerLevelInfo)
}

// GetLogLevel returns the current logging level.
func GetLogLevel() int {
	return int(C.foreman_logger_getlevel(C.foreman_logger_getsharedinstance()))
}

// SetLogFile adds a specified file outputter.
func SetLogFile(file string) {
	if file == "-" {
		SetLogToStdout()
	} else {
		cfile := C.CString(file)
		defer C.free(unsafe.Pointer(cfile))
		C.foreman_logger_addfileoutputter(C.foreman_logger_getsharedinstance(), cfile)
	}
}

// SetLogToStdout add a standard outputter.
func SetLogToStdout() {
	C.foreman_logger_addstdoutputter(C.foreman_logger_getsharedinstance())
}

// SetVerbose enables the verbose logging level.
func SetVerbose(verbose bool) {
	if verbose {
		SetLogLevel(LoggerLevelTrace)
		Warn("Verbose logging enabled.")
	}
}

// Trace outputs a trace message to the shared logger.
func Trace(format string, args ...interface{}) int {
	cmsg := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cmsg))
	return int(C.foreman_logger_trace(C.foreman_logger_getsharedinstance(), cmsg))
}

// Info outputs an information message to the shared logger.
func Info(format string, args ...interface{}) int {
	cmsg := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cmsg))
	return int(C.foreman_logger_info(C.foreman_logger_getsharedinstance(), cmsg))
}

// Warn outputs a warning message to the shared logger.
func Warn(format string, args ...interface{}) int {
	cmsg := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cmsg))
	return int(C.foreman_logger_warn(C.foreman_logger_getsharedinstance(), cmsg))
}

// Error outputs an error message to the shared logger.
func Error(format string, args ...interface{}) int {
	cmsg := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cmsg))
	return int(C.foreman_logger_error(C.foreman_logger_getsharedinstance(), cmsg))
}

// Fatal outputs a fatal message to the shared logger.
func Fatal(format string, args ...interface{}) int {
	cmsg := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cmsg))
	return int(C.foreman_logger_fatal(C.foreman_logger_getsharedinstance(), cmsg))
}
