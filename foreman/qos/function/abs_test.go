// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package function

import (
	"testing"
)

func TestAbsFunction(t *testing.T) {
	fn := NewAbs()

	params := [][]interface{}{
		{1.0},
		{-1.0},
		{1},
		{-1},
	}

	values := []float64{
		1,
		1,
		1,
		1,
	}

	for n, param := range params {
		ret, err := fn.Execute(param)
		if err != nil {
			t.Error(err)
			break
		}
		if ret != values[n] {
			t.Errorf("%s%v != %f", fn.GetName(), param, values[n])
		}
	}
}
