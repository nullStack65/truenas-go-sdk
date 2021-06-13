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

// VmStart1 struct for VmStart1
type VmStart1 struct {
	Overcommit *bool `json:"overcommit,omitempty"`
}

// NewVmStart1 instantiates a new VmStart1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmStart1() *VmStart1 {
	this := VmStart1{}
	return &this
}

// NewVmStart1WithDefaults instantiates a new VmStart1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmStart1WithDefaults() *VmStart1 {
	this := VmStart1{}
	return &this
}

// GetOvercommit returns the Overcommit field value if set, zero value otherwise.
func (o *VmStart1) GetOvercommit() bool {
	if o == nil || o.Overcommit == nil {
		var ret bool
		return ret
	}
	return *o.Overcommit
}

// GetOvercommitOk returns a tuple with the Overcommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmStart1) GetOvercommitOk() (*bool, bool) {
	if o == nil || o.Overcommit == nil {
		return nil, false
	}
	return o.Overcommit, true
}

// HasOvercommit returns a boolean if a field has been set.
func (o *VmStart1) HasOvercommit() bool {
	if o != nil && o.Overcommit != nil {
		return true
	}

	return false
}

// SetOvercommit gets a reference to the given bool and assigns it to the Overcommit field.
func (o *VmStart1) SetOvercommit(v bool) {
	o.Overcommit = &v
}

func (o VmStart1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Overcommit != nil {
		toSerialize["overcommit"] = o.Overcommit
	}
	return json.Marshal(toSerialize)
}

type NullableVmStart1 struct {
	value *VmStart1
	isSet bool
}

func (v NullableVmStart1) Get() *VmStart1 {
	return v.value
}

func (v *NullableVmStart1) Set(val *VmStart1) {
	v.value = val
	v.isSet = true
}

func (v NullableVmStart1) IsSet() bool {
	return v.isSet
}

func (v *NullableVmStart1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmStart1(val *VmStart1) *NullableVmStart1 {
	return &NullableVmStart1{value: val, isSet: true}
}

func (v NullableVmStart1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmStart1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


