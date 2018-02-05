// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"errors"
	"fmt"
	"strings"
)

const (
	PathSep = "/"
)

// Path represents a p object like XPath of XML
type Path struct {
	rootObject interface{}
}

// NewPathWithObject return a path object.
func NewPathWithObject(obj interface{}) *Path {
	p := &Path{
		rootObject: obj,
	}
	return p
}

// newPath return a path object.
func newPath() *Path {
	return NewPathWithObject(nil)
}

// GetRootObject returns the root object.
func (p *Path) GetRootObject() (interface{}, error) {
	if p.rootObject == nil {
		return nil, errors.New(errorPathEmptyRootObject)
	}
	return p.rootObject, nil
}

// GetPathObject returns an object by the given p.
func (p *Path) GetPathObject(path string) (interface{}, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorDecorderEmptyKey)
	}
	return p.getLastPathObject(paths)
}

// GetPathString returns an object by the given p.
func (p *Path) GetPathString(path string) (string, error) {
	obj, err := p.GetPathObject(path)
	if err != nil {
		return "", err
	}
	value, ok := obj.(string)
	if !ok {

	}
	return value, nil
}

// GetPathInteger returns an integer value by the given p.
func (p *Path) GetPathInteger(path string) (int, error) {
	obj, err := p.GetPathObject(path)
	if err != nil {
		return 0, err
	}

	var value int
	switch obj.(type) {
	case int:
		value = obj.(int)
	case int32:
		value = int(obj.(int32))
	case int64:
		value = int(obj.(int64))
	default:
		return 0, fmt.Errorf(errorDecorderInvalidKeyObjectType, p)
	}

	return value, nil
}

// GetPathFloat returns a float value by the given p.
func (p *Path) GetPathFloat(path string) (float64, error) {
	obj, err := p.GetPathObject(path)
	if err != nil {
		return 0, err
	}

	var value float64
	switch obj.(type) {
	case float64:
		value = obj.(float64)
	case float32:
		value = float64(obj.(float32))
	default:
		return 0, fmt.Errorf(errorDecorderInvalidKeyObjectType, p)
	}

	return value, nil
}

// getLastPathObject returns a last object the given paths.
func (p *Path) getLastPathObject(paths []string) (interface{}, error) {
	if p.rootObject == nil {
		return nil, errors.New(errorPathEmptyRootObject)
	}

	lastObj := p.rootObject
	for _, path := range paths {
		childObj, err := p.getChildObject(lastObj, path)
		if err != nil {
			return nil, err
		}
		lastObj = childObj
	}

	return lastObj, nil
}

// getChildObject returns an object the given key.
func (p *Path) getChildObject(obj interface{}, key string) (interface{}, error) {
	switch obj.(type) {
	case map[string]interface{}:
		jsonDir, _ := obj.(map[string]interface{})
		keyObj, hasKey := jsonDir[key]
		if !hasKey {
			return nil, fmt.Errorf(errorDecorderNotFoundKey, key)
		}
		switch keyObj.(type) {
		case string, float64, map[string]interface{}:
			return keyObj, nil
		default:
			return nil, fmt.Errorf(errorDecorderInvalidKeyObjectType, key)
		}
	}
	return "", nil
}
