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

// PluginAvailable0 struct for PluginAvailable0
type PluginAvailable0 struct {
	Cache *bool `json:"cache,omitempty"`
	PluginRepository *string `json:"plugin_repository,omitempty"`
	Branch *string `json:"branch,omitempty"`
}

// NewPluginAvailable0 instantiates a new PluginAvailable0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginAvailable0() *PluginAvailable0 {
	this := PluginAvailable0{}
	return &this
}

// NewPluginAvailable0WithDefaults instantiates a new PluginAvailable0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginAvailable0WithDefaults() *PluginAvailable0 {
	this := PluginAvailable0{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *PluginAvailable0) GetCache() bool {
	if o == nil || o.Cache == nil {
		var ret bool
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginAvailable0) GetCacheOk() (*bool, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *PluginAvailable0) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given bool and assigns it to the Cache field.
func (o *PluginAvailable0) SetCache(v bool) {
	o.Cache = &v
}

// GetPluginRepository returns the PluginRepository field value if set, zero value otherwise.
func (o *PluginAvailable0) GetPluginRepository() string {
	if o == nil || o.PluginRepository == nil {
		var ret string
		return ret
	}
	return *o.PluginRepository
}

// GetPluginRepositoryOk returns a tuple with the PluginRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginAvailable0) GetPluginRepositoryOk() (*string, bool) {
	if o == nil || o.PluginRepository == nil {
		return nil, false
	}
	return o.PluginRepository, true
}

// HasPluginRepository returns a boolean if a field has been set.
func (o *PluginAvailable0) HasPluginRepository() bool {
	if o != nil && o.PluginRepository != nil {
		return true
	}

	return false
}

// SetPluginRepository gets a reference to the given string and assigns it to the PluginRepository field.
func (o *PluginAvailable0) SetPluginRepository(v string) {
	o.PluginRepository = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *PluginAvailable0) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginAvailable0) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *PluginAvailable0) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *PluginAvailable0) SetBranch(v string) {
	o.Branch = &v
}

func (o PluginAvailable0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.PluginRepository != nil {
		toSerialize["plugin_repository"] = o.PluginRepository
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	return json.Marshal(toSerialize)
}

type NullablePluginAvailable0 struct {
	value *PluginAvailable0
	isSet bool
}

func (v NullablePluginAvailable0) Get() *PluginAvailable0 {
	return v.value
}

func (v *NullablePluginAvailable0) Set(val *PluginAvailable0) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginAvailable0) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginAvailable0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginAvailable0(val *PluginAvailable0) *NullablePluginAvailable0 {
	return &NullablePluginAvailable0{value: val, isSet: true}
}

func (v NullablePluginAvailable0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginAvailable0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


