# AgentsSchemasAgentViewSchemaMany200Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **NullableString** | Pass this value as \&quot;cursor\&quot; on your next request, to get the next page of results (Will be \&quot;null\&quot; when last page reached) | [optional] 
**TotalItems** | **int32** | Total number of items found matching your query | 

## Methods

### NewAgentsSchemasAgentViewSchemaMany200Pagination

`func NewAgentsSchemasAgentViewSchemaMany200Pagination(totalItems int32, ) *AgentsSchemasAgentViewSchemaMany200Pagination`

NewAgentsSchemasAgentViewSchemaMany200Pagination instantiates a new AgentsSchemasAgentViewSchemaMany200Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentViewSchemaMany200PaginationWithDefaults

`func NewAgentsSchemasAgentViewSchemaMany200PaginationWithDefaults() *AgentsSchemasAgentViewSchemaMany200Pagination`

NewAgentsSchemasAgentViewSchemaMany200PaginationWithDefaults instantiates a new AgentsSchemasAgentViewSchemaMany200Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTotalItems

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *AgentsSchemasAgentViewSchemaMany200Pagination) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


