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

// CoreJobUpdate struct for CoreJobUpdate
type CoreJobUpdate struct {
	Id *int32 `json:"id,omitempty"`
	JobUpdate *CoreJobUpdate1 `json:"job-update,omitempty"`
}

// NewCoreJobUpdate instantiates a new CoreJobUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreJobUpdate() *CoreJobUpdate {
	this := CoreJobUpdate{}
	var jobUpdate CoreJobUpdate1 = {}
	this.JobUpdate = &jobUpdate
	return &this
}

// NewCoreJobUpdateWithDefaults instantiates a new CoreJobUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreJobUpdateWithDefaults() *CoreJobUpdate {
	this := CoreJobUpdate{}
	var jobUpdate CoreJobUpdate1 = {}
	this.JobUpdate = &jobUpdate
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreJobUpdate) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreJobUpdate) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreJobUpdate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CoreJobUpdate) SetId(v int32) {
	o.Id = &v
}

// GetJobUpdate returns the JobUpdate field value if set, zero value otherwise.
func (o *CoreJobUpdate) GetJobUpdate() CoreJobUpdate1 {
	if o == nil || o.JobUpdate == nil {
		var ret CoreJobUpdate1
		return ret
	}
	return *o.JobUpdate
}

// GetJobUpdateOk returns a tuple with the JobUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreJobUpdate) GetJobUpdateOk() (*CoreJobUpdate1, bool) {
	if o == nil || o.JobUpdate == nil {
		return nil, false
	}
	return o.JobUpdate, true
}

// HasJobUpdate returns a boolean if a field has been set.
func (o *CoreJobUpdate) HasJobUpdate() bool {
	if o != nil && o.JobUpdate != nil {
		return true
	}

	return false
}

// SetJobUpdate gets a reference to the given CoreJobUpdate1 and assigns it to the JobUpdate field.
func (o *CoreJobUpdate) SetJobUpdate(v CoreJobUpdate1) {
	o.JobUpdate = &v
}

func (o CoreJobUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.JobUpdate != nil {
		toSerialize["job-update"] = o.JobUpdate
	}
	return json.Marshal(toSerialize)
}

type NullableCoreJobUpdate struct {
	value *CoreJobUpdate
	isSet bool
}

func (v NullableCoreJobUpdate) Get() *CoreJobUpdate {
	return v.value
}

func (v *NullableCoreJobUpdate) Set(val *CoreJobUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreJobUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreJobUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreJobUpdate(val *CoreJobUpdate) *NullableCoreJobUpdate {
	return &NullableCoreJobUpdate{value: val, isSet: true}
}

func (v NullableCoreJobUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreJobUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


