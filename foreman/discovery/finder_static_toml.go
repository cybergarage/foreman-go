// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"fmt"

	"github.com/cybergarage/foreman-go/foreman/logging"
	"github.com/cybergarage/foreman-go/foreman/node"
	"github.com/BurntSushi/toml"
)

// StaticFinder represents a simple static finder.
type StaticTOMLFinder struct {
	*StaticFinder
}

// NewStaticFinderWithConfig returns a new static finder with specified nodes.
func NewStaticFinderWithConfig(config FinderConfig) Finder {
	nodes := []Node{}
	for _, host := range config.Hosts {
		node := node.NewBaseNode()
		node.Name = host
		nodes = append(nodes, node)
	}
	return NewStaticFinderWithNodes(nodes)
}

// NewStaticFinderWithTOML returns a new static finder with specified nodes.
func NewStaticFinderWithTOML(filename string) (Finder, error) {
	conf := Config{}
	if filename != "" {
		logging.Trace("TOML Config file path: %s", filename)
		_, err := toml.DecodeFile(filename, &conf)
		if err != nil {
			return nil, err
		}
		logging.Trace("Got config: %s", filename)
	}
	return NewStaticFinderWithConfig(conf.Finder), nil
}

// String returns the description
func (finder *StaticTOMLFinder) String() string {
	return fmt.Sprintf("%s", FinderStaticToml)
}
