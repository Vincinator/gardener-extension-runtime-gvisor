// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package healthcheck

import (
	"time"

	healthcheckconfig "github.com/gardener/gardener/extensions/pkg/apis/config/v1alpha1"
	"github.com/gardener/gardener/extensions/pkg/controller/healthcheck"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener-extension-runtime-gvisor/pkg/gvisor"
)

var (
	defaultSyncPeriod = time.Second * 30
	// AddOptions are the default DefaultAddArgs for AddToManager.
	AddOptions = healthcheck.DefaultAddArgs{
		HealthCheckConfig: healthcheckconfig.HealthCheckConfig{SyncPeriod: metav1.Duration{Duration: defaultSyncPeriod}},
	}
)

// RegisterHealthChecks adds a controller with the given Options to the manager.
// The opts.Reconciler is being set with a newly instantiated Actuator.
func RegisterHealthChecks(mgr manager.Manager, opts healthcheck.DefaultAddArgs) error {
	return healthcheck.DefaultRegistration(
		gvisor.Type,
		extensionsv1alpha1.SchemeGroupVersion.WithKind(extensionsv1alpha1.ContainerRuntimeResource),
		func() client.ObjectList { return &extensionsv1alpha1.ContainerRuntimeList{} },
		func() extensionsv1alpha1.Object { return &extensionsv1alpha1.ContainerRuntime{} },
		mgr,
		opts,
		nil,
		[]healthcheck.ConditionTypeToHealthCheck{},
		sets.New[gardencorev1beta1.ConditionType](),
	)
}

// AddToManager adds a controller with the default Options.
func AddToManager(mgr manager.Manager) error {
	return RegisterHealthChecks(mgr, AddOptions)
}
