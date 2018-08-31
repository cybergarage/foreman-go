// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// Manager represents an abstract interface to control managers.
type Manager interface {
	Start() error
	Stop() error
}
