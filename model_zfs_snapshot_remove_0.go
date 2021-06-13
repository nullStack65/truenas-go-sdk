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

// ZfsSnapshotRemove0 struct for ZfsSnapshotRemove0
type ZfsSnapshotRemove0 struct {
	Dataset *string `json:"dataset,omitempty"`
	Name *string `json:"name,omitempty"`
	DeferDelete *bool `json:"defer_delete,omitempty"`
}

// NewZfsSnapshotRemove0 instantiates a new ZfsSnapshotRemove0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZfsSnapshotRemove0() *ZfsSnapshotRemove0 {
	this := ZfsSnapshotRemove0{}
	return &this
}

// NewZfsSnapshotRemove0WithDefaults instantiates a new ZfsSnapshotRemove0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZfsSnapshotRemove0WithDefaults() *ZfsSnapshotRemove0 {
	this := ZfsSnapshotRemove0{}
	return &this
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *ZfsSnapshotRemove0) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotRemove0) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *ZfsSnapshotRemove0) HasDataset() bool {
	if o != nil && o.Dataset != nil {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *ZfsSnapshotRemove0) SetDataset(v string) {
	o.Dataset = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZfsSnapshotRemove0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotRemove0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZfsSnapshotRemove0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZfsSnapshotRemove0) SetName(v string) {
	o.Name = &v
}

// GetDeferDelete returns the DeferDelete field value if set, zero value otherwise.
func (o *ZfsSnapshotRemove0) GetDeferDelete() bool {
	if o == nil || o.DeferDelete == nil {
		var ret bool
		return ret
	}
	return *o.DeferDelete
}

// GetDeferDeleteOk returns a tuple with the DeferDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZfsSnapshotRemove0) GetDeferDeleteOk() (*bool, bool) {
	if o == nil || o.DeferDelete == nil {
		return nil, false
	}
	return o.DeferDelete, true
}

// HasDeferDelete returns a boolean if a field has been set.
func (o *ZfsSnapshotRemove0) HasDeferDelete() bool {
	if o != nil && o.DeferDelete != nil {
		return true
	}

	return false
}

// SetDeferDelete gets a reference to the given bool and assigns it to the DeferDelete field.
func (o *ZfsSnapshotRemove0) SetDeferDelete(v bool) {
	o.DeferDelete = &v
}

func (o ZfsSnapshotRemove0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DeferDelete != nil {
		toSerialize["defer_delete"] = o.DeferDelete
	}
	return json.Marshal(toSerialize)
}

type NullableZfsSnapshotRemove0 struct {
	value *ZfsSnapshotRemove0
	isSet bool
}

func (v NullableZfsSnapshotRemove0) Get() *ZfsSnapshotRemove0 {
	return v.value
}

func (v *NullableZfsSnapshotRemove0) Set(val *ZfsSnapshotRemove0) {
	v.value = val
	v.isSet = true
}

func (v NullableZfsSnapshotRemove0) IsSet() bool {
	return v.isSet
}

func (v *NullableZfsSnapshotRemove0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZfsSnapshotRemove0(val *ZfsSnapshotRemove0) *NullableZfsSnapshotRemove0 {
	return &NullableZfsSnapshotRemove0{value: val, isSet: true}
}

func (v NullableZfsSnapshotRemove0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZfsSnapshotRemove0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


