// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd

// FailureDetectorListener a listener for detector.
type FailureDetectorListener interface {
	// FailureDetectorNodeAdded is called when a new node added in the cluster.
	FailureDetectorNodeAdded(Node)
	// FailureDetectorNodeRemoved is called when the node is removed in the cluster.
	FailureDetectorNodeRemoved(Node)
	// FailureDetectorNodeStatusChanged is called when a status of the node is changed  in the cluster.
	FailureDetectorNodeStatusChanged(Node)
	// FailureDetectorNodeOutOfDate is called when the node is out of date in the cluster.
	FailureDetectorNodeOutOfDate(Node)
}
