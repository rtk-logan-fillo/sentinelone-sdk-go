# AgentsSchemasAgentViewSchemaMany200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]map[string]interface{}** | Errors | [optional] 
**Data** | Pointer to [**[]AgentsSchemasAgentViewSchemaMany200DataInner**](AgentsSchemasAgentViewSchemaMany200DataInner.md) | Response data | [optional] 
**Pagination** | [**AgentsSchemasAgentViewSchemaMany200Pagination**](AgentsSchemasAgentViewSchemaMany200Pagination.md) |  | 

## Methods

### NewAgentsSchemasAgentViewSchemaMany200

`func NewAgentsSchemasAgentViewSchemaMany200(pagination AgentsSchemasAgentViewSchemaMany200Pagination, ) *AgentsSchemasAgentViewSchemaMany200`

NewAgentsSchemasAgentViewSchemaMany200 instantiates a new AgentsSchemasAgentViewSchemaMany200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentViewSchemaMany200WithDefaults

`func NewAgentsSchemasAgentViewSchemaMany200WithDefaults() *AgentsSchemasAgentViewSchemaMany200`

NewAgentsSchemasAgentViewSchemaMany200WithDefaults instantiates a new AgentsSchemasAgentViewSchemaMany200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *AgentsSchemasAgentViewSchemaMany200) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AgentsSchemasAgentViewSchemaMany200) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AgentsSchemasAgentViewSchemaMany200) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AgentsSchemasAgentViewSchemaMany200) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *AgentsSchemasAgentViewSchemaMany200) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *AgentsSchemasAgentViewSchemaMany200) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetData

`func (o *AgentsSchemasAgentViewSchemaMany200) GetData() []AgentsSchemasAgentViewSchemaMany200DataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentsSchemasAgentViewSchemaMany200) GetDataOk() (*[]AgentsSchemasAgentViewSchemaMany200DataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentsSchemasAgentViewSchemaMany200) SetData(v []AgentsSchemasAgentViewSchemaMany200DataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AgentsSchemasAgentViewSchemaMany200) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *AgentsSchemasAgentViewSchemaMany200) GetPagination() AgentsSchemasAgentViewSchemaMany200Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AgentsSchemasAgentViewSchemaMany200) GetPaginationOk() (*AgentsSchemasAgentViewSchemaMany200Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AgentsSchemasAgentViewSchemaMany200) SetPagination(v AgentsSchemasAgentViewSchemaMany200Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


