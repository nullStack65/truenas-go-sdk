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

// ServiceUpdate1 struct for ServiceUpdate1
type ServiceUpdate1 struct {
	Enable *bool `json:"enable,omitempty"`
}

// NewServiceUpdate1 instantiates a new ServiceUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUpdate1() *ServiceUpdate1 {
	this := ServiceUpdate1{}
	return &this
}

// NewServiceUpdate1WithDefaults instantiates a new ServiceUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUpdate1WithDefaults() *ServiceUpdate1 {
	this := ServiceUpdate1{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ServiceUpdate1) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUpdate1) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ServiceUpdate1) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ServiceUpdate1) SetEnable(v bool) {
	o.Enable = &v
}

func (o ServiceUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	return json.Marshal(toSerialize)
}

type NullableServiceUpdate1 struct {
	value *ServiceUpdate1
	isSet bool
}

func (v NullableServiceUpdate1) Get() *ServiceUpdate1 {
	return v.value
}

func (v *NullableServiceUpdate1) Set(val *ServiceUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUpdate1(val *ServiceUpdate1) *NullableServiceUpdate1 {
	return &NullableServiceUpdate1{value: val, isSet: true}
}

func (v NullableServiceUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


