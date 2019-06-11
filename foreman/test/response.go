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

// NewResponse create a scenario query response.
func NewResponse() *Response {
	return NewResponseWithParameters("", 0, nil)
}

// NewResponseWithParameters create a scenario query response with the specified parameters.
func NewResponseWithParameters(query string, code int, content interface{}) *Response {
	res := &Response{
		Query:      query,
		StatusCode: code,
		Content:    content,
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
