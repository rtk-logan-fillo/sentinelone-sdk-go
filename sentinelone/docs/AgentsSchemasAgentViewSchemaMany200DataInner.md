# AgentsSchemasAgentViewSchemaMany200DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | A reference to the containing site | [optional] 
**StorageName** | Pointer to **string** | Storage Name | [optional] 
**Id** | Pointer to **string** | Agent ID | [optional] 
**AccountName** | Pointer to **string** | Name of the containing account | [optional] 
**AllowRemoteShell** | Pointer to **bool** | Agent is capable and policy enabled for remote shell | [optional] 
**ConsoleMigrationStatus** | Pointer to **string** | What step the agent is at in the process of migrating to another console, if any | [optional] 
**LastIpToMgmt** | Pointer to **string** | The last ip used to connect to the Management console | [optional] 
**UserActionsNeeded** | Pointer to **[]string** | A list of pending user actions. List items possible values: \&quot;none, reboot_needed, user_action_needed, upgrade_needed, incompatible_os, unprotected, user_action_needed_fda, user_action_needed_rs_fda, user_action_needed_network, rebootless_without_dynamic_detection, extended_exclusions_partially_accepted, user_action_needed_bluetooth_per\&quot;. | [optional] 
**EncryptedApplications** | Pointer to **bool** | Disk encryption status | [optional] 
**IsPendingUninstall** | Pointer to **bool** | Agent with a pending uninstall request | [optional] 
**IsUpToDate** | Pointer to **bool** | Indicates if the agent version is up to date | [optional] 
**Domain** | Pointer to **string** | Network domain | [optional] 
**ComputerName** | Pointer to **string** | Computer name | [optional] 
**ScanStartedAt** | Pointer to **time.Time** | Start time of last scan | [optional] 
**ActiveThreats** | Pointer to **int32** | Current number of active threats | [optional] 
**ModelName** | Pointer to **string** | Device model | [optional] 
**RemoteProfilingState** | Pointer to **string** | Agent remote profiling state | [optional] 
**AgentVersion** | Pointer to **string** | Agent version | [optional] 
**GroupIp** | Pointer to **string** | IP Address subnet | [optional] 
**InRemoteShellSession** | Pointer to **bool** | Is the Agent in a remote shell session | [optional] 
**ScanFinishedAt** | Pointer to **time.Time** | Finish time of last scan (If applicable) | [optional] 
**CpuCount** | Pointer to **int32** | Number of CPUs | [optional] 
**RangerVersion** | Pointer to **string** | The version of Ranger | [optional] 
**CoreCount** | Pointer to **int32** | CPU cores | [optional] 
**IsDecommissioned** | Pointer to **bool** | Is Agent decommissioned | [optional] 
**OsRevision** | Pointer to **string** | Os revision | [optional] 
**OperationalStateExpiration** | Pointer to **time.Time** | Agent operational state expiration | [optional] 
**IsActive** | Pointer to **bool** | Indicates if the agent was recently active | [optional] 
**StorageType** | Pointer to **string** | Storage Type | [optional] 
**InstallerType** | Pointer to **string** | Installer package type (file extension) | [optional] 
**GroupUpdatedAt** | Pointer to **time.Time** | Group updated at | [optional] 
**IsUninstalled** | Pointer to **bool** | Indicates if Agent was removed from the device | [optional] 
**DetectionState** | Pointer to **string** | Detection State | [optional] 
**TotalMemory** | Pointer to **int32** | Memory size (MB) | [optional] 
**AppsVulnerabilityStatus** | Pointer to **string** | Apps vulnerability status | [optional] 
**GroupId** | Pointer to **string** | A reference to the containing network group | [optional] 
**ActiveDirectory** | Pointer to [**AgentsSchemasAgentViewSchemaMany200DataInnerActiveDirectory**](AgentsSchemasAgentViewSchemaMany200DataInnerActiveDirectory.md) |  | [optional] 
**ScanStatus** | Pointer to **string** | Last scan status | [optional] 
**OsType** | Pointer to **string** | OS type | [optional] 
**AccountId** | Pointer to **string** | A reference to the containing account | [optional] 
**MitigationMode** | Pointer to **string** | Agent mitigation mode policy | [optional] 
**CloudProviders** | Pointer to [**map[string]AgentsSchemasAgentViewSchemaMany200DataInnerCloudProvidersValue**](AgentsSchemasAgentViewSchemaMany200DataInnerCloudProvidersValue.md) | Cloud providers for this agent | [optional] 
**OsArch** | Pointer to **string** | Os arch | [optional] 
**CpuId** | Pointer to **string** | CPU model | [optional] 
**Tags** | Pointer to [**AgentsSchemasAgentViewSchemaMany200DataInnerTags**](AgentsSchemasAgentViewSchemaMany200DataInnerTags.md) |  | [optional] 
**OsUsername** | Pointer to **string** | Os username | [optional] 
**RemoteProfilingStateExpiration** | Pointer to **time.Time** | Agent remote profiling state expiration inseconds | [optional] 
**OperationalState** | Pointer to **string** | Agent operational state | [optional] 
**MachineType** | Pointer to **string** | Machine type | [optional] 
**ThreatRebootRequired** | Pointer to **bool** | Has at least one threat with at least one mitigation action that is pending reboot to succeed | [optional] 
**ExternalId** | Pointer to **string** | External id set by customer | [optional] 
**LicenseKey** | Pointer to **string** | License key | [optional] 
**ExternalIp** | Pointer to **string** | External IPv4 address | [optional] 
**Locations** | Pointer to [**[]AgentsSchemasAgentViewSchemaMany200DataInnerLocationsInner**](AgentsSchemasAgentViewSchemaMany200DataInnerLocationsInner.md) | A list of locations reported by the Agent | [optional] 
**Infected** | Pointer to **bool** | Indicates if the Agent has active threats | [optional] [readonly] 
**RangerStatus** | Pointer to **string** | Is Agent disabled as a Ranger | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number of the endpoint | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at | [optional] 
**LocationType** | Pointer to **string** | Reported location type | [optional] 
**Uuid** | Pointer to **string** | Agent&#39;s universally unique identifier | [optional] 
**NetworkQuarantineEnabled** | Pointer to **bool** | Network quarantine enabled | [optional] 
**OsStartTime** | Pointer to **time.Time** | Last boot time | [optional] 
**LastLoggedInUserName** | Pointer to **string** | Last logged in user name | [optional] 
**GroupName** | Pointer to **string** | Name of the containing network group | [optional] 
**MitigationModeSuspicious** | Pointer to **string** | Mitigation mode policy for suspicious activity | [optional] 
**PolicyUpdatedAt** | Pointer to **time.Time** | Policy updated at | [optional] 
**LastActiveDate** | Pointer to **time.Time** | Last active date | [optional] 
**NetworkInterfaces** | Pointer to [**[]AgentsSchemasAgentViewSchemaMany200DataInnerNetworkInterfacesInner**](AgentsSchemasAgentViewSchemaMany200DataInnerNetworkInterfacesInner.md) | Device&#39;s network interfaces | [optional] 
**SiteName** | Pointer to **string** | Name of the containing site | [optional] 
**ScanAbortedAt** | Pointer to **time.Time** | Abort time of last scan (If applicable) | [optional] 
**NetworkStatus** | Pointer to **string** | Agent&#39;s network connectivity status | [optional] 
**RegisteredAt** | Pointer to **time.Time** | Time of first registration to management console (similar to createdAt) | [optional] 
**OsName** | Pointer to **string** | Os name | [optional] 
**FirewallEnabled** | Pointer to **bool** | Firewall enabled | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at | [optional] 
**LocationEnabled** | Pointer to **bool** | Location enabled | [optional] 

