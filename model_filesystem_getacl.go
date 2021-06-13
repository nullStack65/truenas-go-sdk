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

// FilesystemGetacl struct for FilesystemGetacl
type FilesystemGetacl struct {
	Path *string `json:"path,omitempty"`
	Simplified *bool `json:"simplified,omitempty"`
}

// NewFilesystemGetacl instantiates a new FilesystemGetacl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesystemGetacl() *FilesystemGetacl {
	this := FilesystemGetacl{}
	var simplified bool = true
	this.Simplified = &simplified
	return &this
}

// NewFilesystemGetaclWithDefaults instantiates a new FilesystemGetacl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesystemGetaclWithDefaults() *FilesystemGetacl {
	this := FilesystemGetacl{}
	var simplified bool = true
	this.Simplified = &simplified
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FilesystemGetacl) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemGetacl) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FilesystemGetacl) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FilesystemGetacl) SetPath(v string) {
	o.Path = &v
}

// GetSimplified returns the Simplified field value if set, zero value otherwise.
func (o *FilesystemGetacl) GetSimplified() bool {
	if o == nil || o.Simplified == nil {
		var ret bool
		return ret
	}
	return *o.Simplified
}

// GetSimplifiedOk returns a tuple with the Simplified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesystemGetacl) GetSimplifiedOk() (*bool, bool) {
	if o == nil || o.Simplified == nil {
		return nil, false
	}
	return o.Simplified, true
}

// HasSimplified returns a boolean if a field has been set.
func (o *FilesystemGetacl) HasSimplified() bool {
	if o != nil && o.Simplified != nil {
		return true
	}

	return false
}

// SetSimplified gets a reference to the given bool and assigns it to the Simplified field.
func (o *FilesystemGetacl) SetSimplified(v bool) {
	o.Simplified = &v
}

func (o FilesystemGetacl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Simplified != nil {
		toSerialize["simplified"] = o.Simplified
	}
	return json.Marshal(toSerialize)
}

type NullableFilesystemGetacl struct {
	value *FilesystemGetacl
	isSet bool
}

func (v NullableFilesystemGetacl) Get() *FilesystemGetacl {
	return v.value
}

func (v *NullableFilesystemGetacl) Set(val *FilesystemGetacl) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesystemGetacl) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesystemGetacl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesystemGetacl(val *FilesystemGetacl) *NullableFilesystemGetacl {
	return &NullableFilesystemGetacl{value: val, isSet: true}
}

func (v NullableFilesystemGetacl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesystemGetacl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


