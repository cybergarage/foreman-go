// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/logging"
)

func (server *Server) Trace(format string, a ...interface{}) int {
	return logging.Trace(format, a)
}

func (server *Server) Info(format string, a ...interface{}) int {
	return logging.Info(format, a)
}

func (server *Server) Warn(format string, a ...interface{}) int {
	return logging.Warn(format, a)
}

func (server *Server) Error(format string, a ...interface{}) int {
	return logging.Error(format, a)
}

func (server *Server) Fatal(format string, a ...interface{}) int {
	return logging.Fatal(format, a)
}
