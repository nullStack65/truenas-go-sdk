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

// PoolDetach1 struct for PoolDetach1
type PoolDetach1 struct {
	Label *string `json:"label,omitempty"`
}

// NewPoolDetach1 instantiates a new PoolDetach1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDetach1() *PoolDetach1 {
	this := PoolDetach1{}
	return &this
}

// NewPoolDetach1WithDefaults instantiates a new PoolDetach1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDetach1WithDefaults() *PoolDetach1 {
	this := PoolDetach1{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PoolDetach1) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDetach1) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PoolDetach1) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PoolDetach1) SetLabel(v string) {
	o.Label = &v
}

func (o PoolDetach1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDetach1 struct {
	value *PoolDetach1
	isSet bool
}

func (v NullablePoolDetach1) Get() *PoolDetach1 {
	return v.value
}

func (v *NullablePoolDetach1) Set(val *PoolDetach1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDetach1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDetach1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDetach1(val *PoolDetach1) *NullablePoolDetach1 {
	return &NullablePoolDetach1{value: val, isSet: true}
}

func (v NullablePoolDetach1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDetach1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


