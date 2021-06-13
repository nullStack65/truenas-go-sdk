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

// PoolUpdate1 struct for PoolUpdate1
type PoolUpdate1 struct {
	EncryptionOptions *PoolCreate0EncryptionOptions `json:"encryption_options,omitempty"`
	Topology *PoolCreate0Topology `json:"topology,omitempty"`
	Autotrim *string `json:"autotrim,omitempty"`
}

// NewPoolUpdate1 instantiates a new PoolUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolUpdate1() *PoolUpdate1 {
	this := PoolUpdate1{}
	return &this
}

// NewPoolUpdate1WithDefaults instantiates a new PoolUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolUpdate1WithDefaults() *PoolUpdate1 {
	this := PoolUpdate1{}
	return &this
}

// GetEncryptionOptions returns the EncryptionOptions field value if set, zero value otherwise.
func (o *PoolUpdate1) GetEncryptionOptions() PoolCreate0EncryptionOptions {
	if o == nil || o.EncryptionOptions == nil {
		var ret PoolCreate0EncryptionOptions
		return ret
	}
	return *o.EncryptionOptions
}

// GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUpdate1) GetEncryptionOptionsOk() (*PoolCreate0EncryptionOptions, bool) {
	if o == nil || o.EncryptionOptions == nil {
		return nil, false
	}
	return o.EncryptionOptions, true
}

// HasEncryptionOptions returns a boolean if a field has been set.
func (o *PoolUpdate1) HasEncryptionOptions() bool {
	if o != nil && o.EncryptionOptions != nil {
		return true
	}

	return false
}

// SetEncryptionOptions gets a reference to the given PoolCreate0EncryptionOptions and assigns it to the EncryptionOptions field.
func (o *PoolUpdate1) SetEncryptionOptions(v PoolCreate0EncryptionOptions) {
	o.EncryptionOptions = &v
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *PoolUpdate1) GetTopology() PoolCreate0Topology {
	if o == nil || o.Topology == nil {
		var ret PoolCreate0Topology
		return ret
	}
	return *o.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUpdate1) GetTopologyOk() (*PoolCreate0Topology, bool) {
	if o == nil || o.Topology == nil {
		return nil, false
	}
	return o.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *PoolUpdate1) HasTopology() bool {
	if o != nil && o.Topology != nil {
		return true
	}

	return false
}

// SetTopology gets a reference to the given PoolCreate0Topology and assigns it to the Topology field.
func (o *PoolUpdate1) SetTopology(v PoolCreate0Topology) {
	o.Topology = &v
}

// GetAutotrim returns the Autotrim field value if set, zero value otherwise.
func (o *PoolUpdate1) GetAutotrim() string {
	if o == nil || o.Autotrim == nil {
		var ret string
		return ret
	}
	return *o.Autotrim
}

// GetAutotrimOk returns a tuple with the Autotrim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUpdate1) GetAutotrimOk() (*string, bool) {
	if o == nil || o.Autotrim == nil {
		return nil, false
	}
	return o.Autotrim, true
}

// HasAutotrim returns a boolean if a field has been set.
func (o *PoolUpdate1) HasAutotrim() bool {
	if o != nil && o.Autotrim != nil {
		return true
	}

	return false
}

// SetAutotrim gets a reference to the given string and assigns it to the Autotrim field.
func (o *PoolUpdate1) SetAutotrim(v string) {
	o.Autotrim = &v
}

func (o PoolUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionOptions != nil {
		toSerialize["encryption_options"] = o.EncryptionOptions
	}
	if o.Topology != nil {
		toSerialize["topology"] = o.Topology
	}
	if o.Autotrim != nil {
		toSerialize["autotrim"] = o.Autotrim
	}
	return json.Marshal(toSerialize)
}

type NullablePoolUpdate1 struct {
	value *PoolUpdate1
	isSet bool
}

func (v NullablePoolUpdate1) Get() *PoolUpdate1 {
	return v.value
}

func (v *NullablePoolUpdate1) Set(val *PoolUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolUpdate1(val *PoolUpdate1) *NullablePoolUpdate1 {
	return &NullablePoolUpdate1{value: val, isSet: true}
}

func (v NullablePoolUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


