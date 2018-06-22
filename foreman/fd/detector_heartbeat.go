// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"time"
)

const (
	HeartbeatDetectorDefaultInterval = time.Second * 1
)

// HeartbeatDetectorExecutor represents an abstract interface
type HeartbeatDetectorExecutor interface {
	ExecuteFailureDetection() error
}

// HeartbeatDetector represents a heartbeat based detector.
type HeartbeatDetector struct {
	*baseDetector
	intervalDuration time.Duration
	intervalFuncStop chan bool
}

// NewHeartbeatDetector returns a new gossip detector.
func NewHeartbeatDetector() *HeartbeatDetector {
	detector := &HeartbeatDetector{
		baseDetector:     newBaseDetector(),
		intervalDuration: HeartbeatDetectorDefaultInterval,
		intervalFuncStop: nil,
	}
	return detector
}

// SetInterval sets a heartbeat interval
func (detector *HeartbeatDetector) SetInterval(d time.Duration) {
	detector.intervalDuration = d
}

// Start starts the instance
func (detector *HeartbeatDetector) Start() error {
	err := detector.Stop()
	if err != nil {
		return err
	}

	heartbeatDetectorExecuter(detector)

	return nil
}

// Stop starts the instance
func (detector *HeartbeatDetector) Stop() error {
	if detector.intervalFuncStop != nil {
		detector.intervalFuncStop <- true
	}
	detector.intervalFuncStop = nil
	return nil
}

func heartbeatDetectorExecuter(detector *HeartbeatDetector) {
	executor, ok := (interface{}(detector)).(HeartbeatDetectorExecutor)
	if !ok {
		detector.Stop()
		return
	}

	detector.intervalFuncStop = make(chan bool)

	go func() {
		for {
			executor.ExecuteFailureDetection()
			select {
			case <-time.After(detector.intervalDuration):
			case <-detector.intervalFuncStop:
				return
			}
		}
	}()
}
