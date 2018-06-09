// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"testing"
)

func detectorTest(t *testing.T, detector Detector) {
	err := detector.Start()
	if err != nil {
		t.Error(err)
		detector.Stop()
	}

	err = detector.Stop()
	if err != nil {
		t.Error(err)
	}
}

func TestNewGossipDetector(t *testing.T) {
	detector := NewGossipDetector()
	detectorTest(t, detector)
}
