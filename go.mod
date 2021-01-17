module github.com/aeraki-framework/aeraki

go 1.15

replace github.com/spf13/viper => github.com/istio/viper v1.3.3-0.20190515210538-2789fed3109c

// Old version had no license
replace github.com/chzyer/logex => github.com/chzyer/logex v1.1.11-0.20170329064859-445be9e134b2

// Avoid pulling in incompatible libraries
replace github.com/docker/distribution => github.com/docker/distribution v2.7.1+incompatible

// Avoid pulling in kubernetes/kubernetes
replace github.com/Microsoft/hcsshim => github.com/Microsoft/hcsshim v0.8.8-0.20200421182805-c3e488f0d815

// Client-go does not handle different versions of mergo due to some breaking changes - use the matching version
replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

// See https://github.com/kubernetes/kubernetes/issues/92867, there is a bug in the library
replace github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190815234213-e83c0a1c26c8

// There are some bugs in the Istio 1.8.0
// https://github.com/istio/istio/pull/29209
// https://github.com/istio/istio/pull/29296
replace istio.io/istio => github.com/zhaohuabing/istio v0.0.0-20201201123742-0738ca6370f3

// https://github.com/istio/api/pull/1774 add destination port support for envoyfilter
replace istio.io/api => github.com/istio/api v0.0.0-20201217155105-21c3bd1ba1d3

// Add tRPC config api
replace github.com/envoyproxy/go-control-plane => github.com/zhaohuabing/go-control-plane v0.9.8-0.20210115032445-00648eaf19d3

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/envoyproxy/go-control-plane v0.9.8-0.20201019204000-12785f608982
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.3
	github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20200904194848-62affa334b73
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	istio.io/api v0.0.0-20201125194658-3cee6a1d3ab4
	istio.io/istio v0.0.0-20201118224433-c87a4c874df2
	istio.io/pkg v0.0.0-20201112235759-c861803834b2
)