## Methods

### NewAgentsSchemasAgentViewSchemaMany200DataInner

`func NewAgentsSchemasAgentViewSchemaMany200DataInner() *AgentsSchemasAgentViewSchemaMany200DataInner`

NewAgentsSchemasAgentViewSchemaMany200DataInner instantiates a new AgentsSchemasAgentViewSchemaMany200DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentViewSchemaMany200DataInnerWithDefaults

`func NewAgentsSchemasAgentViewSchemaMany200DataInnerWithDefaults() *AgentsSchemasAgentViewSchemaMany200DataInner`

NewAgentsSchemasAgentViewSchemaMany200DataInnerWithDefaults instantiates a new AgentsSchemasAgentViewSchemaMany200DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStorageName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetStorageName() string`

GetStorageName returns the StorageName field if non-nil, zero value otherwise.

### GetStorageNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetStorageNameOk() (*string, bool)`

GetStorageNameOk returns a tuple with the StorageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetStorageName(v string)`

SetStorageName sets StorageName field to given value.

### HasStorageName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasStorageName() bool`

HasStorageName returns a boolean if a field has been set.

### GetId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAllowRemoteShell

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAllowRemoteShell() bool`

GetAllowRemoteShell returns the AllowRemoteShell field if non-nil, zero value otherwise.

### GetAllowRemoteShellOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAllowRemoteShellOk() (*bool, bool)`

