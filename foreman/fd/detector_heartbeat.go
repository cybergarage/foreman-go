// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

import (
	"fmt"
	"time"
)

const (
	HeartbeatDetectorDefaultIntervalTime = time.Minute * 1
)

const (
	errorDetectorNotFoundExecutor = "Executor Not Found"
)

// HeartbeatDetector represents a heartbeat based detector.
type HeartbeatDetector struct {
	*baseDetector
	executor         FailureDetectionExecutor
	intervalDuration time.Duration
	intervalFuncStop chan bool
}

// NewHeartbeatDetector returns a new gossip detector.
func NewHeartbeatDetector() *HeartbeatDetector {
	detector := &HeartbeatDetector{
		baseDetector:     newBaseDetector(),
		executor:         nil,
		intervalDuration: HeartbeatDetectorDefaultIntervalTime,
		intervalFuncStop: nil,
	}
	return detector
}

// SetDetectionInterval sets a heartbeat interval
func (detector *HeartbeatDetector) SetDetectionInterval(d time.Duration) {
	detector.intervalDuration = d
}

// GetDetectionInterval returns a current heartbeat interval
func (detector *HeartbeatDetector) GetDetectionInterval() time.Duration {
	return detector.intervalDuration
}

// SetExecutor sets a executor
func (detector *HeartbeatDetector) SetExecutor(e FailureDetectionExecutor) error {
	detector.executor = e
	return nil
}

// GetExecutor returns a current executor
func (detector *HeartbeatDetector) GetExecutor() (FailureDetectionExecutor, error) {
	if detector.listener == nil {
		return nil, fmt.Errorf(errorDetectorNotFoundExecutor)
	}
	return detector.executor, nil
}

// Start starts the instance
func (detector *HeartbeatDetector) Start() error {
	err := detector.Stop()
	if err != nil {
		return err
	}

	finder, err := detector.GetFinder()
	if err != nil {
		return err
	}
	if finder == nil {
		return fmt.Errorf(errorDetectorNotFoundFinder)

	}

	if detector.executor == nil {
		return fmt.Errorf(errorDetectorNotFoundExecutor)

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
	detector.intervalFuncStop = make(chan bool)

	go func() {
		for {
			detector.executor.ExecuteFailureDetection(detector)
			select {
			case <-time.After(detector.intervalDuration):
			case <-detector.intervalFuncStop:
				return
			}
		}
	}()
}
