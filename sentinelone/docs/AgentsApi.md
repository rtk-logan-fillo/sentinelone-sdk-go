# \AgentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAgents**](AgentsApi.md#GetAgents) | **Get** /web/api/v2.1/agents | Get Agents



## GetAgents

> AgentsSchemasAgentViewSchemaMany200 GetAgents(ctx).ScanStatusesNin(scanStatusesNin).Query(query).DecommissionedAtGte(decommissionedAtGte).AwsSecurityGroupsContains(awsSecurityGroupsContains).ThreatMitigationStatus(threatMitigationStatus).RegisteredAtLte(registeredAtLte).UpdatedAtGt(updatedAtGt).Domains(domains).NetworkQuarantineEnabled(networkQuarantineEnabled).MigrationStatus(migrationStatus).AccountIds(accountIds).Ids(ids).NetworkInterfaceInetContains(networkInterfaceInetContains).AgentVersionsNin(agentVersionsNin).CloudProvider(cloudProvider).SortOrder(sortOrder).OsTypesNin(osTypesNin).InstallerTypesNin(installerTypesNin).MachineTypes(machineTypes).EncryptedApplications(encryptedApplications).IsPendingUninstall(isPendingUninstall).HasLocalConfiguration(hasLocalConfiguration).IsUpToDate(isUpToDate).IsDecommissioned(isDecommissioned).ComputerName(computerName).LastLoggedInUserNameContains(lastLoggedInUserNameContains).AppsVulnerabilityStatuses(appsVulnerabilityStatuses).K8sVersionContains(k8sVersionContains).FirewallEnabled(firewallEnabled).DecommissionedAtLt(decommissionedAtLt).ActiveThreats(activeThreats).AwsSubnetIdsContains(awsSubnetIdsContains).CreatedAtGte(createdAtGte).K8sNodeNameContains(k8sNodeNameContains).TotalMemoryGte(totalMemoryGte).MachineTypesNin(machineTypesNin).LastActiveDateBetween(lastActiveDateBetween).NetworkInterfaceGatewayMacAddressContains(networkInterfaceGatewayMacAddressContains).AwsRoleContains(awsRoleContains).RemoteProfilingStatesNin(remoteProfilingStatesNin).LastActiveDateLte(lastActiveDateLte).CreatedAtLte(createdAtLte).ConsoleMigrationStatusesNin(consoleMigrationStatusesNin).SortBy(sortBy).FilteredSiteIds(filteredSiteIds).AdComputerNameContains(adComputerNameContains).RegisteredAtGt(registeredAtGt).CloudTagsContains(cloudTagsContains).AdUserQueryContains(adUserQueryContains).ThreatContentHash(threatContentHash).ScanStatuses(scanStatuses).OperationalStatesNin(operationalStatesNin).ThreatRebootRequired(threatRebootRequired).IsActive(isActive).TotalMemoryLt(totalMemoryLt).AdComputerQueryContains(adComputerQueryContains).UserActionsNeeded(userActionsNeeded).UpdatedAtBetween(updatedAtBetween).CpuCountGte(cpuCountGte).Uuids(uuids).SerialNumberContains(serialNumberContains).RegisteredAtBetween(registeredAtBetween).ScanStatus(scanStatus).ExternalIdContains(externalIdContains).TotalMemoryLte(totalMemoryLte).CountOnly(countOnly).LocationIds(locationIds).CloudNetworkContains(cloudNetworkContains).AzureResourceGroupContains(azureResourceGroupContains).CoreCountBetween(coreCountBetween).IsUninstalled(isUninstalled).FilterId(filterId).CpuIdContains(cpuIdContains).K8sTypeContains(k8sTypeContains).CloudProviderNin(cloudProviderNin).MitigationMode(mitigationMode).CloudAccountContains(cloudAccountContains).AdComputerMemberContains(adComputerMemberContains).ConsoleMigrationStatuses(consoleMigrationStatuses).LastActiveDateGt(lastActiveDateGt).OsArch(osArch).AgentContentUpdateIdContains(agentContentUpdateIdContains).RegisteredAtGte(registeredAtGte).CoreCountGt(coreCountGt).CpuCountLte(cpuCountLte).AgentPodNameContains(agentPodNameContains).Limit(limit).DecommissionedAtBetween(decommissionedAtBetween).ThreatCreatedAtLt(threatCreatedAtLt).CoreCountGte(coreCountGte).ClusterNameContains(clusterNameContains).CpuCountGt(cpuCountGt).AdQueryContains(adQueryContains).RangerVersions(rangerVersions).UserActionsNeededNin(userActionsNeededNin).ThreatCreatedAtGt(threatCreatedAtGt).CloudInstanceSizeContains(cloudInstanceSizeContains).AdQuery(adQuery).ExternalIpContains(externalIpContains).LocationEnabled(locationEnabled).TotalMemoryGt(totalMemoryGt).GatewayIp(gatewayIp).RegisteredAtLt(registeredAtLt).AgentVersions(agentVersions).InstallerTypes(installerTypes).CoreCountLte(coreCountLte).Cursor(cursor).TotalMemoryBetween(totalMemoryBetween).LastActiveDateGte(lastActiveDateGte).ComputerNameContains(computerNameContains).LocationIdsNin(locationIdsNin).ThreatCreatedAtLte(threatCreatedAtLte).HasTags(hasTags).UpdatedAtGte(updatedAtGte).ThreatResolved(threatResolved).DecommissionedAtGt(decommissionedAtGt).ThreatHidden(threatHidden).Infected(infected).UuidContains(uuidContains).NetworkStatusesNin(networkStatusesNin).CloudImageContains(cloudImageContains).SiteIds(siteIds).RangerStatus(rangerStatus).DomainsNin(domainsNin).ThreatCreatedAtGte(threatCreatedAtGte).CsvFilterId(csvFilterId).AdUserNameContains(adUserNameContains).AppsVulnerabilityStatusesNin(appsVulnerabilityStatusesNin).Uuid(uuid).CoreCountLt(coreCountLt).Skip(skip).MitigationModeSuspicious(mitigationModeSuspicious).RangerStatusesNin(rangerStatusesNin).CreatedAtGt(createdAtGt).RangerVersionsNin(rangerVersionsNin).FilteredGroupIds(filteredGroupIds).DecommissionedAtLte(decommissionedAtLte).OperationalStates(operationalStates).OsVersionContains(osVersionContains).UpdatedAtLte(updatedAtLte).UpdatedAtLt(updatedAtLt).K8sNodeLabelsContains(k8sNodeLabelsContains).AdUserMemberContains(adUserMemberContains).CloudLocationContains(cloudLocationContains).NetworkStatuses(networkStatuses).RemoteProfilingStates(remoteProfilingStates).NetworkInterfacePhysicalContains(networkInterfacePhysicalContains).AgentNamespaceContains(agentNamespaceContains).RangerStatuses(rangerStatuses).SkipCount(skipCount).ActiveThreatsGt(activeThreatsGt).LastActiveDateLt(lastActiveDateLt).CloudInstanceIdContains(cloudInstanceIdContains).CreatedAtLt(createdAtLt).CpuCountBetween(cpuCountBetween).GcpServiceAccountContains(gcpServiceAccountContains).ThreatCreatedAtBetween(threatCreatedAtBetween).OsTypes(osTypes).CpuCountLt(cpuCountLt).GroupIds(groupIds).ComputerNameLike(computerNameLike).TagsData(tagsData).CreatedAtBetween(createdAtBetween).Execute()