GetAllowRemoteShellOk returns a tuple with the AllowRemoteShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteShell

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetAllowRemoteShell(v bool)`

SetAllowRemoteShell sets AllowRemoteShell field to given value.

### HasAllowRemoteShell

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasAllowRemoteShell() bool`

HasAllowRemoteShell returns a boolean if a field has been set.

### GetConsoleMigrationStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetConsoleMigrationStatus() string`

GetConsoleMigrationStatus returns the ConsoleMigrationStatus field if non-nil, zero value otherwise.

### GetConsoleMigrationStatusOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetConsoleMigrationStatusOk() (*string, bool)`

GetConsoleMigrationStatusOk returns a tuple with the ConsoleMigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleMigrationStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetConsoleMigrationStatus(v string)`

SetConsoleMigrationStatus sets ConsoleMigrationStatus field to given value.

### HasConsoleMigrationStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasConsoleMigrationStatus() bool`

HasConsoleMigrationStatus returns a boolean if a field has been set.

### GetLastIpToMgmt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastIpToMgmt() string`

GetLastIpToMgmt returns the LastIpToMgmt field if non-nil, zero value otherwise.

### GetLastIpToMgmtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastIpToMgmtOk() (*string, bool)`

GetLastIpToMgmtOk returns a tuple with the LastIpToMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIpToMgmt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLastIpToMgmt(v string)`

SetLastIpToMgmt sets LastIpToMgmt field to given value.

### HasLastIpToMgmt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLastIpToMgmt() bool`

HasLastIpToMgmt returns a boolean if a field has been set.

### GetUserActionsNeeded

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUserActionsNeeded() []string`

GetUserActionsNeeded returns the UserActionsNeeded field if non-nil, zero value otherwise.

### GetUserActionsNeededOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUserActionsNeededOk() (*[]string, bool)`

GetUserActionsNeededOk returns a tuple with the UserActionsNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionsNeeded

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetUserActionsNeeded(v []string)`

SetUserActionsNeeded sets UserActionsNeeded field to given value.

### HasUserActionsNeeded

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasUserActionsNeeded() bool`

HasUserActionsNeeded returns a boolean if a field has been set.

### GetEncryptedApplications

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetEncryptedApplications() bool`

GetEncryptedApplications returns the EncryptedApplications field if non-nil, zero value otherwise.

### GetEncryptedApplicationsOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetEncryptedApplicationsOk() (*bool, bool)`

GetEncryptedApplicationsOk returns a tuple with the EncryptedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedApplications

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetEncryptedApplications(v bool)`

SetEncryptedApplications sets EncryptedApplications field to given value.

### HasEncryptedApplications

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasEncryptedApplications() bool`

HasEncryptedApplications returns a boolean if a field has been set.

### GetIsPendingUninstall

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsPendingUninstall() bool`

GetIsPendingUninstall returns the IsPendingUninstall field if non-nil, zero value otherwise.

### GetIsPendingUninstallOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsPendingUninstallOk() (*bool, bool)`

GetIsPendingUninstallOk returns a tuple with the IsPendingUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPendingUninstall

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetIsPendingUninstall(v bool)`

SetIsPendingUninstall sets IsPendingUninstall field to given value.

### HasIsPendingUninstall

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasIsPendingUninstall() bool`

HasIsPendingUninstall returns a boolean if a field has been set.

### GetIsUpToDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsUpToDate() bool`

GetIsUpToDate returns the IsUpToDate field if non-nil, zero value otherwise.

### GetIsUpToDateOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsUpToDateOk() (*bool, bool)`

