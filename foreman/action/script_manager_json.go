// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"encoding/base64"
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorMethodInvalidJSONObject = "Invalid JSON Object : %v"
	errorMethodNotFound          = "Method (%s) Not Found"
)

func (mgr *ScriptManager) importMethodJSONString(jsonStr string) error {
	jsonDecorder := json.NewDecorder()
	jsonObj, err := jsonDecorder.Decode(jsonStr)
	if err != nil {
		return err
	}

	return mgr.importMethodJSONObject(jsonObj)
}

func (mgr *ScriptManager) importMethodJSONObject(jsonObj interface{}) error {
	methodsMap, ok := jsonObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf(errorMethodInvalidJSONObject, jsonObj)
	}

	for name, methodMapObj := range methodsMap {
		methodMap := json.NewPathWithObject(methodMapObj)

		lang, err := methodMap.GetPathString(MethodColumnLanguage)
		if err != nil {
			return fmt.Errorf(errorMethodInvalidJSONObject, jsonObj)
		}

		code, err := methodMap.GetPathString(MethodColumnCode)
		if err != nil {
			return fmt.Errorf(errorMethodInvalidJSONObject, jsonObj)
		}

		enc, _ := methodMap.GetPathString(MethodColumnEncoding)

		err = mgr.importMethod(name, lang, code, enc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mgr *ScriptManager) importMethod(name string, lang string, code string, enc string) error {
	m := NewMethodWithLanguage(lang)
	m.Name = name
	if 0 < len(enc) && enc == MethodEncodingBase64 {
		code, err := base64.StdEncoding.DecodeString(code)
		if err != nil {
			return err
		}
		m.Code = code
	} else {
		m.Code = []byte(code)
	}

	err := mgr.AddMethod(m)
	if err != nil {
		return err
	}

	return nil
}

func (mgr *ScriptManager) exportMethodJSONObjectWithName(name string) (interface{}, error) {
	hasName := false
	if 0 < len(name) {
		hasName = true
	}

	methods := map[string]interface{}{}

	method := mgr.GetFirstMethod()
	for method != nil {
		if hasName {
			if method.Name != name {
				method = mgr.GetNextMethod(method)
				continue
			}
		}

		methodMap := map[string]interface{}{}
		methodMap[MethodColumnLanguage] = method.Language
		methodMap[MethodColumnCode] = base64.StdEncoding.EncodeToString(method.Code)
		methodMap[MethodColumnEncoding] = MethodEncodingBase64

		methods[method.Name] = methodMap

		method = mgr.GetNextMethod(method)
	}

	return methods, nil
}

func (mgr *ScriptManager) exportMethodJSONObject() (interface{}, error) {
	return mgr.exportMethodJSONObjectWithName("")
}
