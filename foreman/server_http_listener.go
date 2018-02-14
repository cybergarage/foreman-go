// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"net/http"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	httpRequestQueryParam       = "q"
	httpResponseContentType     = "Content-Type"
	httpResponseContentTypeJSON = "application/json"
)

// HTTPRequestReceived is a listener for FQL
func (server *Server) HTTPRequestReceived(r *http.Request, w http.ResponseWriter) {
	switch r.URL.Path {
	case HttpServerFqlPath:
		server.fqlRequestReceived(r, w)
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// fqlRequestReceived handles FQL requests
func (server *Server) fqlRequestReceived(r *http.Request, w http.ResponseWriter) {
	queryString := r.FormValue(httpRequestQueryParam)
	server.Info(queryString)

	parser := fql.NewParser()
	queries, err := parser.ParseString(queryString)
	if err != nil {
		parserErr := errors.NewErrorWithCode(errors.ErrorCodeQueryInvalid)
		parserErr.DetailMessage = err.Error()
		server.httpResponseJSONError(r, w, parserErr)
		return
	}

	if len(queries) <= 0 {
		server.badRequestReceived(r, w)
		return
	}

	// FIXME: Exceute all queries
	query := queries[0]

	queryResult, queryErr := server.ExecuteQuery(query)
	if queryErr != nil {
		server.httpResponseJSONError(r, w, queryErr)
		return
	}

	w.Header().Set(httpResponseContentType, httpResponseContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	if queryResult == nil {
		return
	}

	encorder := json.NewEncorder()
	content, ok := encorder.Encode(queryResult)
	if ok != nil {
		return
	}
	w.Write([]byte(content))
}

// badRequestReceived returns http.StatusBadRequest
func (server *Server) badRequestReceived(r *http.Request, w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// httpResponseJSONError returns the specified error
func (server *Server) httpResponseJSONError(r *http.Request, w http.ResponseWriter, err *errors.Error) {
	w.Header().Set(httpResponseContentType, httpResponseContentTypeJSON)
	w.WriteHeader(http.StatusBadRequest)

	encorder := json.NewEncorder()
	content, ok := encorder.Encode(json.NewErrorWithError(err))
	if ok != nil {
		return
	}

	w.Write([]byte(content))
}
