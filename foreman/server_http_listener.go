// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package foreman provides interfaces for Foreman.
package foreman

import (
	"net/http"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	httpResponseContentType     = "Content-Type"
	httpResponseContentTypeJSON = "application/json"
)

// HTTPRequestReceived is a listener for FQL
func (server *Server) HTTPRequestReceived(r *http.Request, w http.ResponseWriter) {
	switch r.URL.Path {
	case HttpRequestFqlPath:
		server.fqlRequestReceived(r, w)
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

// fqlRequestReceived handles FQL requests
func (server *Server) fqlRequestReceived(r *http.Request, w http.ResponseWriter) {
	queryString := r.FormValue(HttpRequestFqlQueryParam)

	parser := fql.NewParser()
	queries, err := parser.ParseString(queryString)
	if err != nil {
		parserErr := errors.NewErrorWithCode(errors.ErrorCodeQueryInvalid)
		parserErr.DetailMessage = err.Error()
		server.httpResponseJSONError(r, w, parserErr)
		return
	}

	queryCnt := len(queries)
	if queryCnt <= 0 {
		server.badRequestReceived(r, w)
		return
	}

	retransParam := r.FormValue(HttpRequestFqlRetransmissionParam)
	if 0 <= len(retransParam) {
		for _, query := range queries {
			query.SetRetransmissionFlag(true)
		}
	}

	var queryAllResponse interface{}
	var queryResponses []interface{}
	if 1 < queryCnt {
		queryResponses = make([]interface{}, queryCnt)
		queryAllResponse = queryResponses
	}

	for n, query := range queries {
		queryResponse, queryErr := server.ExecuteQuery(query)
		if queryErr == nil {
			server.httpLogMessage(r, http.StatusOK, query.String())
		} else {
			server.httpLogMessage(r, http.StatusBadRequest, queryErr.String())
			server.httpResponseJSONError(r, w, queryErr)
			return
		}
		if queryCnt == 1 {
			queryAllResponse = queryResponse
		} else {
			queryResponses[n] = queryResponse
		}
	}

	w.Header().Set(httpResponseContentType, httpResponseContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	if queryAllResponse == nil {
		return
	}

	encorder := json.NewEncorder()
	content, ok := encorder.Encode(queryAllResponse)
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

// httpLogMessage outputs a log message for the specified HTTP request.
func (server *Server) httpLogMessage(r *http.Request, statusCode int, msg string) {
	logging.Info("[%d] : %s", statusCode, msg)
}