GetIsUpToDateOk returns a tuple with the IsUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpToDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetIsUpToDate(v bool)`

SetIsUpToDate sets IsUpToDate field to given value.

### HasIsUpToDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasIsUpToDate() bool`

HasIsUpToDate returns a boolean if a field has been set.

### GetDomain

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetComputerName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetComputerName() string`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetComputerNameOk() (*string, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetComputerName(v string)`

SetComputerName sets ComputerName field to given value.

### HasComputerName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasComputerName() bool`

HasComputerName returns a boolean if a field has been set.

### GetScanStartedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanStartedAt() time.Time`

GetScanStartedAt returns the ScanStartedAt field if non-nil, zero value otherwise.

### GetScanStartedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanStartedAtOk() (*time.Time, bool)`

GetScanStartedAtOk returns a tuple with the ScanStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStartedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetScanStartedAt(v time.Time)`

SetScanStartedAt sets ScanStartedAt field to given value.

### HasScanStartedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasScanStartedAt() bool`

HasScanStartedAt returns a boolean if a field has been set.

### GetActiveThreats

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetActiveThreats() int32`

GetActiveThreats returns the ActiveThreats field if non-nil, zero value otherwise.

### GetActiveThreatsOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetActiveThreatsOk() (*int32, bool)`

GetActiveThreatsOk returns a tuple with the ActiveThreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveThreats

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetActiveThreats(v int32)`

SetActiveThreats sets ActiveThreats field to given value.

### HasActiveThreats

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasActiveThreats() bool`

HasActiveThreats returns a boolean if a field has been set.

### GetModelName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetRemoteProfilingState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRemoteProfilingState() string`

GetRemoteProfilingState returns the RemoteProfilingState field if non-nil, zero value otherwise.

### GetRemoteProfilingStateOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRemoteProfilingStateOk() (*string, bool)`

GetRemoteProfilingStateOk returns a tuple with the RemoteProfilingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetRemoteProfilingState(v string)`

SetRemoteProfilingState sets RemoteProfilingState field to given value.

### HasRemoteProfilingState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasRemoteProfilingState() bool`

HasRemoteProfilingState returns a boolean if a field has been set.

### GetAgentVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetGroupIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupIp() string`

GetGroupIp returns the GroupIp field if non-nil, zero value otherwise.

### GetGroupIpOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupIpOk() (*string, bool)`

GetGroupIpOk returns a tuple with the GroupIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetGroupIp(v string)`

SetGroupIp sets GroupIp field to given value.

### HasGroupIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasGroupIp() bool`

HasGroupIp returns a boolean if a field has been set.

### GetInRemoteShellSession

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInRemoteShellSession() bool`

GetInRemoteShellSession returns the InRemoteShellSession field if non-nil, zero value otherwise.

### GetInRemoteShellSessionOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInRemoteShellSessionOk() (*bool, bool)`

GetInRemoteShellSessionOk returns a tuple with the InRemoteShellSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInRemoteShellSession

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetInRemoteShellSession(v bool)`

SetInRemoteShellSession sets InRemoteShellSession field to given value.

### HasInRemoteShellSession

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasInRemoteShellSession() bool`

HasInRemoteShellSession returns a boolean if a field has been set.

### GetScanFinishedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanFinishedAt() time.Time`

GetScanFinishedAt returns the ScanFinishedAt field if non-nil, zero value otherwise.

### GetScanFinishedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanFinishedAtOk() (*time.Time, bool)`

GetScanFinishedAtOk returns a tuple with the ScanFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFinishedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetScanFinishedAt(v time.Time)`

SetScanFinishedAt sets ScanFinishedAt field to given value.

### HasScanFinishedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasScanFinishedAt() bool`

HasScanFinishedAt returns a boolean if a field has been set.

### GetCpuCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetRangerVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRangerVersion() string`

GetRangerVersion returns the RangerVersion field if non-nil, zero value otherwise.

### GetRangerVersionOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRangerVersionOk() (*string, bool)`

GetRangerVersionOk returns a tuple with the RangerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetRangerVersion(v string)`

SetRangerVersion sets RangerVersion field to given value.

### HasRangerVersion

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasRangerVersion() bool`

HasRangerVersion returns a boolean if a field has been set.

### GetCoreCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCoreCount() int32`

