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

// ExecuteQuery must return the result as a standard array or map.
func (conf *Config) ExecuteQuery(q fql.Query) (interface{}, *errors.Error) {
	if q.GetType() != fql.QueryTypeSelect {
		return nil, errors.NewErrorWithCode(errors.ErrorCodeQueryMethodNotSupported)
	}

	configMap := map[string]interface{}{}

	configMap[ConfigProductKey] = ProductName
	configMap[ConfigVersionKey] = Version

	configMap[rpc.ConfigHttpPortKey] = conf.Server.HTTPPort
	configMap[rpc.ConfigCarbonPortKey] = conf.Server.CarbonPort

	configMap[rpc.ConfigLogLevelKey] = conf.Log.Level

	configMap[rpc.ConfigBoostrapKey] = conf.Server.Boostrap

	configMap[rpc.ConfigMetricsStore] = conf.Metrics.Store

	configContainer := map[string]interface{}{}
	configContainer[strings.ToLower(fql.QueryTargetConfig)] = configMap

	return configContainer, nil
}
