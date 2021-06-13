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

// SharingAfpCreate0 struct for SharingAfpCreate0
type SharingAfpCreate0 struct {
	Path *string `json:"path,omitempty"`
	Home *bool `json:"home,omitempty"`
	Name *string `json:"name,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Allow *[]interface{} `json:"allow,omitempty"`
	Deny *[]interface{} `json:"deny,omitempty"`
	Ro *[]interface{} `json:"ro,omitempty"`
	Rw *[]interface{} `json:"rw,omitempty"`
	Timemachine *bool `json:"timemachine,omitempty"`
	TimemachineQuota *int32 `json:"timemachine_quota,omitempty"`
	Nodev *bool `json:"nodev,omitempty"`
	Nostat *bool `json:"nostat,omitempty"`
	Upriv *bool `json:"upriv,omitempty"`
	Fperm *string `json:"fperm,omitempty"`
	Dperm *string `json:"dperm,omitempty"`
	Umask *string `json:"umask,omitempty"`
	Hostsallow *[]interface{} `json:"hostsallow,omitempty"`
	Hostsdeny *[]interface{} `json:"hostsdeny,omitempty"`
	Vuid NullableString `json:"vuid,omitempty"`
	Auxparams *string `json:"auxparams,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSharingAfpCreate0 instantiates a new SharingAfpCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingAfpCreate0() *SharingAfpCreate0 {
	this := SharingAfpCreate0{}
	return &this
}

// NewSharingAfpCreate0WithDefaults instantiates a new SharingAfpCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingAfpCreate0WithDefaults() *SharingAfpCreate0 {
	this := SharingAfpCreate0{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SharingAfpCreate0) SetPath(v string) {
	o.Path = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetHome() bool {
	if o == nil || o.Home == nil {
		var ret bool
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetHomeOk() (*bool, bool) {
	if o == nil || o.Home == nil {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given bool and assigns it to the Home field.
func (o *SharingAfpCreate0) SetHome(v bool) {
	o.Home = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SharingAfpCreate0) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SharingAfpCreate0) SetComment(v string) {
	o.Comment = &v
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetAllow() []interface{} {
	if o == nil || o.Allow == nil {
		var ret []interface{}
		return ret
	}
	return *o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetAllowOk() (*[]interface{}, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasAllow() bool {
	if o != nil && o.Allow != nil {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []interface{} and assigns it to the Allow field.
func (o *SharingAfpCreate0) SetAllow(v []interface{}) {
	o.Allow = &v
}

// GetDeny returns the Deny field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetDeny() []interface{} {
	if o == nil || o.Deny == nil {
		var ret []interface{}
		return ret
	}
	return *o.Deny
}

// GetDenyOk returns a tuple with the Deny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetDenyOk() (*[]interface{}, bool) {
	if o == nil || o.Deny == nil {
		return nil, false
	}
	return o.Deny, true
}

// HasDeny returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasDeny() bool {
	if o != nil && o.Deny != nil {
		return true
	}

	return false
}

// SetDeny gets a reference to the given []interface{} and assigns it to the Deny field.
func (o *SharingAfpCreate0) SetDeny(v []interface{}) {
	o.Deny = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetRo() []interface{} {
	if o == nil || o.Ro == nil {
		var ret []interface{}
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetRoOk() (*[]interface{}, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given []interface{} and assigns it to the Ro field.
func (o *SharingAfpCreate0) SetRo(v []interface{}) {
	o.Ro = &v
}

// GetRw returns the Rw field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetRw() []interface{} {
	if o == nil || o.Rw == nil {
		var ret []interface{}
		return ret
	}
	return *o.Rw
}

// GetRwOk returns a tuple with the Rw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetRwOk() (*[]interface{}, bool) {
	if o == nil || o.Rw == nil {
		return nil, false
	}
	return o.Rw, true
}

// HasRw returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasRw() bool {
	if o != nil && o.Rw != nil {
		return true
	}

	return false
}

// SetRw gets a reference to the given []interface{} and assigns it to the Rw field.
func (o *SharingAfpCreate0) SetRw(v []interface{}) {
	o.Rw = &v
}

// GetTimemachine returns the Timemachine field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetTimemachine() bool {
	if o == nil || o.Timemachine == nil {
		var ret bool
		return ret
	}
	return *o.Timemachine
}

// GetTimemachineOk returns a tuple with the Timemachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetTimemachineOk() (*bool, bool) {
	if o == nil || o.Timemachine == nil {
		return nil, false
	}
	return o.Timemachine, true
}

// HasTimemachine returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasTimemachine() bool {
	if o != nil && o.Timemachine != nil {
		return true
	}

	return false
}

// SetTimemachine gets a reference to the given bool and assigns it to the Timemachine field.
func (o *SharingAfpCreate0) SetTimemachine(v bool) {
	o.Timemachine = &v
}

// GetTimemachineQuota returns the TimemachineQuota field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetTimemachineQuota() int32 {
	if o == nil || o.TimemachineQuota == nil {
		var ret int32
		return ret
	}
	return *o.TimemachineQuota
}

// GetTimemachineQuotaOk returns a tuple with the TimemachineQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetTimemachineQuotaOk() (*int32, bool) {
	if o == nil || o.TimemachineQuota == nil {
		return nil, false
	}
	return o.TimemachineQuota, true
}

// HasTimemachineQuota returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasTimemachineQuota() bool {
	if o != nil && o.TimemachineQuota != nil {
		return true
	}

	return false
}

// SetTimemachineQuota gets a reference to the given int32 and assigns it to the TimemachineQuota field.
func (o *SharingAfpCreate0) SetTimemachineQuota(v int32) {
	o.TimemachineQuota = &v
}

// GetNodev returns the Nodev field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetNodev() bool {
	if o == nil || o.Nodev == nil {
		var ret bool
		return ret
	}
	return *o.Nodev
}