GetCoreCount returns the CoreCount field if non-nil, zero value otherwise.

### GetCoreCountOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCoreCountOk() (*int32, bool)`

GetCoreCountOk returns a tuple with the CoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetCoreCount(v int32)`

SetCoreCount sets CoreCount field to given value.

### HasCoreCount

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasCoreCount() bool`

HasCoreCount returns a boolean if a field has been set.

### GetIsDecommissioned

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsDecommissioned() bool`

GetIsDecommissioned returns the IsDecommissioned field if non-nil, zero value otherwise.

### GetIsDecommissionedOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsDecommissionedOk() (*bool, bool)`

GetIsDecommissionedOk returns a tuple with the IsDecommissioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDecommissioned

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetIsDecommissioned(v bool)`

SetIsDecommissioned sets IsDecommissioned field to given value.

### HasIsDecommissioned

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasIsDecommissioned() bool`

HasIsDecommissioned returns a boolean if a field has been set.

### GetOsRevision

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsRevision() string`

GetOsRevision returns the OsRevision field if non-nil, zero value otherwise.

### GetOsRevisionOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsRevisionOk() (*string, bool)`

GetOsRevisionOk returns a tuple with the OsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRevision

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsRevision(v string)`

SetOsRevision sets OsRevision field to given value.

### HasOsRevision

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsRevision() bool`

HasOsRevision returns a boolean if a field has been set.

### GetOperationalStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOperationalStateExpiration() time.Time`

GetOperationalStateExpiration returns the OperationalStateExpiration field if non-nil, zero value otherwise.

### GetOperationalStateExpirationOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOperationalStateExpirationOk() (*time.Time, bool)`

GetOperationalStateExpirationOk returns a tuple with the OperationalStateExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOperationalStateExpiration(v time.Time)`

SetOperationalStateExpiration sets OperationalStateExpiration field to given value.

### HasOperationalStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOperationalStateExpiration() bool`

HasOperationalStateExpiration returns a boolean if a field has been set.

### GetIsActive

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetStorageType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetInstallerType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInstallerType() string`

GetInstallerType returns the InstallerType field if non-nil, zero value otherwise.

### GetInstallerTypeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInstallerTypeOk() (*string, bool)`

GetInstallerTypeOk returns a tuple with the InstallerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallerType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetInstallerType(v string)`

SetInstallerType sets InstallerType field to given value.

### HasInstallerType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasInstallerType() bool`

HasInstallerType returns a boolean if a field has been set.

### GetGroupUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupUpdatedAt() time.Time`

GetGroupUpdatedAt returns the GroupUpdatedAt field if non-nil, zero value otherwise.

### GetGroupUpdatedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupUpdatedAtOk() (*time.Time, bool)`

GetGroupUpdatedAtOk returns a tuple with the GroupUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetGroupUpdatedAt(v time.Time)`

SetGroupUpdatedAt sets GroupUpdatedAt field to given value.

### HasGroupUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasGroupUpdatedAt() bool`

HasGroupUpdatedAt returns a boolean if a field has been set.

### GetIsUninstalled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsUninstalled() bool`

GetIsUninstalled returns the IsUninstalled field if non-nil, zero value otherwise.

### GetIsUninstalledOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetIsUninstalledOk() (*bool, bool)`

GetIsUninstalledOk returns a tuple with the IsUninstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUninstalled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetIsUninstalled(v bool)`

SetIsUninstalled sets IsUninstalled field to given value.

### HasIsUninstalled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasIsUninstalled() bool`

HasIsUninstalled returns a boolean if a field has been set.

### GetDetectionState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetDetectionState() string`

GetDetectionState returns the DetectionState field if non-nil, zero value otherwise.

### GetDetectionStateOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetDetectionStateOk() (*string, bool)`

GetDetectionStateOk returns a tuple with the DetectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetDetectionState(v string)`

SetDetectionState sets DetectionState field to given value.

### HasDetectionState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasDetectionState() bool`

HasDetectionState returns a boolean if a field has been set.

### GetTotalMemory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetTotalMemory() int32`

GetTotalMemory returns the TotalMemory field if non-nil, zero value otherwise.

### GetTotalMemoryOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetTotalMemoryOk() (*int32, bool)`

