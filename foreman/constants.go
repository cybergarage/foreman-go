// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foreman

const (
	ProductName                       = "foreman"
	DefaultServerHost                 = "localhost"
	DefaultCarbonPort                 = 2003
	DefaultHttpPort                   = 8188
	DefaultConnectionTimeout          = 60
	DefaultRpcProtocol                = "http"
	DefaultLogFile                    = "-"
	DefaultLogLevel                   = "INFO"
	DefaultMetricsStore               = "sqlite"
	DefaultMetricsInterval            = 300
	DefaultMetricsPeriod              = 3600
	DefaultCluster                    = "foreman"
	DefaultBootstrapMode              = 0
	DefaultFinder                     = "echonet"
	HttpRequestFqlPath                = "/fql"
	HttpRequestFqlQueryParam          = "q"
	HttpRequestFqlRetransmissionParam = "r"
	ConfigCategoryKey                 = "config"
	ConfigProductKey                  = "product"
	ConfigVersionKey                  = "version"
	ConfigClusterKey                  = "cluster"
	ConfigNameKey                     = "name"
	ConfigAddressKey                  = "address"
	ConfigHttpPortKey                 = "http_port"
	ConfigCarbonPortKey               = "carbon_port"
	ConfigTimeoutKey                  = "timeout"
	ConfigRenderPortKey               = "render_port"
	ConfigFqlPathKey                  = "fql_path"
	ConfigLogFileKey                  = "log_file"
	ConfigLogLevelKey                 = "log_level"
	ConfigBootstrapKey                = "boostrap"
	ConfigBoostrapQueryKey            = "boostrap_query"
	ConfigFinderKey                   = "finder"
	ConfigMetricsStoreKey             = "metrics_store"
	FinderEchonet                     = "echonet"
	FinderShared                      = "shared"
	FinderStatic                      = "static"
	FinderNodeCluster                 = "cluster"
	FinderNodeName                    = "name"
	FinderNodeAddress                 = "address"
	FinderNodeRpcPort                 = "rpc_port"
	FinderNodeCarbonPort              = "carbon_port"
	FinderNodeRenderPort              = "render_port"
	RpcError                          = "error"
	RpcErrorCode                      = "code"
	RpcErrorMessage                   = "message"
	RpcErrorDetail                    = "detail"
	RpcErrorDetailCode                = "code"
	RpcErrorDetailMessage             = "message"
	DefaultProfilePort                = 6060
)
