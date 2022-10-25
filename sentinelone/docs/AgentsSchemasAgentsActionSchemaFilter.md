# AgentsSchemasAgentsActionSchemaFilter

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
**GroupIds** | Pointer to **[]string** | List of Group IDs to filter by | [optional] 
**ComputerNameLike** | Pointer to **string** | Match computer name partially (substring) | [optional] 
**TagsData** | Pointer to **string** | A JSON formatted string, where each key represents a tag key, and each value represents a list of string values | [optional] 
**CreatedAtBetween** | Pointer to **string** | Date range for creation time (format: &lt;from_timestamp&gt;-&lt;to_timestamp&gt;, inclusive) | [optional] 

## Methods

### NewAgentsSchemasAgentsActionSchemaFilter

`func NewAgentsSchemasAgentsActionSchemaFilter() *AgentsSchemasAgentsActionSchemaFilter`

NewAgentsSchemasAgentsActionSchemaFilter instantiates a new AgentsSchemasAgentsActionSchemaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentsActionSchemaFilterWithDefaults

`func NewAgentsSchemasAgentsActionSchemaFilterWithDefaults() *AgentsSchemasAgentsActionSchemaFilter`

NewAgentsSchemasAgentsActionSchemaFilterWithDefaults instantiates a new AgentsSchemasAgentsActionSchemaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatusesNin() []string`

GetScanStatusesNin returns the ScanStatusesNin field if non-nil, zero value otherwise.

### GetScanStatusesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatusesNinOk() (*[]string, bool)`

GetScanStatusesNinOk returns a tuple with the ScanStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetScanStatusesNin(v []string)`

SetScanStatusesNin sets ScanStatusesNin field to given value.

### HasScanStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasScanStatusesNin() bool`

HasScanStatusesNin returns a boolean if a field has been set.

### GetQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDecommissionedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtGte() time.Time`

GetDecommissionedAtGte returns the DecommissionedAtGte field if non-nil, zero value otherwise.

### GetDecommissionedAtGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtGteOk() (*time.Time, bool)`

GetDecommissionedAtGteOk returns a tuple with the DecommissionedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDecommissionedAtGte(v time.Time)`

SetDecommissionedAtGte sets DecommissionedAtGte field to given value.

### HasDecommissionedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDecommissionedAtGte() bool`

HasDecommissionedAtGte returns a boolean if a field has been set.

### GetAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsSecurityGroupsContains() []string`

GetAwsSecurityGroupsContains returns the AwsSecurityGroupsContains field if non-nil, zero value otherwise.

### GetAwsSecurityGroupsContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsSecurityGroupsContainsOk() (*[]string, bool)`

GetAwsSecurityGroupsContainsOk returns a tuple with the AwsSecurityGroupsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAwsSecurityGroupsContains(v []string)`

SetAwsSecurityGroupsContains sets AwsSecurityGroupsContains field to given value.

### HasAwsSecurityGroupsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAwsSecurityGroupsContains() bool`

HasAwsSecurityGroupsContains returns a boolean if a field has been set.

### GetThreatMitigationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatMitigationStatus() string`

GetThreatMitigationStatus returns the ThreatMitigationStatus field if non-nil, zero value otherwise.

### GetThreatMitigationStatusOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatMitigationStatusOk() (*string, bool)`

GetThreatMitigationStatusOk returns a tuple with the ThreatMitigationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatMitigationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatMitigationStatus(v string)`

SetThreatMitigationStatus sets ThreatMitigationStatus field to given value.

### HasThreatMitigationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatMitigationStatus() bool`

HasThreatMitigationStatus returns a boolean if a field has been set.

### GetRegisteredAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtLte() time.Time`

GetRegisteredAtLte returns the RegisteredAtLte field if non-nil, zero value otherwise.

### GetRegisteredAtLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtLteOk() (*time.Time, bool)`

GetRegisteredAtLteOk returns a tuple with the RegisteredAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRegisteredAtLte(v time.Time)`

SetRegisteredAtLte sets RegisteredAtLte field to given value.

### HasRegisteredAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRegisteredAtLte() bool`

HasRegisteredAtLte returns a boolean if a field has been set.

### GetUpdatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtGt() time.Time`

GetUpdatedAtGt returns the UpdatedAtGt field if non-nil, zero value otherwise.

### GetUpdatedAtGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtGtOk() (*time.Time, bool)`

GetUpdatedAtGtOk returns a tuple with the UpdatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUpdatedAtGt(v time.Time)`

SetUpdatedAtGt sets UpdatedAtGt field to given value.

### HasUpdatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUpdatedAtGt() bool`

HasUpdatedAtGt returns a boolean if a field has been set.

### GetDomains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkQuarantineEnabled() []bool`

GetNetworkQuarantineEnabled returns the NetworkQuarantineEnabled field if non-nil, zero value otherwise.

### GetNetworkQuarantineEnabledOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkQuarantineEnabledOk() (*[]bool, bool)`

GetNetworkQuarantineEnabledOk returns a tuple with the NetworkQuarantineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkQuarantineEnabled(v []bool)`

SetNetworkQuarantineEnabled sets NetworkQuarantineEnabled field to given value.

### HasNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkQuarantineEnabled() bool`

HasNetworkQuarantineEnabled returns a boolean if a field has been set.

### GetMigrationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMigrationStatus() string`

GetMigrationStatus returns the MigrationStatus field if non-nil, zero value otherwise.

### GetMigrationStatusOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMigrationStatusOk() (*string, bool)`

GetMigrationStatusOk returns a tuple with the MigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetMigrationStatus(v string)`

SetMigrationStatus sets MigrationStatus field to given value.

### HasMigrationStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasMigrationStatus() bool`

HasMigrationStatus returns a boolean if a field has been set.

### GetAccountIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfaceInetContains() []string`

GetNetworkInterfaceInetContains returns the NetworkInterfaceInetContains field if non-nil, zero value otherwise.

### GetNetworkInterfaceInetContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfaceInetContainsOk() (*[]string, bool)`

GetNetworkInterfaceInetContainsOk returns a tuple with the NetworkInterfaceInetContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkInterfaceInetContains(v []string)`

SetNetworkInterfaceInetContains sets NetworkInterfaceInetContains field to given value.

### HasNetworkInterfaceInetContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkInterfaceInetContains() bool`

HasNetworkInterfaceInetContains returns a boolean if a field has been set.

### GetAgentVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentVersionsNin() []string`

GetAgentVersionsNin returns the AgentVersionsNin field if non-nil, zero value otherwise.

### GetAgentVersionsNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentVersionsNinOk() (*[]string, bool)`

GetAgentVersionsNinOk returns a tuple with the AgentVersionsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAgentVersionsNin(v []string)`

SetAgentVersionsNin sets AgentVersionsNin field to given value.

### HasAgentVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAgentVersionsNin() bool`

HasAgentVersionsNin returns a boolean if a field has been set.

