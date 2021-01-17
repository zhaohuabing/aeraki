package trpc

import (
	"fmt"

	"github.com/aeraki-framework/aeraki/pkg/model"
	envoy "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	networking "istio.io/api/networking/v1alpha3"
)

var (
	// TODO: In the current version of Envoy, MaxProgramSize has been deprecated. However even if we do not send
	// MaxProgramSize, Envoy is enforcing max size of 100 via runtime.
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/type/matcher/v3/regex.proto.html#type-matcher-v3-regexmatcher-googlere2.
	regexEngine = &matcher.RegexMatcher_GoogleRe2{GoogleRe2: &matcher.RegexMatcher_GoogleRE2{}}
)

func buildOutboundRouteConfig(context *model.EnvoyFilterContext) (*routev3.RouteConfiguration, error) {
	// trpc service interface should be passed in via serviceentry annotation
	var tRPCService string
	var exist bool
	if tRPCService, exist = context.ServiceEntry.Annotations["tRPCService"]; !exist {
		err := fmt.Errorf("no tRPCService annotation")
		return nil, err
	}

	var routes []*routev3.Route
	clusterName := model.BuildClusterName(model.TrafficDirectionOutbound, "",
		context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number))

	if context.VirtualService == nil {
		routes = []*routev3.Route{defaultRoute(clusterName, tRPCService)}
	} else {
		routes = []*routev3.Route{buildRoute(context, tRPCService)}
	}

	return &routev3.RouteConfiguration{
		Name: clusterName,
		VirtualHosts: []*routev3.VirtualHost{
			{
				Name:    "default",
				Domains: []string{"*"},
				Routes:  routes,
			},
		},
	}, nil
}

func buildInboundRouteConfig(context *model.EnvoyFilterContext) (*routev3.RouteConfiguration, error) {
	// trpc service interface should be passed in via serviceentry annotation
	var tRPCService string
	var exist bool
	if tRPCService, exist = context.ServiceEntry.Annotations["tRPCService"]; !exist {
		err := fmt.Errorf("no tRPCService annotation")
		return nil, err
	}
	clusterName := model.BuildClusterName(model.TrafficDirectionInbound, "",
		"", int(context.ServiceEntry.Spec.Ports[0].Number))
	routes := []*routev3.Route{defaultRoute(clusterName, tRPCService)}
	return &routev3.RouteConfiguration{
		Name: clusterName,
		VirtualHosts: []*routev3.VirtualHost{
			{
				Name:    "default",
				Domains: []string{"*"},
				Routes:  routes,
			},
		},
	}, nil
}

func defaultRoute(clusterName string, tRPCService string) *routev3.Route {
	return &routev3.Route{
		Match: &routev3.RouteMatch{
			PathSpecifier: &routev3.RouteMatch_Prefix{
				Prefix: "/" + tRPCService,
			},
		},
		Action: &routev3.Route_Route{
			Route: &routev3.RouteAction{
				ClusterSpecifier: &routev3.RouteAction_Cluster{Cluster: clusterName},
			},
		},
	}
}

func buildRoute(context *model.EnvoyFilterContext, tRPCService string) *routev3.Route {
	service := context.ServiceEntry.Spec
	vs := context.VirtualService.Spec

	if len(vs.Http) == 0 {
		generatorLog.Errorf("Can't find HTTP Route in virtualService %v, fail back to default route", context.VirtualService)
		clusterName := model.BuildClusterName(model.TrafficDirectionOutbound, "",
			context.ServiceEntry.Spec.Hosts[0], int(context.ServiceEntry.Spec.Ports[0].Number))
		return defaultRoute(clusterName, tRPCService)
	}

	//Now we only support one HTTPRoute in a virtual service, more HTTPRoutes can be added if required
	http := vs.Http[0]
	var routeAction *routev3.RouteAction

	if len(http.Route) > 1 {
		routeAction = buildWeightedCluster(http, service)
	} else {
		routeAction = buildSingleCluster(http, service)
	}

	return &routev3.Route{
		Match: &routev3.RouteMatch{
			PathSpecifier: &routev3.RouteMatch_Prefix{
				Prefix: "/" + tRPCService,
			},
		},
		Action: &routev3.Route_Route{
			Route: routeAction,
		},
	}
}

func buildSingleCluster(http *networking.HTTPRoute, service *networking.ServiceEntry) *routev3.RouteAction {
	clusterName := model.BuildClusterName(model.TrafficDirectionOutbound, http.Route[0].Destination.Subset,
		service.Hosts[0], int(service.Ports[0].Number))
	return &routev3.RouteAction{
		ClusterSpecifier: &routev3.RouteAction_Cluster{
			Cluster: clusterName,
		},
	}
}

func buildWeightedCluster(http *networking.HTTPRoute, service *networking.ServiceEntry) *routev3.RouteAction {
	var clusterWeights []*envoy.WeightedCluster_ClusterWeight
	var totalWeight uint32

	for _, route := range http.Route {
		clusterName := model.BuildClusterName(model.TrafficDirectionOutbound, route.Destination.Subset,
			service.Hosts[0], int(service.Ports[0].Number))
		clusterWeight := &envoy.WeightedCluster_ClusterWeight{
			Name:   clusterName,
			Weight: &wrappers.UInt32Value{Value: uint32(route.Weight)},
		}
		clusterWeights = append(clusterWeights, clusterWeight)
		totalWeight += uint32(route.Weight)
	}

	return &routev3.RouteAction{
		ClusterSpecifier: &routev3.RouteAction_WeightedClusters{
			WeightedClusters: &envoy.WeightedCluster{
				Clusters:    clusterWeights,
				TotalWeight: &wrappers.UInt32Value{Value: totalWeight},
			},
		},
	}
}
