// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qos

import (
	"testing"
)

func TestQoSJSONExportImport(t *testing.T) {
	testQoSJSONCount := 6
	testQoSJSON := "{" +
		"\"qos1\" : \"x == 10.0\", " +
		"\"qos2\": \"x != 1.0\", " +
		"\"qos3\": \"x > 5.0\", " +
		"\"qos4\": \"x >= 10.0\", " +
		"\"qos5\": \"x < 20.0\", " +
		"\"qos6\": \"x <= 10.0\"" +
		"}"

	// Import (JSON String)

	mgr := NewManager()

	err := mgr.importQoSJSONString(testQoSJSON)
	if err != nil {
		t.Error(err)
		return
	}

	rules := mgr.GetRules()
	if len(rules) != testQoSJSONCount {
		t.Errorf("%d != %d", len(rules), testQoSJSONCount)
	}

	// Export (JSON Object)

	qosJSONObj, err := mgr.exportQoSJSONObject()
	if err != nil {
		t.Error(err)
		return
	}

	// Import (JSON Object)

	mgr = NewManager()

	err = mgr.ImportQoSJSONObject(qosJSONObj)
	if err != nil {
		t.Error(err)
	}

	rules = mgr.GetRules()
	if len(rules) != testQoSJSONCount {
		t.Errorf("%d != %d", len(rules), testQoSJSONCount)
	}
}
