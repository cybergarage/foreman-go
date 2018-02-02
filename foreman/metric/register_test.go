// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/register"
)

func TestNewRegister(t *testing.T) {
	NewRegisterWithStore(register.NewStore())
}
