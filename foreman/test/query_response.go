// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

// QueryResponse represents a scenario query response.
type QueryResponse struct {
	*Response
}

// NewQueryResponse create an scenario query response.
func NewQueryResponse() *QueryResponse {
	res := &QueryResponse{
		Response: NewResponse(),
	}
	return res
}
