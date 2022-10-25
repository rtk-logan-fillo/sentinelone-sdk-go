# AgentsSchemasAgentsDangerousActionSchemaFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanStatusesNin** | Pointer to **[]string** | Not included scan statuses | [optional] 
**Query** | Pointer to **string** | A free-text search term, will match applicable attributes (sub-string match). Note: Device&#39;s physical addresses will be matched if they start with the search term only (no match if they contain the term). | [optional] 
**DecommissionedAtGte** | Pointer to **time.Time** | Agents decommissioned after or at this timestamp | [optional] 
**AwsSecurityGroupsContains** | Pointer to **[]string** | Free-text filter by aws securityGroups(supports multiple values) | [optional] 
**ThreatMitigationStatus** | Pointer to **string** | Include only Agents that have threats with this mitigation status | [optional] 
**RegisteredAtLte** | Pointer to **time.Time** | Agents registered before or at this time | [optional] 
**UpdatedAtGt** | Pointer to **time.Time** | Agents updated after this timestamp | [optional] 
**Domains** | Pointer to **[]string** | Included network domains | [optional] 
**NetworkQuarantineEnabled** | Pointer to **[]bool** | The agents supports Network Quarantine Control and its enabled for the agent&#39;s group | [optional] 
**MigrationStatus** | Pointer to **string** | Migration status | [optional] 
**AccountIds** | Pointer to **[]string** | List of Account IDs to filter by | [optional] 
**Ids** | Pointer to **[]string** | A list of Agent IDs | [optional] 
**NetworkInterfaceInetContains** | Pointer to **[]string** | Free-text filter by local IP (supports multiple values) | [optional] 
**AgentVersionsNin** | Pointer to **[]string** | Agent versions not to include | [optional] 
**CloudProvider** | Pointer to **[]string** | Agents from which cloud provider | [optional] 
**OsTypesNin** | Pointer to **[]string** | Not included OS types | [optional] 
**InstallerTypesNin** | Pointer to **[]string** | Exclude Agents installed with these package types | [optional] 
**MachineTypes** | Pointer to **[]string** | Included machine types | [optional] 
**EncryptedApplications** | Pointer to **bool** | Disk encryption status | [optional] 
**IsPendingUninstall** | Pointer to **bool** | Include only Agents with pending uninstall requests | [optional] 
**HasLocalConfiguration** | Pointer to **bool** | Agent has a local configuration set | [optional] 
**IsUpToDate** | Pointer to **bool** | Include only Agents with updated software | [optional] 
**IsDecommissioned** | Pointer to **[]bool** | Include active, decommissioned or both | [optional] 
**ComputerName** | Pointer to **string** | Computer name | [optional] 
**LastLoggedInUserNameContains** | Pointer to **[]string** | Free-text filter by username (supports multiple values) | [optional] 
**AppsVulnerabilityStatuses** | Pointer to **[]string** | Apps vulnerability status in | [optional] 
**K8sVersionContains** | Pointer to **[]string** | Free-text filter by K8s version (supports multiple values) | [optional] 
**FirewallEnabled** | Pointer to **[]bool** | The agents supports Firewall Control and it is enabled for the agent&#39;s group | [optional] 
**DecommissionedAtLt** | Pointer to **time.Time** | Agents decommissioned before this timestamp | [optional] 
**ActiveThreats** | Pointer to **int32** | Include Agents with this amount of active threats | [optional] 
**AwsSubnetIdsContains** | Pointer to **[]string** | Free-text filter by aws subnet ids (supports multiple values) | [optional] 
**CreatedAtGte** | Pointer to **time.Time** | Agents created after or at this timestamp | [optional] 
**K8sNodeNameContains** | Pointer to **[]string** | Free-text filter by K8s node name (supports multiple values) | [optional] 
**TotalMemoryGte** | Pointer to **int32** | Memory size (MB, more than or equal) | [optional] 
**MachineTypesNin** | Pointer to **[]string** | Not included machine types | [optional] 
**LastActiveDateBetween** | Pointer to **string** | Date range for last active date(format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 
**NetworkInterfaceGatewayMacAddressContains** | Pointer to **[]string** | Free-text filter by Gateway MAC address (supports multiple values) | [optional] 
**AwsRoleContains** | Pointer to **[]string** | Free-text filter by aws role(supports multiple values) | [optional] 
**RemoteProfilingStatesNin** | Pointer to **[]string** | Do not include these Agent remote profiling states | [optional] 
**LastActiveDateLte** | Pointer to **time.Time** | Agents last active before or at this time | [optional] 
**CreatedAtLte** | Pointer to **time.Time** | Agents created before or at this timestamp | [optional] 
**ConsoleMigrationStatusesNin** | Pointer to **[]string** | Migration status nin | [optional] 
**FilteredSiteIds** | Pointer to **[]string** | List of Site IDs to filter by | [optional] 
**AdComputerNameContains** | Pointer to **[]string** | Free-text filter by Active Directory computer name string (supports multiple values) | [optional] 
**RegisteredAtGt** | Pointer to **time.Time** | Agents registered after this time | [optional] 
**CloudTagsContains** | Pointer to **[]string** | Free-text filter by cloud tags (supports multiple values) | [optional] 
**AdUserQueryContains** | Pointer to **[]string** | Free-text filter by Active Directory computer name or its groups (supports multiple values) | [optional] 
**ThreatContentHash** | Pointer to **string** | Include only Agents that have at least one threat with this content hash | [optional] 
**ScanStatuses** | Pointer to **[]string** | Included scan statuses | [optional] 
**OperationalStatesNin** | Pointer to **[]string** | Do not include these Agent operational states | [optional] 
**ThreatRebootRequired** | Pointer to **[]bool** | Has at least one threat with at least one mitigation action pending reboot to succeed | [optional] 
**IsActive** | Pointer to **bool** | Include only active Agents | [optional] 
**TotalMemoryLt** | Pointer to **int32** | Memory size (MB, less than) | [optional] 
**AdComputerQueryContains** | Pointer to **[]string** | Free-text filter by Active Directory computer name or its groups (supports multiple values) | [optional] 
**UserActionsNeeded** | Pointer to **[]string** | Included pending user actions | [optional] 
**UpdatedAtBetween** | Pointer to **string** | Date range for update time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 
**CpuCountGte** | Pointer to **int32** | Number of CPUs (more than or equal) | [optional] 
**Uuids** | Pointer to **[]string** | A list of included UUIDs | [optional] 
**SerialNumberContains** | Pointer to **[]string** | Free-text filter by Serial Number (supports multiple values) | [optional] 
**RegisteredAtBetween** | Pointer to **string** | Date range for first registration time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 
**ScanStatus** | Pointer to **string** | Scan status | [optional] 
**ExternalIdContains** | Pointer to **[]string** | Free-text filter by external ID (Customer ID) | [optional] 
**TotalMemoryLte** | Pointer to **int32** | Memory size (MB, less than or equal) | [optional] 
**LocationIds** | Pointer to **[]string** | Include only Agents reporting these locations | [optional] 
**CloudNetworkContains** | Pointer to **[]string** | Free-text filter by cloud network (supports multiple values) | [optional] 
**AzureResourceGroupContains** | Pointer to **[]string** | Free-text filter by azure resource group(supports multiple values) | [optional] 
**CoreCountBetween** | Pointer to **string** | Possible number of CPU cores (inclusive) | [optional] 
**IsUninstalled** | Pointer to **[]bool** | Include installed, uninstalled or both | [optional] 
**FilterId** | Pointer to **string** | Include all Agents matching this saved filter | [optional] 
**CpuIdContains** | Pointer to **[]string** | Free-text filter by CPU name (supports multiple values) | [optional] 
**K8sTypeContains** | Pointer to **[]string** | Free-text filter by K8s type(supports multiple values) | [optional] 
**CloudProviderNin** | Pointer to **[]string** | Exclude Agents from these cloud provider | [optional] 
**MitigationMode** | Pointer to **string** | Agent mitigation mode policy | [optional] 
**CloudAccountContains** | Pointer to **[]string** | Free-text filter by cloud account (supports multiple values) | [optional] 
**AdComputerMemberContains** | Pointer to **[]string** | Free-text filter by Active Directory computer groups string (supports multiple values) | [optional] 
**ConsoleMigrationStatuses** | Pointer to **[]string** | Migration status in | [optional] 
**LastActiveDateGt** | Pointer to **time.Time** | Agents last active after this time | [optional] 
**OsArch** | Pointer to **string** | OS architecture | [optional] 
**AgentContentUpdateIdContains** | Pointer to **[]string** | Free-text filter by content update ID (supports multiple values) | [optional] 
**RegisteredAtGte** | Pointer to **time.Time** | Agents registered after or at this time | [optional] 
**CoreCountGt** | Pointer to **int32** | CPU cores (more than) | [optional] 
**CpuCountLte** | Pointer to **int32** | Number of CPUs (less than or equal) | [optional] 
**AgentPodNameContains** | Pointer to **[]string** | Free-text filter by agent pod name (supports multiple values) | [optional] 
**ThreatCreatedAtLt** | Pointer to **time.Time** | Agents with threats reported before this time | [optional] 
**DecommissionedAtBetween** | Pointer to **string** | Date range for decommission time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 
**CoreCountGte** | Pointer to **int32** | CPU cores (more than or equal) | [optional] 
**ClusterNameContains** | Pointer to **[]string** | Free-text filter by cluster name (supports multiple values) | [optional] 
**CpuCountGt** | Pointer to **int32** | Number of CPUs (more than) | [optional] 
**AdQueryContains** | Pointer to **[]string** | Free-text filter by Active Directory string (supports multiple values) | [optional] 
**RangerVersions** | Pointer to **[]string** | Ranger versions to include | [optional] 
**UserActionsNeededNin** | Pointer to **[]string** | Excluded pending user actions | [optional] 
**ThreatCreatedAtGt** | Pointer to **time.Time** | Agents with threats reported after this time | [optional] 
**CloudInstanceSizeContains** | Pointer to **[]string** | Free-text filter by cloud instance size(supports multiple values) | [optional] 
**AdQuery** | Pointer to **string** | An Active Directory query string | [optional] 
**ExternalIpContains** | Pointer to **[]string** | Free-text filter by visible IP (supports multiple values) | [optional] 
**LocationEnabled** | Pointer to **[]bool** | The agents supports Location Awareness and it is enabled for the agent&#39;s group | [optional] 
**TotalMemoryGt** | Pointer to **int32** | Memory size (MB, more than) | [optional] 
**GatewayIp** | Pointer to **string** | Gateway ip | [optional] 
**RegisteredAtLt** | Pointer to **time.Time** | Agents registered before this time | [optional] 
**AgentVersions** | Pointer to **[]string** | Agent versions to include | [optional] 
**InstallerTypes** | Pointer to **[]string** | Include only Agents installed with these package types | [optional] 
**CoreCountLte** | Pointer to **int32** | CPU cores (less than or equal) | [optional] 
**TotalMemoryBetween** | Pointer to **string** | Total memory range (GB, inclusive) | [optional] 
**LastActiveDateGte** | Pointer to **time.Time** | Agents last active after or at this time | [optional] 
**ComputerNameContains** | Pointer to **[]string** | Free-text filter by computer name (supports multiple values) | [optional] 
**LocationIdsNin** | Pointer to **[]string** | Do not include only Agents reporting these locations | [optional] 
**ThreatCreatedAtLte** | Pointer to **time.Time** | Agents with threats reported before or at this time | [optional] 
**HasTags** | Pointer to **bool** | Include only Agents that have tags | [optional] 
**UpdatedAtGte** | Pointer to **time.Time** | Agents updated after or at this timestamp | [optional] 
**ThreatResolved** | Pointer to **bool** | Include only Agents with at least one resolved threat | [optional] 
**DecommissionedAtGt** | Pointer to **time.Time** | Agents decommissioned after this timestamp | [optional] 
**ThreatHidden** | Pointer to **bool** | Include only Agents with at least one hidden threat | [optional] 
**Infected** | Pointer to **bool** | Include only Agents with at least one active threat | [optional] 
**UuidContains** | Pointer to **[]string** | Free-text filter by Agent UUID (supports multiple values) | [optional] 
**NetworkStatusesNin** | Pointer to **[]string** | Included network statuses | [optional] 
**CloudImageContains** | Pointer to **[]string** | Free-text filter by cloud image (supports multiple values) | [optional] 
**SiteIds** | Pointer to **[]string** | List of Site IDs to filter by | [optional] 
**RangerStatus** | Pointer to **string** | [DEPRECATED] Use rangerStatuses. | [optional] 
**DomainsNin** | Pointer to **[]string** | Not included network domains | [optional] 
**ThreatCreatedAtGte** | Pointer to **time.Time** | Agents with threats reported after or at this time | [optional] 
**CsvFilterId** | Pointer to **string** | The ID of the CSV file to filter by | [optional] 
**AdUserNameContains** | Pointer to **[]string** | Free-text filter by Active Directory username string (supports multiple values) | [optional] 
**AppsVulnerabilityStatusesNin** | Pointer to **[]string** | Apps vulnerability status nin | [optional] 
**Uuid** | Pointer to **string** | Agent&#39;s universally unique identifier | [optional] 
**CoreCountLt** | Pointer to **int32** | CPU cores (less than) | [optional] 
**MitigationModeSuspicious** | Pointer to **string** | Mitigation mode policy for suspicious activity | [optional] 
**RangerStatusesNin** | Pointer to **[]string** | Do not include these Ranger Statuses | [optional] 
**CreatedAtGt** | Pointer to **time.Time** | Agents created after this timestamp | [optional] 
**RangerVersionsNin** | Pointer to **[]string** | Ranger versions not to include | [optional] 
**FilteredGroupIds** | Pointer to **[]string** | List of Group IDs to filter by | [optional] 
**DecommissionedAtLte** | Pointer to **time.Time** | Agents decommissioned before this timestamp | [optional] 
**OperationalStates** | Pointer to **[]string** | Agent operational state | [optional] 
**OsVersionContains** | Pointer to **[]string** | Free-text filter by OS full name and version (supports multiple values) | [optional] 
**UpdatedAtLte** | Pointer to **time.Time** | Agents updated before or at this timestamp | [optional] 
**UpdatedAtLt** | Pointer to **time.Time** | Agents updated before this timestamp | [optional] 
**K8sNodeLabelsContains** | Pointer to **[]string** | Free-text filter by K8s node labels (supports multiple values) | [optional] 
**AdUserMemberContains** | Pointer to **[]string** | Free-text filter by Active Directory user groups string (supports multiple values) | [optional] 
**CloudLocationContains** | Pointer to **[]string** | Free-text filter by cloud location (supports multiple values) | [optional] 
**NetworkStatuses** | Pointer to **[]string** | Included network statuses | [optional] 
**RemoteProfilingStates** | Pointer to **[]string** | Agent remote profiling state | [optional] 
**NetworkInterfacePhysicalContains** | Pointer to **[]string** | Free-text filter by MAC address (supports multiple values) | [optional] 
**AgentNamespaceContains** | Pointer to **[]string** | Free-text filter by agent namespace (supports multiple values) | [optional] 
**RangerStatuses** | Pointer to **[]string** | Status of Ranger | [optional] 
**LastActiveDateLt** | Pointer to **time.Time** | Agents last active before this time | [optional] 
**ActiveThreatsGt** | Pointer to **int32** | Include Agents with at least this amount of active threats | [optional] 
**CloudInstanceIdContains** | Pointer to **[]string** | Free-text filter by cloud instance id(supports multiple values) | [optional] 
**CreatedAtLt** | Pointer to **time.Time** | Agents created before this timestamp | [optional] 
**CpuCountBetween** | Pointer to **string** | Possible number of CPU cores (inclusive) | [optional] 
**GcpServiceAccountContains** | Pointer to **[]string** | Free-text filter by gcp service account (supports multiple values) | [optional] 
**ThreatCreatedAtBetween** | Pointer to **string** | Agents with threats reported in a date range (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 
**OsTypes** | Pointer to **[]string** | Included OS types | [optional] 
**CpuCountLt** | Pointer to **int32** | Number of CPUs (less than) | [optional] 
**GroupIds** | Pointer to **[]string** | List of network groups | [optional] 
**ComputerNameLike** | Pointer to **string** | Match computer name partially (substring) | [optional] 
**TagsData** | Pointer to **string** | A JSON formatted string, where each key represents a tag key, and each value represents a list of string values | [optional] 
**CreatedAtBetween** | Pointer to **string** | Date range for creation time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 

## Methods

### NewAgentsSchemasAgentsDangerousActionSchemaFilter

`func NewAgentsSchemasAgentsDangerousActionSchemaFilter() *AgentsSchemasAgentsDangerousActionSchemaFilter`

NewAgentsSchemasAgentsDangerousActionSchemaFilter instantiates a new AgentsSchemasAgentsDangerousActionSchemaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentsDangerousActionSchemaFilterWithDefaults

`func NewAgentsSchemasAgentsDangerousActionSchemaFilterWithDefaults() *AgentsSchemasAgentsDangerousActionSchemaFilter`

NewAgentsSchemasAgentsDangerousActionSchemaFilterWithDefaults instantiates a new AgentsSchemasAgentsDangerousActionSchemaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatusesNin() []string`

GetScanStatusesNin returns the ScanStatusesNin field if non-nil, zero value otherwise.

### GetScanStatusesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatusesNinOk() (*[]string, bool)`

GetScanStatusesNinOk returns a tuple with the ScanStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetScanStatusesNin(v []string)`

SetScanStatusesNin sets ScanStatusesNin field to given value.

### HasScanStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasScanStatusesNin() bool`

HasScanStatusesNin returns a boolean if a field has been set.

### GetQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDecommissionedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtGte() time.Time`

GetDecommissionedAtGte returns the DecommissionedAtGte field if non-nil, zero value otherwise.

### GetDecommissionedAtGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtGteOk() (*time.Time, bool)`

GetDecommissionedAtGteOk returns a tuple with the DecommissionedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDecommissionedAtGte(v time.Time)`

SetDecommissionedAtGte sets DecommissionedAtGte field to given value.

### HasDecommissionedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDecommissionedAtGte() bool`

HasDecommissionedAtGte returns a boolean if a field has been set.

### GetAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsSecurityGroupsContains() []string`

GetAwsSecurityGroupsContains returns the AwsSecurityGroupsContains field if non-nil, zero value otherwise.

### GetAwsSecurityGroupsContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsSecurityGroupsContainsOk() (*[]string, bool)`

GetAwsSecurityGroupsContainsOk returns a tuple with the AwsSecurityGroupsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAwsSecurityGroupsContains(v []string)`

SetAwsSecurityGroupsContains sets AwsSecurityGroupsContains field to given value.

### HasAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAwsSecurityGroupsContains() bool`

HasAwsSecurityGroupsContains returns a boolean if a field has been set.

### GetThreatMitigationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatMitigationStatus() string`

GetThreatMitigationStatus returns the ThreatMitigationStatus field if non-nil, zero value otherwise.

### GetThreatMitigationStatusOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatMitigationStatusOk() (*string, bool)`

GetThreatMitigationStatusOk returns a tuple with the ThreatMitigationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMitigationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatMitigationStatus(v string)`

SetThreatMitigationStatus sets ThreatMitigationStatus field to given value.

### HasThreatMitigationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatMitigationStatus() bool`

HasThreatMitigationStatus returns a boolean if a field has been set.

### GetRegisteredAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtLte() time.Time`

GetRegisteredAtLte returns the RegisteredAtLte field if non-nil, zero value otherwise.

### GetRegisteredAtLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtLteOk() (*time.Time, bool)`

GetRegisteredAtLteOk returns a tuple with the RegisteredAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRegisteredAtLte(v time.Time)`

SetRegisteredAtLte sets RegisteredAtLte field to given value.

### HasRegisteredAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRegisteredAtLte() bool`

HasRegisteredAtLte returns a boolean if a field has been set.

### GetUpdatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtGt() time.Time`

GetUpdatedAtGt returns the UpdatedAtGt field if non-nil, zero value otherwise.

### GetUpdatedAtGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtGtOk() (*time.Time, bool)`

GetUpdatedAtGtOk returns a tuple with the UpdatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUpdatedAtGt(v time.Time)`

SetUpdatedAtGt sets UpdatedAtGt field to given value.

### HasUpdatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUpdatedAtGt() bool`

HasUpdatedAtGt returns a boolean if a field has been set.

### GetDomains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkQuarantineEnabled() []bool`

GetNetworkQuarantineEnabled returns the NetworkQuarantineEnabled field if non-nil, zero value otherwise.

### GetNetworkQuarantineEnabledOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkQuarantineEnabledOk() (*[]bool, bool)`

GetNetworkQuarantineEnabledOk returns a tuple with the NetworkQuarantineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkQuarantineEnabled(v []bool)`

SetNetworkQuarantineEnabled sets NetworkQuarantineEnabled field to given value.

### HasNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkQuarantineEnabled() bool`

HasNetworkQuarantineEnabled returns a boolean if a field has been set.

### GetMigrationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMigrationStatus() string`

GetMigrationStatus returns the MigrationStatus field if non-nil, zero value otherwise.

### GetMigrationStatusOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMigrationStatusOk() (*string, bool)`

GetMigrationStatusOk returns a tuple with the MigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetMigrationStatus(v string)`

SetMigrationStatus sets MigrationStatus field to given value.

### HasMigrationStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasMigrationStatus() bool`

HasMigrationStatus returns a boolean if a field has been set.

### GetAccountIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfaceInetContains() []string`

GetNetworkInterfaceInetContains returns the NetworkInterfaceInetContains field if non-nil, zero value otherwise.

### GetNetworkInterfaceInetContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfaceInetContainsOk() (*[]string, bool)`

GetNetworkInterfaceInetContainsOk returns a tuple with the NetworkInterfaceInetContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkInterfaceInetContains(v []string)`

SetNetworkInterfaceInetContains sets NetworkInterfaceInetContains field to given value.

### HasNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkInterfaceInetContains() bool`

HasNetworkInterfaceInetContains returns a boolean if a field has been set.

### GetAgentVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentVersionsNin() []string`

GetAgentVersionsNin returns the AgentVersionsNin field if non-nil, zero value otherwise.

### GetAgentVersionsNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentVersionsNinOk() (*[]string, bool)`

GetAgentVersionsNinOk returns a tuple with the AgentVersionsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAgentVersionsNin(v []string)`

SetAgentVersionsNin sets AgentVersionsNin field to given value.

### HasAgentVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAgentVersionsNin() bool`

HasAgentVersionsNin returns a boolean if a field has been set.

### GetCloudProvider

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudProvider() []string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudProviderOk() (*[]string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudProvider(v []string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetOsTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsTypesNin() []string`

GetOsTypesNin returns the OsTypesNin field if non-nil, zero value otherwise.

### GetOsTypesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsTypesNinOk() (*[]string, bool)`

GetOsTypesNinOk returns a tuple with the OsTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOsTypesNin(v []string)`

SetOsTypesNin sets OsTypesNin field to given value.

### HasOsTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOsTypesNin() bool`

HasOsTypesNin returns a boolean if a field has been set.

### GetInstallerTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInstallerTypesNin() []string`

GetInstallerTypesNin returns the InstallerTypesNin field if non-nil, zero value otherwise.

### GetInstallerTypesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInstallerTypesNinOk() (*[]string, bool)`

GetInstallerTypesNinOk returns a tuple with the InstallerTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetInstallerTypesNin(v []string)`

SetInstallerTypesNin sets InstallerTypesNin field to given value.

### HasInstallerTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasInstallerTypesNin() bool`

HasInstallerTypesNin returns a boolean if a field has been set.

### GetMachineTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMachineTypes() []string`

GetMachineTypes returns the MachineTypes field if non-nil, zero value otherwise.

### GetMachineTypesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMachineTypesOk() (*[]string, bool)`

GetMachineTypesOk returns a tuple with the MachineTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetMachineTypes(v []string)`

SetMachineTypes sets MachineTypes field to given value.

### HasMachineTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasMachineTypes() bool`

HasMachineTypes returns a boolean if a field has been set.

### GetEncryptedApplications

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetEncryptedApplications() bool`

GetEncryptedApplications returns the EncryptedApplications field if non-nil, zero value otherwise.

### GetEncryptedApplicationsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetEncryptedApplicationsOk() (*bool, bool)`

GetEncryptedApplicationsOk returns a tuple with the EncryptedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedApplications

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetEncryptedApplications(v bool)`

SetEncryptedApplications sets EncryptedApplications field to given value.

### HasEncryptedApplications

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasEncryptedApplications() bool`

HasEncryptedApplications returns a boolean if a field has been set.

### GetIsPendingUninstall

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsPendingUninstall() bool`

GetIsPendingUninstall returns the IsPendingUninstall field if non-nil, zero value otherwise.

### GetIsPendingUninstallOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsPendingUninstallOk() (*bool, bool)`

GetIsPendingUninstallOk returns a tuple with the IsPendingUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPendingUninstall

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIsPendingUninstall(v bool)`

SetIsPendingUninstall sets IsPendingUninstall field to given value.

### HasIsPendingUninstall

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIsPendingUninstall() bool`

HasIsPendingUninstall returns a boolean if a field has been set.

### GetHasLocalConfiguration

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetHasLocalConfiguration() bool`

GetHasLocalConfiguration returns the HasLocalConfiguration field if non-nil, zero value otherwise.

### GetHasLocalConfigurationOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetHasLocalConfigurationOk() (*bool, bool)`

GetHasLocalConfigurationOk returns a tuple with the HasLocalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLocalConfiguration

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetHasLocalConfiguration(v bool)`

SetHasLocalConfiguration sets HasLocalConfiguration field to given value.

### HasHasLocalConfiguration

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasHasLocalConfiguration() bool`

HasHasLocalConfiguration returns a boolean if a field has been set.

### GetIsUpToDate

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsUpToDate() bool`

GetIsUpToDate returns the IsUpToDate field if non-nil, zero value otherwise.

### GetIsUpToDateOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsUpToDateOk() (*bool, bool)`

GetIsUpToDateOk returns a tuple with the IsUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpToDate

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIsUpToDate(v bool)`

SetIsUpToDate sets IsUpToDate field to given value.

### HasIsUpToDate

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIsUpToDate() bool`

HasIsUpToDate returns a boolean if a field has been set.

### GetIsDecommissioned

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsDecommissioned() []bool`

GetIsDecommissioned returns the IsDecommissioned field if non-nil, zero value otherwise.

### GetIsDecommissionedOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsDecommissionedOk() (*[]bool, bool)`

GetIsDecommissionedOk returns a tuple with the IsDecommissioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecommissioned

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIsDecommissioned(v []bool)`

SetIsDecommissioned sets IsDecommissioned field to given value.

### HasIsDecommissioned

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIsDecommissioned() bool`

HasIsDecommissioned returns a boolean if a field has been set.

### GetComputerName

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastLoggedInUserNameContains() []string`

GetLastLoggedInUserNameContains returns the LastLoggedInUserNameContains field if non-nil, zero value otherwise.

### GetLastLoggedInUserNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastLoggedInUserNameContainsOk() (*[]string, bool)`

GetLastLoggedInUserNameContainsOk returns a tuple with the LastLoggedInUserNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastLoggedInUserNameContains(v []string)`

SetLastLoggedInUserNameContains sets LastLoggedInUserNameContains field to given value.

### HasLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastLoggedInUserNameContains() bool`

HasLastLoggedInUserNameContains returns a boolean if a field has been set.

### GetAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAppsVulnerabilityStatuses() []string`

GetAppsVulnerabilityStatuses returns the AppsVulnerabilityStatuses field if non-nil, zero value otherwise.

### GetAppsVulnerabilityStatusesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAppsVulnerabilityStatusesOk() (*[]string, bool)`

GetAppsVulnerabilityStatusesOk returns a tuple with the AppsVulnerabilityStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAppsVulnerabilityStatuses(v []string)`

SetAppsVulnerabilityStatuses sets AppsVulnerabilityStatuses field to given value.

### HasAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAppsVulnerabilityStatuses() bool`

HasAppsVulnerabilityStatuses returns a boolean if a field has been set.

### GetK8sVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sVersionContains() []string`

GetK8sVersionContains returns the K8sVersionContains field if non-nil, zero value otherwise.

### GetK8sVersionContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sVersionContainsOk() (*[]string, bool)`

GetK8sVersionContainsOk returns a tuple with the K8sVersionContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetK8sVersionContains(v []string)`

SetK8sVersionContains sets K8sVersionContains field to given value.

### HasK8sVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasK8sVersionContains() bool`

HasK8sVersionContains returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFirewallEnabled() []bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFirewallEnabledOk() (*[]bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetFirewallEnabled(v []bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetDecommissionedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtLt() time.Time`

GetDecommissionedAtLt returns the DecommissionedAtLt field if non-nil, zero value otherwise.

### GetDecommissionedAtLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtLtOk() (*time.Time, bool)`

GetDecommissionedAtLtOk returns a tuple with the DecommissionedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDecommissionedAtLt(v time.Time)`

SetDecommissionedAtLt sets DecommissionedAtLt field to given value.

### HasDecommissionedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDecommissionedAtLt() bool`

HasDecommissionedAtLt returns a boolean if a field has been set.

### GetActiveThreats

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetActiveThreats() int32`

GetActiveThreats returns the ActiveThreats field if non-nil, zero value otherwise.

### GetActiveThreatsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetActiveThreatsOk() (*int32, bool)`

GetActiveThreatsOk returns a tuple with the ActiveThreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreats

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetActiveThreats(v int32)`

SetActiveThreats sets ActiveThreats field to given value.

### HasActiveThreats

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasActiveThreats() bool`

HasActiveThreats returns a boolean if a field has been set.

### GetAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsSubnetIdsContains() []string`

GetAwsSubnetIdsContains returns the AwsSubnetIdsContains field if non-nil, zero value otherwise.

### GetAwsSubnetIdsContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsSubnetIdsContainsOk() (*[]string, bool)`

GetAwsSubnetIdsContainsOk returns a tuple with the AwsSubnetIdsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAwsSubnetIdsContains(v []string)`

SetAwsSubnetIdsContains sets AwsSubnetIdsContains field to given value.

### HasAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAwsSubnetIdsContains() bool`

HasAwsSubnetIdsContains returns a boolean if a field has been set.

### GetCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtGte() time.Time`

GetCreatedAtGte returns the CreatedAtGte field if non-nil, zero value otherwise.

### GetCreatedAtGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtGteOk() (*time.Time, bool)`

GetCreatedAtGteOk returns a tuple with the CreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCreatedAtGte(v time.Time)`

SetCreatedAtGte sets CreatedAtGte field to given value.

### HasCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCreatedAtGte() bool`

HasCreatedAtGte returns a boolean if a field has been set.

### GetK8sNodeNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sNodeNameContains() []string`

GetK8sNodeNameContains returns the K8sNodeNameContains field if non-nil, zero value otherwise.

### GetK8sNodeNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sNodeNameContainsOk() (*[]string, bool)`

GetK8sNodeNameContainsOk returns a tuple with the K8sNodeNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetK8sNodeNameContains(v []string)`

SetK8sNodeNameContains sets K8sNodeNameContains field to given value.

### HasK8sNodeNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasK8sNodeNameContains() bool`

HasK8sNodeNameContains returns a boolean if a field has been set.

### GetTotalMemoryGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryGte() int32`

GetTotalMemoryGte returns the TotalMemoryGte field if non-nil, zero value otherwise.

### GetTotalMemoryGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryGteOk() (*int32, bool)`

GetTotalMemoryGteOk returns a tuple with the TotalMemoryGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTotalMemoryGte(v int32)`

SetTotalMemoryGte sets TotalMemoryGte field to given value.

### HasTotalMemoryGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTotalMemoryGte() bool`

HasTotalMemoryGte returns a boolean if a field has been set.

### GetMachineTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMachineTypesNin() []string`

GetMachineTypesNin returns the MachineTypesNin field if non-nil, zero value otherwise.

### GetMachineTypesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMachineTypesNinOk() (*[]string, bool)`

GetMachineTypesNinOk returns a tuple with the MachineTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetMachineTypesNin(v []string)`

SetMachineTypesNin sets MachineTypesNin field to given value.

### HasMachineTypesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasMachineTypesNin() bool`

HasMachineTypesNin returns a boolean if a field has been set.

### GetLastActiveDateBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateBetween() string`

GetLastActiveDateBetween returns the LastActiveDateBetween field if non-nil, zero value otherwise.

### GetLastActiveDateBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateBetweenOk() (*string, bool)`

GetLastActiveDateBetweenOk returns a tuple with the LastActiveDateBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastActiveDateBetween(v string)`

SetLastActiveDateBetween sets LastActiveDateBetween field to given value.

### HasLastActiveDateBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastActiveDateBetween() bool`

HasLastActiveDateBetween returns a boolean if a field has been set.

### GetNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfaceGatewayMacAddressContains() []string`

GetNetworkInterfaceGatewayMacAddressContains returns the NetworkInterfaceGatewayMacAddressContains field if non-nil, zero value otherwise.

### GetNetworkInterfaceGatewayMacAddressContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfaceGatewayMacAddressContainsOk() (*[]string, bool)`

GetNetworkInterfaceGatewayMacAddressContainsOk returns a tuple with the NetworkInterfaceGatewayMacAddressContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkInterfaceGatewayMacAddressContains(v []string)`

SetNetworkInterfaceGatewayMacAddressContains sets NetworkInterfaceGatewayMacAddressContains field to given value.

### HasNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkInterfaceGatewayMacAddressContains() bool`

HasNetworkInterfaceGatewayMacAddressContains returns a boolean if a field has been set.

### GetAwsRoleContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsRoleContains() []string`

GetAwsRoleContains returns the AwsRoleContains field if non-nil, zero value otherwise.

### GetAwsRoleContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAwsRoleContainsOk() (*[]string, bool)`

GetAwsRoleContainsOk returns a tuple with the AwsRoleContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAwsRoleContains(v []string)`

SetAwsRoleContains sets AwsRoleContains field to given value.

### HasAwsRoleContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAwsRoleContains() bool`

HasAwsRoleContains returns a boolean if a field has been set.

### GetRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRemoteProfilingStatesNin() []string`

GetRemoteProfilingStatesNin returns the RemoteProfilingStatesNin field if non-nil, zero value otherwise.

### GetRemoteProfilingStatesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRemoteProfilingStatesNinOk() (*[]string, bool)`

GetRemoteProfilingStatesNinOk returns a tuple with the RemoteProfilingStatesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRemoteProfilingStatesNin(v []string)`

SetRemoteProfilingStatesNin sets RemoteProfilingStatesNin field to given value.

### HasRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRemoteProfilingStatesNin() bool`

HasRemoteProfilingStatesNin returns a boolean if a field has been set.

### GetLastActiveDateLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateLte() time.Time`

GetLastActiveDateLte returns the LastActiveDateLte field if non-nil, zero value otherwise.

### GetLastActiveDateLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateLteOk() (*time.Time, bool)`

GetLastActiveDateLteOk returns a tuple with the LastActiveDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastActiveDateLte(v time.Time)`

SetLastActiveDateLte sets LastActiveDateLte field to given value.

### HasLastActiveDateLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastActiveDateLte() bool`

HasLastActiveDateLte returns a boolean if a field has been set.

### GetCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtLte() time.Time`

GetCreatedAtLte returns the CreatedAtLte field if non-nil, zero value otherwise.

### GetCreatedAtLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtLteOk() (*time.Time, bool)`

GetCreatedAtLteOk returns a tuple with the CreatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCreatedAtLte(v time.Time)`

SetCreatedAtLte sets CreatedAtLte field to given value.

### HasCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCreatedAtLte() bool`

HasCreatedAtLte returns a boolean if a field has been set.

### GetConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetConsoleMigrationStatusesNin() []string`

GetConsoleMigrationStatusesNin returns the ConsoleMigrationStatusesNin field if non-nil, zero value otherwise.

### GetConsoleMigrationStatusesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetConsoleMigrationStatusesNinOk() (*[]string, bool)`

GetConsoleMigrationStatusesNinOk returns a tuple with the ConsoleMigrationStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetConsoleMigrationStatusesNin(v []string)`

SetConsoleMigrationStatusesNin sets ConsoleMigrationStatusesNin field to given value.

### HasConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasConsoleMigrationStatusesNin() bool`

HasConsoleMigrationStatusesNin returns a boolean if a field has been set.

### GetFilteredSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilteredSiteIds() []string`

GetFilteredSiteIds returns the FilteredSiteIds field if non-nil, zero value otherwise.

### GetFilteredSiteIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilteredSiteIdsOk() (*[]string, bool)`

GetFilteredSiteIdsOk returns a tuple with the FilteredSiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetFilteredSiteIds(v []string)`

SetFilteredSiteIds sets FilteredSiteIds field to given value.

### HasFilteredSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasFilteredSiteIds() bool`

HasFilteredSiteIds returns a boolean if a field has been set.

### GetAdComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerNameContains() []string`

GetAdComputerNameContains returns the AdComputerNameContains field if non-nil, zero value otherwise.

### GetAdComputerNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerNameContainsOk() (*[]string, bool)`

GetAdComputerNameContainsOk returns a tuple with the AdComputerNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdComputerNameContains(v []string)`

SetAdComputerNameContains sets AdComputerNameContains field to given value.

### HasAdComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdComputerNameContains() bool`

HasAdComputerNameContains returns a boolean if a field has been set.

### GetRegisteredAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtGt() time.Time`

GetRegisteredAtGt returns the RegisteredAtGt field if non-nil, zero value otherwise.

### GetRegisteredAtGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtGtOk() (*time.Time, bool)`

GetRegisteredAtGtOk returns a tuple with the RegisteredAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRegisteredAtGt(v time.Time)`

SetRegisteredAtGt sets RegisteredAtGt field to given value.

### HasRegisteredAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRegisteredAtGt() bool`

HasRegisteredAtGt returns a boolean if a field has been set.

### GetCloudTagsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudTagsContains() []string`

GetCloudTagsContains returns the CloudTagsContains field if non-nil, zero value otherwise.

### GetCloudTagsContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudTagsContainsOk() (*[]string, bool)`

GetCloudTagsContainsOk returns a tuple with the CloudTagsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTagsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudTagsContains(v []string)`

SetCloudTagsContains sets CloudTagsContains field to given value.

### HasCloudTagsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudTagsContains() bool`

HasCloudTagsContains returns a boolean if a field has been set.

### GetAdUserQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserQueryContains() []string`

GetAdUserQueryContains returns the AdUserQueryContains field if non-nil, zero value otherwise.

### GetAdUserQueryContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserQueryContainsOk() (*[]string, bool)`

GetAdUserQueryContainsOk returns a tuple with the AdUserQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdUserQueryContains(v []string)`

SetAdUserQueryContains sets AdUserQueryContains field to given value.

### HasAdUserQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdUserQueryContains() bool`

HasAdUserQueryContains returns a boolean if a field has been set.

### GetThreatContentHash

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatContentHash() string`

GetThreatContentHash returns the ThreatContentHash field if non-nil, zero value otherwise.

### GetThreatContentHashOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatContentHashOk() (*string, bool)`

GetThreatContentHashOk returns a tuple with the ThreatContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatContentHash

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatContentHash(v string)`

SetThreatContentHash sets ThreatContentHash field to given value.

### HasThreatContentHash

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatContentHash() bool`

HasThreatContentHash returns a boolean if a field has been set.

### GetScanStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatuses() []string`

GetScanStatuses returns the ScanStatuses field if non-nil, zero value otherwise.

### GetScanStatusesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatusesOk() (*[]string, bool)`

GetScanStatusesOk returns a tuple with the ScanStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetScanStatuses(v []string)`

SetScanStatuses sets ScanStatuses field to given value.

### HasScanStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasScanStatuses() bool`

HasScanStatuses returns a boolean if a field has been set.

### GetOperationalStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOperationalStatesNin() []string`

GetOperationalStatesNin returns the OperationalStatesNin field if non-nil, zero value otherwise.

### GetOperationalStatesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOperationalStatesNinOk() (*[]string, bool)`

GetOperationalStatesNinOk returns a tuple with the OperationalStatesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOperationalStatesNin(v []string)`

SetOperationalStatesNin sets OperationalStatesNin field to given value.

### HasOperationalStatesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOperationalStatesNin() bool`

HasOperationalStatesNin returns a boolean if a field has been set.

### GetThreatRebootRequired

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatRebootRequired() []bool`

GetThreatRebootRequired returns the ThreatRebootRequired field if non-nil, zero value otherwise.

### GetThreatRebootRequiredOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatRebootRequiredOk() (*[]bool, bool)`

GetThreatRebootRequiredOk returns a tuple with the ThreatRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatRebootRequired

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatRebootRequired(v []bool)`

SetThreatRebootRequired sets ThreatRebootRequired field to given value.

### HasThreatRebootRequired

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatRebootRequired() bool`

HasThreatRebootRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTotalMemoryLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryLt() int32`

GetTotalMemoryLt returns the TotalMemoryLt field if non-nil, zero value otherwise.

### GetTotalMemoryLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryLtOk() (*int32, bool)`

GetTotalMemoryLtOk returns a tuple with the TotalMemoryLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTotalMemoryLt(v int32)`

SetTotalMemoryLt sets TotalMemoryLt field to given value.

### HasTotalMemoryLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTotalMemoryLt() bool`

HasTotalMemoryLt returns a boolean if a field has been set.

### GetAdComputerQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerQueryContains() []string`

GetAdComputerQueryContains returns the AdComputerQueryContains field if non-nil, zero value otherwise.

### GetAdComputerQueryContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerQueryContainsOk() (*[]string, bool)`

GetAdComputerQueryContainsOk returns a tuple with the AdComputerQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdComputerQueryContains(v []string)`

SetAdComputerQueryContains sets AdComputerQueryContains field to given value.

### HasAdComputerQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdComputerQueryContains() bool`

HasAdComputerQueryContains returns a boolean if a field has been set.

### GetUserActionsNeeded

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUserActionsNeeded() []string`

GetUserActionsNeeded returns the UserActionsNeeded field if non-nil, zero value otherwise.

### GetUserActionsNeededOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUserActionsNeededOk() (*[]string, bool)`

GetUserActionsNeededOk returns a tuple with the UserActionsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsNeeded

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUserActionsNeeded(v []string)`

SetUserActionsNeeded sets UserActionsNeeded field to given value.

### HasUserActionsNeeded

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUserActionsNeeded() bool`

HasUserActionsNeeded returns a boolean if a field has been set.

### GetUpdatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtBetween() string`

GetUpdatedAtBetween returns the UpdatedAtBetween field if non-nil, zero value otherwise.

### GetUpdatedAtBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtBetweenOk() (*string, bool)`

GetUpdatedAtBetweenOk returns a tuple with the UpdatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUpdatedAtBetween(v string)`

SetUpdatedAtBetween sets UpdatedAtBetween field to given value.

### HasUpdatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUpdatedAtBetween() bool`

HasUpdatedAtBetween returns a boolean if a field has been set.

### GetCpuCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountGte() int32`

GetCpuCountGte returns the CpuCountGte field if non-nil, zero value otherwise.

### GetCpuCountGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountGteOk() (*int32, bool)`

GetCpuCountGteOk returns a tuple with the CpuCountGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuCountGte(v int32)`

SetCpuCountGte sets CpuCountGte field to given value.

### HasCpuCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuCountGte() bool`

HasCpuCountGte returns a boolean if a field has been set.

### GetUuids

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuids() []string`

GetUuids returns the Uuids field if non-nil, zero value otherwise.

### GetUuidsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuidsOk() (*[]string, bool)`

GetUuidsOk returns a tuple with the Uuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuids

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUuids(v []string)`

SetUuids sets Uuids field to given value.

### HasUuids

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUuids() bool`

HasUuids returns a boolean if a field has been set.

### GetSerialNumberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetSerialNumberContains() []string`

GetSerialNumberContains returns the SerialNumberContains field if non-nil, zero value otherwise.

### GetSerialNumberContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetSerialNumberContainsOk() (*[]string, bool)`

GetSerialNumberContainsOk returns a tuple with the SerialNumberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetSerialNumberContains(v []string)`

SetSerialNumberContains sets SerialNumberContains field to given value.

### HasSerialNumberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasSerialNumberContains() bool`

HasSerialNumberContains returns a boolean if a field has been set.

### GetRegisteredAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtBetween() string`

GetRegisteredAtBetween returns the RegisteredAtBetween field if non-nil, zero value otherwise.

### GetRegisteredAtBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtBetweenOk() (*string, bool)`

GetRegisteredAtBetweenOk returns a tuple with the RegisteredAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRegisteredAtBetween(v string)`

SetRegisteredAtBetween sets RegisteredAtBetween field to given value.

### HasRegisteredAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRegisteredAtBetween() bool`

HasRegisteredAtBetween returns a boolean if a field has been set.

### GetScanStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetExternalIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetExternalIdContains() []string`

GetExternalIdContains returns the ExternalIdContains field if non-nil, zero value otherwise.

### GetExternalIdContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetExternalIdContainsOk() (*[]string, bool)`

GetExternalIdContainsOk returns a tuple with the ExternalIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetExternalIdContains(v []string)`

SetExternalIdContains sets ExternalIdContains field to given value.

### HasExternalIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasExternalIdContains() bool`

HasExternalIdContains returns a boolean if a field has been set.

### GetTotalMemoryLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryLte() int32`

GetTotalMemoryLte returns the TotalMemoryLte field if non-nil, zero value otherwise.

### GetTotalMemoryLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryLteOk() (*int32, bool)`

GetTotalMemoryLteOk returns a tuple with the TotalMemoryLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTotalMemoryLte(v int32)`

SetTotalMemoryLte sets TotalMemoryLte field to given value.

### HasTotalMemoryLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTotalMemoryLte() bool`

HasTotalMemoryLte returns a boolean if a field has been set.

### GetLocationIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetCloudNetworkContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudNetworkContains() []string`

GetCloudNetworkContains returns the CloudNetworkContains field if non-nil, zero value otherwise.

### GetCloudNetworkContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudNetworkContainsOk() (*[]string, bool)`

GetCloudNetworkContainsOk returns a tuple with the CloudNetworkContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudNetworkContains(v []string)`

SetCloudNetworkContains sets CloudNetworkContains field to given value.

### HasCloudNetworkContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudNetworkContains() bool`

HasCloudNetworkContains returns a boolean if a field has been set.

### GetAzureResourceGroupContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAzureResourceGroupContains() []string`

GetAzureResourceGroupContains returns the AzureResourceGroupContains field if non-nil, zero value otherwise.

### GetAzureResourceGroupContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAzureResourceGroupContainsOk() (*[]string, bool)`

GetAzureResourceGroupContainsOk returns a tuple with the AzureResourceGroupContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroupContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAzureResourceGroupContains(v []string)`

SetAzureResourceGroupContains sets AzureResourceGroupContains field to given value.

### HasAzureResourceGroupContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAzureResourceGroupContains() bool`

HasAzureResourceGroupContains returns a boolean if a field has been set.

### GetCoreCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountBetween() string`

GetCoreCountBetween returns the CoreCountBetween field if non-nil, zero value otherwise.

### GetCoreCountBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountBetweenOk() (*string, bool)`

GetCoreCountBetweenOk returns a tuple with the CoreCountBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCoreCountBetween(v string)`

SetCoreCountBetween sets CoreCountBetween field to given value.

### HasCoreCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCoreCountBetween() bool`

HasCoreCountBetween returns a boolean if a field has been set.

### GetIsUninstalled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsUninstalled() []bool`

GetIsUninstalled returns the IsUninstalled field if non-nil, zero value otherwise.

### GetIsUninstalledOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetIsUninstalledOk() (*[]bool, bool)`

GetIsUninstalledOk returns a tuple with the IsUninstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUninstalled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetIsUninstalled(v []bool)`

SetIsUninstalled sets IsUninstalled field to given value.

### HasIsUninstalled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasIsUninstalled() bool`

HasIsUninstalled returns a boolean if a field has been set.

### GetFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetCpuIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuIdContains() []string`

GetCpuIdContains returns the CpuIdContains field if non-nil, zero value otherwise.

### GetCpuIdContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuIdContainsOk() (*[]string, bool)`

GetCpuIdContainsOk returns a tuple with the CpuIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuIdContains(v []string)`

SetCpuIdContains sets CpuIdContains field to given value.

### HasCpuIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuIdContains() bool`

HasCpuIdContains returns a boolean if a field has been set.

### GetK8sTypeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sTypeContains() []string`

GetK8sTypeContains returns the K8sTypeContains field if non-nil, zero value otherwise.

### GetK8sTypeContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sTypeContainsOk() (*[]string, bool)`

GetK8sTypeContainsOk returns a tuple with the K8sTypeContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTypeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetK8sTypeContains(v []string)`

SetK8sTypeContains sets K8sTypeContains field to given value.

### HasK8sTypeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasK8sTypeContains() bool`

HasK8sTypeContains returns a boolean if a field has been set.

### GetCloudProviderNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudProviderNin() []string`

GetCloudProviderNin returns the CloudProviderNin field if non-nil, zero value otherwise.

### GetCloudProviderNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudProviderNinOk() (*[]string, bool)`

GetCloudProviderNinOk returns a tuple with the CloudProviderNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudProviderNin(v []string)`

SetCloudProviderNin sets CloudProviderNin field to given value.

### HasCloudProviderNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudProviderNin() bool`

HasCloudProviderNin returns a boolean if a field has been set.

### GetMitigationMode

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMitigationMode() string`

GetMitigationMode returns the MitigationMode field if non-nil, zero value otherwise.

### GetMitigationModeOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMitigationModeOk() (*string, bool)`

GetMitigationModeOk returns a tuple with the MitigationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationMode

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetMitigationMode(v string)`

SetMitigationMode sets MitigationMode field to given value.

### HasMitigationMode

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasMitigationMode() bool`

HasMitigationMode returns a boolean if a field has been set.

### GetCloudAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudAccountContains() []string`

GetCloudAccountContains returns the CloudAccountContains field if non-nil, zero value otherwise.

### GetCloudAccountContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudAccountContainsOk() (*[]string, bool)`

GetCloudAccountContainsOk returns a tuple with the CloudAccountContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudAccountContains(v []string)`

SetCloudAccountContains sets CloudAccountContains field to given value.

### HasCloudAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudAccountContains() bool`

HasCloudAccountContains returns a boolean if a field has been set.

### GetAdComputerMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerMemberContains() []string`

GetAdComputerMemberContains returns the AdComputerMemberContains field if non-nil, zero value otherwise.

### GetAdComputerMemberContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdComputerMemberContainsOk() (*[]string, bool)`

GetAdComputerMemberContainsOk returns a tuple with the AdComputerMemberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdComputerMemberContains(v []string)`

SetAdComputerMemberContains sets AdComputerMemberContains field to given value.

### HasAdComputerMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdComputerMemberContains() bool`

HasAdComputerMemberContains returns a boolean if a field has been set.

### GetConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetConsoleMigrationStatuses() []string`

GetConsoleMigrationStatuses returns the ConsoleMigrationStatuses field if non-nil, zero value otherwise.

### GetConsoleMigrationStatusesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetConsoleMigrationStatusesOk() (*[]string, bool)`

GetConsoleMigrationStatusesOk returns a tuple with the ConsoleMigrationStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetConsoleMigrationStatuses(v []string)`

SetConsoleMigrationStatuses sets ConsoleMigrationStatuses field to given value.

### HasConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasConsoleMigrationStatuses() bool`

HasConsoleMigrationStatuses returns a boolean if a field has been set.

### GetLastActiveDateGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateGt() time.Time`

GetLastActiveDateGt returns the LastActiveDateGt field if non-nil, zero value otherwise.

### GetLastActiveDateGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateGtOk() (*time.Time, bool)`

GetLastActiveDateGtOk returns a tuple with the LastActiveDateGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastActiveDateGt(v time.Time)`

SetLastActiveDateGt sets LastActiveDateGt field to given value.

### HasLastActiveDateGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastActiveDateGt() bool`

HasLastActiveDateGt returns a boolean if a field has been set.

### GetOsArch

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### GetAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentContentUpdateIdContains() []string`

GetAgentContentUpdateIdContains returns the AgentContentUpdateIdContains field if non-nil, zero value otherwise.

### GetAgentContentUpdateIdContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentContentUpdateIdContainsOk() (*[]string, bool)`

GetAgentContentUpdateIdContainsOk returns a tuple with the AgentContentUpdateIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAgentContentUpdateIdContains(v []string)`

SetAgentContentUpdateIdContains sets AgentContentUpdateIdContains field to given value.

### HasAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAgentContentUpdateIdContains() bool`

HasAgentContentUpdateIdContains returns a boolean if a field has been set.

### GetRegisteredAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtGte() time.Time`

GetRegisteredAtGte returns the RegisteredAtGte field if non-nil, zero value otherwise.

### GetRegisteredAtGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtGteOk() (*time.Time, bool)`

GetRegisteredAtGteOk returns a tuple with the RegisteredAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRegisteredAtGte(v time.Time)`

SetRegisteredAtGte sets RegisteredAtGte field to given value.

### HasRegisteredAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRegisteredAtGte() bool`

HasRegisteredAtGte returns a boolean if a field has been set.

### GetCoreCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountGt() int32`

GetCoreCountGt returns the CoreCountGt field if non-nil, zero value otherwise.

### GetCoreCountGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountGtOk() (*int32, bool)`

GetCoreCountGtOk returns a tuple with the CoreCountGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCoreCountGt(v int32)`

SetCoreCountGt sets CoreCountGt field to given value.

### HasCoreCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCoreCountGt() bool`

HasCoreCountGt returns a boolean if a field has been set.

### GetCpuCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountLte() int32`

GetCpuCountLte returns the CpuCountLte field if non-nil, zero value otherwise.

### GetCpuCountLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountLteOk() (*int32, bool)`

GetCpuCountLteOk returns a tuple with the CpuCountLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuCountLte(v int32)`

SetCpuCountLte sets CpuCountLte field to given value.

### HasCpuCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuCountLte() bool`

HasCpuCountLte returns a boolean if a field has been set.

### GetAgentPodNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentPodNameContains() []string`

GetAgentPodNameContains returns the AgentPodNameContains field if non-nil, zero value otherwise.

### GetAgentPodNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentPodNameContainsOk() (*[]string, bool)`

GetAgentPodNameContainsOk returns a tuple with the AgentPodNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPodNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAgentPodNameContains(v []string)`

SetAgentPodNameContains sets AgentPodNameContains field to given value.

### HasAgentPodNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAgentPodNameContains() bool`

HasAgentPodNameContains returns a boolean if a field has been set.

### GetThreatCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtLt() time.Time`

GetThreatCreatedAtLt returns the ThreatCreatedAtLt field if non-nil, zero value otherwise.

### GetThreatCreatedAtLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtLtOk() (*time.Time, bool)`

GetThreatCreatedAtLtOk returns a tuple with the ThreatCreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatCreatedAtLt(v time.Time)`

SetThreatCreatedAtLt sets ThreatCreatedAtLt field to given value.

### HasThreatCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatCreatedAtLt() bool`

HasThreatCreatedAtLt returns a boolean if a field has been set.

### GetDecommissionedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtBetween() string`

GetDecommissionedAtBetween returns the DecommissionedAtBetween field if non-nil, zero value otherwise.

### GetDecommissionedAtBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtBetweenOk() (*string, bool)`

GetDecommissionedAtBetweenOk returns a tuple with the DecommissionedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDecommissionedAtBetween(v string)`

SetDecommissionedAtBetween sets DecommissionedAtBetween field to given value.

### HasDecommissionedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDecommissionedAtBetween() bool`

HasDecommissionedAtBetween returns a boolean if a field has been set.

### GetCoreCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountGte() int32`

GetCoreCountGte returns the CoreCountGte field if non-nil, zero value otherwise.

### GetCoreCountGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountGteOk() (*int32, bool)`

GetCoreCountGteOk returns a tuple with the CoreCountGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCoreCountGte(v int32)`

SetCoreCountGte sets CoreCountGte field to given value.

### HasCoreCountGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCoreCountGte() bool`

HasCoreCountGte returns a boolean if a field has been set.

### GetClusterNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetClusterNameContains() []string`

GetClusterNameContains returns the ClusterNameContains field if non-nil, zero value otherwise.

### GetClusterNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetClusterNameContainsOk() (*[]string, bool)`

GetClusterNameContainsOk returns a tuple with the ClusterNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetClusterNameContains(v []string)`

SetClusterNameContains sets ClusterNameContains field to given value.

### HasClusterNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasClusterNameContains() bool`

HasClusterNameContains returns a boolean if a field has been set.

### GetCpuCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountGt() int32`

GetCpuCountGt returns the CpuCountGt field if non-nil, zero value otherwise.

### GetCpuCountGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountGtOk() (*int32, bool)`

GetCpuCountGtOk returns a tuple with the CpuCountGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuCountGt(v int32)`

SetCpuCountGt sets CpuCountGt field to given value.

### HasCpuCountGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuCountGt() bool`

HasCpuCountGt returns a boolean if a field has been set.

### GetAdQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdQueryContains() []string`

GetAdQueryContains returns the AdQueryContains field if non-nil, zero value otherwise.

### GetAdQueryContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdQueryContainsOk() (*[]string, bool)`

GetAdQueryContainsOk returns a tuple with the AdQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdQueryContains(v []string)`

SetAdQueryContains sets AdQueryContains field to given value.

### HasAdQueryContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdQueryContains() bool`

HasAdQueryContains returns a boolean if a field has been set.

### GetRangerVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerVersions() []string`

GetRangerVersions returns the RangerVersions field if non-nil, zero value otherwise.

### GetRangerVersionsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerVersionsOk() (*[]string, bool)`

GetRangerVersionsOk returns a tuple with the RangerVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRangerVersions(v []string)`

SetRangerVersions sets RangerVersions field to given value.

### HasRangerVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRangerVersions() bool`

HasRangerVersions returns a boolean if a field has been set.

### GetUserActionsNeededNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUserActionsNeededNin() []string`

GetUserActionsNeededNin returns the UserActionsNeededNin field if non-nil, zero value otherwise.

### GetUserActionsNeededNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUserActionsNeededNinOk() (*[]string, bool)`

GetUserActionsNeededNinOk returns a tuple with the UserActionsNeededNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsNeededNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUserActionsNeededNin(v []string)`

SetUserActionsNeededNin sets UserActionsNeededNin field to given value.

### HasUserActionsNeededNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUserActionsNeededNin() bool`

HasUserActionsNeededNin returns a boolean if a field has been set.

### GetThreatCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtGt() time.Time`

GetThreatCreatedAtGt returns the ThreatCreatedAtGt field if non-nil, zero value otherwise.

### GetThreatCreatedAtGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtGtOk() (*time.Time, bool)`

GetThreatCreatedAtGtOk returns a tuple with the ThreatCreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatCreatedAtGt(v time.Time)`

SetThreatCreatedAtGt sets ThreatCreatedAtGt field to given value.

### HasThreatCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatCreatedAtGt() bool`

HasThreatCreatedAtGt returns a boolean if a field has been set.

### GetCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudInstanceSizeContains() []string`

GetCloudInstanceSizeContains returns the CloudInstanceSizeContains field if non-nil, zero value otherwise.

### GetCloudInstanceSizeContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudInstanceSizeContainsOk() (*[]string, bool)`

GetCloudInstanceSizeContainsOk returns a tuple with the CloudInstanceSizeContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudInstanceSizeContains(v []string)`

SetCloudInstanceSizeContains sets CloudInstanceSizeContains field to given value.

### HasCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudInstanceSizeContains() bool`

HasCloudInstanceSizeContains returns a boolean if a field has been set.

### GetAdQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdQuery() string`

GetAdQuery returns the AdQuery field if non-nil, zero value otherwise.

### GetAdQueryOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdQueryOk() (*string, bool)`

GetAdQueryOk returns a tuple with the AdQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdQuery(v string)`

SetAdQuery sets AdQuery field to given value.

### HasAdQuery

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdQuery() bool`

HasAdQuery returns a boolean if a field has been set.

### GetExternalIpContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetExternalIpContains() []string`

GetExternalIpContains returns the ExternalIpContains field if non-nil, zero value otherwise.

### GetExternalIpContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetExternalIpContainsOk() (*[]string, bool)`

GetExternalIpContainsOk returns a tuple with the ExternalIpContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetExternalIpContains(v []string)`

SetExternalIpContains sets ExternalIpContains field to given value.

### HasExternalIpContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasExternalIpContains() bool`

HasExternalIpContains returns a boolean if a field has been set.

### GetLocationEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationEnabled() []bool`

GetLocationEnabled returns the LocationEnabled field if non-nil, zero value otherwise.

### GetLocationEnabledOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationEnabledOk() (*[]bool, bool)`

GetLocationEnabledOk returns a tuple with the LocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLocationEnabled(v []bool)`

SetLocationEnabled sets LocationEnabled field to given value.

### HasLocationEnabled

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLocationEnabled() bool`

HasLocationEnabled returns a boolean if a field has been set.

### GetTotalMemoryGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryGt() int32`

GetTotalMemoryGt returns the TotalMemoryGt field if non-nil, zero value otherwise.

### GetTotalMemoryGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryGtOk() (*int32, bool)`

GetTotalMemoryGtOk returns a tuple with the TotalMemoryGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTotalMemoryGt(v int32)`

SetTotalMemoryGt sets TotalMemoryGt field to given value.

### HasTotalMemoryGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTotalMemoryGt() bool`

HasTotalMemoryGt returns a boolean if a field has been set.

### GetGatewayIp

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetRegisteredAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtLt() time.Time`

GetRegisteredAtLt returns the RegisteredAtLt field if non-nil, zero value otherwise.

### GetRegisteredAtLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRegisteredAtLtOk() (*time.Time, bool)`

GetRegisteredAtLtOk returns a tuple with the RegisteredAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRegisteredAtLt(v time.Time)`

SetRegisteredAtLt sets RegisteredAtLt field to given value.

### HasRegisteredAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRegisteredAtLt() bool`

HasRegisteredAtLt returns a boolean if a field has been set.

### GetAgentVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentVersions() []string`

GetAgentVersions returns the AgentVersions field if non-nil, zero value otherwise.

### GetAgentVersionsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentVersionsOk() (*[]string, bool)`

GetAgentVersionsOk returns a tuple with the AgentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAgentVersions(v []string)`

SetAgentVersions sets AgentVersions field to given value.

### HasAgentVersions

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAgentVersions() bool`

HasAgentVersions returns a boolean if a field has been set.

### GetInstallerTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInstallerTypes() []string`

GetInstallerTypes returns the InstallerTypes field if non-nil, zero value otherwise.

### GetInstallerTypesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInstallerTypesOk() (*[]string, bool)`

GetInstallerTypesOk returns a tuple with the InstallerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetInstallerTypes(v []string)`

SetInstallerTypes sets InstallerTypes field to given value.

### HasInstallerTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasInstallerTypes() bool`

HasInstallerTypes returns a boolean if a field has been set.

### GetCoreCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountLte() int32`

GetCoreCountLte returns the CoreCountLte field if non-nil, zero value otherwise.

### GetCoreCountLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountLteOk() (*int32, bool)`

GetCoreCountLteOk returns a tuple with the CoreCountLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCoreCountLte(v int32)`

SetCoreCountLte sets CoreCountLte field to given value.

### HasCoreCountLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCoreCountLte() bool`

HasCoreCountLte returns a boolean if a field has been set.

### GetTotalMemoryBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryBetween() string`

GetTotalMemoryBetween returns the TotalMemoryBetween field if non-nil, zero value otherwise.

### GetTotalMemoryBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTotalMemoryBetweenOk() (*string, bool)`

GetTotalMemoryBetweenOk returns a tuple with the TotalMemoryBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTotalMemoryBetween(v string)`

SetTotalMemoryBetween sets TotalMemoryBetween field to given value.

### HasTotalMemoryBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTotalMemoryBetween() bool`

HasTotalMemoryBetween returns a boolean if a field has been set.

### GetLastActiveDateGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateGte() time.Time`

GetLastActiveDateGte returns the LastActiveDateGte field if non-nil, zero value otherwise.

### GetLastActiveDateGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateGteOk() (*time.Time, bool)`

GetLastActiveDateGteOk returns a tuple with the LastActiveDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastActiveDateGte(v time.Time)`

SetLastActiveDateGte sets LastActiveDateGte field to given value.

### HasLastActiveDateGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastActiveDateGte() bool`

HasLastActiveDateGte returns a boolean if a field has been set.

### GetComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerNameContains() []string`

GetComputerNameContains returns the ComputerNameContains field if non-nil, zero value otherwise.

### GetComputerNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerNameContainsOk() (*[]string, bool)`

GetComputerNameContainsOk returns a tuple with the ComputerNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetComputerNameContains(v []string)`

SetComputerNameContains sets ComputerNameContains field to given value.

### HasComputerNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasComputerNameContains() bool`

HasComputerNameContains returns a boolean if a field has been set.

### GetLocationIdsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationIdsNin() []string`

GetLocationIdsNin returns the LocationIdsNin field if non-nil, zero value otherwise.

### GetLocationIdsNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLocationIdsNinOk() (*[]string, bool)`

GetLocationIdsNinOk returns a tuple with the LocationIdsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIdsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLocationIdsNin(v []string)`

SetLocationIdsNin sets LocationIdsNin field to given value.

### HasLocationIdsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLocationIdsNin() bool`

HasLocationIdsNin returns a boolean if a field has been set.

### GetThreatCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtLte() time.Time`

GetThreatCreatedAtLte returns the ThreatCreatedAtLte field if non-nil, zero value otherwise.

### GetThreatCreatedAtLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtLteOk() (*time.Time, bool)`

GetThreatCreatedAtLteOk returns a tuple with the ThreatCreatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatCreatedAtLte(v time.Time)`

SetThreatCreatedAtLte sets ThreatCreatedAtLte field to given value.

### HasThreatCreatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatCreatedAtLte() bool`

HasThreatCreatedAtLte returns a boolean if a field has been set.

### GetHasTags

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetHasTags() bool`

GetHasTags returns the HasTags field if non-nil, zero value otherwise.

### GetHasTagsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetHasTagsOk() (*bool, bool)`

GetHasTagsOk returns a tuple with the HasTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTags

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetHasTags(v bool)`

SetHasTags sets HasTags field to given value.

### HasHasTags

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasHasTags() bool`

HasHasTags returns a boolean if a field has been set.

### GetUpdatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtGte() time.Time`

GetUpdatedAtGte returns the UpdatedAtGte field if non-nil, zero value otherwise.

### GetUpdatedAtGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtGteOk() (*time.Time, bool)`

GetUpdatedAtGteOk returns a tuple with the UpdatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUpdatedAtGte(v time.Time)`

SetUpdatedAtGte sets UpdatedAtGte field to given value.

### HasUpdatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUpdatedAtGte() bool`

HasUpdatedAtGte returns a boolean if a field has been set.

### GetThreatResolved

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatResolved() bool`

GetThreatResolved returns the ThreatResolved field if non-nil, zero value otherwise.

### GetThreatResolvedOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatResolvedOk() (*bool, bool)`

GetThreatResolvedOk returns a tuple with the ThreatResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatResolved

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatResolved(v bool)`

SetThreatResolved sets ThreatResolved field to given value.

### HasThreatResolved

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatResolved() bool`

HasThreatResolved returns a boolean if a field has been set.

### GetDecommissionedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtGt() time.Time`

GetDecommissionedAtGt returns the DecommissionedAtGt field if non-nil, zero value otherwise.

### GetDecommissionedAtGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtGtOk() (*time.Time, bool)`

GetDecommissionedAtGtOk returns a tuple with the DecommissionedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDecommissionedAtGt(v time.Time)`

SetDecommissionedAtGt sets DecommissionedAtGt field to given value.

### HasDecommissionedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDecommissionedAtGt() bool`

HasDecommissionedAtGt returns a boolean if a field has been set.

### GetThreatHidden

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatHidden() bool`

GetThreatHidden returns the ThreatHidden field if non-nil, zero value otherwise.

### GetThreatHiddenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatHiddenOk() (*bool, bool)`

GetThreatHiddenOk returns a tuple with the ThreatHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatHidden

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatHidden(v bool)`

SetThreatHidden sets ThreatHidden field to given value.

### HasThreatHidden

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatHidden() bool`

HasThreatHidden returns a boolean if a field has been set.

### GetInfected

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInfected() bool`

GetInfected returns the Infected field if non-nil, zero value otherwise.

### GetInfectedOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetInfectedOk() (*bool, bool)`

GetInfectedOk returns a tuple with the Infected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfected

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetInfected(v bool)`

SetInfected sets Infected field to given value.

### HasInfected

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasInfected() bool`

HasInfected returns a boolean if a field has been set.

### GetUuidContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuidContains() []string`

GetUuidContains returns the UuidContains field if non-nil, zero value otherwise.

### GetUuidContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuidContainsOk() (*[]string, bool)`

GetUuidContainsOk returns a tuple with the UuidContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUuidContains(v []string)`

SetUuidContains sets UuidContains field to given value.

### HasUuidContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUuidContains() bool`

HasUuidContains returns a boolean if a field has been set.

### GetNetworkStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkStatusesNin() []string`

GetNetworkStatusesNin returns the NetworkStatusesNin field if non-nil, zero value otherwise.

### GetNetworkStatusesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkStatusesNinOk() (*[]string, bool)`

GetNetworkStatusesNinOk returns a tuple with the NetworkStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkStatusesNin(v []string)`

SetNetworkStatusesNin sets NetworkStatusesNin field to given value.

### HasNetworkStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkStatusesNin() bool`

HasNetworkStatusesNin returns a boolean if a field has been set.

### GetCloudImageContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudImageContains() []string`

GetCloudImageContains returns the CloudImageContains field if non-nil, zero value otherwise.

### GetCloudImageContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudImageContainsOk() (*[]string, bool)`

GetCloudImageContainsOk returns a tuple with the CloudImageContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudImageContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudImageContains(v []string)`

SetCloudImageContains sets CloudImageContains field to given value.

### HasCloudImageContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudImageContains() bool`

HasCloudImageContains returns a boolean if a field has been set.

### GetSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetRangerStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatus() string`

GetRangerStatus returns the RangerStatus field if non-nil, zero value otherwise.

### GetRangerStatusOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatusOk() (*string, bool)`

GetRangerStatusOk returns a tuple with the RangerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRangerStatus(v string)`

SetRangerStatus sets RangerStatus field to given value.

### HasRangerStatus

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRangerStatus() bool`

HasRangerStatus returns a boolean if a field has been set.

### GetDomainsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDomainsNin() []string`

GetDomainsNin returns the DomainsNin field if non-nil, zero value otherwise.

### GetDomainsNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDomainsNinOk() (*[]string, bool)`

GetDomainsNinOk returns a tuple with the DomainsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDomainsNin(v []string)`

SetDomainsNin sets DomainsNin field to given value.

### HasDomainsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDomainsNin() bool`

HasDomainsNin returns a boolean if a field has been set.

### GetThreatCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtGte() time.Time`

GetThreatCreatedAtGte returns the ThreatCreatedAtGte field if non-nil, zero value otherwise.

### GetThreatCreatedAtGteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtGteOk() (*time.Time, bool)`

GetThreatCreatedAtGteOk returns a tuple with the ThreatCreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatCreatedAtGte(v time.Time)`

SetThreatCreatedAtGte sets ThreatCreatedAtGte field to given value.

### HasThreatCreatedAtGte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatCreatedAtGte() bool`

HasThreatCreatedAtGte returns a boolean if a field has been set.

### GetCsvFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCsvFilterId() string`

GetCsvFilterId returns the CsvFilterId field if non-nil, zero value otherwise.

### GetCsvFilterIdOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCsvFilterIdOk() (*string, bool)`

GetCsvFilterIdOk returns a tuple with the CsvFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCsvFilterId(v string)`

SetCsvFilterId sets CsvFilterId field to given value.

### HasCsvFilterId

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCsvFilterId() bool`

HasCsvFilterId returns a boolean if a field has been set.

### GetAdUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserNameContains() []string`

GetAdUserNameContains returns the AdUserNameContains field if non-nil, zero value otherwise.

### GetAdUserNameContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserNameContainsOk() (*[]string, bool)`

GetAdUserNameContainsOk returns a tuple with the AdUserNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdUserNameContains(v []string)`

SetAdUserNameContains sets AdUserNameContains field to given value.

### HasAdUserNameContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdUserNameContains() bool`

HasAdUserNameContains returns a boolean if a field has been set.

### GetAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAppsVulnerabilityStatusesNin() []string`

GetAppsVulnerabilityStatusesNin returns the AppsVulnerabilityStatusesNin field if non-nil, zero value otherwise.

### GetAppsVulnerabilityStatusesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAppsVulnerabilityStatusesNinOk() (*[]string, bool)`

GetAppsVulnerabilityStatusesNinOk returns a tuple with the AppsVulnerabilityStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAppsVulnerabilityStatusesNin(v []string)`

SetAppsVulnerabilityStatusesNin sets AppsVulnerabilityStatusesNin field to given value.

### HasAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAppsVulnerabilityStatusesNin() bool`

HasAppsVulnerabilityStatusesNin returns a boolean if a field has been set.

### GetUuid

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCoreCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountLt() int32`

GetCoreCountLt returns the CoreCountLt field if non-nil, zero value otherwise.

### GetCoreCountLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCoreCountLtOk() (*int32, bool)`

GetCoreCountLtOk returns a tuple with the CoreCountLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCoreCountLt(v int32)`

SetCoreCountLt sets CoreCountLt field to given value.

### HasCoreCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCoreCountLt() bool`

HasCoreCountLt returns a boolean if a field has been set.

### GetMitigationModeSuspicious

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMitigationModeSuspicious() string`

GetMitigationModeSuspicious returns the MitigationModeSuspicious field if non-nil, zero value otherwise.

### GetMitigationModeSuspiciousOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetMitigationModeSuspiciousOk() (*string, bool)`

GetMitigationModeSuspiciousOk returns a tuple with the MitigationModeSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationModeSuspicious

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetMitigationModeSuspicious(v string)`

SetMitigationModeSuspicious sets MitigationModeSuspicious field to given value.

### HasMitigationModeSuspicious

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasMitigationModeSuspicious() bool`

HasMitigationModeSuspicious returns a boolean if a field has been set.

### GetRangerStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatusesNin() []string`

GetRangerStatusesNin returns the RangerStatusesNin field if non-nil, zero value otherwise.

### GetRangerStatusesNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatusesNinOk() (*[]string, bool)`

GetRangerStatusesNinOk returns a tuple with the RangerStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRangerStatusesNin(v []string)`

SetRangerStatusesNin sets RangerStatusesNin field to given value.

### HasRangerStatusesNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRangerStatusesNin() bool`

HasRangerStatusesNin returns a boolean if a field has been set.

### GetCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtGt() time.Time`

GetCreatedAtGt returns the CreatedAtGt field if non-nil, zero value otherwise.

### GetCreatedAtGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtGtOk() (*time.Time, bool)`

GetCreatedAtGtOk returns a tuple with the CreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCreatedAtGt(v time.Time)`

SetCreatedAtGt sets CreatedAtGt field to given value.

### HasCreatedAtGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCreatedAtGt() bool`

HasCreatedAtGt returns a boolean if a field has been set.

### GetRangerVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerVersionsNin() []string`

GetRangerVersionsNin returns the RangerVersionsNin field if non-nil, zero value otherwise.

### GetRangerVersionsNinOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerVersionsNinOk() (*[]string, bool)`

GetRangerVersionsNinOk returns a tuple with the RangerVersionsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRangerVersionsNin(v []string)`

SetRangerVersionsNin sets RangerVersionsNin field to given value.

### HasRangerVersionsNin

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRangerVersionsNin() bool`

HasRangerVersionsNin returns a boolean if a field has been set.

### GetFilteredGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilteredGroupIds() []string`

GetFilteredGroupIds returns the FilteredGroupIds field if non-nil, zero value otherwise.

### GetFilteredGroupIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetFilteredGroupIdsOk() (*[]string, bool)`

GetFilteredGroupIdsOk returns a tuple with the FilteredGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetFilteredGroupIds(v []string)`

SetFilteredGroupIds sets FilteredGroupIds field to given value.

### HasFilteredGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasFilteredGroupIds() bool`

HasFilteredGroupIds returns a boolean if a field has been set.

### GetDecommissionedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtLte() time.Time`

GetDecommissionedAtLte returns the DecommissionedAtLte field if non-nil, zero value otherwise.

### GetDecommissionedAtLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetDecommissionedAtLteOk() (*time.Time, bool)`

GetDecommissionedAtLteOk returns a tuple with the DecommissionedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetDecommissionedAtLte(v time.Time)`

SetDecommissionedAtLte sets DecommissionedAtLte field to given value.

### HasDecommissionedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasDecommissionedAtLte() bool`

HasDecommissionedAtLte returns a boolean if a field has been set.

### GetOperationalStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOperationalStates() []string`

GetOperationalStates returns the OperationalStates field if non-nil, zero value otherwise.

### GetOperationalStatesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOperationalStatesOk() (*[]string, bool)`

GetOperationalStatesOk returns a tuple with the OperationalStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOperationalStates(v []string)`

SetOperationalStates sets OperationalStates field to given value.

### HasOperationalStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOperationalStates() bool`

HasOperationalStates returns a boolean if a field has been set.

### GetOsVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsVersionContains() []string`

GetOsVersionContains returns the OsVersionContains field if non-nil, zero value otherwise.

### GetOsVersionContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsVersionContainsOk() (*[]string, bool)`

GetOsVersionContainsOk returns a tuple with the OsVersionContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOsVersionContains(v []string)`

SetOsVersionContains sets OsVersionContains field to given value.

### HasOsVersionContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOsVersionContains() bool`

HasOsVersionContains returns a boolean if a field has been set.

### GetUpdatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtLte() time.Time`

GetUpdatedAtLte returns the UpdatedAtLte field if non-nil, zero value otherwise.

### GetUpdatedAtLteOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtLteOk() (*time.Time, bool)`

GetUpdatedAtLteOk returns a tuple with the UpdatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUpdatedAtLte(v time.Time)`

SetUpdatedAtLte sets UpdatedAtLte field to given value.

### HasUpdatedAtLte

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUpdatedAtLte() bool`

HasUpdatedAtLte returns a boolean if a field has been set.

### GetUpdatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtLt() time.Time`

GetUpdatedAtLt returns the UpdatedAtLt field if non-nil, zero value otherwise.

### GetUpdatedAtLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetUpdatedAtLtOk() (*time.Time, bool)`

GetUpdatedAtLtOk returns a tuple with the UpdatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetUpdatedAtLt(v time.Time)`

SetUpdatedAtLt sets UpdatedAtLt field to given value.

### HasUpdatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasUpdatedAtLt() bool`

HasUpdatedAtLt returns a boolean if a field has been set.

### GetK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sNodeLabelsContains() []string`

GetK8sNodeLabelsContains returns the K8sNodeLabelsContains field if non-nil, zero value otherwise.

### GetK8sNodeLabelsContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetK8sNodeLabelsContainsOk() (*[]string, bool)`

GetK8sNodeLabelsContainsOk returns a tuple with the K8sNodeLabelsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetK8sNodeLabelsContains(v []string)`

SetK8sNodeLabelsContains sets K8sNodeLabelsContains field to given value.

### HasK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasK8sNodeLabelsContains() bool`

HasK8sNodeLabelsContains returns a boolean if a field has been set.

### GetAdUserMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserMemberContains() []string`

GetAdUserMemberContains returns the AdUserMemberContains field if non-nil, zero value otherwise.

### GetAdUserMemberContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAdUserMemberContainsOk() (*[]string, bool)`

GetAdUserMemberContainsOk returns a tuple with the AdUserMemberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAdUserMemberContains(v []string)`

SetAdUserMemberContains sets AdUserMemberContains field to given value.

### HasAdUserMemberContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAdUserMemberContains() bool`

HasAdUserMemberContains returns a boolean if a field has been set.

### GetCloudLocationContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudLocationContains() []string`

GetCloudLocationContains returns the CloudLocationContains field if non-nil, zero value otherwise.

### GetCloudLocationContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudLocationContainsOk() (*[]string, bool)`

GetCloudLocationContainsOk returns a tuple with the CloudLocationContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLocationContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudLocationContains(v []string)`

SetCloudLocationContains sets CloudLocationContains field to given value.

### HasCloudLocationContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudLocationContains() bool`

HasCloudLocationContains returns a boolean if a field has been set.

### GetNetworkStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkStatuses() []string`

GetNetworkStatuses returns the NetworkStatuses field if non-nil, zero value otherwise.

### GetNetworkStatusesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkStatusesOk() (*[]string, bool)`

GetNetworkStatusesOk returns a tuple with the NetworkStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkStatuses(v []string)`

SetNetworkStatuses sets NetworkStatuses field to given value.

### HasNetworkStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkStatuses() bool`

HasNetworkStatuses returns a boolean if a field has been set.

### GetRemoteProfilingStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRemoteProfilingStates() []string`

GetRemoteProfilingStates returns the RemoteProfilingStates field if non-nil, zero value otherwise.

### GetRemoteProfilingStatesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRemoteProfilingStatesOk() (*[]string, bool)`

GetRemoteProfilingStatesOk returns a tuple with the RemoteProfilingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRemoteProfilingStates(v []string)`

SetRemoteProfilingStates sets RemoteProfilingStates field to given value.

### HasRemoteProfilingStates

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRemoteProfilingStates() bool`

HasRemoteProfilingStates returns a boolean if a field has been set.

### GetNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfacePhysicalContains() []string`

GetNetworkInterfacePhysicalContains returns the NetworkInterfacePhysicalContains field if non-nil, zero value otherwise.

### GetNetworkInterfacePhysicalContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetNetworkInterfacePhysicalContainsOk() (*[]string, bool)`

GetNetworkInterfacePhysicalContainsOk returns a tuple with the NetworkInterfacePhysicalContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetNetworkInterfacePhysicalContains(v []string)`

SetNetworkInterfacePhysicalContains sets NetworkInterfacePhysicalContains field to given value.

### HasNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasNetworkInterfacePhysicalContains() bool`

HasNetworkInterfacePhysicalContains returns a boolean if a field has been set.

### GetAgentNamespaceContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentNamespaceContains() []string`

GetAgentNamespaceContains returns the AgentNamespaceContains field if non-nil, zero value otherwise.

### GetAgentNamespaceContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetAgentNamespaceContainsOk() (*[]string, bool)`

GetAgentNamespaceContainsOk returns a tuple with the AgentNamespaceContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentNamespaceContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetAgentNamespaceContains(v []string)`

SetAgentNamespaceContains sets AgentNamespaceContains field to given value.

### HasAgentNamespaceContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasAgentNamespaceContains() bool`

HasAgentNamespaceContains returns a boolean if a field has been set.

### GetRangerStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatuses() []string`

GetRangerStatuses returns the RangerStatuses field if non-nil, zero value otherwise.

### GetRangerStatusesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetRangerStatusesOk() (*[]string, bool)`

GetRangerStatusesOk returns a tuple with the RangerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetRangerStatuses(v []string)`

SetRangerStatuses sets RangerStatuses field to given value.

### HasRangerStatuses

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasRangerStatuses() bool`

HasRangerStatuses returns a boolean if a field has been set.

### GetLastActiveDateLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateLt() time.Time`

GetLastActiveDateLt returns the LastActiveDateLt field if non-nil, zero value otherwise.

### GetLastActiveDateLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetLastActiveDateLtOk() (*time.Time, bool)`

GetLastActiveDateLtOk returns a tuple with the LastActiveDateLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetLastActiveDateLt(v time.Time)`

SetLastActiveDateLt sets LastActiveDateLt field to given value.

### HasLastActiveDateLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasLastActiveDateLt() bool`

HasLastActiveDateLt returns a boolean if a field has been set.

### GetActiveThreatsGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetActiveThreatsGt() int32`

GetActiveThreatsGt returns the ActiveThreatsGt field if non-nil, zero value otherwise.

### GetActiveThreatsGtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetActiveThreatsGtOk() (*int32, bool)`

GetActiveThreatsGtOk returns a tuple with the ActiveThreatsGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreatsGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetActiveThreatsGt(v int32)`

SetActiveThreatsGt sets ActiveThreatsGt field to given value.

### HasActiveThreatsGt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasActiveThreatsGt() bool`

HasActiveThreatsGt returns a boolean if a field has been set.

### GetCloudInstanceIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudInstanceIdContains() []string`

GetCloudInstanceIdContains returns the CloudInstanceIdContains field if non-nil, zero value otherwise.

### GetCloudInstanceIdContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCloudInstanceIdContainsOk() (*[]string, bool)`

GetCloudInstanceIdContainsOk returns a tuple with the CloudInstanceIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCloudInstanceIdContains(v []string)`

SetCloudInstanceIdContains sets CloudInstanceIdContains field to given value.

### HasCloudInstanceIdContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCloudInstanceIdContains() bool`

HasCloudInstanceIdContains returns a boolean if a field has been set.

### GetCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### GetCpuCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountBetween() string`

GetCpuCountBetween returns the CpuCountBetween field if non-nil, zero value otherwise.

### GetCpuCountBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountBetweenOk() (*string, bool)`

GetCpuCountBetweenOk returns a tuple with the CpuCountBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuCountBetween(v string)`

SetCpuCountBetween sets CpuCountBetween field to given value.

### HasCpuCountBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuCountBetween() bool`

HasCpuCountBetween returns a boolean if a field has been set.

### GetGcpServiceAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGcpServiceAccountContains() []string`

GetGcpServiceAccountContains returns the GcpServiceAccountContains field if non-nil, zero value otherwise.

### GetGcpServiceAccountContainsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGcpServiceAccountContainsOk() (*[]string, bool)`

GetGcpServiceAccountContainsOk returns a tuple with the GcpServiceAccountContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetGcpServiceAccountContains(v []string)`

SetGcpServiceAccountContains sets GcpServiceAccountContains field to given value.

### HasGcpServiceAccountContains

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasGcpServiceAccountContains() bool`

HasGcpServiceAccountContains returns a boolean if a field has been set.

### GetThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtBetween() string`

GetThreatCreatedAtBetween returns the ThreatCreatedAtBetween field if non-nil, zero value otherwise.

### GetThreatCreatedAtBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetThreatCreatedAtBetweenOk() (*string, bool)`

GetThreatCreatedAtBetweenOk returns a tuple with the ThreatCreatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetThreatCreatedAtBetween(v string)`

SetThreatCreatedAtBetween sets ThreatCreatedAtBetween field to given value.

### HasThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasThreatCreatedAtBetween() bool`

HasThreatCreatedAtBetween returns a boolean if a field has been set.

### GetOsTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsTypes() []string`

GetOsTypes returns the OsTypes field if non-nil, zero value otherwise.

### GetOsTypesOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetOsTypesOk() (*[]string, bool)`

GetOsTypesOk returns a tuple with the OsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetOsTypes(v []string)`

SetOsTypes sets OsTypes field to given value.

### HasOsTypes

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasOsTypes() bool`

HasOsTypes returns a boolean if a field has been set.

### GetCpuCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountLt() int32`

GetCpuCountLt returns the CpuCountLt field if non-nil, zero value otherwise.

### GetCpuCountLtOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCpuCountLtOk() (*int32, bool)`

GetCpuCountLtOk returns a tuple with the CpuCountLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCpuCountLt(v int32)`

SetCpuCountLt sets CpuCountLt field to given value.

### HasCpuCountLt

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCpuCountLt() bool`

HasCpuCountLt returns a boolean if a field has been set.

### GetGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetComputerNameLike

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerNameLike() string`

GetComputerNameLike returns the ComputerNameLike field if non-nil, zero value otherwise.

### GetComputerNameLikeOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetComputerNameLikeOk() (*string, bool)`

GetComputerNameLikeOk returns a tuple with the ComputerNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerNameLike

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetComputerNameLike(v string)`

SetComputerNameLike sets ComputerNameLike field to given value.

### HasComputerNameLike

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasComputerNameLike() bool`

HasComputerNameLike returns a boolean if a field has been set.

### GetTagsData

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTagsData() string`

GetTagsData returns the TagsData field if non-nil, zero value otherwise.

### GetTagsDataOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetTagsDataOk() (*string, bool)`

GetTagsDataOk returns a tuple with the TagsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsData

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetTagsData(v string)`

SetTagsData sets TagsData field to given value.

### HasTagsData

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasTagsData() bool`

HasTagsData returns a boolean if a field has been set.

### GetCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtBetween() string`

GetCreatedAtBetween returns the CreatedAtBetween field if non-nil, zero value otherwise.

### GetCreatedAtBetweenOk

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) GetCreatedAtBetweenOk() (*string, bool)`

GetCreatedAtBetweenOk returns a tuple with the CreatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) SetCreatedAtBetween(v string)`

SetCreatedAtBetween sets CreatedAtBetween field to given value.

### HasCreatedAtBetween

`func (o *AgentsSchemasAgentsDangerousActionSchemaFilter) HasCreatedAtBetween() bool`

HasCreatedAtBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


