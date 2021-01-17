// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trpc

import (
	"github.com/aeraki-framework/aeraki/pkg/model"
	trpc "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/trpc_proxy/v3"
)

func buildOutboundProxy(context *model.EnvoyFilterContext) *trpc.TrpcProxy {
	route, err := buildOutboundRouteConfig(context)
	if err != nil {
		generatorLog.Errorf("Failed to generate tRPC EnvoyFilter: %v, %v", context.ServiceEntry, err)
		return nil
	}

	return &trpc.TrpcProxy{
		StatPrefix: model.BuildClusterName(model.TrafficDirectionOutbound, "",
			context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number)),
		RouteSpecifier: &trpc.TrpcProxy_RouteConfig{
			RouteConfig: route,
		},
	}
}
func buildInboundProxy(context *model.EnvoyFilterContext) *trpc.TrpcProxy {
	route, err := buildInboundRouteConfig(context)
	if err != nil {
		generatorLog.Errorf("Failed to generate tRPC EnvoyFilter: %v, %v", context.ServiceEntry, err)
		return nil
	}

	return &trpc.TrpcProxy{
		StatPrefix: model.BuildClusterName(model.TrafficDirectionInbound, "",
			context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number)),
		RouteSpecifier: &trpc.TrpcProxy_RouteConfig{
			RouteConfig: route,
		},
	}
}
