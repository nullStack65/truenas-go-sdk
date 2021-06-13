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

// FailoverUpgrade0 struct for FailoverUpgrade0
type FailoverUpgrade0 struct {
	Train *string `json:"train,omitempty"`
}

// NewFailoverUpgrade0 instantiates a new FailoverUpgrade0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverUpgrade0() *FailoverUpgrade0 {
	this := FailoverUpgrade0{}
	return &this
}

// NewFailoverUpgrade0WithDefaults instantiates a new FailoverUpgrade0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverUpgrade0WithDefaults() *FailoverUpgrade0 {
	this := FailoverUpgrade0{}
	return &this
}

// GetTrain returns the Train field value if set, zero value otherwise.
func (o *FailoverUpgrade0) GetTrain() string {
	if o == nil || o.Train == nil {
		var ret string
		return ret
	}
	return *o.Train
}

// GetTrainOk returns a tuple with the Train field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverUpgrade0) GetTrainOk() (*string, bool) {
	if o == nil || o.Train == nil {
		return nil, false
	}
	return o.Train, true
}

// HasTrain returns a boolean if a field has been set.
func (o *FailoverUpgrade0) HasTrain() bool {
	if o != nil && o.Train != nil {
		return true
	}

	return false
}

// SetTrain gets a reference to the given string and assigns it to the Train field.
func (o *FailoverUpgrade0) SetTrain(v string) {
	o.Train = &v
}

func (o FailoverUpgrade0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Train != nil {
		toSerialize["train"] = o.Train
	}
	return json.Marshal(toSerialize)
}

type NullableFailoverUpgrade0 struct {
	value *FailoverUpgrade0
	isSet bool
}

func (v NullableFailoverUpgrade0) Get() *FailoverUpgrade0 {
	return v.value
}

func (v *NullableFailoverUpgrade0) Set(val *FailoverUpgrade0) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverUpgrade0) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverUpgrade0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverUpgrade0(val *FailoverUpgrade0) *NullableFailoverUpgrade0 {
	return &NullableFailoverUpgrade0{value: val, isSet: true}
}

func (v NullableFailoverUpgrade0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverUpgrade0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


