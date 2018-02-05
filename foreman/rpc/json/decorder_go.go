// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	PathSep = "/"
)

// goDecorder represents an goDecorder for JSON
type goDecorder struct {
	Decorder
	rootObj interface{}
}

// NewDecorder return a JSON Decorder.
func NewDecorder() Decorder {
	d := &goDecorder{
		rootObj: nil,
	}
	return d
}

// Decode decordes the specified JSON string.
func (d *goDecorder) Decode(jsonString string) error {
	return json.Unmarshal([]byte(jsonString), &d.rootObj)
}

func (d *goDecorder) GetRootObject() (interface{}, error) {
	if d.rootObj == nil {
		return nil, errors.New(errorDecorderEmptyRootObject)
	}
	return d.rootObj, nil
}

// GetPathObject returns an object by the given path.
func (d *goDecorder) GetPathObject(path string) (interface{}, error) {
	paths := strings.Split(path, PathSep)
	if len(paths) <= 0 {
		return "", errors.New(errorDecorderEmptyKey)
	}
	return d.getLastPathObject(paths)
}

// GetPathString returns an object by the given path.
func (d *goDecorder) GetPathString(path string) (string, error) {
	obj, err := d.GetPathObject(path)
	if err != nil {
		return "", err
	}
	value, ok := obj.(string)
	if !ok {

	}
	return value, nil
}

// GetPathInteger returns an integer value by the given path.
func (d *goDecorder) GetPathInteger(path string) (int, error) {
	obj, err := d.GetPathObject(path)
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
		return 0, fmt.Errorf(errorDecorderInvalidKeyObjectType, path)
	}

	return value, nil
}

// GetPathFloat returns a float value by the given path.
func (d *goDecorder) GetPathFloat(path string) (float64, error) {
	obj, err := d.GetPathObject(path)
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
		return 0, fmt.Errorf(errorDecorderInvalidKeyObjectType, path)
	}

	return value, nil
}

// getLastPathObject returns a last object the given paths.
func (d *goDecorder) getLastPathObject(paths []string) (interface{}, error) {
	if d.rootObj == nil {
		return nil, errors.New(errorDecorderEmptyRootObject)
	}

	lastObj := d.rootObj
	for _, path := range paths {
		childObj, err := d.getChildObject(lastObj, path)
		if err != nil {
			return nil, err
		}
		lastObj = childObj
	}

	return lastObj, nil
}

// getChildObject returns an object the given key.
func (d *goDecorder) getChildObject(obj interface{}, key string) (interface{}, error) {
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