Get Agents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    scanStatusesNin := []string{"ScanStatusesNin_example"} // []string | Not included scan statuses. Example: \"started,aborted\". (optional)
    query := "query_example" // string | A free-text search term, will match applicable attributes (sub-string match). Note: Device's physical addresses will be matched if they start with the search term only (no match if they contain the term). Example: \"Linux\". (optional)
    decommissionedAtGte := time.Now() // time.Time | Agents decommissioned after or at this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    awsSecurityGroupsContains := []string{"Inner_example"} // []string | Free-text filter by aws securityGroups(supports multiple values) (optional)
    threatMitigationStatus := "threatMitigationStatus_example" // string | Include only Agents that have threats with this mitigation status. Example: \"mitigated\". (optional)
    registeredAtLte := time.Now() // time.Time | Agents registered before or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    updatedAtGt := time.Now() // time.Time | Agents updated after this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    domains := []string{"Inner_example"} // []string | Included network domains. Example: \"mybusiness.net,workgroup\". (optional)
    networkQuarantineEnabled := []bool{false} // []bool | The agents supports Network Quarantine Control and its enabled for the agent's group (optional)
    migrationStatus := "migrationStatus_example" // string | Migration status. Example: \"N/A\". (optional)
    accountIds := []string{"Inner_example"} // []string | List of Account IDs to filter by. Example: \"225494730938493804,225494730938493915\". (optional)
    ids := []string{"Inner_example"} // []string | A list of Agent IDs. Example: \"225494730938493804,225494730938493915\". (optional)
    networkInterfaceInetContains := []string{"Inner_example"} // []string | Free-text filter by local IP (supports multiple values). Example: \"192,10.0.0\". (optional)
    agentVersionsNin := []string{"Inner_example"} // []string | Agent versions not to include. Example: \"2.0.0.0,2.1.5.144\". (optional)
    cloudProvider := []string{"Inner_example"} // []string | Agents from which cloud provider (optional)
    sortOrder := "sortOrder_example" // string | Sort direction. Example: \"asc\". (optional)
    osTypesNin := []string{"OsTypesNin_example"} // []string | Not included OS types. Example: \"windows\". (optional)
    installerTypesNin := []string{"InstallerTypesNin_example"} // []string | Exclude Agents installed with these package types. Example: \".msi\". (optional)
    machineTypes := []string{"MachineTypes_example"} // []string | Included machine types. Example: \"laptop,desktop\". (optional)
    encryptedApplications := true // bool | Disk encryption status (optional)
    isPendingUninstall := true // bool | Include only Agents with pending uninstall requests (optional)
    hasLocalConfiguration := true // bool | Agent has a local configuration set (optional)
    isUpToDate := true // bool | Include only Agents with updated software (optional)
    isDecommissioned := []bool{false} // []bool | Include active, decommissioned or both. Example: \"True,False\". (optional)
    computerName := "computerName_example" // string | Computer name. Example: \"My Office Desktop\". (optional)
    lastLoggedInUserNameContains := []string{"Inner_example"} // []string | Free-text filter by username (supports multiple values). Example: \"admin,johnd1\". (optional)
    appsVulnerabilityStatuses := []string{"AppsVulnerabilityStatuses_example"} // []string | Apps vulnerability status in. Example: \"patch_required\". (optional)
    k8sVersionContains := []string{"Inner_example"} // []string | Free-text filter by K8s version (supports multiple values) (optional)
    firewallEnabled := []bool{false} // []bool | The agents supports Firewall Control and it is enabled for the agent's group (optional)
    decommissionedAtLt := time.Now() // time.Time | Agents decommissioned before this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    activeThreats := int32(56) // int32 | Include Agents with this amount of active threats. Example: \"3\". (optional)
    awsSubnetIdsContains := []string{"Inner_example"} // []string | Free-text filter by aws subnet ids (supports multiple values) (optional)
    createdAtGte := time.Now() // time.Time | Agents created after or at this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    k8sNodeNameContains := []string{"Inner_example"} // []string | Free-text filter by K8s node name (supports multiple values) (optional)
    totalMemoryGte := int32(56) // int32 | Memory size (MB, more than or equal) (optional)
    machineTypesNin := []string{"MachineTypesNin_example"} // []string | Not included machine types. Example: \"laptop,desktop\". (optional)
    lastActiveDateBetween := "lastActiveDateBetween_example" // string | Date range for last active date(format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978764288-1514978999999\". (optional)
    networkInterfaceGatewayMacAddressContains := []string{"Inner_example"} // []string | Free-text filter by Gateway MAC address (supports multiple values). Example: \"aa:0f,:41:\". (optional)
    awsRoleContains := []string{"Inner_example"} // []string | Free-text filter by aws role(supports multiple values) (optional)
    remoteProfilingStatesNin := []string{"Inner_example"} // []string | Do not include these Agent remote profiling states (optional)
    lastActiveDateLte := time.Now() // time.Time | Agents last active before or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    createdAtLte := time.Now() // time.Time | Agents created before or at this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    consoleMigrationStatusesNin := []string{"ConsoleMigrationStatusesNin_example"} // []string | Migration status nin. Example: \"N/A\". (optional)
    sortBy := "sortBy_example" // string | The column to sort the results by. Example: \"id\". (optional)
    filteredSiteIds := []string{"Inner_example"} // []string | List of Site IDs to filter by. Example: \"225494730938493804,225494730938493915\". (optional)
    adComputerNameContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory computer name string (supports multiple values). Example: \"DC=sentinelone\". (optional)
    registeredAtGt := time.Now() // time.Time | Agents registered after this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    cloudTagsContains := []string{"Inner_example"} // []string | Free-text filter by cloud tags (supports multiple values) (optional)
    adUserQueryContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory computer name or its groups (supports multiple values). Example: \"DC=sentinelone,John\". (optional)
    threatContentHash := "threatContentHash_example" // string | Include only Agents that have at least one threat with this content hash. Example: \"cf23df2207d99a74fbe169e3eba035e633b65d94\". (optional)
    scanStatuses := []string{"ScanStatuses_example"} // []string | Included scan statuses. Example: \"started,aborted\". (optional)
    operationalStatesNin := []string{"Inner_example"} // []string | Do not include these Agent operational states (optional)
    threatRebootRequired := []bool{false} // []bool | Has at least one threat with at least one mitigation action pending reboot to succeed (optional)
    isActive := true // bool | Include only active Agents (optional)
    totalMemoryLt := int32(56) // int32 | Memory size (MB, less than) (optional)
    adComputerQueryContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory computer name or its groups (supports multiple values). Example: \"DC=sentinelone,Windows\". (optional)
    userActionsNeeded := []string{"UserActionsNeeded_example"} // []string | Included pending user actions. Example: \"reboot_needed,upgrade_needed\". (optional)
    updatedAtBetween := "updatedAtBetween_example" // string | Date range for update time (format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978890136-1514978650130\". (optional)
    cpuCountGte := int32(56) // int32 | Number of CPUs (more than or equal) (optional)
    uuids := []string{"Inner_example"} // []string | A list of included UUIDs. Example: \"ff819e70af13be381993075eb0ce5f2f6de05b11,ff819e70af13be381993075eb0ce5f2f6de05c22\". (optional)
    serialNumberContains := []string{"Inner_example"} // []string | Free-text filter by Serial Number (supports multiple values) (optional)
    registeredAtBetween := "registeredAtBetween_example" // string | Date range for first registration time (format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978764288-1514978999999\". (optional)
    scanStatus := "scanStatus_example" // string | Scan status. Example: \"none\". (optional)
    externalIdContains := []string{"Inner_example"} // []string | Free-text filter by external ID (Customer ID). Example: \"Tag#1 - monitoring,Performance machine\". (optional)
    totalMemoryLte := int32(56) // int32 | Memory size (MB, less than or equal) (optional)
    countOnly := true // bool | If true, only total number of items will be returned, without any of the actual objects. (optional) (default to false)
    locationIds := []string{"Inner_example"} // []string | Include only Agents reporting these locations. Example: \"225494730938493804,225494730938493915\". (optional)
    cloudNetworkContains := []string{"Inner_example"} // []string | Free-text filter by cloud network (supports multiple values) (optional)
    azureResourceGroupContains := []string{"Inner_example"} // []string | Free-text filter by azure resource group(supports multiple values) (optional)
    coreCountBetween := "coreCountBetween_example" // string | Possible number of CPU cores (inclusive). Example: \"2-8\". (optional)
    isUninstalled := []bool{false} // []bool | Include installed, uninstalled or both. Example: \"True,False\". (optional)
    filterId := "filterId_example" // string | Include all Agents matching this saved filter. Example: \"225494730938493804\". (optional)
    cpuIdContains := []string{"Inner_example"} // []string | Free-text filter by CPU name (supports multiple values). Example: \"Intel,AMD\". (optional)
    k8sTypeContains := []string{"Inner_example"} // []string | Free-text filter by K8s type(supports multiple values) (optional)
    cloudProviderNin := []string{"Inner_example"} // []string | Exclude Agents from these cloud provider (optional)
    mitigationMode := "mitigationMode_example" // string | Agent mitigation mode policy. Example: \"detect\". (optional)
    cloudAccountContains := []string{"Inner_example"} // []string | Free-text filter by cloud account (supports multiple values) (optional)
    adComputerMemberContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory computer groups string (supports multiple values). Example: \"DC=sentinelone\". (optional)
    consoleMigrationStatuses := []string{"ConsoleMigrationStatuses_example"} // []string | Migration status in. Example: \"N/A\". (optional)
    lastActiveDateGt := time.Now() // time.Time | Agents last active after this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    osArch := "osArch_example" // string | OS architecture. Example: \"32 bit\". (optional)
    agentContentUpdateIdContains := []string{"Inner_example"} // []string | Free-text filter by content update ID (supports multiple values) (optional)
    registeredAtGte := time.Now() // time.Time | Agents registered after or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    coreCountGt := int32(56) // int32 | CPU cores (more than) (optional)
    cpuCountLte := int32(56) // int32 | Number of CPUs (less than or equal) (optional)
    agentPodNameContains := []string{"Inner_example"} // []string | Free-text filter by agent pod name (supports multiple values) (optional)
    limit := int32(56) // int32 | Limit number of returned items (1-1000). Example: \"10\". (optional) (default to 10)
    decommissionedAtBetween := "decommissionedAtBetween_example" // string | Date range for decommission time (format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978890136-1514978650130\". (optional)
    threatCreatedAtLt := time.Now() // time.Time | Agents with threats reported before this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    coreCountGte := int32(56) // int32 | CPU cores (more than or equal) (optional)
    clusterNameContains := []string{"Inner_example"} // []string | Free-text filter by cluster name (supports multiple values) (optional)
    cpuCountGt := int32(56) // int32 | Number of CPUs (more than) (optional)
    adQueryContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory string (supports multiple values). Example: \"DC=sentinelone\". (optional)
    rangerVersions := []string{"Inner_example"} // []string | Ranger versions to include. Example: \"2.0.0.0,2.1.5.144\". (optional)
    userActionsNeededNin := []string{"UserActionsNeededNin_example"} // []string | Excluded pending user actions. Example: \"reboot_needed,upgrade_needed\". (optional)
    threatCreatedAtGt := time.Now() // time.Time | Agents with threats reported after this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    cloudInstanceSizeContains := []string{"Inner_example"} // []string | Free-text filter by cloud instance size(supports multiple values) (optional)
    adQuery := "adQuery_example" // string | An Active Directory query string. Example: \"CN=Managers,DC=sentinelone,DC=com\". (optional)
    externalIpContains := []string{"Inner_example"} // []string | Free-text filter by visible IP (supports multiple values). Example: \"205,127.0\". (optional)
    locationEnabled := []bool{false} // []bool | The agents supports Location Awareness and it is enabled for the agent's group (optional)
    totalMemoryGt := int32(56) // int32 | Memory size (MB, more than) (optional)
    gatewayIp := "gatewayIp_example" // string | Gateway ip. Example: \"192.168.0.1\". (optional)
    registeredAtLt := time.Now() // time.Time | Agents registered before this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    agentVersions := []string{"Inner_example"} // []string | Agent versions to include. Example: \"2.0.0.0,2.1.5.144\". (optional)
    installerTypes := []string{"InstallerTypes_example"} // []string | Include only Agents installed with these package types. Example: \".msi\". (optional)
    coreCountLte := int32(56) // int32 | CPU cores (less than or equal) (optional)
    cursor := "cursor_example" // string | Cursor position returned by the last request. Use to iterate over more than 1000 items. Example: \"YWdlbnRfaWQ6NTgwMjkzODE=\". (optional)
    totalMemoryBetween := "totalMemoryBetween_example" // string | Total memory range (GB, inclusive). Example: \"4-8\". (optional)
    lastActiveDateGte := time.Now() // time.Time | Agents last active after or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    computerNameContains := []string{"Inner_example"} // []string | Free-text filter by computer name (supports multiple values). Example: \"john-office,WIN\". (optional)
    locationIdsNin := []string{"Inner_example"} // []string | Do not include only Agents reporting these locations. Example: \"225494730938493804,225494730938493915\". (optional)
    threatCreatedAtLte := time.Now() // time.Time | Agents with threats reported before or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    hasTags := true // bool | Include only Agents that have tags (optional)
    updatedAtGte := time.Now() // time.Time | Agents updated after or at this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    threatResolved := true // bool | Include only Agents with at least one resolved threat (optional)
    decommissionedAtGt := time.Now() // time.Time | Agents decommissioned after this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    threatHidden := true // bool | Include only Agents with at least one hidden threat (optional)
    infected := true // bool | Include only Agents with at least one active threat (optional)
    uuidContains := []string{"Inner_example"} // []string | Free-text filter by Agent UUID (supports multiple values). Example: \"e92-01928,b055\". (optional)
    networkStatusesNin := []string{"NetworkStatusesNin_example"} // []string | Included network statuses. Example: \"connected,connecting\". (optional)
    cloudImageContains := []string{"Inner_example"} // []string | Free-text filter by cloud image (supports multiple values) (optional)
    siteIds := []string{"Inner_example"} // []string | List of Site IDs to filter by. Example: \"225494730938493804,225494730938493915\". (optional)
    rangerStatus := "rangerStatus_example" // string | [DEPRECATED] Use rangerStatuses. Example: \"NotApplicable\". (optional)
    domainsNin := []string{"Inner_example"} // []string | Not included network domains. Example: \"mybusiness.net,workgroup\". (optional)
    threatCreatedAtGte := time.Now() // time.Time | Agents with threats reported after or at this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    csvFilterId := "csvFilterId_example" // string | The ID of the CSV file to filter by. Example: \"225494730938493804\". (optional)
    adUserNameContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory username string (supports multiple values). Example: \"DC=sentinelone\". (optional)
    appsVulnerabilityStatusesNin := []string{"AppsVulnerabilityStatusesNin_example"} // []string | Apps vulnerability status nin. Example: \"patch_required\". (optional)
    uuid := "uuid_example" // string | Agent's universally unique identifier. Example: \"ff819e70af13be381993075eb0ce5f2f6de05be2\". (optional)
    coreCountLt := int32(56) // int32 | CPU cores (less than) (optional)
    skip := int32(56) // int32 | Skip first number of items (0-1000). To iterate over more than 1000 items,  use \"cursor\". Example: \"150\". (optional)
    mitigationModeSuspicious := "mitigationModeSuspicious_example" // string | Mitigation mode policy for suspicious activity. Example: \"detect\". (optional)
    rangerStatusesNin := []string{"RangerStatusesNin_example"} // []string | Do not include these Ranger Statuses. Example: \"NotApplicable\". (optional)
    createdAtGt := time.Now() // time.Time | Agents created after this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    rangerVersionsNin := []string{"Inner_example"} // []string | Ranger versions not to include. Example: \"2.0.0.0,2.1.5.144\". (optional)
    filteredGroupIds := []string{"Inner_example"} // []string | List of Group IDs to filter by. Example: \"225494730938493804,225494730938493915\". (optional)
    decommissionedAtLte := time.Now() // time.Time | Agents decommissioned before this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    operationalStates := []string{"Inner_example"} // []string | Agent operational state (optional)
    osVersionContains := []string{"Inner_example"} // []string | Free-text filter by OS full name and version (supports multiple values). Example: \"Service Pack 1\". (optional)
    updatedAtLte := time.Now() // time.Time | Agents updated before or at this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    updatedAtLt := time.Now() // time.Time | Agents updated before this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    k8sNodeLabelsContains := []string{"Inner_example"} // []string | Free-text filter by K8s node labels (supports multiple values) (optional)
    adUserMemberContains := []string{"Inner_example"} // []string | Free-text filter by Active Directory user groups string (supports multiple values). Example: \"DC=sentinelone\". (optional)
    cloudLocationContains := []string{"Inner_example"} // []string | Free-text filter by cloud location (supports multiple values) (optional)
    networkStatuses := []string{"NetworkStatuses_example"} // []string | Included network statuses. Example: \"connected,connecting\". (optional)
    remoteProfilingStates := []string{"Inner_example"} // []string | Agent remote profiling state (optional)
    networkInterfacePhysicalContains := []string{"Inner_example"} // []string | Free-text filter by MAC address (supports multiple values). Example: \"aa:0f,:41:\". (optional)
    agentNamespaceContains := []string{"Inner_example"} // []string | Free-text filter by agent namespace (supports multiple values) (optional)
    rangerStatuses := []string{"RangerStatuses_example"} // []string | Status of Ranger. Example: \"NotApplicable\". (optional)
    skipCount := true // bool | If true, total number of items will not be calculated, which speeds up execution time. (optional)
    activeThreatsGt := int32(56) // int32 | Include Agents with at least this amount of active threats. Example: \"5\". (optional)
    lastActiveDateLt := time.Now() // time.Time | Agents last active before this time. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    cloudInstanceIdContains := []string{"Inner_example"} // []string | Free-text filter by cloud instance id(supports multiple values) (optional)
    createdAtLt := time.Now() // time.Time | Agents created before this timestamp. Example: \"2018-02-27T04:49:26.257525Z\". (optional)
    cpuCountBetween := "cpuCountBetween_example" // string | Possible number of CPU cores (inclusive). Example: \"2-8\". (optional)
    gcpServiceAccountContains := []string{"Inner_example"} // []string | Free-text filter by gcp service account (supports multiple values) (optional)
    threatCreatedAtBetween := "threatCreatedAtBetween_example" // string | Agents with threats reported in a date range (format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978764288-1514978999999\". (optional)
    osTypes := []string{"OsTypes_example"} // []string | Included OS types. Example: \"windows\". (optional)
    cpuCountLt := int32(56) // int32 | Number of CPUs (less than) (optional)
    groupIds := []string{"Inner_example"} // []string | List of Group IDs to filter by. Example: \"225494730938493804,225494730938493915\". (optional)
    computerNameLike := "computerNameLike_example" // string | Match computer name partially (substring). Example: \"Lab1\". (optional)
    tagsData := "tagsData_example" // string | A JSON formatted string, where each key represents a tag key, and each value represents a list of string values. Example: \"{\"key1\": [\"value1_1\", \"value1_2\"], \"key2\": [\"value2\"]}\". (optional)
    createdAtBetween := "createdAtBetween_example" // string | Date range for creation time (format: <from_timestamp>-<to_timestamp>, inclusive). Example: \"1514978890136-1514978650130\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.GetAgents(context.Background()).ScanStatusesNin(scanStatusesNin).Query(query).DecommissionedAtGte(decommissionedAtGte).AwsSecurityGroupsContains(awsSecurityGroupsContains).ThreatMitigationStatus(threatMitigationStatus).RegisteredAtLte(registeredAtLte).UpdatedAtGt(updatedAtGt).Domains(domains).NetworkQuarantineEnabled(networkQuarantineEnabled).MigrationStatus(migrationStatus).AccountIds(accountIds).Ids(ids).NetworkInterfaceInetContains(networkInterfaceInetContains).AgentVersionsNin(agentVersionsNin).CloudProvider(cloudProvider).SortOrder(sortOrder).OsTypesNin(osTypesNin).InstallerTypesNin(installerTypesNin).MachineTypes(machineTypes).EncryptedApplications(encryptedApplications).IsPendingUninstall(isPendingUninstall).HasLocalConfiguration(hasLocalConfiguration).IsUpToDate(isUpToDate).IsDecommissioned(isDecommissioned).ComputerName(computerName).LastLoggedInUserNameContains(lastLoggedInUserNameContains).AppsVulnerabilityStatuses(appsVulnerabilityStatuses).K8sVersionContains(k8sVersionContains).FirewallEnabled(firewallEnabled).DecommissionedAtLt(decommissionedAtLt).ActiveThreats(activeThreats).AwsSubnetIdsContains(awsSubnetIdsContains).CreatedAtGte(createdAtGte).K8sNodeNameContains(k8sNodeNameContains).TotalMemoryGte(totalMemoryGte).MachineTypesNin(machineTypesNin).LastActiveDateBetween(lastActiveDateBetween).NetworkInterfaceGatewayMacAddressContains(networkInterfaceGatewayMacAddressContains).AwsRoleContains(awsRoleContains).RemoteProfilingStatesNin(remoteProfilingStatesNin).LastActiveDateLte(lastActiveDateLte).CreatedAtLte(createdAtLte).ConsoleMigrationStatusesNin(consoleMigrationStatusesNin).SortBy(sortBy).FilteredSiteIds(filteredSiteIds).AdComputerNameContains(adComputerNameContains).RegisteredAtGt(registeredAtGt).CloudTagsContains(cloudTagsContains).AdUserQueryContains(adUserQueryContains).ThreatContentHash(threatContentHash).ScanStatuses(scanStatuses).OperationalStatesNin(operationalStatesNin).ThreatRebootRequired(threatRebootRequired).IsActive(isActive).TotalMemoryLt(totalMemoryLt).AdComputerQueryContains(adComputerQueryContains).UserActionsNeeded(userActionsNeeded).UpdatedAtBetween(updatedAtBetween).CpuCountGte(cpuCountGte).Uuids(uuids).SerialNumberContains(serialNumberContains).RegisteredAtBetween(registeredAtBetween).ScanStatus(scanStatus).ExternalIdContains(externalIdContains).TotalMemoryLte(totalMemoryLte).CountOnly(countOnly).LocationIds(locationIds).CloudNetworkContains(cloudNetworkContains).AzureResourceGroupContains(azureResourceGroupContains).CoreCountBetween(coreCountBetween).IsUninstalled(isUninstalled).FilterId(filterId).CpuIdContains(cpuIdContains).K8sTypeContains(k8sTypeContains).CloudProviderNin(cloudProviderNin).MitigationMode(mitigationMode).CloudAccountContains(cloudAccountContains).AdComputerMemberContains(adComputerMemberContains).ConsoleMigrationStatuses(consoleMigrationStatuses).LastActiveDateGt(lastActiveDateGt).OsArch(osArch).AgentContentUpdateIdContains(agentContentUpdateIdContains).RegisteredAtGte(registeredAtGte).CoreCountGt(coreCountGt).CpuCountLte(cpuCountLte).AgentPodNameContains(agentPodNameContains).Limit(limit).DecommissionedAtBetween(decommissionedAtBetween).ThreatCreatedAtLt(threatCreatedAtLt).CoreCountGte(coreCountGte).ClusterNameContains(clusterNameContains).CpuCountGt(cpuCountGt).AdQueryContains(adQueryContains).RangerVersions(rangerVersions).UserActionsNeededNin(userActionsNeededNin).ThreatCreatedAtGt(threatCreatedAtGt).CloudInstanceSizeContains(cloudInstanceSizeContains).AdQuery(adQuery).ExternalIpContains(externalIpContains).LocationEnabled(locationEnabled).TotalMemoryGt(totalMemoryGt).GatewayIp(gatewayIp).RegisteredAtLt(registeredAtLt).AgentVersions(agentVersions).InstallerTypes(installerTypes).CoreCountLte(coreCountLte).Cursor(cursor).TotalMemoryBetween(totalMemoryBetween).LastActiveDateGte(lastActiveDateGte).ComputerNameContains(computerNameContains).LocationIdsNin(locationIdsNin).ThreatCreatedAtLte(threatCreatedAtLte).HasTags(hasTags).UpdatedAtGte(updatedAtGte).ThreatResolved(threatResolved).DecommissionedAtGt(decommissionedAtGt).ThreatHidden(threatHidden).Infected(infected).UuidContains(uuidContains).NetworkStatusesNin(networkStatusesNin).CloudImageContains(cloudImageContains).SiteIds(siteIds).RangerStatus(rangerStatus).DomainsNin(domainsNin).ThreatCreatedAtGte(threatCreatedAtGte).CsvFilterId(csvFilterId).AdUserNameContains(adUserNameContains).AppsVulnerabilityStatusesNin(appsVulnerabilityStatusesNin).Uuid(uuid).CoreCountLt(coreCountLt).Skip(skip).MitigationModeSuspicious(mitigationModeSuspicious).RangerStatusesNin(rangerStatusesNin).CreatedAtGt(createdAtGt).RangerVersionsNin(rangerVersionsNin).FilteredGroupIds(filteredGroupIds).DecommissionedAtLte(decommissionedAtLte).OperationalStates(operationalStates).OsVersionContains(osVersionContains).UpdatedAtLte(updatedAtLte).UpdatedAtLt(updatedAtLt).K8sNodeLabelsContains(k8sNodeLabelsContains).AdUserMemberContains(adUserMemberContains).CloudLocationContains(cloudLocationContains).NetworkStatuses(networkStatuses).RemoteProfilingStates(remoteProfilingStates).NetworkInterfacePhysicalContains(networkInterfacePhysicalContains).AgentNamespaceContains(agentNamespaceContains).RangerStatuses(rangerStatuses).SkipCount(skipCount).ActiveThreatsGt(activeThreatsGt).LastActiveDateLt(lastActiveDateLt).CloudInstanceIdContains(cloudInstanceIdContains).CreatedAtLt(createdAtLt).CpuCountBetween(cpuCountBetween).GcpServiceAccountContains(gcpServiceAccountContains).ThreatCreatedAtBetween(threatCreatedAtBetween).OsTypes(osTypes).CpuCountLt(cpuCountLt).GroupIds(groupIds).ComputerNameLike(computerNameLike).TagsData(tagsData).CreatedAtBetween(createdAtBetween).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgents`: AgentsSchemasAgentViewSchemaMany200
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scanStatusesNin** | **[]string** | Not included scan statuses. Example: \&quot;started,aborted\&quot;. | 
 **query** | **string** | A free-text search term, will match applicable attributes (sub-string match). Note: Device&#39;s physical addresses will be matched if they start with the search term only (no match if they contain the term). Example: \&quot;Linux\&quot;. | 
 **decommissionedAtGte** | **time.Time** | Agents decommissioned after or at this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **awsSecurityGroupsContains** | **[]string** | Free-text filter by aws securityGroups(supports multiple values) | 
 **threatMitigationStatus** | **string** | Include only Agents that have threats with this mitigation status. Example: \&quot;mitigated\&quot;. | 
 **registeredAtLte** | **time.Time** | Agents registered before or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **updatedAtGt** | **time.Time** | Agents updated after this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **domains** | **[]string** | Included network domains. Example: \&quot;mybusiness.net,workgroup\&quot;. | 
 **networkQuarantineEnabled** | **[]bool** | The agents supports Network Quarantine Control and its enabled for the agent&#39;s group | 
 **migrationStatus** | **string** | Migration status. Example: \&quot;N/A\&quot;. | 
 **accountIds** | **[]string** | List of Account IDs to filter by. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **ids** | **[]string** | A list of Agent IDs. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **networkInterfaceInetContains** | **[]string** | Free-text filter by local IP (supports multiple values). Example: \&quot;192,10.0.0\&quot;. | 
 **agentVersionsNin** | **[]string** | Agent versions not to include. Example: \&quot;2.0.0.0,2.1.5.144\&quot;. | 
 **cloudProvider** | **[]string** | Agents from which cloud provider | 
 **sortOrder** | **string** | Sort direction. Example: \&quot;asc\&quot;. | 
 **osTypesNin** | **[]string** | Not included OS types. Example: \&quot;windows\&quot;. | 
 **installerTypesNin** | **[]string** | Exclude Agents installed with these package types. Example: \&quot;.msi\&quot;. | 
 **machineTypes** | **[]string** | Included machine types. Example: \&quot;laptop,desktop\&quot;. | 
 **encryptedApplications** | **bool** | Disk encryption status | 
 **isPendingUninstall** | **bool** | Include only Agents with pending uninstall requests | 
 **hasLocalConfiguration** | **bool** | Agent has a local configuration set | 
 **isUpToDate** | **bool** | Include only Agents with updated software | 
 **isDecommissioned** | **[]bool** | Include active, decommissioned or both. Example: \&quot;True,False\&quot;. | 
 **computerName** | **string** | Computer name. Example: \&quot;My Office Desktop\&quot;. | 
 **lastLoggedInUserNameContains** | **[]string** | Free-text filter by username (supports multiple values). Example: \&quot;admin,johnd1\&quot;. | 
 **appsVulnerabilityStatuses** | **[]string** | Apps vulnerability status in. Example: \&quot;patch_required\&quot;. | 
 **k8sVersionContains** | **[]string** | Free-text filter by K8s version (supports multiple values) | 
 **firewallEnabled** | **[]bool** | The agents supports Firewall Control and it is enabled for the agent&#39;s group | 
 **decommissionedAtLt** | **time.Time** | Agents decommissioned before this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **activeThreats** | **int32** | Include Agents with this amount of active threats. Example: \&quot;3\&quot;. | 
 **awsSubnetIdsContains** | **[]string** | Free-text filter by aws subnet ids (supports multiple values) | 
 **createdAtGte** | **time.Time** | Agents created after or at this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **k8sNodeNameContains** | **[]string** | Free-text filter by K8s node name (supports multiple values) | 
 **totalMemoryGte** | **int32** | Memory size (MB, more than or equal) | 
 **machineTypesNin** | **[]string** | Not included machine types. Example: \&quot;laptop,desktop\&quot;. | 
 **lastActiveDateBetween** | **string** | Date range for last active date(format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978764288-1514978999999\&quot;. | 
 **networkInterfaceGatewayMacAddressContains** | **[]string** | Free-text filter by Gateway MAC address (supports multiple values). Example: \&quot;aa:0f,:41:\&quot;. | 
 **awsRoleContains** | **[]string** | Free-text filter by aws role(supports multiple values) | 
 **remoteProfilingStatesNin** | **[]string** | Do not include these Agent remote profiling states | 
 **lastActiveDateLte** | **time.Time** | Agents last active before or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **createdAtLte** | **time.Time** | Agents created before or at this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **consoleMigrationStatusesNin** | **[]string** | Migration status nin. Example: \&quot;N/A\&quot;. | 
 **sortBy** | **string** | The column to sort the results by. Example: \&quot;id\&quot;. | 
 **filteredSiteIds** | **[]string** | List of Site IDs to filter by. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **adComputerNameContains** | **[]string** | Free-text filter by Active Directory computer name string (supports multiple values). Example: \&quot;DC&#x3D;sentinelone\&quot;. | 
 **registeredAtGt** | **time.Time** | Agents registered after this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **cloudTagsContains** | **[]string** | Free-text filter by cloud tags (supports multiple values) | 
 **adUserQueryContains** | **[]string** | Free-text filter by Active Directory computer name or its groups (supports multiple values). Example: \&quot;DC&#x3D;sentinelone,John\&quot;. | 
 **threatContentHash** | **string** | Include only Agents that have at least one threat with this content hash. Example: \&quot;cf23df2207d99a74fbe169e3eba035e633b65d94\&quot;. | 
 **scanStatuses** | **[]string** | Included scan statuses. Example: \&quot;started,aborted\&quot;. | 
 **operationalStatesNin** | **[]string** | Do not include these Agent operational states | 
 **threatRebootRequired** | **[]bool** | Has at least one threat with at least one mitigation action pending reboot to succeed | 
 **isActive** | **bool** | Include only active Agents | 
 **totalMemoryLt** | **int32** | Memory size (MB, less than) | 
 **adComputerQueryContains** | **[]string** | Free-text filter by Active Directory computer name or its groups (supports multiple values). Example: \&quot;DC&#x3D;sentinelone,Windows\&quot;. | 
 **userActionsNeeded** | **[]string** | Included pending user actions. Example: \&quot;reboot_needed,upgrade_needed\&quot;. | 
 **updatedAtBetween** | **string** | Date range for update time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978890136-1514978650130\&quot;. | 
 **cpuCountGte** | **int32** | Number of CPUs (more than or equal) | 
 **uuids** | **[]string** | A list of included UUIDs. Example: \&quot;ff819e70af13be381993075eb0ce5f2f6de05b11,ff819e70af13be381993075eb0ce5f2f6de05c22\&quot;. | 
 **serialNumberContains** | **[]string** | Free-text filter by Serial Number (supports multiple values) | 
 **registeredAtBetween** | **string** | Date range for first registration time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978764288-1514978999999\&quot;. | 
 **scanStatus** | **string** | Scan status. Example: \&quot;none\&quot;. | 
 **externalIdContains** | **[]string** | Free-text filter by external ID (Customer ID). Example: \&quot;Tag#1 - monitoring,Performance machine\&quot;. | 
 **totalMemoryLte** | **int32** | Memory size (MB, less than or equal) | 
 **countOnly** | **bool** | If true, only total number of items will be returned, without any of the actual objects. | [default to false]
 **locationIds** | **[]string** | Include only Agents reporting these locations. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **cloudNetworkContains** | **[]string** | Free-text filter by cloud network (supports multiple values) | 
 **azureResourceGroupContains** | **[]string** | Free-text filter by azure resource group(supports multiple values) | 
 **coreCountBetween** | **string** | Possible number of CPU cores (inclusive). Example: \&quot;2-8\&quot;. | 
 **isUninstalled** | **[]bool** | Include installed, uninstalled or both. Example: \&quot;True,False\&quot;. | 
 **filterId** | **string** | Include all Agents matching this saved filter. Example: \&quot;225494730938493804\&quot;. | 
 **cpuIdContains** | **[]string** | Free-text filter by CPU name (supports multiple values). Example: \&quot;Intel,AMD\&quot;. | 
 **k8sTypeContains** | **[]string** | Free-text filter by K8s type(supports multiple values) | 
 **cloudProviderNin** | **[]string** | Exclude Agents from these cloud provider | 
 **mitigationMode** | **string** | Agent mitigation mode policy. Example: \&quot;detect\&quot;. | 
 **cloudAccountContains** | **[]string** | Free-text filter by cloud account (supports multiple values) | 
 **adComputerMemberContains** | **[]string** | Free-text filter by Active Directory computer groups string (supports multiple values). Example: \&quot;DC&#x3D;sentinelone\&quot;. | 
 **consoleMigrationStatuses** | **[]string** | Migration status in. Example: \&quot;N/A\&quot;. | 
 **lastActiveDateGt** | **time.Time** | Agents last active after this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **osArch** | **string** | OS architecture. Example: \&quot;32 bit\&quot;. | 
 **agentContentUpdateIdContains** | **[]string** | Free-text filter by content update ID (supports multiple values) | 
 **registeredAtGte** | **time.Time** | Agents registered after or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **coreCountGt** | **int32** | CPU cores (more than) | 
 **cpuCountLte** | **int32** | Number of CPUs (less than or equal) | 
 **agentPodNameContains** | **[]string** | Free-text filter by agent pod name (supports multiple values) | 
 **limit** | **int32** | Limit number of returned items (1-1000). Example: \&quot;10\&quot;. | [default to 10]
 **decommissionedAtBetween** | **string** | Date range for decommission time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978890136-1514978650130\&quot;. | 
 **threatCreatedAtLt** | **time.Time** | Agents with threats reported before this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **coreCountGte** | **int32** | CPU cores (more than or equal) | 
 **clusterNameContains** | **[]string** | Free-text filter by cluster name (supports multiple values) | 
 **cpuCountGt** | **int32** | Number of CPUs (more than) | 
 **adQueryContains** | **[]string** | Free-text filter by Active Directory string (supports multiple values). Example: \&quot;DC&#x3D;sentinelone\&quot;. | 
 **rangerVersions** | **[]string** | Ranger versions to include. Example: \&quot;2.0.0.0,2.1.5.144\&quot;. | 
 **userActionsNeededNin** | **[]string** | Excluded pending user actions. Example: \&quot;reboot_needed,upgrade_needed\&quot;. | 
 **threatCreatedAtGt** | **time.Time** | Agents with threats reported after this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **cloudInstanceSizeContains** | **[]string** | Free-text filter by cloud instance size(supports multiple values) | 
 **adQuery** | **string** | An Active Directory query string. Example: \&quot;CN&#x3D;Managers,DC&#x3D;sentinelone,DC&#x3D;com\&quot;. | 
 **externalIpContains** | **[]string** | Free-text filter by visible IP (supports multiple values). Example: \&quot;205,127.0\&quot;. | 
 **locationEnabled** | **[]bool** | The agents supports Location Awareness and it is enabled for the agent&#39;s group | 
 **totalMemoryGt** | **int32** | Memory size (MB, more than) | 
 **gatewayIp** | **string** | Gateway ip. Example: \&quot;192.168.0.1\&quot;. | 
 **registeredAtLt** | **time.Time** | Agents registered before this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **agentVersions** | **[]string** | Agent versions to include. Example: \&quot;2.0.0.0,2.1.5.144\&quot;. | 
 **installerTypes** | **[]string** | Include only Agents installed with these package types. Example: \&quot;.msi\&quot;. | 
 **coreCountLte** | **int32** | CPU cores (less than or equal) | 
 **cursor** | **string** | Cursor position returned by the last request. Use to iterate over more than 1000 items. Example: \&quot;YWdlbnRfaWQ6NTgwMjkzODE&#x3D;\&quot;. | 
 **totalMemoryBetween** | **string** | Total memory range (GB, inclusive). Example: \&quot;4-8\&quot;. | 
 **lastActiveDateGte** | **time.Time** | Agents last active after or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **computerNameContains** | **[]string** | Free-text filter by computer name (supports multiple values). Example: \&quot;john-office,WIN\&quot;. | 
 **locationIdsNin** | **[]string** | Do not include only Agents reporting these locations. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **threatCreatedAtLte** | **time.Time** | Agents with threats reported before or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **hasTags** | **bool** | Include only Agents that have tags | 
 **updatedAtGte** | **time.Time** | Agents updated after or at this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **threatResolved** | **bool** | Include only Agents with at least one resolved threat | 
 **decommissionedAtGt** | **time.Time** | Agents decommissioned after this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **threatHidden** | **bool** | Include only Agents with at least one hidden threat | 
 **infected** | **bool** | Include only Agents with at least one active threat | 
 **uuidContains** | **[]string** | Free-text filter by Agent UUID (supports multiple values). Example: \&quot;e92-01928,b055\&quot;. | 
 **networkStatusesNin** | **[]string** | Included network statuses. Example: \&quot;connected,connecting\&quot;. | 
 **cloudImageContains** | **[]string** | Free-text filter by cloud image (supports multiple values) | 
 **siteIds** | **[]string** | List of Site IDs to filter by. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **rangerStatus** | **string** | [DEPRECATED] Use rangerStatuses. Example: \&quot;NotApplicable\&quot;. | 
 **domainsNin** | **[]string** | Not included network domains. Example: \&quot;mybusiness.net,workgroup\&quot;. | 
 **threatCreatedAtGte** | **time.Time** | Agents with threats reported after or at this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **csvFilterId** | **string** | The ID of the CSV file to filter by. Example: \&quot;225494730938493804\&quot;. | 
 **adUserNameContains** | **[]string** | Free-text filter by Active Directory username string (supports multiple values). Example: \&quot;DC&#x3D;sentinelone\&quot;. | 
 **appsVulnerabilityStatusesNin** | **[]string** | Apps vulnerability status nin. Example: \&quot;patch_required\&quot;. | 
 **uuid** | **string** | Agent&#39;s universally unique identifier. Example: \&quot;ff819e70af13be381993075eb0ce5f2f6de05be2\&quot;. | 
 **coreCountLt** | **int32** | CPU cores (less than) | 
 **skip** | **int32** | Skip first number of items (0-1000). To iterate over more than 1000 items,  use \&quot;cursor\&quot;. Example: \&quot;150\&quot;. | 
 **mitigationModeSuspicious** | **string** | Mitigation mode policy for suspicious activity. Example: \&quot;detect\&quot;. | 
 **rangerStatusesNin** | **[]string** | Do not include these Ranger Statuses. Example: \&quot;NotApplicable\&quot;. | 
 **createdAtGt** | **time.Time** | Agents created after this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **rangerVersionsNin** | **[]string** | Ranger versions not to include. Example: \&quot;2.0.0.0,2.1.5.144\&quot;. | 
 **filteredGroupIds** | **[]string** | List of Group IDs to filter by. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **decommissionedAtLte** | **time.Time** | Agents decommissioned before this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **operationalStates** | **[]string** | Agent operational state | 
 **osVersionContains** | **[]string** | Free-text filter by OS full name and version (supports multiple values). Example: \&quot;Service Pack 1\&quot;. | 
 **updatedAtLte** | **time.Time** | Agents updated before or at this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **updatedAtLt** | **time.Time** | Agents updated before this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **k8sNodeLabelsContains** | **[]string** | Free-text filter by K8s node labels (supports multiple values) | 
 **adUserMemberContains** | **[]string** | Free-text filter by Active Directory user groups string (supports multiple values). Example: \&quot;DC&#x3D;sentinelone\&quot;. | 
 **cloudLocationContains** | **[]string** | Free-text filter by cloud location (supports multiple values) | 
 **networkStatuses** | **[]string** | Included network statuses. Example: \&quot;connected,connecting\&quot;. | 
 **remoteProfilingStates** | **[]string** | Agent remote profiling state | 
 **networkInterfacePhysicalContains** | **[]string** | Free-text filter by MAC address (supports multiple values). Example: \&quot;aa:0f,:41:\&quot;. | 
 **agentNamespaceContains** | **[]string** | Free-text filter by agent namespace (supports multiple values) | 
 **rangerStatuses** | **[]string** | Status of Ranger. Example: \&quot;NotApplicable\&quot;. | 
 **skipCount** | **bool** | If true, total number of items will not be calculated, which speeds up execution time. | 
 **activeThreatsGt** | **int32** | Include Agents with at least this amount of active threats. Example: \&quot;5\&quot;. | 
 **lastActiveDateLt** | **time.Time** | Agents last active before this time. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **cloudInstanceIdContains** | **[]string** | Free-text filter by cloud instance id(supports multiple values) | 
 **createdAtLt** | **time.Time** | Agents created before this timestamp. Example: \&quot;2018-02-27T04:49:26.257525Z\&quot;. | 
 **cpuCountBetween** | **string** | Possible number of CPU cores (inclusive). Example: \&quot;2-8\&quot;. | 
 **gcpServiceAccountContains** | **[]string** | Free-text filter by gcp service account (supports multiple values) | 
 **threatCreatedAtBetween** | **string** | Agents with threats reported in a date range (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978764288-1514978999999\&quot;. | 
 **osTypes** | **[]string** | Included OS types. Example: \&quot;windows\&quot;. | 
 **cpuCountLt** | **int32** | Number of CPUs (less than) | 
 **groupIds** | **[]string** | List of Group IDs to filter by. Example: \&quot;225494730938493804,225494730938493915\&quot;. | 
 **computerNameLike** | **string** | Match computer name partially (substring). Example: \&quot;Lab1\&quot;. | 
 **tagsData** | **string** | A JSON formatted string, where each key represents a tag key, and each value represents a list of string values. Example: \&quot;{\&quot;key1\&quot;: [\&quot;value1_1\&quot;, \&quot;value1_2\&quot;], \&quot;key2\&quot;: [\&quot;value2\&quot;]}\&quot;. | 
 **createdAtBetween** | **string** | Date range for creation time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive). Example: \&quot;1514978890136-1514978650130\&quot;. | 

### Return type

[**AgentsSchemasAgentViewSchemaMany200**](AgentsSchemasAgentViewSchemaMany200.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

