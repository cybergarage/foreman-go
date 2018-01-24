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
	testObjectInvalidTimestamp = "%s is invalid timestamp (%s)"
)

func testStore(t *testing.T, store *Store) {
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
		obj := NewObject()
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

		objData, ok := obj.GetData().(string)
		if !ok {
			t.Error(fmt.Errorf(testObjectInvalidDataType, key))
		}
		if objData != data {
			t.Error(fmt.Errorf(testObjectInvalidData, key, objData, data))
		}

		ver := obj.GetVersion()
		if ver <= 0 {
			t.Error(fmt.Errorf(testObjectInvalidVersion, key, ver))
		}

		ts := obj.GetTimestamp()
		if ts.Sub(now) < 0 {
			t.Error(fmt.Errorf(testObjectInvalidTimestamp, key, ts))
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
