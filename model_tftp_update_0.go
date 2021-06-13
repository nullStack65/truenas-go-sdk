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

// TftpUpdate0 struct for TftpUpdate0
type TftpUpdate0 struct {
	Newfiles *bool `json:"newfiles,omitempty"`
	Directory *string `json:"directory,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Options *string `json:"options,omitempty"`
	Umask *string `json:"umask,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewTftpUpdate0 instantiates a new TftpUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTftpUpdate0() *TftpUpdate0 {
	this := TftpUpdate0{}
	return &this
}

// NewTftpUpdate0WithDefaults instantiates a new TftpUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTftpUpdate0WithDefaults() *TftpUpdate0 {
	this := TftpUpdate0{}
	return &this
}

// GetNewfiles returns the Newfiles field value if set, zero value otherwise.
func (o *TftpUpdate0) GetNewfiles() bool {
	if o == nil || o.Newfiles == nil {
		var ret bool
		return ret
	}
	return *o.Newfiles
}

// GetNewfilesOk returns a tuple with the Newfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetNewfilesOk() (*bool, bool) {
	if o == nil || o.Newfiles == nil {
		return nil, false
	}
	return o.Newfiles, true
}

// HasNewfiles returns a boolean if a field has been set.
func (o *TftpUpdate0) HasNewfiles() bool {
	if o != nil && o.Newfiles != nil {
		return true
	}

	return false
}

// SetNewfiles gets a reference to the given bool and assigns it to the Newfiles field.
func (o *TftpUpdate0) SetNewfiles(v bool) {
	o.Newfiles = &v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *TftpUpdate0) GetDirectory() string {
	if o == nil || o.Directory == nil {
		var ret string
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetDirectoryOk() (*string, bool) {
	if o == nil || o.Directory == nil {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *TftpUpdate0) HasDirectory() bool {
	if o != nil && o.Directory != nil {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given string and assigns it to the Directory field.
func (o *TftpUpdate0) SetDirectory(v string) {
	o.Directory = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *TftpUpdate0) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *TftpUpdate0) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *TftpUpdate0) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *TftpUpdate0) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *TftpUpdate0) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *TftpUpdate0) SetPort(v int32) {
	o.Port = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TftpUpdate0) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TftpUpdate0) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *TftpUpdate0) SetOptions(v string) {
	o.Options = &v
}

// GetUmask returns the Umask field value if set, zero value otherwise.
func (o *TftpUpdate0) GetUmask() string {
	if o == nil || o.Umask == nil {
		var ret string
		return ret
	}
	return *o.Umask
}

// GetUmaskOk returns a tuple with the Umask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetUmaskOk() (*string, bool) {
	if o == nil || o.Umask == nil {
		return nil, false
	}
	return o.Umask, true
}

// HasUmask returns a boolean if a field has been set.
func (o *TftpUpdate0) HasUmask() bool {
	if o != nil && o.Umask != nil {
		return true
	}

	return false
}

// SetUmask gets a reference to the given string and assigns it to the Umask field.
func (o *TftpUpdate0) SetUmask(v string) {
	o.Umask = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TftpUpdate0) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TftpUpdate0) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TftpUpdate0) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TftpUpdate0) SetUsername(v string) {
	o.Username = &v
}

func (o TftpUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Newfiles != nil {
		toSerialize["newfiles"] = o.Newfiles
	}
	if o.Directory != nil {
		toSerialize["directory"] = o.Directory
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Umask != nil {
		toSerialize["umask"] = o.Umask
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableTftpUpdate0 struct {
	value *TftpUpdate0
	isSet bool
}

func (v NullableTftpUpdate0) Get() *TftpUpdate0 {
	return v.value
}

func (v *NullableTftpUpdate0) Set(val *TftpUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableTftpUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableTftpUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTftpUpdate0(val *TftpUpdate0) *NullableTftpUpdate0 {
	return &NullableTftpUpdate0{value: val, isSet: true}
}

func (v NullableTftpUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTftpUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


