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

// PoolDatasetEncryptionSummary struct for PoolDatasetEncryptionSummary
type PoolDatasetEncryptionSummary struct {
	Id *string `json:"id,omitempty"`
	EncryptionRootSummaryOptions *PoolDatasetEncryptionSummary1 `json:"encryption_root_summary_options,omitempty"`
}

// NewPoolDatasetEncryptionSummary instantiates a new PoolDatasetEncryptionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetEncryptionSummary() *PoolDatasetEncryptionSummary {
	this := PoolDatasetEncryptionSummary{}
	var encryptionRootSummaryOptions PoolDatasetEncryptionSummary1 = {}
	this.EncryptionRootSummaryOptions = &encryptionRootSummaryOptions
	return &this
}

// NewPoolDatasetEncryptionSummaryWithDefaults instantiates a new PoolDatasetEncryptionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetEncryptionSummaryWithDefaults() *PoolDatasetEncryptionSummary {
	this := PoolDatasetEncryptionSummary{}
	var encryptionRootSummaryOptions PoolDatasetEncryptionSummary1 = {}
	this.EncryptionRootSummaryOptions = &encryptionRootSummaryOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolDatasetEncryptionSummary) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetEncryptionSummary) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolDatasetEncryptionSummary) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolDatasetEncryptionSummary) SetId(v string) {
	o.Id = &v
}

// GetEncryptionRootSummaryOptions returns the EncryptionRootSummaryOptions field value if set, zero value otherwise.
func (o *PoolDatasetEncryptionSummary) GetEncryptionRootSummaryOptions() PoolDatasetEncryptionSummary1 {
	if o == nil || o.EncryptionRootSummaryOptions == nil {
		var ret PoolDatasetEncryptionSummary1
		return ret
	}
	return *o.EncryptionRootSummaryOptions
}

// GetEncryptionRootSummaryOptionsOk returns a tuple with the EncryptionRootSummaryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetEncryptionSummary) GetEncryptionRootSummaryOptionsOk() (*PoolDatasetEncryptionSummary1, bool) {
	if o == nil || o.EncryptionRootSummaryOptions == nil {
		return nil, false
	}
	return o.EncryptionRootSummaryOptions, true
}

// HasEncryptionRootSummaryOptions returns a boolean if a field has been set.
func (o *PoolDatasetEncryptionSummary) HasEncryptionRootSummaryOptions() bool {
	if o != nil && o.EncryptionRootSummaryOptions != nil {
		return true
	}

	return false
}

// SetEncryptionRootSummaryOptions gets a reference to the given PoolDatasetEncryptionSummary1 and assigns it to the EncryptionRootSummaryOptions field.
func (o *PoolDatasetEncryptionSummary) SetEncryptionRootSummaryOptions(v PoolDatasetEncryptionSummary1) {
	o.EncryptionRootSummaryOptions = &v
}

func (o PoolDatasetEncryptionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EncryptionRootSummaryOptions != nil {
		toSerialize["encryption_root_summary_options"] = o.EncryptionRootSummaryOptions
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetEncryptionSummary struct {
	value *PoolDatasetEncryptionSummary
	isSet bool
}

func (v NullablePoolDatasetEncryptionSummary) Get() *PoolDatasetEncryptionSummary {
	return v.value
}

func (v *NullablePoolDatasetEncryptionSummary) Set(val *PoolDatasetEncryptionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetEncryptionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetEncryptionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetEncryptionSummary(val *PoolDatasetEncryptionSummary) *NullablePoolDatasetEncryptionSummary {
	return &NullablePoolDatasetEncryptionSummary{value: val, isSet: true}
}

func (v NullablePoolDatasetEncryptionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetEncryptionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


