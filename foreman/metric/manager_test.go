// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"testing"

	"github.com/cybergarage/foreman-go/foreman/register"
)

func TestNewManager(t *testing.T) {
	mgr := NewManager()
	mgr.SetRegisterStore(register.NewManager().GetStore())
	err := mgr.Start()
	if err != nil {
		t.Error(err)
	}

	err = mgr.Stop()
	if err != nil {
		t.Error(err)
	}
}
