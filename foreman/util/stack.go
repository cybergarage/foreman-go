// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

// Stack is a stack for any parser objects.
type Stack struct {
	objects []interface{}
}

// NewStack creates a new stack.
func NewStack() *Stack {
	s := &Stack{
		objects: make([]interface{}, 0),
	}
	return s
}

// Push adds a parser object to the stack.
func (s *Stack) Push(obj interface{}) {
	s.objects = append(s.objects, obj)
}

// Peek returns a top parser object
func (s *Stack) Peek() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	return s.objects[objectCount-1]
}

// Pop removes and returns a top parser object
func (s *Stack) Pop() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	obj := s.objects[objectCount-1]
	s.objects = s.objects[0:(objectCount - 1)]
	return obj
}

// Size return a count of the stack objects
func (s *Stack) Size() int {
	return len(s.objects)
}
