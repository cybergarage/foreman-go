// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"net/http"

	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	httpRequestQueryParam   = "q"
	httpResponseContentType = "application/json"
)

// HTTPRequestReceived is a listener for FQL
func (server *Server) HTTPRequestReceived(r *http.Request, w http.ResponseWriter) {
	switch r.URL.Path {
	case serverFQLPath:
		server.fqlRequestReceived(r, w)
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// fqlRequestReceived handles FQL requests
func (server *Server) fqlRequestReceived(r *http.Request, w http.ResponseWriter) {
	queryString := r.FormValue(httpRequestQueryParam)
	parser := fql.NewParser()
	queries, err := parser.ParseString(queryString)
	if err != nil {
		server.badRequestReceived(r, w)
		return
	}

	query := queries[0]

	resultObj, err := server.ExecuteQuery(query)
	if err != nil {
		server.badRequestReceived(r, w)
		return
	}

	e := json.NewEncorder()
	_, err = e.Encode(resultObj)
	if err != nil {
		server.badRequestReceived(r, w)
		return
	}

	server.badRequestReceived(r, w)
}

// badRequestReceived returns http.StatusBadRequest
func (server *Server) badRequestReceived(r *http.Request, w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
