// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

// ParserObjectStack is a stack for any parser objects.
type ParserObjectStack struct {
	objects []interface{}
}

// NewParserObjectStack creates a new stack.
func NewParserObjectStack() *ParserObjectStack {
	s := &ParserObjectStack{
		objects: make([]interface{}, 0),
	}
	return s
}

// PushObject adds a parser object to the stack.
func (s *ParserObjectStack) PushObject(obj interface{}) {
	s.objects = append(s.objects, obj)
}

// PeekObject returns a top parser object
func (s *ParserObjectStack) PeekObject() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	return s.objects[objectCount-1]
}

// PopObject removes and returns a top parser object
func (s *ParserObjectStack) PopObject() interface{} {
	objectCount := len(s.objects)
	if objectCount <= 0 {
		return nil
	}
	obj := s.objects[objectCount-1]
	s.objects = s.objects[0:(objectCount - 1)]
	return obj
}

// Size return a count of the stack objects
func (s *ParserObjectStack) Size() int {
	return len(s.objects)
}
