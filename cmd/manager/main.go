/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	clusterapis "sigs.k8s.io/cluster-api/pkg/apis"
	"sigs.k8s.io/cluster-api/pkg/apis/cluster/common"
	capicluster "sigs.k8s.io/cluster-api/pkg/controller/cluster"
	capimachine "sigs.k8s.io/cluster-api/pkg/controller/machine"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"

	"sigs.k8s.io/cluster-api-provider-skeleton/pkg/cloud/skeleton/actuators/cluster"
	"sigs.k8s.io/cluster-api-provider-skeleton/pkg/cloud/skeleton/actuators/machine"
	"sigs.k8s.io/cluster-api-provider-skeleton/pkg/cloud/skeleton/providerconfig"
)

func main() {
	cfg := config.GetConfigOrDie()

	flag.Parse()
	log := logf.Log.WithName("skeleton-controller-manager")
	logf.SetLogger(logf.ZapLogger(false))
	entryLog := log.WithName("entrypoint")

	// Setup a Manager
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		entryLog.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	// Set up the actuators used by the manager
	clusterActuator, _ := cluster.NewActuator(cluster.ActuatorParams{})
	machineActuator, _ := machine.NewActuator(machine.ActuatorParams{})

	// The machineActuator implements ClusterProvider, see details in machine/actuator.go
	common.RegisterClusterProvisioner("skeleton", machineActuator)
	if err := providerconfig.AddToScheme(mgr.GetScheme()); err != nil {
		panic(err)
	}

	// Add the clusterapis to the manager scheme
	if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		panic(err)
	}

	// Add the machine controller to the manager and give it the actuator
	capimachine.AddWithActuator(mgr, machineActuator)

	// Add the cluster controller to the manager along with the actuator
	capicluster.AddWithActuator(mgr, clusterActuator)

	// Run the manager
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		entryLog.Error(err, "unable to run manager")
		os.Exit(1)
	}
}
