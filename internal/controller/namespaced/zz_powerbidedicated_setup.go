// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	powerbiembedded "github.com/upbound/provider-azure/internal/controller/namespaced/powerbidedicated/powerbiembedded"
)

// Setup_powerbidedicated creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_powerbidedicated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		powerbiembedded.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_powerbidedicated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_powerbidedicated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		powerbiembedded.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
