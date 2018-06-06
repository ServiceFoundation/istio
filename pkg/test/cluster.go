//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package test

import (
	"k8s.io/client-go/rest"

	"istio.io/istio/pkg/test/impl/apiserver"
)

// Cluster dependency.
var Cluster = &ClusterDependency{exclusive: false}

// ExclusiveCluster dependency.
var ExclusiveCluster = &ClusterDependency{exclusive: true}

// GKE dependency
var GKE = &ClusterDependency{gke: true}

// ExclusiveGKE dependency
var ExclusiveGKE = &ClusterDependency{exclusive: true, gke: true}

// ClusterDependency represents a typed ClusterDependency dependency.
type ClusterDependency struct {
	exclusive bool
	gke       bool
}

// Dependency is the default dependency
var _ Dependency = &ClusterDependency{}

func (a *ClusterDependency) String() string {
	return "cluster"
}

func (a *ClusterDependency) Initialize() (interface{}, error) {
	// Use the underlying platform-specific implementation to instantiate a new APIServer instance.
	return apiserver.New()
}

func (a *ClusterDependency) Reset(interface{}) error {
	return nil
}

func (a *ClusterDependency) Cleanup(interface{}) {
}

// GetConfig returns the configuration
func (a *ClusterDependency) GetConfig() *rest.Config {
	return nil
}
