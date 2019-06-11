// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/test"
)

// QueryEvent represents a scenario query event.
type QueryEvent struct {
	test.Executor
	*test.QueryEvent
	*Client
}

// NewQueryEventWithEvent create a scenario query event with the specified base event.
func NewQueryEventWithEvent(e *test.Event) (*QueryEvent, error) {
	tq, err := test.NewQueryEventWithEvent(e)
	if err != nil {
		return nil, err
	}
	q := &QueryEvent{
		QueryEvent: tq,
		Client:     nil,
	}
	return q, nil
}

// NewQueryEventWithEventAndClient create a scenario query event with the specified base event and client.
func NewQueryEventWithEventAndClient(e *test.Event, client *Client) (*QueryEvent, error) {
	q, err := NewQueryEventWithEvent(e)
	if err != nil {
		return nil, err
	}
	q.SetClient(client)
	return q, nil
}

// SetClient sets a query client.
func (q *QueryEvent) SetClient(client *Client) {
	q.Client = client
}

// Execute runs this query event.
func (q *QueryEvent) Execute() (*test.Response, *errors.Error) {
	resObj, resCode, err := q.PostQueryOverHTTP(q.Query)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	res := test.NewQueryResponse()
	res.Query = q.Query
	res.StatusCode = resCode
	res.Content = resObj

	err = q.VerifyResponse(res, q.Expectation)
	if err != nil {
		return nil, errors.NewErrorWithError(err)
	}

	return res.Response, nil
}
