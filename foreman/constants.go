// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

const (
	ProductName                       = "foreman"
	DefaultServerHost                 = "localhost"
	DefaultCarbonPort                 = 2003
	DefaultHttpPort                   = 8188
	DefaultRpcProtocol                = "http"
	DefaultLogFile                    = "-"
	DefaultLogLevel                   = "INFO"
	DefaultMetricsStore               = "sqlite"
	DefaultMetricsInterval            = 300
	DefaultMetricsPeriod              = 3600
	DefaultCluster                    = "foreman"
	DefaultBoostrapMode               = 0
	DefaultFinder                     = "echonet"
	HttpRequestFqlPath                = "/fql"
	HttpRequestFqlQueryParam          = "q"
	HttpRequestFqlRetransmissionParam = "r"
	ConfigCategoryKey                 = "config"
	ConfigProductKey                  = "product"
	ConfigVersionKey                  = "version"
	ConfigHostKey                     = "host"
	ConfigCarbonPortKey               = "carbon_port"
	ConfigHttpPortKey                 = "http_port"
	ConfigFqlPathKey                  = "fql_path"
	ConfigFqlQueryKey                 = "fql_query"
	ConfigLogFileKey                  = "log_file"
	ConfigLogLevelKey                 = "log_level"
	ConfigBoostrapKey                 = "boostrap"
	ConfigFinderKey                   = "finder"
	ConfigMetricsStore                = "metrics_store"
	FinderEchonet                     = "echonet"
	RpcError                          = "error"
	RpcErrorCode                      = "code"
	RpcErrorMessage                   = "message"
	RpcErrorDetail                    = "detail"
	RpcErrorDetailCode                = "code"
	RpcErrorDetailMessage             = "message"
)
