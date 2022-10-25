# AgentsSchemasAgentsActionSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**AgentsSchemasAgentsActionSchemaFilter**](AgentsSchemasAgentsActionSchemaFilter.md) |  | 
**Data** | Pointer to **map[string]interface{}** | Data | [optional] 

## Methods

### NewAgentsSchemasAgentsActionSchema

`func NewAgentsSchemasAgentsActionSchema(filter AgentsSchemasAgentsActionSchemaFilter, ) *AgentsSchemasAgentsActionSchema`

NewAgentsSchemasAgentsActionSchema instantiates a new AgentsSchemasAgentsActionSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsSchemasAgentsActionSchemaWithDefaults

`func NewAgentsSchemasAgentsActionSchemaWithDefaults() *AgentsSchemasAgentsActionSchema`

NewAgentsSchemasAgentsActionSchemaWithDefaults instantiates a new AgentsSchemasAgentsActionSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AgentsSchemasAgentsActionSchema) GetFilter() AgentsSchemasAgentsActionSchemaFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AgentsSchemasAgentsActionSchema) GetFilterOk() (*AgentsSchemasAgentsActionSchemaFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AgentsSchemasAgentsActionSchema) SetFilter(v AgentsSchemasAgentsActionSchemaFilter)`

SetFilter sets Filter field to given value.


### GetData

`func (o *AgentsSchemasAgentsActionSchema) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentsSchemasAgentsActionSchema) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentsSchemasAgentsActionSchema) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AgentsSchemasAgentsActionSchema) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AgentsSchemasAgentsActionSchema) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AgentsSchemasAgentsActionSchema) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