// GetNodevOk returns a tuple with the Nodev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetNodevOk() (*bool, bool) {
	if o == nil || o.Nodev == nil {
		return nil, false
	}
	return o.Nodev, true
}

// HasNodev returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasNodev() bool {
	if o != nil && o.Nodev != nil {
		return true
	}

	return false
}

// SetNodev gets a reference to the given bool and assigns it to the Nodev field.
func (o *SharingAfpCreate0) SetNodev(v bool) {
	o.Nodev = &v
}

// GetNostat returns the Nostat field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetNostat() bool {
	if o == nil || o.Nostat == nil {
		var ret bool
		return ret
	}
	return *o.Nostat
}

// GetNostatOk returns a tuple with the Nostat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetNostatOk() (*bool, bool) {
	if o == nil || o.Nostat == nil {
		return nil, false
	}
	return o.Nostat, true
}

// HasNostat returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasNostat() bool {
	if o != nil && o.Nostat != nil {
		return true
	}

	return false
}

// SetNostat gets a reference to the given bool and assigns it to the Nostat field.
func (o *SharingAfpCreate0) SetNostat(v bool) {
	o.Nostat = &v
}

// GetUpriv returns the Upriv field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetUpriv() bool {
	if o == nil || o.Upriv == nil {
		var ret bool
		return ret
	}
	return *o.Upriv
}

// GetUprivOk returns a tuple with the Upriv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetUprivOk() (*bool, bool) {
	if o == nil || o.Upriv == nil {
		return nil, false
	}
	return o.Upriv, true
}

// HasUpriv returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasUpriv() bool {
	if o != nil && o.Upriv != nil {
		return true
	}

	return false
}

// SetUpriv gets a reference to the given bool and assigns it to the Upriv field.
func (o *SharingAfpCreate0) SetUpriv(v bool) {
	o.Upriv = &v
}

// GetFperm returns the Fperm field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetFperm() string {
	if o == nil || o.Fperm == nil {
		var ret string
		return ret
	}
	return *o.Fperm
}

// GetFpermOk returns a tuple with the Fperm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetFpermOk() (*string, bool) {
	if o == nil || o.Fperm == nil {
		return nil, false
	}
	return o.Fperm, true
}

// HasFperm returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasFperm() bool {
	if o != nil && o.Fperm != nil {
		return true
	}

	return false
}

// SetFperm gets a reference to the given string and assigns it to the Fperm field.
func (o *SharingAfpCreate0) SetFperm(v string) {
	o.Fperm = &v
}

// GetDperm returns the Dperm field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetDperm() string {
	if o == nil || o.Dperm == nil {
		var ret string
		return ret
	}
	return *o.Dperm
}

// GetDpermOk returns a tuple with the Dperm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetDpermOk() (*string, bool) {
	if o == nil || o.Dperm == nil {
		return nil, false
	}
	return o.Dperm, true
}

// HasDperm returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasDperm() bool {
	if o != nil && o.Dperm != nil {
		return true
	}

	return false
}

// SetDperm gets a reference to the given string and assigns it to the Dperm field.
func (o *SharingAfpCreate0) SetDperm(v string) {
	o.Dperm = &v
}

// GetUmask returns the Umask field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetUmask() string {
	if o == nil || o.Umask == nil {
		var ret string
		return ret
	}
	return *o.Umask
}

// GetUmaskOk returns a tuple with the Umask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetUmaskOk() (*string, bool) {
	if o == nil || o.Umask == nil {
		return nil, false
	}
	return o.Umask, true
}

// HasUmask returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasUmask() bool {
	if o != nil && o.Umask != nil {
		return true
	}

	return false
}

// SetUmask gets a reference to the given string and assigns it to the Umask field.
func (o *SharingAfpCreate0) SetUmask(v string) {
	o.Umask = &v
}

// GetHostsallow returns the Hostsallow field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetHostsallow() []interface{} {
	if o == nil || o.Hostsallow == nil {
		var ret []interface{}
		return ret
	}
	return *o.Hostsallow
}

// GetHostsallowOk returns a tuple with the Hostsallow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetHostsallowOk() (*[]interface{}, bool) {
	if o == nil || o.Hostsallow == nil {
		return nil, false
	}
	return o.Hostsallow, true
}

