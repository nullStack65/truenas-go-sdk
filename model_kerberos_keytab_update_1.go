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

// KerberosKeytabUpdate1 struct for KerberosKeytabUpdate1
type KerberosKeytabUpdate1 struct {
	File *string `json:"file,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewKerberosKeytabUpdate1 instantiates a new KerberosKeytabUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosKeytabUpdate1() *KerberosKeytabUpdate1 {
	this := KerberosKeytabUpdate1{}
	return &this
}

// NewKerberosKeytabUpdate1WithDefaults instantiates a new KerberosKeytabUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosKeytabUpdate1WithDefaults() *KerberosKeytabUpdate1 {
	this := KerberosKeytabUpdate1{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *KerberosKeytabUpdate1) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKeytabUpdate1) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *KerberosKeytabUpdate1) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *KerberosKeytabUpdate1) SetFile(v string) {
	o.File = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KerberosKeytabUpdate1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKeytabUpdate1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KerberosKeytabUpdate1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KerberosKeytabUpdate1) SetName(v string) {
	o.Name = &v
}

func (o KerberosKeytabUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableKerberosKeytabUpdate1 struct {
	value *KerberosKeytabUpdate1
	isSet bool
}

func (v NullableKerberosKeytabUpdate1) Get() *KerberosKeytabUpdate1 {
	return v.value
}

func (v *NullableKerberosKeytabUpdate1) Set(val *KerberosKeytabUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosKeytabUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosKeytabUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosKeytabUpdate1(val *KerberosKeytabUpdate1) *NullableKerberosKeytabUpdate1 {
	return &NullableKerberosKeytabUpdate1{value: val, isSet: true}
}

func (v NullableKerberosKeytabUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosKeytabUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


