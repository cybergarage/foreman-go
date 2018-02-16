// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// Response represents a scenario query response.
type Response struct {
	Query      string
	StatusCode int
	Content    interface{}
}

// NewResponse create an scenario query response.
func NewResponse() *Response {
	res := &Response{
		Query:      "",
		StatusCode: 0,
		Content:    nil,
	}
	return res
}

// GetQuery returns the stored query.
func (res *Response) GetQuery() string {
	return res.Query
}

// GetStatusCode returns the stored status code.
func (res *Response) GetStatusCode() int {
	return res.StatusCode
}

// GetContent returns the stored content.
func (res *Response) GetContent() interface{} {
	return res.Content
}
