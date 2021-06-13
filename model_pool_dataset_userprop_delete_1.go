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

// PoolDatasetUserpropDelete1 struct for PoolDatasetUserpropDelete1
type PoolDatasetUserpropDelete1 struct {
	Name *string `json:"name,omitempty"`
}

// NewPoolDatasetUserpropDelete1 instantiates a new PoolDatasetUserpropDelete1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetUserpropDelete1() *PoolDatasetUserpropDelete1 {
	this := PoolDatasetUserpropDelete1{}
	return &this
}

// NewPoolDatasetUserpropDelete1WithDefaults instantiates a new PoolDatasetUserpropDelete1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetUserpropDelete1WithDefaults() *PoolDatasetUserpropDelete1 {
	this := PoolDatasetUserpropDelete1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PoolDatasetUserpropDelete1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetUserpropDelete1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PoolDatasetUserpropDelete1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PoolDatasetUserpropDelete1) SetName(v string) {
	o.Name = &v
}

func (o PoolDatasetUserpropDelete1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetUserpropDelete1 struct {
	value *PoolDatasetUserpropDelete1
	isSet bool
}

func (v NullablePoolDatasetUserpropDelete1) Get() *PoolDatasetUserpropDelete1 {
	return v.value
}

func (v *NullablePoolDatasetUserpropDelete1) Set(val *PoolDatasetUserpropDelete1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetUserpropDelete1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetUserpropDelete1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetUserpropDelete1(val *PoolDatasetUserpropDelete1) *NullablePoolDatasetUserpropDelete1 {
	return &NullablePoolDatasetUserpropDelete1{value: val, isSet: true}
}

func (v NullablePoolDatasetUserpropDelete1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetUserpropDelete1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


