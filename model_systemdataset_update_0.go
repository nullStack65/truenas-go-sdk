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

// SystemdatasetUpdate0 struct for SystemdatasetUpdate0
type SystemdatasetUpdate0 struct {
	Pool NullableString `json:"pool,omitempty"`
	PoolExclude NullableString `json:"pool_exclude,omitempty"`
	Syslog *bool `json:"syslog,omitempty"`
}

// NewSystemdatasetUpdate0 instantiates a new SystemdatasetUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemdatasetUpdate0() *SystemdatasetUpdate0 {
	this := SystemdatasetUpdate0{}
	return &this
}

// NewSystemdatasetUpdate0WithDefaults instantiates a new SystemdatasetUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemdatasetUpdate0WithDefaults() *SystemdatasetUpdate0 {
	this := SystemdatasetUpdate0{}
	return &this
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemdatasetUpdate0) GetPool() string {
	if o == nil || o.Pool.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemdatasetUpdate0) GetPoolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *SystemdatasetUpdate0) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableString and assigns it to the Pool field.
func (o *SystemdatasetUpdate0) SetPool(v string) {
	o.Pool.Set(&v)
}
// SetPoolNil sets the value for Pool to be an explicit nil
func (o *SystemdatasetUpdate0) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *SystemdatasetUpdate0) UnsetPool() {
	o.Pool.Unset()
}

// GetPoolExclude returns the PoolExclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemdatasetUpdate0) GetPoolExclude() string {
	if o == nil || o.PoolExclude.Get() == nil {
		var ret string
		return ret
	}
	return *o.PoolExclude.Get()
}

// GetPoolExcludeOk returns a tuple with the PoolExclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemdatasetUpdate0) GetPoolExcludeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PoolExclude.Get(), o.PoolExclude.IsSet()
}

// HasPoolExclude returns a boolean if a field has been set.
func (o *SystemdatasetUpdate0) HasPoolExclude() bool {
	if o != nil && o.PoolExclude.IsSet() {
		return true
	}

	return false
}

// SetPoolExclude gets a reference to the given NullableString and assigns it to the PoolExclude field.
func (o *SystemdatasetUpdate0) SetPoolExclude(v string) {
	o.PoolExclude.Set(&v)
}
// SetPoolExcludeNil sets the value for PoolExclude to be an explicit nil
func (o *SystemdatasetUpdate0) SetPoolExcludeNil() {
	o.PoolExclude.Set(nil)
}

// UnsetPoolExclude ensures that no value is present for PoolExclude, not even an explicit nil
func (o *SystemdatasetUpdate0) UnsetPoolExclude() {
	o.PoolExclude.Unset()
}

// GetSyslog returns the Syslog field value if set, zero value otherwise.
func (o *SystemdatasetUpdate0) GetSyslog() bool {
	if o == nil || o.Syslog == nil {
		var ret bool
		return ret
	}
	return *o.Syslog
}

// GetSyslogOk returns a tuple with the Syslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemdatasetUpdate0) GetSyslogOk() (*bool, bool) {
	if o == nil || o.Syslog == nil {
		return nil, false
	}
	return o.Syslog, true
}

// HasSyslog returns a boolean if a field has been set.
func (o *SystemdatasetUpdate0) HasSyslog() bool {
	if o != nil && o.Syslog != nil {
		return true
	}

	return false
}

// SetSyslog gets a reference to the given bool and assigns it to the Syslog field.
func (o *SystemdatasetUpdate0) SetSyslog(v bool) {
	o.Syslog = &v
}

func (o SystemdatasetUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pool.IsSet() {
		toSerialize["pool"] = o.Pool.Get()
	}
	if o.PoolExclude.IsSet() {
		toSerialize["pool_exclude"] = o.PoolExclude.Get()
	}
	if o.Syslog != nil {
		toSerialize["syslog"] = o.Syslog
	}
	return json.Marshal(toSerialize)
}

type NullableSystemdatasetUpdate0 struct {
	value *SystemdatasetUpdate0
	isSet bool
}

func (v NullableSystemdatasetUpdate0) Get() *SystemdatasetUpdate0 {
	return v.value
}

func (v *NullableSystemdatasetUpdate0) Set(val *SystemdatasetUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemdatasetUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemdatasetUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemdatasetUpdate0(val *SystemdatasetUpdate0) *NullableSystemdatasetUpdate0 {
	return &NullableSystemdatasetUpdate0{value: val, isSet: true}
}

func (v NullableSystemdatasetUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemdatasetUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


