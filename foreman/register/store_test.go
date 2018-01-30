// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package register

import (
	"fmt"
	"testing"
	"time"
)

const (
	testStoreLoopCount  = 100
	testStoreKeyFormat  = "key%d"
	testStoreDataFormat = "data%d"
)

const (
	testObjectInvalidDataSize  = "Invalid data count (%d != %d)"
	testObjectNotFound         = "%s is not found"
	testObjectInvalidDataType  = "%s is invalid data type"
	testObjectInvalidData      = "%s is invalid data (%s != %s)"
	testObjectInvalidVersion   = "%s is invalid version (%d)"
	testObjectInvalidTimestamp = "%s is invalid timestamp (%s < %s)"
)

func testStore(t *testing.T, store Store) {
	now := time.Now()

	// Open

	err := store.Open()
	if err != nil {
		t.Error(err)
	}

	// Clear

	err = store.Clear()
	if err != nil {
		t.Error(err)
	}

	if store.Size() != 0 {
		t.Errorf(testObjectInvalidDataSize, store.Size(), 0)
	}

	// Add objects

	for n := 0; n < testStoreLoopCount; n++ {
		obj := newTestObject()
		obj.SetName(fmt.Sprintf(testStoreKeyFormat, n))
		obj.SetData(fmt.Sprintf(testStoreDataFormat, n))
		err = store.SetObject(obj)
		if err != nil {
			t.Error(err)
		}
	}

	if store.Size() != testStoreLoopCount {
		t.Errorf(testObjectInvalidDataSize, store.Size(), testStoreLoopCount)
	}

	// Get objects

	for n := 0; n < testStoreLoopCount; n++ {
		key := fmt.Sprintf(testStoreKeyFormat, n)
		data := fmt.Sprintf(testStoreDataFormat, n)
		obj, ok := store.GetObject(key)
		if !ok {
			t.Error(fmt.Errorf(testObjectNotFound, key))
		}

		objRawData, err := obj.GetData()
		if err != nil {
			t.Error(err)
		}

		objData, _ := objRawData.(string)
		if objData != data {
			t.Error(fmt.Errorf(testObjectInvalidData, key, objData, data))
		}

		err = obj.UpdateVersion()
		if err != nil {
			t.Error(err)
		}

		ver, err := obj.GetVersion()
		if err != nil {
			t.Error(err)
		}
		if ver <= 0 {
			t.Error(fmt.Errorf(testObjectInvalidVersion, key, ver))
		}

		ts, err := obj.GetTimestamp()
		if err != nil {
			t.Error(err)
		}
		if (ts.Sub(now) / time.Second) < 0 {
			t.Error(fmt.Errorf(testObjectInvalidTimestamp, key, ts, now))
		}
	}

	if store.Size() != testStoreLoopCount {
		t.Errorf(testObjectInvalidDataSize, store.Size(), testStoreLoopCount)
	}

	// Clear

	err = store.Clear()
	if err != nil {
		t.Error(err)
	}

	if store.Size() != 0 {
		t.Errorf(testObjectInvalidDataSize, store.Size(), 0)
	}

	// Close

	err = store.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestDefaultStore(t *testing.T) {
	store := NewStore()
	testStore(t, store)
}
