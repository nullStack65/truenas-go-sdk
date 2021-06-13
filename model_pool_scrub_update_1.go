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

// PoolScrubUpdate1 struct for PoolScrubUpdate1
type PoolScrubUpdate1 struct {
	Pool *int32 `json:"pool,omitempty"`
	Threshold *int32 `json:"threshold,omitempty"`
	Description *string `json:"description,omitempty"`
	Schedule *CloudsyncCreate0Schedule `json:"schedule,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewPoolScrubUpdate1 instantiates a new PoolScrubUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolScrubUpdate1() *PoolScrubUpdate1 {
	this := PoolScrubUpdate1{}
	return &this
}

// NewPoolScrubUpdate1WithDefaults instantiates a new PoolScrubUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolScrubUpdate1WithDefaults() *PoolScrubUpdate1 {
	this := PoolScrubUpdate1{}
	return &this
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *PoolScrubUpdate1) GetPool() int32 {
	if o == nil || o.Pool == nil {
		var ret int32
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolScrubUpdate1) GetPoolOk() (*int32, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *PoolScrubUpdate1) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given int32 and assigns it to the Pool field.
func (o *PoolScrubUpdate1) SetPool(v int32) {
	o.Pool = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *PoolScrubUpdate1) GetThreshold() int32 {
	if o == nil || o.Threshold == nil {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolScrubUpdate1) GetThresholdOk() (*int32, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *PoolScrubUpdate1) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *PoolScrubUpdate1) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PoolScrubUpdate1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolScrubUpdate1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PoolScrubUpdate1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PoolScrubUpdate1) SetDescription(v string) {
	o.Description = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PoolScrubUpdate1) GetSchedule() CloudsyncCreate0Schedule {
	if o == nil || o.Schedule == nil {
		var ret CloudsyncCreate0Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolScrubUpdate1) GetScheduleOk() (*CloudsyncCreate0Schedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PoolScrubUpdate1) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CloudsyncCreate0Schedule and assigns it to the Schedule field.
func (o *PoolScrubUpdate1) SetSchedule(v CloudsyncCreate0Schedule) {
	o.Schedule = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PoolScrubUpdate1) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolScrubUpdate1) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PoolScrubUpdate1) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PoolScrubUpdate1) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o PoolScrubUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pool != nil {
		toSerialize["pool"] = o.Pool
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullablePoolScrubUpdate1 struct {
	value *PoolScrubUpdate1
	isSet bool
}

func (v NullablePoolScrubUpdate1) Get() *PoolScrubUpdate1 {
	return v.value
}

func (v *NullablePoolScrubUpdate1) Set(val *PoolScrubUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolScrubUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolScrubUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolScrubUpdate1(val *PoolScrubUpdate1) *NullablePoolScrubUpdate1 {
	return &NullablePoolScrubUpdate1{value: val, isSet: true}
}

func (v NullablePoolScrubUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolScrubUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


