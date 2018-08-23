// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/rpc/json"
)

const (
	errorQoSInvalidJSONObject = "Invalid JSON Object : %v"
	errorQoSNotFound          = "QoS (%s) Not Found"
)

func (mgr *Manager) importQoSJSONString(jsonStr string) error {
	jsonDecorder := json.NewDecorder()
	jsonObj, err := jsonDecorder.Decode(jsonStr)
	if err != nil {
		return err
	}

	return mgr.ImportQoSJSONObject(jsonObj)
}

func (mgr *Manager) ImportQoSJSONObject(jsonObj interface{}) error {

	switch jsonObj.(type) {
	case map[string]string:
		qosStringMap := jsonObj.(map[string]string)
		for name, formula := range qosStringMap {
			err := mgr.CreateQoS(name, formula)
			if err != nil {
				return err
			}
		}
		return nil
	case map[string]interface{}:
		qosMap := jsonObj.(map[string]interface{})
		for name, formulaObj := range qosMap {
			formula, ok := formulaObj.(string)
			if !ok {
				return fmt.Errorf(errorQoSInvalidJSONObject, jsonObj)
			}
			err := mgr.CreateQoS(name, formula)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf(errorQoSInvalidJSONObject, jsonObj)
}

func (mgr *Manager) exportQoSJSONObjectWithName(name string) (interface{}, error) {
	hasName := false
	if 0 < len(name) {
		hasName = true
	}

	ruleMap := map[string]string{}
	for _, rule := range mgr.GetRules() {
		ruleName := rule.GetName()
		if hasName {
			if ruleName != name {
				continue
			}
		}
		ruleMap[ruleName] = rule.String()
	}

	if hasName && (len(ruleMap) <= 0) {
		return nil, fmt.Errorf(errorQoSNotFound, name)
	}

	return ruleMap, nil
}

func (mgr *Manager) exportQoSJSONObject() (interface{}, error) {
	return mgr.exportQoSJSONObjectWithName("")
}
