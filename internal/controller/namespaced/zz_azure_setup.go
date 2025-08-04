// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	resourcegroup "github.com/upbound/provider-azure/internal/controller/namespaced/azure/resourcegroup"
	resourceproviderregistration "github.com/upbound/provider-azure/internal/controller/namespaced/azure/resourceproviderregistration"
	subscription "github.com/upbound/provider-azure/internal/controller/namespaced/azure/subscription"
)

// Setup_azure creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_azure(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.Setup,
		resourceproviderregistration.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_azure creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_azure(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		resourcegroup.SetupGated,
		resourceproviderregistration.SetupGated,
		subscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