GetTotalMemoryOk returns a tuple with the TotalMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetTotalMemory(v int32)`

SetTotalMemory sets TotalMemory field to given value.

### HasTotalMemory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasTotalMemory() bool`

HasTotalMemory returns a boolean if a field has been set.

### GetAppsVulnerabilityStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAppsVulnerabilityStatus() string`

GetAppsVulnerabilityStatus returns the AppsVulnerabilityStatus field if non-nil, zero value otherwise.

### GetAppsVulnerabilityStatusOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAppsVulnerabilityStatusOk() (*string, bool)`

GetAppsVulnerabilityStatusOk returns a tuple with the AppsVulnerabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsVulnerabilityStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetAppsVulnerabilityStatus(v string)`

SetAppsVulnerabilityStatus sets AppsVulnerabilityStatus field to given value.

### HasAppsVulnerabilityStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasAppsVulnerabilityStatus() bool`

HasAppsVulnerabilityStatus returns a boolean if a field has been set.

### GetGroupId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetActiveDirectory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetActiveDirectory() AgentsSchemasAgentViewSchemaMany200DataInnerActiveDirectory`

GetActiveDirectory returns the ActiveDirectory field if non-nil, zero value otherwise.

### GetActiveDirectoryOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetActiveDirectoryOk() (*AgentsSchemasAgentViewSchemaMany200DataInnerActiveDirectory, bool)`

GetActiveDirectoryOk returns a tuple with the ActiveDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetActiveDirectory(v AgentsSchemasAgentViewSchemaMany200DataInnerActiveDirectory)`

SetActiveDirectory sets ActiveDirectory field to given value.

### HasActiveDirectory

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasActiveDirectory() bool`

HasActiveDirectory returns a boolean if a field has been set.

### GetScanStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanStatus() string`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanStatusOk() (*string, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetScanStatus(v string)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetOsType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetAccountId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetMitigationMode

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMitigationMode() string`

GetMitigationMode returns the MitigationMode field if non-nil, zero value otherwise.

### GetMitigationModeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMitigationModeOk() (*string, bool)`

GetMitigationModeOk returns a tuple with the MitigationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationMode

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetMitigationMode(v string)`

SetMitigationMode sets MitigationMode field to given value.

### HasMitigationMode

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasMitigationMode() bool`

HasMitigationMode returns a boolean if a field has been set.

### GetCloudProviders

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCloudProviders() map[string]AgentsSchemasAgentViewSchemaMany200DataInnerCloudProvidersValue`

GetCloudProviders returns the CloudProviders field if non-nil, zero value otherwise.

### GetCloudProvidersOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCloudProvidersOk() (*map[string]AgentsSchemasAgentViewSchemaMany200DataInnerCloudProvidersValue, bool)`

GetCloudProvidersOk returns a tuple with the CloudProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviders

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetCloudProviders(v map[string]AgentsSchemasAgentViewSchemaMany200DataInnerCloudProvidersValue)`

SetCloudProviders sets CloudProviders field to given value.

### HasCloudProviders

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasCloudProviders() bool`

HasCloudProviders returns a boolean if a field has been set.

### GetOsArch

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### GetCpuId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCpuId() string`

GetCpuId returns the CpuId field if non-nil, zero value otherwise.

### GetCpuIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCpuIdOk() (*string, bool)`

GetCpuIdOk returns a tuple with the CpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetCpuId(v string)`

SetCpuId sets CpuId field to given value.

### HasCpuId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasCpuId() bool`

HasCpuId returns a boolean if a field has been set.

