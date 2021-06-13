/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// PoolDatasetUserpropUpdate1 struct for PoolDatasetUserpropUpdate1
type PoolDatasetUserpropUpdate1 struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPoolDatasetUserpropUpdate1 instantiates a new PoolDatasetUserpropUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetUserpropUpdate1() *PoolDatasetUserpropUpdate1 {
	this := PoolDatasetUserpropUpdate1{}
	return &this
}

// NewPoolDatasetUserpropUpdate1WithDefaults instantiates a new PoolDatasetUserpropUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetUserpropUpdate1WithDefaults() *PoolDatasetUserpropUpdate1 {
	this := PoolDatasetUserpropUpdate1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolDatasetUserpropUpdate1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetUserpropUpdate1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolDatasetUserpropUpdate1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolDatasetUserpropUpdate1) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PoolDatasetUserpropUpdate1) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetUserpropUpdate1) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PoolDatasetUserpropUpdate1) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PoolDatasetUserpropUpdate1) SetValue(v string) {
	o.Value = &v
}

func (o PoolDatasetUserpropUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetUserpropUpdate1 struct {
	value *PoolDatasetUserpropUpdate1
	isSet bool
}

func (v NullablePoolDatasetUserpropUpdate1) Get() *PoolDatasetUserpropUpdate1 {
	return v.value
}

func (v *NullablePoolDatasetUserpropUpdate1) Set(val *PoolDatasetUserpropUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetUserpropUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetUserpropUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetUserpropUpdate1(val *PoolDatasetUserpropUpdate1) *NullablePoolDatasetUserpropUpdate1 {
	return &NullablePoolDatasetUserpropUpdate1{value: val, isSet: true}
}

func (v NullablePoolDatasetUserpropUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetUserpropUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


