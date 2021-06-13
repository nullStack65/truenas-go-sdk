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

// ReplicationRestore1 struct for ReplicationRestore1
type ReplicationRestore1 struct {
	Name *string `json:"name,omitempty"`
	TargetDataset *string `json:"target_dataset,omitempty"`
}

// NewReplicationRestore1 instantiates a new ReplicationRestore1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicationRestore1() *ReplicationRestore1 {
	this := ReplicationRestore1{}
	return &this
}

// NewReplicationRestore1WithDefaults instantiates a new ReplicationRestore1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicationRestore1WithDefaults() *ReplicationRestore1 {
	this := ReplicationRestore1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReplicationRestore1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationRestore1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReplicationRestore1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReplicationRestore1) SetName(v string) {
	o.Name = &v
}

// GetTargetDataset returns the TargetDataset field value if set, zero value otherwise.
func (o *ReplicationRestore1) GetTargetDataset() string {
	if o == nil || o.TargetDataset == nil {
		var ret string
		return ret
	}
	return *o.TargetDataset
}

// GetTargetDatasetOk returns a tuple with the TargetDataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicationRestore1) GetTargetDatasetOk() (*string, bool) {
	if o == nil || o.TargetDataset == nil {
		return nil, false
	}
	return o.TargetDataset, true
}

// HasTargetDataset returns a boolean if a field has been set.
func (o *ReplicationRestore1) HasTargetDataset() bool {
	if o != nil && o.TargetDataset != nil {
		return true
	}

	return false
}

// SetTargetDataset gets a reference to the given string and assigns it to the TargetDataset field.
func (o *ReplicationRestore1) SetTargetDataset(v string) {
	o.TargetDataset = &v
}

func (o ReplicationRestore1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TargetDataset != nil {
		toSerialize["target_dataset"] = o.TargetDataset
	}
	return json.Marshal(toSerialize)
}

type NullableReplicationRestore1 struct {
	value *ReplicationRestore1
	isSet bool
}

func (v NullableReplicationRestore1) Get() *ReplicationRestore1 {
	return v.value
}

func (v *NullableReplicationRestore1) Set(val *ReplicationRestore1) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicationRestore1) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicationRestore1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicationRestore1(val *ReplicationRestore1) *NullableReplicationRestore1 {
	return &NullableReplicationRestore1{value: val, isSet: true}
}

func (v NullableReplicationRestore1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicationRestore1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


