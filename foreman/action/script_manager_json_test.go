// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package action

import (
	"testing"
)

func TestMethodJSONExportImport(t *testing.T) {
	testActionCount := 6
	testActionJSON := "{" +
		"\"action01\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}," +
		"\"action02\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}," +
		"\"action03\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}," +
		"\"action04\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}," +
		"\"action05\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}," +
		"\"action06\" : { \"" + ActionColumnLanguage + "\" : \"" + ActionLanguageSystem + "\", \"" + ActionColumnCode + "\" : \"#!/bin/sh\"}" +
		"}"

	// Import (JSON String)

	mgr := NewScriptManager()

	err := mgr.importMethodJSONString(testActionJSON)
	if err != nil {
		t.Error(err)
		return
	}

	methodCnt := mgr.GetMethodCount()
	if methodCnt != testActionCount {
		t.Errorf("%d != %d", methodCnt, testActionCount)
	}

	// Export (JSON Object)

	actionJSONObj, err := mgr.exportMethodJSONObject()
	if err != nil {
		t.Error(err)
		return
	}

	// Import (JSON Object)

	mgr = NewScriptManager()

	err = mgr.ImportMethodJSONObject(actionJSONObj)
	if err != nil {
		t.Error(err)
	}

	methodCnt = mgr.GetMethodCount()
	if methodCnt != testActionCount {
		t.Errorf("%d != %d", methodCnt, testActionCount)
	}
}
