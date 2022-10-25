# NoDataSchema400ErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | Error code | 
**Detail** | **NullableString** | Error detailed description | 
**Title** | **string** | Error general description | 
**Params** | Pointer to **map[string]interface{}** | Extra data that are sent within the error | [optional] 

## Methods

### NewNoDataSchema400ErrorsInner

`func NewNoDataSchema400ErrorsInner(code int32, detail NullableString, title string, ) *NoDataSchema400ErrorsInner`

NewNoDataSchema400ErrorsInner instantiates a new NoDataSchema400ErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoDataSchema400ErrorsInnerWithDefaults

`func NewNoDataSchema400ErrorsInnerWithDefaults() *NoDataSchema400ErrorsInner`

NewNoDataSchema400ErrorsInnerWithDefaults instantiates a new NoDataSchema400ErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NoDataSchema400ErrorsInner) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NoDataSchema400ErrorsInner) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NoDataSchema400ErrorsInner) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDetail

`func (o *NoDataSchema400ErrorsInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *NoDataSchema400ErrorsInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *NoDataSchema400ErrorsInner) SetDetail(v string)`

SetDetail sets Detail field to given value.


### SetDetailNil

`func (o *NoDataSchema400ErrorsInner) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *NoDataSchema400ErrorsInner) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetTitle

`func (o *NoDataSchema400ErrorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NoDataSchema400ErrorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NoDataSchema400ErrorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetParams

`func (o *NoDataSchema400ErrorsInner) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NoDataSchema400ErrorsInner) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NoDataSchema400ErrorsInner) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *NoDataSchema400ErrorsInner) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


