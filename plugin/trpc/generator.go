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
	"github.com/aeraki-framework/aeraki/pkg/envoyfilter"
	"github.com/aeraki-framework/aeraki/pkg/model"
	"istio.io/pkg/log"
)

var generatorLog = log.RegisterScope("tRPC-generator", "tRPC generator", 0)

// Generator defines a tRPC envoyfilter Generator
type Generator struct {
}

// NewGenerator creates an new tRPC Generator instance
func NewGenerator() *Generator {
	return &Generator{}
}

// Generate create EnvoyFilters for  services
func (*Generator) Generate(context *model.EnvoyFilterContext) []*model.EnvoyFilterWrapper {
	return envoyfilter.GenerateReplaceNetworkFilter(
		context.ServiceEntry.Spec,
		buildOutboundProxy(context),
		buildInboundProxy(context),
		"envoy.filters.network.trpc_proxy",
		"type.googleapis.com/envoy.config.filter.network.trpc_proxy.v3.TrpcProxy")
}
