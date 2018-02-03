// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"fmt"
)

// baseEncorder represents an baseEncorder for JSON
// FIXME : The encorder is not implementended yet.
type baseEncorder struct {
	Encorder
}

// objectToString returns the string reprensation.
func (e *baseEncorder) objectToString(obj interface{}) (string, error) {
	switch obj.(type) {
	case string:
		return fmt.Sprintf("\"%s\"", obj.(string)), nil
	case float64:
		return fmt.Sprintf("%f", obj.(float64)), nil
	case float32:
		return fmt.Sprintf("%f", obj.(float32)), nil
	case int64:
		return fmt.Sprintf("%d", obj.(int64)), nil
	case int32:
		return fmt.Sprintf("%d", obj.(int32)), nil
	case int:
		return fmt.Sprintf("%d", obj.(int)), nil
	}

	return "", fmt.Errorf(errorUnknownObjectType, obj)
}

// appendObject appends a encorded string of the specified object.
func (e *baseEncorder) appendObject(buf *bytes.Buffer, indent int, obj interface{}) error {
	switch obj.(type) {
	case map[string]interface{}:
		objMap := obj.(map[string]interface{})
		objLastIdx := len(objMap) - 1
		n := 0
		buf.WriteString("{")
		for key, mapObj := range objMap {
			buf.WriteString(fmt.Sprintf("\"%s\":", key))
			switch mapObj.(type) {
			case []interface{}, map[string]interface{}:
				e.appendObject(buf, (indent + 1), mapObj)
			default:
				mapObjStr, err := e.objectToString(mapObj)
				if err != nil {
					return err
				}
				buf.WriteString(mapObjStr)
			}
			if n != objLastIdx {
				buf.WriteString(",")
			}
			n++
		}
		buf.WriteString("}")
	case []interface{}:
		objArray := obj.([]interface{})
		objLastIdx := len(objArray) - 1
		buf.WriteString("[")
		for n, childObj := range objArray {
			switch childObj.(type) {
			case []interface{}, map[string]interface{}:
				e.appendObject(buf, (indent + 1), childObj)
			default:
				childObjStr, err := e.objectToString(childObj)
				if err != nil {
					return err
				}
				buf.WriteString(childObjStr)
			}
			if n != objLastIdx {
				buf.WriteString(",")
			}
		}
		buf.WriteString("]")
	default:
		return fmt.Errorf(errorUnknownObjectType, obj)
	}
	return nil
}

// Encode returns a encorded string of the specified object.
func (e *baseEncorder) Encode(obj interface{}) (string, error) {
	var buf bytes.Buffer
	err := e.appendObject(&buf, 0, obj)
	return buf.String(), err
}