### GetTags

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetTags() AgentsSchemasAgentViewSchemaMany200DataInnerTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetTagsOk() (*AgentsSchemasAgentViewSchemaMany200DataInnerTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetTags(v AgentsSchemasAgentViewSchemaMany200DataInnerTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOsUsername

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetRemoteProfilingStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRemoteProfilingStateExpiration() time.Time`

GetRemoteProfilingStateExpiration returns the RemoteProfilingStateExpiration field if non-nil, zero value otherwise.

### GetRemoteProfilingStateExpirationOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRemoteProfilingStateExpirationOk() (*time.Time, bool)`

GetRemoteProfilingStateExpirationOk returns a tuple with the RemoteProfilingStateExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteProfilingStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetRemoteProfilingStateExpiration(v time.Time)`

SetRemoteProfilingStateExpiration sets RemoteProfilingStateExpiration field to given value.

### HasRemoteProfilingStateExpiration

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasRemoteProfilingStateExpiration() bool`

HasRemoteProfilingStateExpiration returns a boolean if a field has been set.

### GetOperationalState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOperationalState() string`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOperationalStateOk() (*string, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOperationalState(v string)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetMachineType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetThreatRebootRequired

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetThreatRebootRequired() bool`

GetThreatRebootRequired returns the ThreatRebootRequired field if non-nil, zero value otherwise.

### GetThreatRebootRequiredOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetThreatRebootRequiredOk() (*bool, bool)`

GetThreatRebootRequiredOk returns a tuple with the ThreatRebootRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatRebootRequired

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetThreatRebootRequired(v bool)`

SetThreatRebootRequired sets ThreatRebootRequired field to given value.

### HasThreatRebootRequired

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasThreatRebootRequired() bool`

HasThreatRebootRequired returns a boolean if a field has been set.

### GetExternalId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLicenseKey

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLicenseKey() string`

GetLicenseKey returns the LicenseKey field if non-nil, zero value otherwise.

### GetLicenseKeyOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLicenseKeyOk() (*string, bool)`

GetLicenseKeyOk returns a tuple with the LicenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseKey

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLicenseKey(v string)`

SetLicenseKey sets LicenseKey field to given value.

### HasLicenseKey

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLicenseKey() bool`

HasLicenseKey returns a boolean if a field has been set.

### GetExternalIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetLocations

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocations() []AgentsSchemasAgentViewSchemaMany200DataInnerLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocationsOk() (*[]AgentsSchemasAgentViewSchemaMany200DataInnerLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLocations(v []AgentsSchemasAgentViewSchemaMany200DataInnerLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetInfected

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInfected() bool`

GetInfected returns the Infected field if non-nil, zero value otherwise.

### GetInfectedOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetInfectedOk() (*bool, bool)`

GetInfectedOk returns a tuple with the Infected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfected

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetInfected(v bool)`

SetInfected sets Infected field to given value.

### HasInfected

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasInfected() bool`

HasInfected returns a boolean if a field has been set.

### GetRangerStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRangerStatus() string`

GetRangerStatus returns the RangerStatus field if non-nil, zero value otherwise.

### GetRangerStatusOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRangerStatusOk() (*string, bool)`

GetRangerStatusOk returns a tuple with the RangerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangerStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetRangerStatus(v string)`

SetRangerStatus sets RangerStatus field to given value.

### HasRangerStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasRangerStatus() bool`

HasRangerStatus returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLocationType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetUuid

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkQuarantineEnabled() bool`

GetNetworkQuarantineEnabled returns the NetworkQuarantineEnabled field if non-nil, zero value otherwise.

### GetNetworkQuarantineEnabledOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkQuarantineEnabledOk() (*bool, bool)`

GetNetworkQuarantineEnabledOk returns a tuple with the NetworkQuarantineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetNetworkQuarantineEnabled(v bool)`

SetNetworkQuarantineEnabled sets NetworkQuarantineEnabled field to given value.

### HasNetworkQuarantineEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasNetworkQuarantineEnabled() bool`

HasNetworkQuarantineEnabled returns a boolean if a field has been set.

### GetOsStartTime

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsStartTime() time.Time`

GetOsStartTime returns the OsStartTime field if non-nil, zero value otherwise.

### GetOsStartTimeOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsStartTimeOk() (*time.Time, bool)`

GetOsStartTimeOk returns a tuple with the OsStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsStartTime

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsStartTime(v time.Time)`

SetOsStartTime sets OsStartTime field to given value.

### HasOsStartTime

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsStartTime() bool`

HasOsStartTime returns a boolean if a field has been set.

### GetLastLoggedInUserName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastLoggedInUserName() string`

GetLastLoggedInUserName returns the LastLoggedInUserName field if non-nil, zero value otherwise.

### GetLastLoggedInUserNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastLoggedInUserNameOk() (*string, bool)`

GetLastLoggedInUserNameOk returns a tuple with the LastLoggedInUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInUserName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLastLoggedInUserName(v string)`

SetLastLoggedInUserName sets LastLoggedInUserName field to given value.

### HasLastLoggedInUserName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLastLoggedInUserName() bool`

HasLastLoggedInUserName returns a boolean if a field has been set.

### GetGroupName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetMitigationModeSuspicious

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMitigationModeSuspicious() string`

GetMitigationModeSuspicious returns the MitigationModeSuspicious field if non-nil, zero value otherwise.

### GetMitigationModeSuspiciousOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetMitigationModeSuspiciousOk() (*string, bool)`

GetMitigationModeSuspiciousOk returns a tuple with the MitigationModeSuspicious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigationModeSuspicious

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetMitigationModeSuspicious(v string)`

SetMitigationModeSuspicious sets MitigationModeSuspicious field to given value.

### HasMitigationModeSuspicious

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasMitigationModeSuspicious() bool`

HasMitigationModeSuspicious returns a boolean if a field has been set.

### GetPolicyUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetPolicyUpdatedAt() time.Time`

GetPolicyUpdatedAt returns the PolicyUpdatedAt field if non-nil, zero value otherwise.

### GetPolicyUpdatedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetPolicyUpdatedAtOk() (*time.Time, bool)`

GetPolicyUpdatedAtOk returns a tuple with the PolicyUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetPolicyUpdatedAt(v time.Time)`

SetPolicyUpdatedAt sets PolicyUpdatedAt field to given value.

### HasPolicyUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasPolicyUpdatedAt() bool`

HasPolicyUpdatedAt returns a boolean if a field has been set.

### GetLastActiveDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastActiveDate() time.Time`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLastActiveDateOk() (*time.Time, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLastActiveDate(v time.Time)`

SetLastActiveDate sets LastActiveDate field to given value.

### HasLastActiveDate

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLastActiveDate() bool`

HasLastActiveDate returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkInterfaces() []AgentsSchemasAgentViewSchemaMany200DataInnerNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkInterfacesOk() (*[]AgentsSchemasAgentViewSchemaMany200DataInnerNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetNetworkInterfaces(v []AgentsSchemasAgentViewSchemaMany200DataInnerNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetSiteName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetScanAbortedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanAbortedAt() time.Time`

GetScanAbortedAt returns the ScanAbortedAt field if non-nil, zero value otherwise.

### GetScanAbortedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetScanAbortedAtOk() (*time.Time, bool)`

GetScanAbortedAtOk returns a tuple with the ScanAbortedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAbortedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetScanAbortedAt(v time.Time)`

SetScanAbortedAt sets ScanAbortedAt field to given value.

### HasScanAbortedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasScanAbortedAt() bool`

HasScanAbortedAt returns a boolean if a field has been set.

### GetNetworkStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkStatus() string`

GetNetworkStatus returns the NetworkStatus field if non-nil, zero value otherwise.

### GetNetworkStatusOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetNetworkStatusOk() (*string, bool)`

GetNetworkStatusOk returns a tuple with the NetworkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetNetworkStatus(v string)`

SetNetworkStatus sets NetworkStatus field to given value.

### HasNetworkStatus

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasNetworkStatus() bool`

HasNetworkStatus returns a boolean if a field has been set.

### GetRegisteredAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRegisteredAt() time.Time`

GetRegisteredAt returns the RegisteredAt field if non-nil, zero value otherwise.

### GetRegisteredAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetRegisteredAtOk() (*time.Time, bool)`

GetRegisteredAtOk returns a tuple with the RegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetRegisteredAt(v time.Time)`

SetRegisteredAt sets RegisteredAt field to given value.

### HasRegisteredAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasRegisteredAt() bool`

HasRegisteredAt returns a boolean if a field has been set.

### GetOsName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLocationEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocationEnabled() bool`

GetLocationEnabled returns the LocationEnabled field if non-nil, zero value otherwise.

### GetLocationEnabledOk

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) GetLocationEnabledOk() (*bool, bool)`

GetLocationEnabledOk returns a tuple with the LocationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) SetLocationEnabled(v bool)`

SetLocationEnabled sets LocationEnabled field to given value.

### HasLocationEnabled

`func (o *AgentsSchemasAgentViewSchemaMany200DataInner) HasLocationEnabled() bool`

HasLocationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


