// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	mssqldatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabase"
	mssqldatabaseextendedauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabaseextendedauditingpolicy"
	mssqldatabasevulnerabilityassessmentrulebaseline "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlelasticpool"
	mssqlfailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfailovergroup"
	mssqlfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlfirewallrule"
	mssqljobagent "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobagent"
	mssqljobcredential "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqljobcredential"
	mssqlmanageddatabase "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancetransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancetransparentdataencryption"
	mssqlmanagedinstancevulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqloutboundfirewallrule"
	mssqlserver "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserver"
	mssqlserverdnsalias "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "github.com/upbound/provider-azure/internal/controller/cluster/sql/mssqlvirtualnetworkrule"
)

// Setup_sql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_sql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		mssqldatabase.Setup,
		mssqldatabaseextendedauditingpolicy.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqljobagent.Setup,
		mssqljobcredential.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancetransparentdataencryption.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
		mssqlservermicrosoftsupportauditingpolicy.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		mssqlvirtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
