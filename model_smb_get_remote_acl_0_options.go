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

// SmbGetRemoteAcl0Options struct for SmbGetRemoteAcl0Options
type SmbGetRemoteAcl0Options struct {
	UseKerberos *bool `json:"use_kerberos,omitempty"`
	OutputFormat *string `json:"output_format,omitempty"`
}

// NewSmbGetRemoteAcl0Options instantiates a new SmbGetRemoteAcl0Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbGetRemoteAcl0Options() *SmbGetRemoteAcl0Options {
	this := SmbGetRemoteAcl0Options{}
	return &this
}

// NewSmbGetRemoteAcl0OptionsWithDefaults instantiates a new SmbGetRemoteAcl0Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbGetRemoteAcl0OptionsWithDefaults() *SmbGetRemoteAcl0Options {
	this := SmbGetRemoteAcl0Options{}
	return &this
}

// GetUseKerberos returns the UseKerberos field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0Options) GetUseKerberos() bool {
	if o == nil || o.UseKerberos == nil {
		var ret bool
		return ret
	}
	return *o.UseKerberos
}

// GetUseKerberosOk returns a tuple with the UseKerberos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0Options) GetUseKerberosOk() (*bool, bool) {
	if o == nil || o.UseKerberos == nil {
		return nil, false
	}
	return o.UseKerberos, true
}

// HasUseKerberos returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0Options) HasUseKerberos() bool {
	if o != nil && o.UseKerberos != nil {
		return true
	}

	return false
}

// SetUseKerberos gets a reference to the given bool and assigns it to the UseKerberos field.
func (o *SmbGetRemoteAcl0Options) SetUseKerberos(v bool) {
	o.UseKerberos = &v
}

// GetOutputFormat returns the OutputFormat field value if set, zero value otherwise.
func (o *SmbGetRemoteAcl0Options) GetOutputFormat() string {
	if o == nil || o.OutputFormat == nil {
		var ret string
		return ret
	}
	return *o.OutputFormat
}

// GetOutputFormatOk returns a tuple with the OutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbGetRemoteAcl0Options) GetOutputFormatOk() (*string, bool) {
	if o == nil || o.OutputFormat == nil {
		return nil, false
	}
	return o.OutputFormat, true
}

// HasOutputFormat returns a boolean if a field has been set.
func (o *SmbGetRemoteAcl0Options) HasOutputFormat() bool {
	if o != nil && o.OutputFormat != nil {
		return true
	}

	return false
}

// SetOutputFormat gets a reference to the given string and assigns it to the OutputFormat field.
func (o *SmbGetRemoteAcl0Options) SetOutputFormat(v string) {
	o.OutputFormat = &v
}

func (o SmbGetRemoteAcl0Options) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UseKerberos != nil {
		toSerialize["use_kerberos"] = o.UseKerberos
	}
	if o.OutputFormat != nil {
		toSerialize["output_format"] = o.OutputFormat
	}
	return json.Marshal(toSerialize)
}

type NullableSmbGetRemoteAcl0Options struct {
	value *SmbGetRemoteAcl0Options
	isSet bool
}

func (v NullableSmbGetRemoteAcl0Options) Get() *SmbGetRemoteAcl0Options {
	return v.value
}

func (v *NullableSmbGetRemoteAcl0Options) Set(val *SmbGetRemoteAcl0Options) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbGetRemoteAcl0Options) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbGetRemoteAcl0Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbGetRemoteAcl0Options(val *SmbGetRemoteAcl0Options) *NullableSmbGetRemoteAcl0Options {
	return &NullableSmbGetRemoteAcl0Options{value: val, isSet: true}
}

func (v NullableSmbGetRemoteAcl0Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbGetRemoteAcl0Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


