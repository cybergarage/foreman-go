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
	errorMethodNotFound          = "Method (%s )Not Found"
)

func (mgr *ScriptManager) CreateMethodJSONString(jsonStr string) error {
	jsonDecorder := json.NewDecorder()
	jsonObj, err := jsonDecorder.Decode(jsonStr)
	if err != nil {
		return err
	}

	return mgr.CreateMethodJSONObject(jsonObj)
}

func (mgr *ScriptManager) CreateMethodJSONObject(jsonObj interface{}) error {
	methodsMap, ok := jsonObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf(errorMethodInvalidJSONObject, jsonObj)
	}

	for name, methodMapObj := range methodsMap {
		methodMap := json.NewPathWithObject(methodMapObj)

		lang, err := methodMap.GetPathString(ActionColumnLanguage)
		if err != nil {
			return fmt.Errorf(errorMethodInvalidJSONObject, methodMapObj)
		}

		code, err := methodMap.GetPathString(ActionColumnCode)
		if err != nil {
			return fmt.Errorf(errorMethodInvalidJSONObject, methodMapObj)
		}

		enc, _ := methodMap.GetPathString(ActionColumnEncoding)

		err = mgr.CreateMethod(name, lang, code, enc)
		if err != nil {
			return err
		}
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
		methodMap[ActionColumnLanguage] = method.Language
		methodMap[ActionColumnCode] = base64.StdEncoding.EncodeToString(method.Code)
		methodMap[ActionColumnEncoding] = ActionEncodingBase64

		methods[method.Name] = methodMap

		method = mgr.GetNextMethod(method)
	}

	if hasName && (len(methods) <= 0) {
		return nil, fmt.Errorf(errorMethodNotFound, name)
	}

	return methods, nil
}

func (mgr *ScriptManager) exportMethodJSONObject() (interface{}, error) {
	return mgr.exportMethodJSONObjectWithName("")
}
