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

// VmDeviceDelete1 struct for VmDeviceDelete1
type VmDeviceDelete1 struct {
	Zvol *bool `json:"zvol,omitempty"`
	RawFile *bool `json:"raw_file,omitempty"`
}

// NewVmDeviceDelete1 instantiates a new VmDeviceDelete1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmDeviceDelete1() *VmDeviceDelete1 {
	this := VmDeviceDelete1{}
	return &this
}

// NewVmDeviceDelete1WithDefaults instantiates a new VmDeviceDelete1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmDeviceDelete1WithDefaults() *VmDeviceDelete1 {
	this := VmDeviceDelete1{}
	return &this
}

// GetZvol returns the Zvol field value if set, zero value otherwise.
func (o *VmDeviceDelete1) GetZvol() bool {
	if o == nil || o.Zvol == nil {
		var ret bool
		return ret
	}
	return *o.Zvol
}

// GetZvolOk returns a tuple with the Zvol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmDeviceDelete1) GetZvolOk() (*bool, bool) {
	if o == nil || o.Zvol == nil {
		return nil, false
	}
	return o.Zvol, true
}

// HasZvol returns a boolean if a field has been set.
func (o *VmDeviceDelete1) HasZvol() bool {
	if o != nil && o.Zvol != nil {
		return true
	}

	return false
}

// SetZvol gets a reference to the given bool and assigns it to the Zvol field.
func (o *VmDeviceDelete1) SetZvol(v bool) {
	o.Zvol = &v
}

// GetRawFile returns the RawFile field value if set, zero value otherwise.
func (o *VmDeviceDelete1) GetRawFile() bool {
	if o == nil || o.RawFile == nil {
		var ret bool
		return ret
	}
	return *o.RawFile
}

// GetRawFileOk returns a tuple with the RawFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmDeviceDelete1) GetRawFileOk() (*bool, bool) {
	if o == nil || o.RawFile == nil {
		return nil, false
	}
	return o.RawFile, true
}

// HasRawFile returns a boolean if a field has been set.
func (o *VmDeviceDelete1) HasRawFile() bool {
	if o != nil && o.RawFile != nil {
		return true
	}

	return false
}

// SetRawFile gets a reference to the given bool and assigns it to the RawFile field.
func (o *VmDeviceDelete1) SetRawFile(v bool) {
	o.RawFile = &v
}

func (o VmDeviceDelete1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Zvol != nil {
		toSerialize["zvol"] = o.Zvol
	}
	if o.RawFile != nil {
		toSerialize["raw_file"] = o.RawFile
	}
	return json.Marshal(toSerialize)
}

type NullableVmDeviceDelete1 struct {
	value *VmDeviceDelete1
	isSet bool
}

func (v NullableVmDeviceDelete1) Get() *VmDeviceDelete1 {
	return v.value
}

func (v *NullableVmDeviceDelete1) Set(val *VmDeviceDelete1) {
	v.value = val
	v.isSet = true
}

func (v NullableVmDeviceDelete1) IsSet() bool {
	return v.isSet
}

func (v *NullableVmDeviceDelete1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmDeviceDelete1(val *VmDeviceDelete1) *NullableVmDeviceDelete1 {
	return &NullableVmDeviceDelete1{value: val, isSet: true}
}

func (v NullableVmDeviceDelete1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmDeviceDelete1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


