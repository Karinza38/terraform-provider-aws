// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backup

// Exports for use in tests only.
var (
	ResourceFramework               = resourceFramework
	ResourceLogicallyAirGappedVault = newLogicallyAirGappedVaultResource

	FindBackupVaultByName                   = findBackupVaultByName
	FindFrameworkByName                     = findFrameworkByName
	FindLogicallyAirGappedBackupVaultByName = findLogicallyAirGappedBackupVaultByName
	FindVaultAccessPolicyByName             = findVaultAccessPolicyByName
)
