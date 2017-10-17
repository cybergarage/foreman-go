// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

// A Problem represents a SAT problem.
type Problem struct {
	Clauses []Clause
}

// NewProblem returns a new null problem.
func NewProblem() *Problem {
	p := &Problem{}
	return p
}