### GetCloudProvider

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudProvider() []string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudProviderOk() (*[]string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudProvider(v []string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetOsTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsTypesNin() []string`

GetOsTypesNin returns the OsTypesNin field if non-nil, zero value otherwise.

### GetOsTypesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsTypesNinOk() (*[]string, bool)`

GetOsTypesNinOk returns a tuple with the OsTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOsTypesNin(v []string)`

SetOsTypesNin sets OsTypesNin field to given value.

### HasOsTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOsTypesNin() bool`

HasOsTypesNin returns a boolean if a field has been set.

### GetInstallerTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInstallerTypesNin() []string`

GetInstallerTypesNin returns the InstallerTypesNin field if non-nil, zero value otherwise.

### GetInstallerTypesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInstallerTypesNinOk() (*[]string, bool)`

GetInstallerTypesNinOk returns a tuple with the InstallerTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetInstallerTypesNin(v []string)`

SetInstallerTypesNin sets InstallerTypesNin field to given value.

### HasInstallerTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasInstallerTypesNin() bool`

HasInstallerTypesNin returns a boolean if a field has been set.

### GetMachineTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMachineTypes() []string`

GetMachineTypes returns the MachineTypes field if non-nil, zero value otherwise.

### GetMachineTypesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMachineTypesOk() (*[]string, bool)`

GetMachineTypesOk returns a tuple with the MachineTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetMachineTypes(v []string)`

SetMachineTypes sets MachineTypes field to given value.

### HasMachineTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasMachineTypes() bool`

HasMachineTypes returns a boolean if a field has been set.

### GetEncryptedApplications

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetEncryptedApplications() bool`

GetEncryptedApplications returns the EncryptedApplications field if non-nil, zero value otherwise.

### GetEncryptedApplicationsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetEncryptedApplicationsOk() (*bool, bool)`

GetEncryptedApplicationsOk returns a tuple with the EncryptedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedApplications

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetEncryptedApplications(v bool)`

SetEncryptedApplications sets EncryptedApplications field to given value.

### HasEncryptedApplications

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasEncryptedApplications() bool`

HasEncryptedApplications returns a boolean if a field has been set.

### GetIsPendingUninstall

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsPendingUninstall() bool`

GetIsPendingUninstall returns the IsPendingUninstall field if non-nil, zero value otherwise.

### GetIsPendingUninstallOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsPendingUninstallOk() (*bool, bool)`

GetIsPendingUninstallOk returns a tuple with the IsPendingUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPendingUninstall

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIsPendingUninstall(v bool)`

SetIsPendingUninstall sets IsPendingUninstall field to given value.

### HasIsPendingUninstall

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIsPendingUninstall() bool`

HasIsPendingUninstall returns a boolean if a field has been set.

### GetHasLocalConfiguration

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetHasLocalConfiguration() bool`

GetHasLocalConfiguration returns the HasLocalConfiguration field if non-nil, zero value otherwise.

### GetHasLocalConfigurationOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetHasLocalConfigurationOk() (*bool, bool)`

GetHasLocalConfigurationOk returns a tuple with the HasLocalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLocalConfiguration

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetHasLocalConfiguration(v bool)`

SetHasLocalConfiguration sets HasLocalConfiguration field to given value.

### HasHasLocalConfiguration

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasHasLocalConfiguration() bool`

HasHasLocalConfiguration returns a boolean if a field has been set.

### GetIsUpToDate

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsUpToDate() bool`

GetIsUpToDate returns the IsUpToDate field if non-nil, zero value otherwise.

### GetIsUpToDateOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsUpToDateOk() (*bool, bool)`

GetIsUpToDateOk returns a tuple with the IsUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpToDate

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIsUpToDate(v bool)`

SetIsUpToDate sets IsUpToDate field to given value.

### HasIsUpToDate

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIsUpToDate() bool`

HasIsUpToDate returns a boolean if a field has been set.

### GetIsDecommissioned

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsDecommissioned() []bool`

GetIsDecommissioned returns the IsDecommissioned field if non-nil, zero value otherwise.

### GetIsDecommissionedOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsDecommissionedOk() (*[]bool, bool)`

GetIsDecommissionedOk returns a tuple with the IsDecommissioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecommissioned

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIsDecommissioned(v []bool)`

SetIsDecommissioned sets IsDecommissioned field to given value.

### HasIsDecommissioned

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIsDecommissioned() bool`

HasIsDecommissioned returns a boolean if a field has been set.

### GetComputerName

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastLoggedInUserNameContains() []string`

GetLastLoggedInUserNameContains returns the LastLoggedInUserNameContains field if non-nil, zero value otherwise.

### GetLastLoggedInUserNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastLoggedInUserNameContainsOk() (*[]string, bool)`

GetLastLoggedInUserNameContainsOk returns a tuple with the LastLoggedInUserNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastLoggedInUserNameContains(v []string)`

SetLastLoggedInUserNameContains sets LastLoggedInUserNameContains field to given value.

### HasLastLoggedInUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastLoggedInUserNameContains() bool`

HasLastLoggedInUserNameContains returns a boolean if a field has been set.

### GetAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAppsVulnerabilityStatuses() []string`

GetAppsVulnerabilityStatuses returns the AppsVulnerabilityStatuses field if non-nil, zero value otherwise.

### GetAppsVulnerabilityStatusesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAppsVulnerabilityStatusesOk() (*[]string, bool)`

GetAppsVulnerabilityStatusesOk returns a tuple with the AppsVulnerabilityStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAppsVulnerabilityStatuses(v []string)`

SetAppsVulnerabilityStatuses sets AppsVulnerabilityStatuses field to given value.

### HasAppsVulnerabilityStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAppsVulnerabilityStatuses() bool`

HasAppsVulnerabilityStatuses returns a boolean if a field has been set.

### GetK8sVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sVersionContains() []string`

GetK8sVersionContains returns the K8sVersionContains field if non-nil, zero value otherwise.

### GetK8sVersionContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sVersionContainsOk() (*[]string, bool)`

GetK8sVersionContainsOk returns a tuple with the K8sVersionContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetK8sVersionContains(v []string)`

SetK8sVersionContains sets K8sVersionContains field to given value.

### HasK8sVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasK8sVersionContains() bool`

HasK8sVersionContains returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFirewallEnabled() []bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFirewallEnabledOk() (*[]bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetFirewallEnabled(v []bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetDecommissionedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtLt() time.Time`

GetDecommissionedAtLt returns the DecommissionedAtLt field if non-nil, zero value otherwise.

### GetDecommissionedAtLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtLtOk() (*time.Time, bool)`

GetDecommissionedAtLtOk returns a tuple with the DecommissionedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDecommissionedAtLt(v time.Time)`

SetDecommissionedAtLt sets DecommissionedAtLt field to given value.

### HasDecommissionedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDecommissionedAtLt() bool`

HasDecommissionedAtLt returns a boolean if a field has been set.

### GetActiveThreats

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetActiveThreats() int32`

GetActiveThreats returns the ActiveThreats field if non-nil, zero value otherwise.

### GetActiveThreatsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetActiveThreatsOk() (*int32, bool)`

GetActiveThreatsOk returns a tuple with the ActiveThreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreats

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetActiveThreats(v int32)`

SetActiveThreats sets ActiveThreats field to given value.

### HasActiveThreats

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasActiveThreats() bool`

HasActiveThreats returns a boolean if a field has been set.

### GetAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsSubnetIdsContains() []string`

GetAwsSubnetIdsContains returns the AwsSubnetIdsContains field if non-nil, zero value otherwise.

### GetAwsSubnetIdsContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsSubnetIdsContainsOk() (*[]string, bool)`

GetAwsSubnetIdsContainsOk returns a tuple with the AwsSubnetIdsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAwsSubnetIdsContains(v []string)`

SetAwsSubnetIdsContains sets AwsSubnetIdsContains field to given value.

### HasAwsSubnetIdsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAwsSubnetIdsContains() bool`

HasAwsSubnetIdsContains returns a boolean if a field has been set.

### GetCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtGte() time.Time`

GetCreatedAtGte returns the CreatedAtGte field if non-nil, zero value otherwise.

### GetCreatedAtGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtGteOk() (*time.Time, bool)`

GetCreatedAtGteOk returns a tuple with the CreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCreatedAtGte(v time.Time)`

SetCreatedAtGte sets CreatedAtGte field to given value.

### HasCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCreatedAtGte() bool`

HasCreatedAtGte returns a boolean if a field has been set.

### GetK8sNodeNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sNodeNameContains() []string`

GetK8sNodeNameContains returns the K8sNodeNameContains field if non-nil, zero value otherwise.

### GetK8sNodeNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sNodeNameContainsOk() (*[]string, bool)`

GetK8sNodeNameContainsOk returns a tuple with the K8sNodeNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetK8sNodeNameContains(v []string)`

SetK8sNodeNameContains sets K8sNodeNameContains field to given value.

### HasK8sNodeNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasK8sNodeNameContains() bool`

HasK8sNodeNameContains returns a boolean if a field has been set.

### GetTotalMemoryGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryGte() int32`

GetTotalMemoryGte returns the TotalMemoryGte field if non-nil, zero value otherwise.

### GetTotalMemoryGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryGteOk() (*int32, bool)`

GetTotalMemoryGteOk returns a tuple with the TotalMemoryGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTotalMemoryGte(v int32)`

SetTotalMemoryGte sets TotalMemoryGte field to given value.

### HasTotalMemoryGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTotalMemoryGte() bool`

HasTotalMemoryGte returns a boolean if a field has been set.

### GetMachineTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMachineTypesNin() []string`

GetMachineTypesNin returns the MachineTypesNin field if non-nil, zero value otherwise.

### GetMachineTypesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMachineTypesNinOk() (*[]string, bool)`

GetMachineTypesNinOk returns a tuple with the MachineTypesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetMachineTypesNin(v []string)`

SetMachineTypesNin sets MachineTypesNin field to given value.

### HasMachineTypesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasMachineTypesNin() bool`

HasMachineTypesNin returns a boolean if a field has been set.

### GetLastActiveDateBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateBetween() string`

GetLastActiveDateBetween returns the LastActiveDateBetween field if non-nil, zero value otherwise.

### GetLastActiveDateBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateBetweenOk() (*string, bool)`

GetLastActiveDateBetweenOk returns a tuple with the LastActiveDateBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastActiveDateBetween(v string)`

SetLastActiveDateBetween sets LastActiveDateBetween field to given value.

### HasLastActiveDateBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastActiveDateBetween() bool`

HasLastActiveDateBetween returns a boolean if a field has been set.

### GetNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfaceGatewayMacAddressContains() []string`

GetNetworkInterfaceGatewayMacAddressContains returns the NetworkInterfaceGatewayMacAddressContains field if non-nil, zero value otherwise.

### GetNetworkInterfaceGatewayMacAddressContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfaceGatewayMacAddressContainsOk() (*[]string, bool)`

GetNetworkInterfaceGatewayMacAddressContainsOk returns a tuple with the NetworkInterfaceGatewayMacAddressContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkInterfaceGatewayMacAddressContains(v []string)`

SetNetworkInterfaceGatewayMacAddressContains sets NetworkInterfaceGatewayMacAddressContains field to given value.

### HasNetworkInterfaceGatewayMacAddressContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkInterfaceGatewayMacAddressContains() bool`

HasNetworkInterfaceGatewayMacAddressContains returns a boolean if a field has been set.

### GetAwsRoleContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsRoleContains() []string`

GetAwsRoleContains returns the AwsRoleContains field if non-nil, zero value otherwise.

### GetAwsRoleContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAwsRoleContainsOk() (*[]string, bool)`

GetAwsRoleContainsOk returns a tuple with the AwsRoleContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRoleContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAwsRoleContains(v []string)`

SetAwsRoleContains sets AwsRoleContains field to given value.

### HasAwsRoleContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAwsRoleContains() bool`

HasAwsRoleContains returns a boolean if a field has been set.

### GetRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRemoteProfilingStatesNin() []string`

GetRemoteProfilingStatesNin returns the RemoteProfilingStatesNin field if non-nil, zero value otherwise.

### GetRemoteProfilingStatesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRemoteProfilingStatesNinOk() (*[]string, bool)`

GetRemoteProfilingStatesNinOk returns a tuple with the RemoteProfilingStatesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRemoteProfilingStatesNin(v []string)`

SetRemoteProfilingStatesNin sets RemoteProfilingStatesNin field to given value.

### HasRemoteProfilingStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRemoteProfilingStatesNin() bool`

HasRemoteProfilingStatesNin returns a boolean if a field has been set.

### GetLastActiveDateLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateLte() time.Time`

GetLastActiveDateLte returns the LastActiveDateLte field if non-nil, zero value otherwise.

### GetLastActiveDateLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateLteOk() (*time.Time, bool)`

GetLastActiveDateLteOk returns a tuple with the LastActiveDateLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastActiveDateLte(v time.Time)`

SetLastActiveDateLte sets LastActiveDateLte field to given value.

### HasLastActiveDateLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastActiveDateLte() bool`

HasLastActiveDateLte returns a boolean if a field has been set.

### GetCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtLte() time.Time`

GetCreatedAtLte returns the CreatedAtLte field if non-nil, zero value otherwise.

### GetCreatedAtLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtLteOk() (*time.Time, bool)`

GetCreatedAtLteOk returns a tuple with the CreatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCreatedAtLte(v time.Time)`

SetCreatedAtLte sets CreatedAtLte field to given value.

### HasCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCreatedAtLte() bool`

HasCreatedAtLte returns a boolean if a field has been set.

### GetConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetConsoleMigrationStatusesNin() []string`

GetConsoleMigrationStatusesNin returns the ConsoleMigrationStatusesNin field if non-nil, zero value otherwise.

### GetConsoleMigrationStatusesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetConsoleMigrationStatusesNinOk() (*[]string, bool)`

GetConsoleMigrationStatusesNinOk returns a tuple with the ConsoleMigrationStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetConsoleMigrationStatusesNin(v []string)`

SetConsoleMigrationStatusesNin sets ConsoleMigrationStatusesNin field to given value.

### HasConsoleMigrationStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasConsoleMigrationStatusesNin() bool`

HasConsoleMigrationStatusesNin returns a boolean if a field has been set.

### GetFilteredSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilteredSiteIds() []string`

GetFilteredSiteIds returns the FilteredSiteIds field if non-nil, zero value otherwise.

### GetFilteredSiteIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilteredSiteIdsOk() (*[]string, bool)`

GetFilteredSiteIdsOk returns a tuple with the FilteredSiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetFilteredSiteIds(v []string)`

SetFilteredSiteIds sets FilteredSiteIds field to given value.

### HasFilteredSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasFilteredSiteIds() bool`

HasFilteredSiteIds returns a boolean if a field has been set.

### GetAdComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerNameContains() []string`

GetAdComputerNameContains returns the AdComputerNameContains field if non-nil, zero value otherwise.

### GetAdComputerNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerNameContainsOk() (*[]string, bool)`

GetAdComputerNameContainsOk returns a tuple with the AdComputerNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdComputerNameContains(v []string)`

SetAdComputerNameContains sets AdComputerNameContains field to given value.

### HasAdComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdComputerNameContains() bool`

HasAdComputerNameContains returns a boolean if a field has been set.

### GetRegisteredAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtGt() time.Time`

GetRegisteredAtGt returns the RegisteredAtGt field if non-nil, zero value otherwise.

### GetRegisteredAtGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtGtOk() (*time.Time, bool)`

GetRegisteredAtGtOk returns a tuple with the RegisteredAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRegisteredAtGt(v time.Time)`

SetRegisteredAtGt sets RegisteredAtGt field to given value.

### HasRegisteredAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRegisteredAtGt() bool`

HasRegisteredAtGt returns a boolean if a field has been set.

### GetCloudTagsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudTagsContains() []string`

GetCloudTagsContains returns the CloudTagsContains field if non-nil, zero value otherwise.

### GetCloudTagsContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudTagsContainsOk() (*[]string, bool)`

GetCloudTagsContainsOk returns a tuple with the CloudTagsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudTagsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudTagsContains(v []string)`

SetCloudTagsContains sets CloudTagsContains field to given value.

### HasCloudTagsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudTagsContains() bool`

HasCloudTagsContains returns a boolean if a field has been set.

### GetAdUserQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserQueryContains() []string`

GetAdUserQueryContains returns the AdUserQueryContains field if non-nil, zero value otherwise.

### GetAdUserQueryContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserQueryContainsOk() (*[]string, bool)`

GetAdUserQueryContainsOk returns a tuple with the AdUserQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdUserQueryContains(v []string)`

SetAdUserQueryContains sets AdUserQueryContains field to given value.

### HasAdUserQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdUserQueryContains() bool`

HasAdUserQueryContains returns a boolean if a field has been set.

### GetThreatContentHash

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatContentHash() string`

GetThreatContentHash returns the ThreatContentHash field if non-nil, zero value otherwise.

### GetThreatContentHashOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatContentHashOk() (*string, bool)`

GetThreatContentHashOk returns a tuple with the ThreatContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatContentHash

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatContentHash(v string)`

SetThreatContentHash sets ThreatContentHash field to given value.

### HasThreatContentHash

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatContentHash() bool`

HasThreatContentHash returns a boolean if a field has been set.

### GetScanStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatuses() []string`

GetScanStatuses returns the ScanStatuses field if non-nil, zero value otherwise.

### GetScanStatusesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatusesOk() (*[]string, bool)`

GetScanStatusesOk returns a tuple with the ScanStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetScanStatuses(v []string)`

SetScanStatuses sets ScanStatuses field to given value.

### HasScanStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasScanStatuses() bool`

HasScanStatuses returns a boolean if a field has been set.

### GetOperationalStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOperationalStatesNin() []string`

GetOperationalStatesNin returns the OperationalStatesNin field if non-nil, zero value otherwise.

### GetOperationalStatesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOperationalStatesNinOk() (*[]string, bool)`

GetOperationalStatesNinOk returns a tuple with the OperationalStatesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOperationalStatesNin(v []string)`

SetOperationalStatesNin sets OperationalStatesNin field to given value.

### HasOperationalStatesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOperationalStatesNin() bool`

HasOperationalStatesNin returns a boolean if a field has been set.

### GetThreatRebootRequired

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatRebootRequired() []bool`

GetThreatRebootRequired returns the ThreatRebootRequired field if non-nil, zero value otherwise.

### GetThreatRebootRequiredOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatRebootRequiredOk() (*[]bool, bool)`

GetThreatRebootRequiredOk returns a tuple with the ThreatRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatRebootRequired

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatRebootRequired(v []bool)`

SetThreatRebootRequired sets ThreatRebootRequired field to given value.

### HasThreatRebootRequired

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatRebootRequired() bool`

HasThreatRebootRequired returns a boolean if a field has been set.

### GetIsActive

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTotalMemoryLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryLt() int32`

GetTotalMemoryLt returns the TotalMemoryLt field if non-nil, zero value otherwise.

### GetTotalMemoryLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryLtOk() (*int32, bool)`

GetTotalMemoryLtOk returns a tuple with the TotalMemoryLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTotalMemoryLt(v int32)`

SetTotalMemoryLt sets TotalMemoryLt field to given value.

### HasTotalMemoryLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTotalMemoryLt() bool`

HasTotalMemoryLt returns a boolean if a field has been set.

### GetAdComputerQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerQueryContains() []string`

GetAdComputerQueryContains returns the AdComputerQueryContains field if non-nil, zero value otherwise.

### GetAdComputerQueryContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerQueryContainsOk() (*[]string, bool)`

GetAdComputerQueryContainsOk returns a tuple with the AdComputerQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdComputerQueryContains(v []string)`

SetAdComputerQueryContains sets AdComputerQueryContains field to given value.

### HasAdComputerQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdComputerQueryContains() bool`

HasAdComputerQueryContains returns a boolean if a field has been set.

### GetUserActionsNeeded

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUserActionsNeeded() []string`

GetUserActionsNeeded returns the UserActionsNeeded field if non-nil, zero value otherwise.

### GetUserActionsNeededOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUserActionsNeededOk() (*[]string, bool)`

GetUserActionsNeededOk returns a tuple with the UserActionsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsNeeded

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUserActionsNeeded(v []string)`

SetUserActionsNeeded sets UserActionsNeeded field to given value.

### HasUserActionsNeeded

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUserActionsNeeded() bool`

HasUserActionsNeeded returns a boolean if a field has been set.

### GetUpdatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtBetween() string`

GetUpdatedAtBetween returns the UpdatedAtBetween field if non-nil, zero value otherwise.

### GetUpdatedAtBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtBetweenOk() (*string, bool)`

GetUpdatedAtBetweenOk returns a tuple with the UpdatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUpdatedAtBetween(v string)`

SetUpdatedAtBetween sets UpdatedAtBetween field to given value.

### HasUpdatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUpdatedAtBetween() bool`

HasUpdatedAtBetween returns a boolean if a field has been set.

### GetCpuCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountGte() int32`

GetCpuCountGte returns the CpuCountGte field if non-nil, zero value otherwise.

### GetCpuCountGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountGteOk() (*int32, bool)`

GetCpuCountGteOk returns a tuple with the CpuCountGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuCountGte(v int32)`

SetCpuCountGte sets CpuCountGte field to given value.

### HasCpuCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuCountGte() bool`

HasCpuCountGte returns a boolean if a field has been set.

### GetUuids

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuids() []string`

GetUuids returns the Uuids field if non-nil, zero value otherwise.

### GetUuidsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuidsOk() (*[]string, bool)`

GetUuidsOk returns a tuple with the Uuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuids

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUuids(v []string)`

SetUuids sets Uuids field to given value.

### HasUuids

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUuids() bool`

HasUuids returns a boolean if a field has been set.

### GetSerialNumberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetSerialNumberContains() []string`

GetSerialNumberContains returns the SerialNumberContains field if non-nil, zero value otherwise.

### GetSerialNumberContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetSerialNumberContainsOk() (*[]string, bool)`

GetSerialNumberContainsOk returns a tuple with the SerialNumberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetSerialNumberContains(v []string)`

SetSerialNumberContains sets SerialNumberContains field to given value.

### HasSerialNumberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasSerialNumberContains() bool`

HasSerialNumberContains returns a boolean if a field has been set.

### GetRegisteredAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtBetween() string`

GetRegisteredAtBetween returns the RegisteredAtBetween field if non-nil, zero value otherwise.

### GetRegisteredAtBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtBetweenOk() (*string, bool)`

GetRegisteredAtBetweenOk returns a tuple with the RegisteredAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRegisteredAtBetween(v string)`

SetRegisteredAtBetween sets RegisteredAtBetween field to given value.

### HasRegisteredAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRegisteredAtBetween() bool`

HasRegisteredAtBetween returns a boolean if a field has been set.

### GetScanStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetExternalIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetExternalIdContains() []string`

GetExternalIdContains returns the ExternalIdContains field if non-nil, zero value otherwise.

### GetExternalIdContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetExternalIdContainsOk() (*[]string, bool)`

GetExternalIdContainsOk returns a tuple with the ExternalIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetExternalIdContains(v []string)`

SetExternalIdContains sets ExternalIdContains field to given value.

### HasExternalIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasExternalIdContains() bool`

HasExternalIdContains returns a boolean if a field has been set.

### GetTotalMemoryLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryLte() int32`

GetTotalMemoryLte returns the TotalMemoryLte field if non-nil, zero value otherwise.

### GetTotalMemoryLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryLteOk() (*int32, bool)`

GetTotalMemoryLteOk returns a tuple with the TotalMemoryLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTotalMemoryLte(v int32)`

SetTotalMemoryLte sets TotalMemoryLte field to given value.

### HasTotalMemoryLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTotalMemoryLte() bool`

HasTotalMemoryLte returns a boolean if a field has been set.

### GetLocationIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationIds() []string`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationIdsOk() (*[]string, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLocationIds(v []string)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetCloudNetworkContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudNetworkContains() []string`

GetCloudNetworkContains returns the CloudNetworkContains field if non-nil, zero value otherwise.

### GetCloudNetworkContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudNetworkContainsOk() (*[]string, bool)`

GetCloudNetworkContainsOk returns a tuple with the CloudNetworkContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudNetworkContains(v []string)`

SetCloudNetworkContains sets CloudNetworkContains field to given value.

### HasCloudNetworkContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudNetworkContains() bool`

HasCloudNetworkContains returns a boolean if a field has been set.

### GetAzureResourceGroupContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAzureResourceGroupContains() []string`

GetAzureResourceGroupContains returns the AzureResourceGroupContains field if non-nil, zero value otherwise.

### GetAzureResourceGroupContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAzureResourceGroupContainsOk() (*[]string, bool)`

GetAzureResourceGroupContainsOk returns a tuple with the AzureResourceGroupContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureResourceGroupContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAzureResourceGroupContains(v []string)`

SetAzureResourceGroupContains sets AzureResourceGroupContains field to given value.

### HasAzureResourceGroupContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAzureResourceGroupContains() bool`

HasAzureResourceGroupContains returns a boolean if a field has been set.

### GetCoreCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountBetween() string`

GetCoreCountBetween returns the CoreCountBetween field if non-nil, zero value otherwise.

### GetCoreCountBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountBetweenOk() (*string, bool)`

GetCoreCountBetweenOk returns a tuple with the CoreCountBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCoreCountBetween(v string)`

SetCoreCountBetween sets CoreCountBetween field to given value.

### HasCoreCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCoreCountBetween() bool`

HasCoreCountBetween returns a boolean if a field has been set.

### GetIsUninstalled

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsUninstalled() []bool`

GetIsUninstalled returns the IsUninstalled field if non-nil, zero value otherwise.

### GetIsUninstalledOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetIsUninstalledOk() (*[]bool, bool)`

GetIsUninstalledOk returns a tuple with the IsUninstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUninstalled

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetIsUninstalled(v []bool)`

SetIsUninstalled sets IsUninstalled field to given value.

### HasIsUninstalled

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasIsUninstalled() bool`

HasIsUninstalled returns a boolean if a field has been set.

### GetFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilterId() string`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilterIdOk() (*string, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetFilterId(v string)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetCpuIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuIdContains() []string`

GetCpuIdContains returns the CpuIdContains field if non-nil, zero value otherwise.

### GetCpuIdContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuIdContainsOk() (*[]string, bool)`

GetCpuIdContainsOk returns a tuple with the CpuIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuIdContains(v []string)`

SetCpuIdContains sets CpuIdContains field to given value.

### HasCpuIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuIdContains() bool`

HasCpuIdContains returns a boolean if a field has been set.

### GetK8sTypeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sTypeContains() []string`

GetK8sTypeContains returns the K8sTypeContains field if non-nil, zero value otherwise.

### GetK8sTypeContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sTypeContainsOk() (*[]string, bool)`

GetK8sTypeContainsOk returns a tuple with the K8sTypeContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sTypeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetK8sTypeContains(v []string)`

SetK8sTypeContains sets K8sTypeContains field to given value.

### HasK8sTypeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasK8sTypeContains() bool`

HasK8sTypeContains returns a boolean if a field has been set.

### GetCloudProviderNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudProviderNin() []string`

GetCloudProviderNin returns the CloudProviderNin field if non-nil, zero value otherwise.

### GetCloudProviderNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudProviderNinOk() (*[]string, bool)`

GetCloudProviderNinOk returns a tuple with the CloudProviderNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudProviderNin(v []string)`

SetCloudProviderNin sets CloudProviderNin field to given value.

### HasCloudProviderNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudProviderNin() bool`

HasCloudProviderNin returns a boolean if a field has been set.

### GetMitigationMode

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMitigationMode() string`

GetMitigationMode returns the MitigationMode field if non-nil, zero value otherwise.

### GetMitigationModeOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMitigationModeOk() (*string, bool)`

GetMitigationModeOk returns a tuple with the MitigationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationMode

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetMitigationMode(v string)`

SetMitigationMode sets MitigationMode field to given value.

### HasMitigationMode

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasMitigationMode() bool`

HasMitigationMode returns a boolean if a field has been set.

### GetCloudAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudAccountContains() []string`

GetCloudAccountContains returns the CloudAccountContains field if non-nil, zero value otherwise.

### GetCloudAccountContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudAccountContainsOk() (*[]string, bool)`

GetCloudAccountContainsOk returns a tuple with the CloudAccountContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudAccountContains(v []string)`

SetCloudAccountContains sets CloudAccountContains field to given value.

### HasCloudAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudAccountContains() bool`

HasCloudAccountContains returns a boolean if a field has been set.

### GetAdComputerMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerMemberContains() []string`

GetAdComputerMemberContains returns the AdComputerMemberContains field if non-nil, zero value otherwise.

### GetAdComputerMemberContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdComputerMemberContainsOk() (*[]string, bool)`

GetAdComputerMemberContainsOk returns a tuple with the AdComputerMemberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdComputerMemberContains(v []string)`

SetAdComputerMemberContains sets AdComputerMemberContains field to given value.

### HasAdComputerMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdComputerMemberContains() bool`

HasAdComputerMemberContains returns a boolean if a field has been set.

### GetConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetConsoleMigrationStatuses() []string`

GetConsoleMigrationStatuses returns the ConsoleMigrationStatuses field if non-nil, zero value otherwise.

### GetConsoleMigrationStatusesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetConsoleMigrationStatusesOk() (*[]string, bool)`

GetConsoleMigrationStatusesOk returns a tuple with the ConsoleMigrationStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetConsoleMigrationStatuses(v []string)`

SetConsoleMigrationStatuses sets ConsoleMigrationStatuses field to given value.

### HasConsoleMigrationStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasConsoleMigrationStatuses() bool`

HasConsoleMigrationStatuses returns a boolean if a field has been set.

### GetLastActiveDateGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateGt() time.Time`

GetLastActiveDateGt returns the LastActiveDateGt field if non-nil, zero value otherwise.

### GetLastActiveDateGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateGtOk() (*time.Time, bool)`

GetLastActiveDateGtOk returns a tuple with the LastActiveDateGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastActiveDateGt(v time.Time)`

SetLastActiveDateGt sets LastActiveDateGt field to given value.

### HasLastActiveDateGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastActiveDateGt() bool`

HasLastActiveDateGt returns a boolean if a field has been set.

### GetOsArch

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### GetAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentContentUpdateIdContains() []string`

GetAgentContentUpdateIdContains returns the AgentContentUpdateIdContains field if non-nil, zero value otherwise.

### GetAgentContentUpdateIdContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentContentUpdateIdContainsOk() (*[]string, bool)`

GetAgentContentUpdateIdContainsOk returns a tuple with the AgentContentUpdateIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAgentContentUpdateIdContains(v []string)`

SetAgentContentUpdateIdContains sets AgentContentUpdateIdContains field to given value.

### HasAgentContentUpdateIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAgentContentUpdateIdContains() bool`

HasAgentContentUpdateIdContains returns a boolean if a field has been set.

### GetRegisteredAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtGte() time.Time`

GetRegisteredAtGte returns the RegisteredAtGte field if non-nil, zero value otherwise.

### GetRegisteredAtGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtGteOk() (*time.Time, bool)`

GetRegisteredAtGteOk returns a tuple with the RegisteredAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRegisteredAtGte(v time.Time)`

SetRegisteredAtGte sets RegisteredAtGte field to given value.

### HasRegisteredAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRegisteredAtGte() bool`

HasRegisteredAtGte returns a boolean if a field has been set.

### GetCoreCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountGt() int32`

GetCoreCountGt returns the CoreCountGt field if non-nil, zero value otherwise.

### GetCoreCountGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountGtOk() (*int32, bool)`

GetCoreCountGtOk returns a tuple with the CoreCountGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCoreCountGt(v int32)`

SetCoreCountGt sets CoreCountGt field to given value.

### HasCoreCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCoreCountGt() bool`

HasCoreCountGt returns a boolean if a field has been set.

### GetCpuCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountLte() int32`

GetCpuCountLte returns the CpuCountLte field if non-nil, zero value otherwise.

### GetCpuCountLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountLteOk() (*int32, bool)`

GetCpuCountLteOk returns a tuple with the CpuCountLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuCountLte(v int32)`

SetCpuCountLte sets CpuCountLte field to given value.

### HasCpuCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuCountLte() bool`

HasCpuCountLte returns a boolean if a field has been set.

### GetAgentPodNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentPodNameContains() []string`

GetAgentPodNameContains returns the AgentPodNameContains field if non-nil, zero value otherwise.

### GetAgentPodNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentPodNameContainsOk() (*[]string, bool)`

GetAgentPodNameContainsOk returns a tuple with the AgentPodNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPodNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAgentPodNameContains(v []string)`

SetAgentPodNameContains sets AgentPodNameContains field to given value.

### HasAgentPodNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAgentPodNameContains() bool`

HasAgentPodNameContains returns a boolean if a field has been set.

### GetThreatCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtLt() time.Time`

GetThreatCreatedAtLt returns the ThreatCreatedAtLt field if non-nil, zero value otherwise.

### GetThreatCreatedAtLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtLtOk() (*time.Time, bool)`

GetThreatCreatedAtLtOk returns a tuple with the ThreatCreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatCreatedAtLt(v time.Time)`

SetThreatCreatedAtLt sets ThreatCreatedAtLt field to given value.

### HasThreatCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatCreatedAtLt() bool`

HasThreatCreatedAtLt returns a boolean if a field has been set.

### GetDecommissionedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtBetween() string`

GetDecommissionedAtBetween returns the DecommissionedAtBetween field if non-nil, zero value otherwise.

### GetDecommissionedAtBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtBetweenOk() (*string, bool)`

GetDecommissionedAtBetweenOk returns a tuple with the DecommissionedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDecommissionedAtBetween(v string)`

SetDecommissionedAtBetween sets DecommissionedAtBetween field to given value.

### HasDecommissionedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDecommissionedAtBetween() bool`

HasDecommissionedAtBetween returns a boolean if a field has been set.

### GetCoreCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountGte() int32`

GetCoreCountGte returns the CoreCountGte field if non-nil, zero value otherwise.

### GetCoreCountGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountGteOk() (*int32, bool)`

GetCoreCountGteOk returns a tuple with the CoreCountGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCoreCountGte(v int32)`

SetCoreCountGte sets CoreCountGte field to given value.

### HasCoreCountGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCoreCountGte() bool`

HasCoreCountGte returns a boolean if a field has been set.

### GetClusterNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetClusterNameContains() []string`

GetClusterNameContains returns the ClusterNameContains field if non-nil, zero value otherwise.

### GetClusterNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetClusterNameContainsOk() (*[]string, bool)`

GetClusterNameContainsOk returns a tuple with the ClusterNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetClusterNameContains(v []string)`

SetClusterNameContains sets ClusterNameContains field to given value.

### HasClusterNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasClusterNameContains() bool`

HasClusterNameContains returns a boolean if a field has been set.

### GetCpuCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountGt() int32`

GetCpuCountGt returns the CpuCountGt field if non-nil, zero value otherwise.

### GetCpuCountGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountGtOk() (*int32, bool)`

GetCpuCountGtOk returns a tuple with the CpuCountGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuCountGt(v int32)`

SetCpuCountGt sets CpuCountGt field to given value.

### HasCpuCountGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuCountGt() bool`

HasCpuCountGt returns a boolean if a field has been set.

### GetAdQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdQueryContains() []string`

GetAdQueryContains returns the AdQueryContains field if non-nil, zero value otherwise.

### GetAdQueryContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdQueryContainsOk() (*[]string, bool)`

GetAdQueryContainsOk returns a tuple with the AdQueryContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdQueryContains(v []string)`

SetAdQueryContains sets AdQueryContains field to given value.

### HasAdQueryContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdQueryContains() bool`

HasAdQueryContains returns a boolean if a field has been set.

### GetRangerVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerVersions() []string`

GetRangerVersions returns the RangerVersions field if non-nil, zero value otherwise.

### GetRangerVersionsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerVersionsOk() (*[]string, bool)`

GetRangerVersionsOk returns a tuple with the RangerVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRangerVersions(v []string)`

SetRangerVersions sets RangerVersions field to given value.

### HasRangerVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRangerVersions() bool`

HasRangerVersions returns a boolean if a field has been set.

### GetUserActionsNeededNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUserActionsNeededNin() []string`

GetUserActionsNeededNin returns the UserActionsNeededNin field if non-nil, zero value otherwise.

### GetUserActionsNeededNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUserActionsNeededNinOk() (*[]string, bool)`

GetUserActionsNeededNinOk returns a tuple with the UserActionsNeededNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsNeededNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUserActionsNeededNin(v []string)`

SetUserActionsNeededNin sets UserActionsNeededNin field to given value.

### HasUserActionsNeededNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUserActionsNeededNin() bool`

HasUserActionsNeededNin returns a boolean if a field has been set.

### GetThreatCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtGt() time.Time`

GetThreatCreatedAtGt returns the ThreatCreatedAtGt field if non-nil, zero value otherwise.

### GetThreatCreatedAtGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtGtOk() (*time.Time, bool)`

GetThreatCreatedAtGtOk returns a tuple with the ThreatCreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatCreatedAtGt(v time.Time)`

SetThreatCreatedAtGt sets ThreatCreatedAtGt field to given value.

### HasThreatCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatCreatedAtGt() bool`

HasThreatCreatedAtGt returns a boolean if a field has been set.

### GetCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudInstanceSizeContains() []string`

GetCloudInstanceSizeContains returns the CloudInstanceSizeContains field if non-nil, zero value otherwise.

### GetCloudInstanceSizeContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudInstanceSizeContainsOk() (*[]string, bool)`

GetCloudInstanceSizeContainsOk returns a tuple with the CloudInstanceSizeContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudInstanceSizeContains(v []string)`

SetCloudInstanceSizeContains sets CloudInstanceSizeContains field to given value.

### HasCloudInstanceSizeContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudInstanceSizeContains() bool`

HasCloudInstanceSizeContains returns a boolean if a field has been set.

### GetAdQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdQuery() string`

GetAdQuery returns the AdQuery field if non-nil, zero value otherwise.

### GetAdQueryOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdQueryOk() (*string, bool)`

GetAdQueryOk returns a tuple with the AdQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdQuery(v string)`

SetAdQuery sets AdQuery field to given value.

### HasAdQuery

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdQuery() bool`

HasAdQuery returns a boolean if a field has been set.

### GetExternalIpContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetExternalIpContains() []string`

GetExternalIpContains returns the ExternalIpContains field if non-nil, zero value otherwise.

### GetExternalIpContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetExternalIpContainsOk() (*[]string, bool)`

GetExternalIpContainsOk returns a tuple with the ExternalIpContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetExternalIpContains(v []string)`

SetExternalIpContains sets ExternalIpContains field to given value.

### HasExternalIpContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasExternalIpContains() bool`

HasExternalIpContains returns a boolean if a field has been set.

### GetLocationEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationEnabled() []bool`

GetLocationEnabled returns the LocationEnabled field if non-nil, zero value otherwise.

### GetLocationEnabledOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationEnabledOk() (*[]bool, bool)`

GetLocationEnabledOk returns a tuple with the LocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLocationEnabled(v []bool)`

SetLocationEnabled sets LocationEnabled field to given value.

### HasLocationEnabled

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLocationEnabled() bool`

HasLocationEnabled returns a boolean if a field has been set.

### GetTotalMemoryGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryGt() int32`

GetTotalMemoryGt returns the TotalMemoryGt field if non-nil, zero value otherwise.

### GetTotalMemoryGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryGtOk() (*int32, bool)`

GetTotalMemoryGtOk returns a tuple with the TotalMemoryGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTotalMemoryGt(v int32)`

SetTotalMemoryGt sets TotalMemoryGt field to given value.

### HasTotalMemoryGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTotalMemoryGt() bool`

HasTotalMemoryGt returns a boolean if a field has been set.

### GetGatewayIp

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetRegisteredAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtLt() time.Time`

GetRegisteredAtLt returns the RegisteredAtLt field if non-nil, zero value otherwise.

### GetRegisteredAtLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRegisteredAtLtOk() (*time.Time, bool)`

GetRegisteredAtLtOk returns a tuple with the RegisteredAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRegisteredAtLt(v time.Time)`

SetRegisteredAtLt sets RegisteredAtLt field to given value.

### HasRegisteredAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRegisteredAtLt() bool`

HasRegisteredAtLt returns a boolean if a field has been set.

### GetAgentVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentVersions() []string`

GetAgentVersions returns the AgentVersions field if non-nil, zero value otherwise.

### GetAgentVersionsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentVersionsOk() (*[]string, bool)`

GetAgentVersionsOk returns a tuple with the AgentVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAgentVersions(v []string)`

SetAgentVersions sets AgentVersions field to given value.

### HasAgentVersions

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAgentVersions() bool`

HasAgentVersions returns a boolean if a field has been set.

### GetInstallerTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInstallerTypes() []string`

GetInstallerTypes returns the InstallerTypes field if non-nil, zero value otherwise.

### GetInstallerTypesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInstallerTypesOk() (*[]string, bool)`

GetInstallerTypesOk returns a tuple with the InstallerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetInstallerTypes(v []string)`

SetInstallerTypes sets InstallerTypes field to given value.

### HasInstallerTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasInstallerTypes() bool`

HasInstallerTypes returns a boolean if a field has been set.

### GetCoreCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountLte() int32`

GetCoreCountLte returns the CoreCountLte field if non-nil, zero value otherwise.

### GetCoreCountLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountLteOk() (*int32, bool)`

GetCoreCountLteOk returns a tuple with the CoreCountLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCoreCountLte(v int32)`

SetCoreCountLte sets CoreCountLte field to given value.

### HasCoreCountLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCoreCountLte() bool`

HasCoreCountLte returns a boolean if a field has been set.

### GetTotalMemoryBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryBetween() string`

GetTotalMemoryBetween returns the TotalMemoryBetween field if non-nil, zero value otherwise.

### GetTotalMemoryBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTotalMemoryBetweenOk() (*string, bool)`

GetTotalMemoryBetweenOk returns a tuple with the TotalMemoryBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTotalMemoryBetween(v string)`

SetTotalMemoryBetween sets TotalMemoryBetween field to given value.

### HasTotalMemoryBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTotalMemoryBetween() bool`

HasTotalMemoryBetween returns a boolean if a field has been set.

### GetLastActiveDateGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateGte() time.Time`

GetLastActiveDateGte returns the LastActiveDateGte field if non-nil, zero value otherwise.

### GetLastActiveDateGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateGteOk() (*time.Time, bool)`

GetLastActiveDateGteOk returns a tuple with the LastActiveDateGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastActiveDateGte(v time.Time)`

SetLastActiveDateGte sets LastActiveDateGte field to given value.

### HasLastActiveDateGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastActiveDateGte() bool`

HasLastActiveDateGte returns a boolean if a field has been set.

### GetComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerNameContains() []string`

GetComputerNameContains returns the ComputerNameContains field if non-nil, zero value otherwise.

### GetComputerNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerNameContainsOk() (*[]string, bool)`

GetComputerNameContainsOk returns a tuple with the ComputerNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetComputerNameContains(v []string)`

SetComputerNameContains sets ComputerNameContains field to given value.

### HasComputerNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasComputerNameContains() bool`

HasComputerNameContains returns a boolean if a field has been set.

### GetLocationIdsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationIdsNin() []string`

GetLocationIdsNin returns the LocationIdsNin field if non-nil, zero value otherwise.

### GetLocationIdsNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLocationIdsNinOk() (*[]string, bool)`

GetLocationIdsNinOk returns a tuple with the LocationIdsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIdsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLocationIdsNin(v []string)`

SetLocationIdsNin sets LocationIdsNin field to given value.

### HasLocationIdsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLocationIdsNin() bool`

HasLocationIdsNin returns a boolean if a field has been set.

### GetThreatCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtLte() time.Time`

GetThreatCreatedAtLte returns the ThreatCreatedAtLte field if non-nil, zero value otherwise.

### GetThreatCreatedAtLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtLteOk() (*time.Time, bool)`

GetThreatCreatedAtLteOk returns a tuple with the ThreatCreatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatCreatedAtLte(v time.Time)`

SetThreatCreatedAtLte sets ThreatCreatedAtLte field to given value.

### HasThreatCreatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatCreatedAtLte() bool`

HasThreatCreatedAtLte returns a boolean if a field has been set.

### GetHasTags

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetHasTags() bool`

GetHasTags returns the HasTags field if non-nil, zero value otherwise.

### GetHasTagsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetHasTagsOk() (*bool, bool)`

GetHasTagsOk returns a tuple with the HasTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTags

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetHasTags(v bool)`

SetHasTags sets HasTags field to given value.

### HasHasTags

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasHasTags() bool`

HasHasTags returns a boolean if a field has been set.

### GetUpdatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtGte() time.Time`

GetUpdatedAtGte returns the UpdatedAtGte field if non-nil, zero value otherwise.

### GetUpdatedAtGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtGteOk() (*time.Time, bool)`

GetUpdatedAtGteOk returns a tuple with the UpdatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUpdatedAtGte(v time.Time)`

SetUpdatedAtGte sets UpdatedAtGte field to given value.

### HasUpdatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUpdatedAtGte() bool`

HasUpdatedAtGte returns a boolean if a field has been set.

### GetThreatResolved

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatResolved() bool`

GetThreatResolved returns the ThreatResolved field if non-nil, zero value otherwise.

### GetThreatResolvedOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatResolvedOk() (*bool, bool)`

GetThreatResolvedOk returns a tuple with the ThreatResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatResolved

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatResolved(v bool)`

SetThreatResolved sets ThreatResolved field to given value.

### HasThreatResolved

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatResolved() bool`

HasThreatResolved returns a boolean if a field has been set.

### GetDecommissionedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtGt() time.Time`

GetDecommissionedAtGt returns the DecommissionedAtGt field if non-nil, zero value otherwise.

### GetDecommissionedAtGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtGtOk() (*time.Time, bool)`

GetDecommissionedAtGtOk returns a tuple with the DecommissionedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDecommissionedAtGt(v time.Time)`

SetDecommissionedAtGt sets DecommissionedAtGt field to given value.

### HasDecommissionedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDecommissionedAtGt() bool`

HasDecommissionedAtGt returns a boolean if a field has been set.

### GetThreatHidden

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatHidden() bool`

GetThreatHidden returns the ThreatHidden field if non-nil, zero value otherwise.

### GetThreatHiddenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatHiddenOk() (*bool, bool)`

GetThreatHiddenOk returns a tuple with the ThreatHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatHidden

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatHidden(v bool)`

SetThreatHidden sets ThreatHidden field to given value.

### HasThreatHidden

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatHidden() bool`

HasThreatHidden returns a boolean if a field has been set.

### GetInfected

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInfected() bool`

GetInfected returns the Infected field if non-nil, zero value otherwise.

### GetInfectedOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetInfectedOk() (*bool, bool)`

GetInfectedOk returns a tuple with the Infected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfected

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetInfected(v bool)`

SetInfected sets Infected field to given value.

### HasInfected

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasInfected() bool`

HasInfected returns a boolean if a field has been set.

### GetUuidContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuidContains() []string`

GetUuidContains returns the UuidContains field if non-nil, zero value otherwise.

### GetUuidContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuidContainsOk() (*[]string, bool)`

GetUuidContainsOk returns a tuple with the UuidContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuidContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUuidContains(v []string)`

SetUuidContains sets UuidContains field to given value.

### HasUuidContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUuidContains() bool`

HasUuidContains returns a boolean if a field has been set.

### GetNetworkStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkStatusesNin() []string`

GetNetworkStatusesNin returns the NetworkStatusesNin field if non-nil, zero value otherwise.

### GetNetworkStatusesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkStatusesNinOk() (*[]string, bool)`

GetNetworkStatusesNinOk returns a tuple with the NetworkStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkStatusesNin(v []string)`

SetNetworkStatusesNin sets NetworkStatusesNin field to given value.

### HasNetworkStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkStatusesNin() bool`

HasNetworkStatusesNin returns a boolean if a field has been set.

### GetCloudImageContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudImageContains() []string`

GetCloudImageContains returns the CloudImageContains field if non-nil, zero value otherwise.

### GetCloudImageContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudImageContainsOk() (*[]string, bool)`

GetCloudImageContainsOk returns a tuple with the CloudImageContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudImageContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudImageContains(v []string)`

SetCloudImageContains sets CloudImageContains field to given value.

### HasCloudImageContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudImageContains() bool`

HasCloudImageContains returns a boolean if a field has been set.

### GetSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetSiteIds() []string`

GetSiteIds returns the SiteIds field if non-nil, zero value otherwise.

### GetSiteIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetSiteIdsOk() (*[]string, bool)`

GetSiteIdsOk returns a tuple with the SiteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetSiteIds(v []string)`

SetSiteIds sets SiteIds field to given value.

### HasSiteIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasSiteIds() bool`

HasSiteIds returns a boolean if a field has been set.

### GetRangerStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatus() string`

GetRangerStatus returns the RangerStatus field if non-nil, zero value otherwise.

### GetRangerStatusOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatusOk() (*string, bool)`

GetRangerStatusOk returns a tuple with the RangerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRangerStatus(v string)`

SetRangerStatus sets RangerStatus field to given value.

### HasRangerStatus

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRangerStatus() bool`

HasRangerStatus returns a boolean if a field has been set.

### GetDomainsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDomainsNin() []string`

GetDomainsNin returns the DomainsNin field if non-nil, zero value otherwise.

### GetDomainsNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDomainsNinOk() (*[]string, bool)`

GetDomainsNinOk returns a tuple with the DomainsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDomainsNin(v []string)`

SetDomainsNin sets DomainsNin field to given value.

### HasDomainsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDomainsNin() bool`

HasDomainsNin returns a boolean if a field has been set.

### GetThreatCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtGte() time.Time`

GetThreatCreatedAtGte returns the ThreatCreatedAtGte field if non-nil, zero value otherwise.

### GetThreatCreatedAtGteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtGteOk() (*time.Time, bool)`

GetThreatCreatedAtGteOk returns a tuple with the ThreatCreatedAtGte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatCreatedAtGte(v time.Time)`

SetThreatCreatedAtGte sets ThreatCreatedAtGte field to given value.

### HasThreatCreatedAtGte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatCreatedAtGte() bool`

HasThreatCreatedAtGte returns a boolean if a field has been set.

### GetCsvFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCsvFilterId() string`

GetCsvFilterId returns the CsvFilterId field if non-nil, zero value otherwise.

### GetCsvFilterIdOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCsvFilterIdOk() (*string, bool)`

GetCsvFilterIdOk returns a tuple with the CsvFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCsvFilterId(v string)`

SetCsvFilterId sets CsvFilterId field to given value.

### HasCsvFilterId

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCsvFilterId() bool`

HasCsvFilterId returns a boolean if a field has been set.

### GetAdUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserNameContains() []string`

GetAdUserNameContains returns the AdUserNameContains field if non-nil, zero value otherwise.

### GetAdUserNameContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserNameContainsOk() (*[]string, bool)`

GetAdUserNameContainsOk returns a tuple with the AdUserNameContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdUserNameContains(v []string)`

SetAdUserNameContains sets AdUserNameContains field to given value.

### HasAdUserNameContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdUserNameContains() bool`

HasAdUserNameContains returns a boolean if a field has been set.

### GetAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAppsVulnerabilityStatusesNin() []string`

GetAppsVulnerabilityStatusesNin returns the AppsVulnerabilityStatusesNin field if non-nil, zero value otherwise.

### GetAppsVulnerabilityStatusesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAppsVulnerabilityStatusesNinOk() (*[]string, bool)`

GetAppsVulnerabilityStatusesNinOk returns a tuple with the AppsVulnerabilityStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAppsVulnerabilityStatusesNin(v []string)`

SetAppsVulnerabilityStatusesNin sets AppsVulnerabilityStatusesNin field to given value.

### HasAppsVulnerabilityStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAppsVulnerabilityStatusesNin() bool`

HasAppsVulnerabilityStatusesNin returns a boolean if a field has been set.

### GetUuid

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCoreCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountLt() int32`

GetCoreCountLt returns the CoreCountLt field if non-nil, zero value otherwise.

### GetCoreCountLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCoreCountLtOk() (*int32, bool)`

GetCoreCountLtOk returns a tuple with the CoreCountLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCoreCountLt(v int32)`

SetCoreCountLt sets CoreCountLt field to given value.

### HasCoreCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCoreCountLt() bool`

HasCoreCountLt returns a boolean if a field has been set.

### GetMitigationModeSuspicious

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMitigationModeSuspicious() string`

GetMitigationModeSuspicious returns the MitigationModeSuspicious field if non-nil, zero value otherwise.

### GetMitigationModeSuspiciousOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetMitigationModeSuspiciousOk() (*string, bool)`

GetMitigationModeSuspiciousOk returns a tuple with the MitigationModeSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationModeSuspicious

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetMitigationModeSuspicious(v string)`

SetMitigationModeSuspicious sets MitigationModeSuspicious field to given value.

### HasMitigationModeSuspicious

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasMitigationModeSuspicious() bool`

HasMitigationModeSuspicious returns a boolean if a field has been set.

### GetRangerStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatusesNin() []string`

GetRangerStatusesNin returns the RangerStatusesNin field if non-nil, zero value otherwise.

### GetRangerStatusesNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatusesNinOk() (*[]string, bool)`

GetRangerStatusesNinOk returns a tuple with the RangerStatusesNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRangerStatusesNin(v []string)`

SetRangerStatusesNin sets RangerStatusesNin field to given value.

### HasRangerStatusesNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRangerStatusesNin() bool`

HasRangerStatusesNin returns a boolean if a field has been set.

### GetCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtGt() time.Time`

GetCreatedAtGt returns the CreatedAtGt field if non-nil, zero value otherwise.

### GetCreatedAtGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtGtOk() (*time.Time, bool)`

GetCreatedAtGtOk returns a tuple with the CreatedAtGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCreatedAtGt(v time.Time)`

SetCreatedAtGt sets CreatedAtGt field to given value.

### HasCreatedAtGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCreatedAtGt() bool`

HasCreatedAtGt returns a boolean if a field has been set.

### GetRangerVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerVersionsNin() []string`

GetRangerVersionsNin returns the RangerVersionsNin field if non-nil, zero value otherwise.

### GetRangerVersionsNinOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerVersionsNinOk() (*[]string, bool)`

GetRangerVersionsNinOk returns a tuple with the RangerVersionsNin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRangerVersionsNin(v []string)`

SetRangerVersionsNin sets RangerVersionsNin field to given value.

### HasRangerVersionsNin

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRangerVersionsNin() bool`

HasRangerVersionsNin returns a boolean if a field has been set.

### GetFilteredGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilteredGroupIds() []string`

GetFilteredGroupIds returns the FilteredGroupIds field if non-nil, zero value otherwise.

### GetFilteredGroupIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetFilteredGroupIdsOk() (*[]string, bool)`

GetFilteredGroupIdsOk returns a tuple with the FilteredGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetFilteredGroupIds(v []string)`

SetFilteredGroupIds sets FilteredGroupIds field to given value.

### HasFilteredGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasFilteredGroupIds() bool`

HasFilteredGroupIds returns a boolean if a field has been set.

### GetDecommissionedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtLte() time.Time`

GetDecommissionedAtLte returns the DecommissionedAtLte field if non-nil, zero value otherwise.

### GetDecommissionedAtLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetDecommissionedAtLteOk() (*time.Time, bool)`

GetDecommissionedAtLteOk returns a tuple with the DecommissionedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecommissionedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetDecommissionedAtLte(v time.Time)`

SetDecommissionedAtLte sets DecommissionedAtLte field to given value.

### HasDecommissionedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasDecommissionedAtLte() bool`

HasDecommissionedAtLte returns a boolean if a field has been set.

### GetOperationalStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOperationalStates() []string`

GetOperationalStates returns the OperationalStates field if non-nil, zero value otherwise.

### GetOperationalStatesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOperationalStatesOk() (*[]string, bool)`

GetOperationalStatesOk returns a tuple with the OperationalStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOperationalStates(v []string)`

SetOperationalStates sets OperationalStates field to given value.

### HasOperationalStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOperationalStates() bool`

HasOperationalStates returns a boolean if a field has been set.

### GetOsVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsVersionContains() []string`

GetOsVersionContains returns the OsVersionContains field if non-nil, zero value otherwise.

### GetOsVersionContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsVersionContainsOk() (*[]string, bool)`

GetOsVersionContainsOk returns a tuple with the OsVersionContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOsVersionContains(v []string)`

SetOsVersionContains sets OsVersionContains field to given value.

### HasOsVersionContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOsVersionContains() bool`

HasOsVersionContains returns a boolean if a field has been set.

### GetUpdatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtLte() time.Time`

GetUpdatedAtLte returns the UpdatedAtLte field if non-nil, zero value otherwise.

### GetUpdatedAtLteOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtLteOk() (*time.Time, bool)`

GetUpdatedAtLteOk returns a tuple with the UpdatedAtLte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUpdatedAtLte(v time.Time)`

SetUpdatedAtLte sets UpdatedAtLte field to given value.

### HasUpdatedAtLte

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUpdatedAtLte() bool`

HasUpdatedAtLte returns a boolean if a field has been set.

### GetUpdatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtLt() time.Time`

GetUpdatedAtLt returns the UpdatedAtLt field if non-nil, zero value otherwise.

### GetUpdatedAtLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetUpdatedAtLtOk() (*time.Time, bool)`

GetUpdatedAtLtOk returns a tuple with the UpdatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetUpdatedAtLt(v time.Time)`

SetUpdatedAtLt sets UpdatedAtLt field to given value.

### HasUpdatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasUpdatedAtLt() bool`

HasUpdatedAtLt returns a boolean if a field has been set.

### GetK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sNodeLabelsContains() []string`

GetK8sNodeLabelsContains returns the K8sNodeLabelsContains field if non-nil, zero value otherwise.

### GetK8sNodeLabelsContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetK8sNodeLabelsContainsOk() (*[]string, bool)`

GetK8sNodeLabelsContainsOk returns a tuple with the K8sNodeLabelsContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetK8sNodeLabelsContains(v []string)`

SetK8sNodeLabelsContains sets K8sNodeLabelsContains field to given value.

### HasK8sNodeLabelsContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasK8sNodeLabelsContains() bool`

HasK8sNodeLabelsContains returns a boolean if a field has been set.

### GetAdUserMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserMemberContains() []string`

GetAdUserMemberContains returns the AdUserMemberContains field if non-nil, zero value otherwise.

### GetAdUserMemberContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAdUserMemberContainsOk() (*[]string, bool)`

GetAdUserMemberContainsOk returns a tuple with the AdUserMemberContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAdUserMemberContains(v []string)`

SetAdUserMemberContains sets AdUserMemberContains field to given value.

### HasAdUserMemberContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAdUserMemberContains() bool`

HasAdUserMemberContains returns a boolean if a field has been set.

### GetCloudLocationContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudLocationContains() []string`

GetCloudLocationContains returns the CloudLocationContains field if non-nil, zero value otherwise.

### GetCloudLocationContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudLocationContainsOk() (*[]string, bool)`

GetCloudLocationContainsOk returns a tuple with the CloudLocationContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLocationContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudLocationContains(v []string)`

SetCloudLocationContains sets CloudLocationContains field to given value.

### HasCloudLocationContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudLocationContains() bool`

HasCloudLocationContains returns a boolean if a field has been set.

### GetNetworkStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkStatuses() []string`

GetNetworkStatuses returns the NetworkStatuses field if non-nil, zero value otherwise.

### GetNetworkStatusesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkStatusesOk() (*[]string, bool)`

GetNetworkStatusesOk returns a tuple with the NetworkStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkStatuses(v []string)`

SetNetworkStatuses sets NetworkStatuses field to given value.

### HasNetworkStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkStatuses() bool`

HasNetworkStatuses returns a boolean if a field has been set.

### GetRemoteProfilingStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRemoteProfilingStates() []string`

GetRemoteProfilingStates returns the RemoteProfilingStates field if non-nil, zero value otherwise.

### GetRemoteProfilingStatesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRemoteProfilingStatesOk() (*[]string, bool)`

GetRemoteProfilingStatesOk returns a tuple with the RemoteProfilingStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRemoteProfilingStates(v []string)`

SetRemoteProfilingStates sets RemoteProfilingStates field to given value.

### HasRemoteProfilingStates

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRemoteProfilingStates() bool`

HasRemoteProfilingStates returns a boolean if a field has been set.

### GetNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfacePhysicalContains() []string`

GetNetworkInterfacePhysicalContains returns the NetworkInterfacePhysicalContains field if non-nil, zero value otherwise.

### GetNetworkInterfacePhysicalContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetNetworkInterfacePhysicalContainsOk() (*[]string, bool)`

GetNetworkInterfacePhysicalContainsOk returns a tuple with the NetworkInterfacePhysicalContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetNetworkInterfacePhysicalContains(v []string)`

SetNetworkInterfacePhysicalContains sets NetworkInterfacePhysicalContains field to given value.

### HasNetworkInterfacePhysicalContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasNetworkInterfacePhysicalContains() bool`

HasNetworkInterfacePhysicalContains returns a boolean if a field has been set.

### GetAgentNamespaceContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentNamespaceContains() []string`

GetAgentNamespaceContains returns the AgentNamespaceContains field if non-nil, zero value otherwise.

### GetAgentNamespaceContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetAgentNamespaceContainsOk() (*[]string, bool)`

GetAgentNamespaceContainsOk returns a tuple with the AgentNamespaceContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentNamespaceContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetAgentNamespaceContains(v []string)`

SetAgentNamespaceContains sets AgentNamespaceContains field to given value.

### HasAgentNamespaceContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasAgentNamespaceContains() bool`

HasAgentNamespaceContains returns a boolean if a field has been set.

### GetRangerStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatuses() []string`

GetRangerStatuses returns the RangerStatuses field if non-nil, zero value otherwise.

### GetRangerStatusesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetRangerStatusesOk() (*[]string, bool)`

GetRangerStatusesOk returns a tuple with the RangerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetRangerStatuses(v []string)`

SetRangerStatuses sets RangerStatuses field to given value.

### HasRangerStatuses

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasRangerStatuses() bool`

HasRangerStatuses returns a boolean if a field has been set.

### GetLastActiveDateLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateLt() time.Time`

GetLastActiveDateLt returns the LastActiveDateLt field if non-nil, zero value otherwise.

### GetLastActiveDateLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetLastActiveDateLtOk() (*time.Time, bool)`

GetLastActiveDateLtOk returns a tuple with the LastActiveDateLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDateLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetLastActiveDateLt(v time.Time)`

SetLastActiveDateLt sets LastActiveDateLt field to given value.

### HasLastActiveDateLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasLastActiveDateLt() bool`

HasLastActiveDateLt returns a boolean if a field has been set.

### GetActiveThreatsGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetActiveThreatsGt() int32`

GetActiveThreatsGt returns the ActiveThreatsGt field if non-nil, zero value otherwise.

### GetActiveThreatsGtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetActiveThreatsGtOk() (*int32, bool)`

GetActiveThreatsGtOk returns a tuple with the ActiveThreatsGt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreatsGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetActiveThreatsGt(v int32)`

SetActiveThreatsGt sets ActiveThreatsGt field to given value.

### HasActiveThreatsGt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasActiveThreatsGt() bool`

HasActiveThreatsGt returns a boolean if a field has been set.

### GetCloudInstanceIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudInstanceIdContains() []string`

GetCloudInstanceIdContains returns the CloudInstanceIdContains field if non-nil, zero value otherwise.

### GetCloudInstanceIdContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCloudInstanceIdContainsOk() (*[]string, bool)`

GetCloudInstanceIdContainsOk returns a tuple with the CloudInstanceIdContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInstanceIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCloudInstanceIdContains(v []string)`

SetCloudInstanceIdContains sets CloudInstanceIdContains field to given value.

### HasCloudInstanceIdContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCloudInstanceIdContains() bool`

HasCloudInstanceIdContains returns a boolean if a field has been set.

### GetCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtLt() time.Time`

GetCreatedAtLt returns the CreatedAtLt field if non-nil, zero value otherwise.

### GetCreatedAtLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtLtOk() (*time.Time, bool)`

GetCreatedAtLtOk returns a tuple with the CreatedAtLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCreatedAtLt(v time.Time)`

SetCreatedAtLt sets CreatedAtLt field to given value.

### HasCreatedAtLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCreatedAtLt() bool`

HasCreatedAtLt returns a boolean if a field has been set.

### GetCpuCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountBetween() string`

GetCpuCountBetween returns the CpuCountBetween field if non-nil, zero value otherwise.

### GetCpuCountBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountBetweenOk() (*string, bool)`

GetCpuCountBetweenOk returns a tuple with the CpuCountBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuCountBetween(v string)`

SetCpuCountBetween sets CpuCountBetween field to given value.

### HasCpuCountBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuCountBetween() bool`

HasCpuCountBetween returns a boolean if a field has been set.

### GetGcpServiceAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGcpServiceAccountContains() []string`

GetGcpServiceAccountContains returns the GcpServiceAccountContains field if non-nil, zero value otherwise.

### GetGcpServiceAccountContainsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGcpServiceAccountContainsOk() (*[]string, bool)`

GetGcpServiceAccountContainsOk returns a tuple with the GcpServiceAccountContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetGcpServiceAccountContains(v []string)`

SetGcpServiceAccountContains sets GcpServiceAccountContains field to given value.

### HasGcpServiceAccountContains

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasGcpServiceAccountContains() bool`

HasGcpServiceAccountContains returns a boolean if a field has been set.

### GetThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtBetween() string`

GetThreatCreatedAtBetween returns the ThreatCreatedAtBetween field if non-nil, zero value otherwise.

### GetThreatCreatedAtBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetThreatCreatedAtBetweenOk() (*string, bool)`

GetThreatCreatedAtBetweenOk returns a tuple with the ThreatCreatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetThreatCreatedAtBetween(v string)`

SetThreatCreatedAtBetween sets ThreatCreatedAtBetween field to given value.

### HasThreatCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasThreatCreatedAtBetween() bool`

HasThreatCreatedAtBetween returns a boolean if a field has been set.

### GetOsTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsTypes() []string`

GetOsTypes returns the OsTypes field if non-nil, zero value otherwise.

### GetOsTypesOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetOsTypesOk() (*[]string, bool)`

GetOsTypesOk returns a tuple with the OsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetOsTypes(v []string)`

SetOsTypes sets OsTypes field to given value.

### HasOsTypes

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasOsTypes() bool`

HasOsTypes returns a boolean if a field has been set.

### GetCpuCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountLt() int32`

GetCpuCountLt returns the CpuCountLt field if non-nil, zero value otherwise.

### GetCpuCountLtOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCpuCountLtOk() (*int32, bool)`

GetCpuCountLtOk returns a tuple with the CpuCountLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCpuCountLt(v int32)`

SetCpuCountLt sets CpuCountLt field to given value.

### HasCpuCountLt

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCpuCountLt() bool`

HasCpuCountLt returns a boolean if a field has been set.

### GetGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetComputerNameLike

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerNameLike() string`

GetComputerNameLike returns the ComputerNameLike field if non-nil, zero value otherwise.

### GetComputerNameLikeOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetComputerNameLikeOk() (*string, bool)`

GetComputerNameLikeOk returns a tuple with the ComputerNameLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerNameLike

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetComputerNameLike(v string)`

SetComputerNameLike sets ComputerNameLike field to given value.

### HasComputerNameLike

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasComputerNameLike() bool`

HasComputerNameLike returns a boolean if a field has been set.

### GetTagsData

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTagsData() string`

GetTagsData returns the TagsData field if non-nil, zero value otherwise.

### GetTagsDataOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetTagsDataOk() (*string, bool)`

GetTagsDataOk returns a tuple with the TagsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsData

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetTagsData(v string)`

SetTagsData sets TagsData field to given value.

### HasTagsData

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasTagsData() bool`

HasTagsData returns a boolean if a field has been set.

### GetCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtBetween() string`

GetCreatedAtBetween returns the CreatedAtBetween field if non-nil, zero value otherwise.

### GetCreatedAtBetweenOk

`func (o *AgentsSchemasAgentsActionSchemaFilter) GetCreatedAtBetweenOk() (*string, bool)`

GetCreatedAtBetweenOk returns a tuple with the CreatedAtBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) SetCreatedAtBetween(v string)`

SetCreatedAtBetween sets CreatedAtBetween field to given value.

### HasCreatedAtBetween

`func (o *AgentsSchemasAgentsActionSchemaFilter) HasCreatedAtBetween() bool`

HasCreatedAtBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


