// Copyright 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

// Server represents a Foreman Server.
type Server struct {
}

// NewServer returns a new Server.
func NewServer() *Server {
	server := &Server{}
	return server
}
