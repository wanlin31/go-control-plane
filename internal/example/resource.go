// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package example

import (
	testres "github.com/envoyproxy/go-control-plane/pkg/test/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	
)

const (
	ClusterName  = "example_proxy_cluster"
	RouteName    = "local_route"
	ListenerName = "listener_test"
	ListenerPort = 10000
	UpstreamHost = "10.104.2.7"
	UpstreamPort = 50051
)

func GenerateSnapshot() *cache.Snapshot {
	snap, _ := cache.NewSnapshot("1",
		map[resource.Type][]types.Resource{
			resource.EndpointType: {testres.MakeEndpoint(ClusterName, UpstreamPort)},
			resource.ClusterType:  {testres.MakeCluster(testres.Ads,ClusterName)},
			resource.RouteType:    {testres.MakeRoute(RouteName, ClusterName)},
			resource.ListenerType: {testres.MakeRouteHTTPListener(testres.Ads, ListenerName, ListenerPort, RouteName)},
		},
	)
	return snap
}
