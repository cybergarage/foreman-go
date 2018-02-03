// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"testing"
)

func TestNewError(t *testing.T) {
	errObj := NewError(0, "")
	encorder := NewEncorder()
	_, err := encorder.Encode(errObj)
	if err != nil {
		t.Error(err)
	}
}
