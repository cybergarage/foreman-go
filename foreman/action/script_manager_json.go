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
	errorActionInvalidJSONObject = "Invalid JSON Object : %v"
	errorActionNotFound          = "Action (%s) Not Found"
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
	actionsMap, ok := jsonObj.(map[string]interface{})
	if !ok {
		return fmt.Errorf(errorActionInvalidJSONObject, jsonObj)
	}

	for name, actionMapObj := range actionsMap {
		actionMap := json.NewPathWithObject(actionMapObj)

		lang, err := actionMap.GetPathString(ActionColumnLanguage)
		if err != nil {
			return fmt.Errorf(errorActionInvalidJSONObject, jsonObj)
		}

		code, err := actionMap.GetPathString(ActionColumnCode)
		if err != nil {
			return fmt.Errorf(errorActionInvalidJSONObject, jsonObj)
		}

		enc, _ := actionMap.GetPathString(ActionColumnEncoding)

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
	if 0 < len(enc) && enc == ActionEncodingBase64 {
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
		methodMap[ActionColumnLanguage] = method.Language
		methodMap[ActionColumnCode] = base64.StdEncoding.EncodeToString(method.Code)
		methodMap[ActionColumnEncoding] = ActionEncodingBase64

		methods[method.Name] = methodMap

		method = mgr.GetNextMethod(method)
	}

	return methods, nil
}

func (mgr *ScriptManager) exportMethodJSONObject() (interface{}, error) {
	return mgr.exportMethodJSONObjectWithName("")
}