// HasHostsallow returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasHostsallow() bool {
	if o != nil && o.Hostsallow != nil {
		return true
	}

	return false
}

// SetHostsallow gets a reference to the given []interface{} and assigns it to the Hostsallow field.
func (o *SharingAfpCreate0) SetHostsallow(v []interface{}) {
	o.Hostsallow = &v
}

// GetHostsdeny returns the Hostsdeny field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetHostsdeny() []interface{} {
	if o == nil || o.Hostsdeny == nil {
		var ret []interface{}
		return ret
	}
	return *o.Hostsdeny
}

// GetHostsdenyOk returns a tuple with the Hostsdeny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetHostsdenyOk() (*[]interface{}, bool) {
	if o == nil || o.Hostsdeny == nil {
		return nil, false
	}
	return o.Hostsdeny, true
}

// HasHostsdeny returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasHostsdeny() bool {
	if o != nil && o.Hostsdeny != nil {
		return true
	}

	return false
}

// SetHostsdeny gets a reference to the given []interface{} and assigns it to the Hostsdeny field.
func (o *SharingAfpCreate0) SetHostsdeny(v []interface{}) {
	o.Hostsdeny = &v
}

// GetVuid returns the Vuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SharingAfpCreate0) GetVuid() string {
	if o == nil || o.Vuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Vuid.Get()
}

// GetVuidOk returns a tuple with the Vuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SharingAfpCreate0) GetVuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vuid.Get(), o.Vuid.IsSet()
}

// HasVuid returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasVuid() bool {
	if o != nil && o.Vuid.IsSet() {
		return true
	}

	return false
}

// SetVuid gets a reference to the given NullableString and assigns it to the Vuid field.
func (o *SharingAfpCreate0) SetVuid(v string) {
	o.Vuid.Set(&v)
}
// SetVuidNil sets the value for Vuid to be an explicit nil
func (o *SharingAfpCreate0) SetVuidNil() {
	o.Vuid.Set(nil)
}

// UnsetVuid ensures that no value is present for Vuid, not even an explicit nil
func (o *SharingAfpCreate0) UnsetVuid() {
	o.Vuid.Unset()
}

// GetAuxparams returns the Auxparams field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetAuxparams() string {
	if o == nil || o.Auxparams == nil {
		var ret string
		return ret
	}
	return *o.Auxparams
}

// GetAuxparamsOk returns a tuple with the Auxparams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetAuxparamsOk() (*string, bool) {
	if o == nil || o.Auxparams == nil {
		return nil, false
	}
	return o.Auxparams, true
}

// HasAuxparams returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasAuxparams() bool {
	if o != nil && o.Auxparams != nil {
		return true
	}

	return false
}

// SetAuxparams gets a reference to the given string and assigns it to the Auxparams field.
func (o *SharingAfpCreate0) SetAuxparams(v string) {
	o.Auxparams = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SharingAfpCreate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingAfpCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SharingAfpCreate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SharingAfpCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SharingAfpCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Home != nil {
		toSerialize["home"] = o.Home
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}
	if o.Deny != nil {
		toSerialize["deny"] = o.Deny
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Rw != nil {
		toSerialize["rw"] = o.Rw
	}
	if o.Timemachine != nil {
		toSerialize["timemachine"] = o.Timemachine
	}
	if o.TimemachineQuota != nil {
		toSerialize["timemachine_quota"] = o.TimemachineQuota
	}
	if o.Nodev != nil {
		toSerialize["nodev"] = o.Nodev
	}
	if o.Nostat != nil {
		toSerialize["nostat"] = o.Nostat
	}
	if o.Upriv != nil {
		toSerialize["upriv"] = o.Upriv
	}
	if o.Fperm != nil {
		toSerialize["fperm"] = o.Fperm
	}
	if o.Dperm != nil {
		toSerialize["dperm"] = o.Dperm
	}
	if o.Umask != nil {
		toSerialize["umask"] = o.Umask
	}
	if o.Hostsallow != nil {
		toSerialize["hostsallow"] = o.Hostsallow
	}
	if o.Hostsdeny != nil {
		toSerialize["hostsdeny"] = o.Hostsdeny
	}
	if o.Vuid.IsSet() {
		toSerialize["vuid"] = o.Vuid.Get()
	}
	if o.Auxparams != nil {
		toSerialize["auxparams"] = o.Auxparams
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSharingAfpCreate0 struct {
	value *SharingAfpCreate0
	isSet bool
}

func (v NullableSharingAfpCreate0) Get() *SharingAfpCreate0 {
	return v.value
}

func (v *NullableSharingAfpCreate0) Set(val *SharingAfpCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingAfpCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingAfpCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingAfpCreate0(val *SharingAfpCreate0) *NullableSharingAfpCreate0 {
	return &NullableSharingAfpCreate0{value: val, isSet: true}
}

func (v NullableSharingAfpCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingAfpCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


