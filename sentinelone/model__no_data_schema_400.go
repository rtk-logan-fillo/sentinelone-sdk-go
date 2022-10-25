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

// NoDataSchema400 struct for NoDataSchema400
type NoDataSchema400 struct {
	// A list of encountered errors
	Errors []NoDataSchema400ErrorsInner `json:"errors,omitempty"`
}

// NewNoDataSchema400 instantiates a new NoDataSchema400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoDataSchema400() *NoDataSchema400 {
	this := NoDataSchema400{}
	return &this
}

// NewNoDataSchema400WithDefaults instantiates a new NoDataSchema400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoDataSchema400WithDefaults() *NoDataSchema400 {
	this := NoDataSchema400{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *NoDataSchema400) GetErrors() []NoDataSchema400ErrorsInner {
	if o == nil || o.Errors == nil {
		var ret []NoDataSchema400ErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoDataSchema400) GetErrorsOk() ([]NoDataSchema400ErrorsInner, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *NoDataSchema400) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []NoDataSchema400ErrorsInner and assigns it to the Errors field.
func (o *NoDataSchema400) SetErrors(v []NoDataSchema400ErrorsInner) {
	o.Errors = v
}

func (o NoDataSchema400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableNoDataSchema400 struct {
	value *NoDataSchema400
	isSet bool
}

func (v NullableNoDataSchema400) Get() *NoDataSchema400 {
	return v.value
}

func (v *NullableNoDataSchema400) Set(val *NoDataSchema400) {
	v.value = val
	v.isSet = true
}

func (v NullableNoDataSchema400) IsSet() bool {
	return v.isSet
}

func (v *NullableNoDataSchema400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoDataSchema400(val *NoDataSchema400) *NullableNoDataSchema400 {
	return &NullableNoDataSchema400{value: val, isSet: true}
}

func (v NullableNoDataSchema400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoDataSchema400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


