// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

import (
	"fmt"
	"sync"
	"testing"

	"github.com/cybergarage/foreman-go/foreman/fql"
)

func severExecuteMultipleActions(t *testing.T, server *Server, waitGroup *sync.WaitGroup, query fql.Query) {
	defer waitGroup.Done()
	_, err := server.ExecuteQuery(query)
	if err != nil {
		t.Error(err)
	}
}

func TestSeverExecuteMultipleActions(t *testing.T) {
	server := NewServer()
	defer server.Stop()

	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	// Set a sleep action

	sleep_action_name := "echo"
	sleep_action_code := "aW1wb3J0IHRpbWUKZGVmIGVjaG8ocGFyYW1zLHJlc3VsdHMpOgoJZm9yIGtleSwgdmFsdWUgaW4gcGFyYW1zLml0ZW1zKCk6CgkJI3ByaW50KGtleSArICIgPSAiICsgdmFsdWUgKyAiIChzZXQpIikKCQlyZXN1bHRzW2tleV0gPSB2YWx1ZQoJCSNwcmludChrZXkgKyAiID0gIiArIHZhbHVlICsgIiAoc2xlZXApIikKCQl0aW1lLnNsZWVwKDEpCgkJI3ByaW50KGtleSArICIgPSAiICsgdmFsdWUgKyAiIChkb25lKSIpCglyZXR1cm4gVHJ1ZQo="

	queryString := fmt.Sprintf("SET (\"%s\",\"python\",\"%s\",\"base64\") INTO ACTION", sleep_action_name, sleep_action_code)

	parser := fql.NewParser()
	queries, err := parser.ParseString(queryString)
	if err != nil {
		t.Error(err)
		return
	}

	for _, query := range queries {
		_, err := server.ExecuteQuery(query)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Execute a sleep actions parallelly

	queries = fql.NewQueries()
	queryCnt := 10
	for n := 0; n < queryCnt; n++ {
		queryString = fmt.Sprintf("EXECUTE %s (message) VALUES (\"hello ([%d/%d])\")", sleep_action_name, (n + 1), queryCnt)

		parser = fql.NewParser()
		parsedQueries, err := parser.ParseString(queryString)
		if err != nil {
			t.Error(err)
			return
		}

		queries = append(queries, parsedQueries[0])
	}

	waitGroup := &sync.WaitGroup{}

	for _, query := range queries {
		waitGroup.Add(1)
		go severExecuteMultipleActions(t, server, waitGroup, query)
	}

	waitGroup.Wait()
}
