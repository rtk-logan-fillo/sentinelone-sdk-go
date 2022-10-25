/*
S1 MGMT API

SentinelOne Management Console API specification. Changes made to original schema (saved as sentinelone_original.json) - Converted to yaml and intially removed all paths, definitions, and tags - Added security definition APIToken - Added global security APIToken - Added tags Agents, Agent Actions - Added /agents path and required definitions - Added /agents/actions/disconnect and required definitions - Added /agents/actions/connect and required definitions - Added operation id for all paths (for more readable generation) - Removed minimum value attribute if value is of type string (to prevents validation warnings) - Removed firstFullModeTime parameter in GetAgents (breaks generation due to malformed datetime serverside)

API version: 2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sentinelone

import (
	"encoding/json"
)

// AgentsSchemasAgentViewSchemaMany200 struct for AgentsSchemasAgentViewSchemaMany200
type AgentsSchemasAgentViewSchemaMany200 struct {
	// Errors
	Errors []map[string]interface{} `json:"errors,omitempty"`
	// Response data
	Data []AgentsSchemasAgentViewSchemaMany200DataInner `json:"data,omitempty"`
	Pagination AgentsSchemasAgentViewSchemaMany200Pagination `json:"pagination"`
}

// NewAgentsSchemasAgentViewSchemaMany200 instantiates a new AgentsSchemasAgentViewSchemaMany200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentsSchemasAgentViewSchemaMany200(pagination AgentsSchemasAgentViewSchemaMany200Pagination) *AgentsSchemasAgentViewSchemaMany200 {
	this := AgentsSchemasAgentViewSchemaMany200{}
	this.Pagination = pagination
	return &this
}

// NewAgentsSchemasAgentViewSchemaMany200WithDefaults instantiates a new AgentsSchemasAgentViewSchemaMany200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentsSchemasAgentViewSchemaMany200WithDefaults() *AgentsSchemasAgentViewSchemaMany200 {
	this := AgentsSchemasAgentViewSchemaMany200{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AgentsSchemasAgentViewSchemaMany200) GetErrors() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentsSchemasAgentViewSchemaMany200) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AgentsSchemasAgentViewSchemaMany200) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *AgentsSchemasAgentViewSchemaMany200) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AgentsSchemasAgentViewSchemaMany200) GetData() []AgentsSchemasAgentViewSchemaMany200DataInner {
	if o == nil || o.Data == nil {
		var ret []AgentsSchemasAgentViewSchemaMany200DataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsSchemasAgentViewSchemaMany200) GetDataOk() ([]AgentsSchemasAgentViewSchemaMany200DataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AgentsSchemasAgentViewSchemaMany200) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AgentsSchemasAgentViewSchemaMany200DataInner and assigns it to the Data field.
func (o *AgentsSchemasAgentViewSchemaMany200) SetData(v []AgentsSchemasAgentViewSchemaMany200DataInner) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *AgentsSchemasAgentViewSchemaMany200) GetPagination() AgentsSchemasAgentViewSchemaMany200Pagination {
	if o == nil {
		var ret AgentsSchemasAgentViewSchemaMany200Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *AgentsSchemasAgentViewSchemaMany200) GetPaginationOk() (*AgentsSchemasAgentViewSchemaMany200Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *AgentsSchemasAgentViewSchemaMany200) SetPagination(v AgentsSchemasAgentViewSchemaMany200Pagination) {
	o.Pagination = v
}

func (o AgentsSchemasAgentViewSchemaMany200) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableAgentsSchemasAgentViewSchemaMany200 struct {
	value *AgentsSchemasAgentViewSchemaMany200
	isSet bool
}

func (v NullableAgentsSchemasAgentViewSchemaMany200) Get() *AgentsSchemasAgentViewSchemaMany200 {
	return v.value
}

func (v *NullableAgentsSchemasAgentViewSchemaMany200) Set(val *AgentsSchemasAgentViewSchemaMany200) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentsSchemasAgentViewSchemaMany200) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentsSchemasAgentViewSchemaMany200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentsSchemasAgentViewSchemaMany200(val *AgentsSchemasAgentViewSchemaMany200) *NullableAgentsSchemasAgentViewSchemaMany200 {
	return &NullableAgentsSchemasAgentViewSchemaMany200{value: val, isSet: true}
}

func (v NullableAgentsSchemasAgentViewSchemaMany200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentsSchemasAgentViewSchemaMany200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


