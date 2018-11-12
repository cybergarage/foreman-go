// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

// #include <foreman/foreman-c.h>
import "C"

import (
	"strings"

	"github.com/cybergarage/foreman-go/foreman/errors"
	"github.com/cybergarage/foreman-go/foreman/fql"
	"github.com/cybergarage/foreman-go/foreman/rpc"
)

// ExecuteConfigQuery must return the result as a standard array or map.
func (server *Server) ExecuteConfigQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	configMap := map[string]interface{}{}

	configMap[ConfigProductKey] = ProductName
	configMap[ConfigVersionKey] = Version

	configMap[rpc.ConfigHostKey] = server.GetName()
	configMap[rpc.ConfigHttpPortKey] = server.GetRPCPort()
	configMap[rpc.ConfigCarbonPortKey] = server.GetCarbonPort()

	configMap[rpc.ConfigLogLevelKey] = server.Log.Level
	configMap[rpc.ConfigBoostrapKey] = server.Server.Boostrap
	configMap[rpc.ConfigMetricsStore] = server.metricMgr.String()

	configContainer := map[string]interface{}{}
	configContainer[strings.ToLower(fql.QueryTargetConfig)] = configMap

	return configContainer, nil
}
