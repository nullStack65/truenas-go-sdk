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

// SharingSmbCreate0 struct for SharingSmbCreate0
type SharingSmbCreate0 struct {
	Purpose *string `json:"purpose,omitempty"`
	Path *string `json:"path,omitempty"`
	PathSuffix *string `json:"path_suffix,omitempty"`
	Home *bool `json:"home,omitempty"`
	Name *string `json:"name,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Ro *bool `json:"ro,omitempty"`
	Browsable *bool `json:"browsable,omitempty"`
	Timemachine *bool `json:"timemachine,omitempty"`
	Recyclebin *bool `json:"recyclebin,omitempty"`
	Guestok *bool `json:"guestok,omitempty"`
	Abe *bool `json:"abe,omitempty"`
	Hostsallow *[]interface{} `json:"hostsallow,omitempty"`
	Hostsdeny *[]interface{} `json:"hostsdeny,omitempty"`
	AaplNameMangling *bool `json:"aapl_name_mangling,omitempty"`
	Acl *bool `json:"acl,omitempty"`
	Durablehandle *bool `json:"durablehandle,omitempty"`
	Shadowcopy *bool `json:"shadowcopy,omitempty"`
	Streams *bool `json:"streams,omitempty"`
	Fsrvp *bool `json:"fsrvp,omitempty"`
	Auxsmbconf *string `json:"auxsmbconf,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSharingSmbCreate0 instantiates a new SharingSmbCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingSmbCreate0() *SharingSmbCreate0 {
	this := SharingSmbCreate0{}
	return &this
}

// NewSharingSmbCreate0WithDefaults instantiates a new SharingSmbCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingSmbCreate0WithDefaults() *SharingSmbCreate0 {
	this := SharingSmbCreate0{}
	return &this
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *SharingSmbCreate0) SetPurpose(v string) {
	o.Purpose = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SharingSmbCreate0) SetPath(v string) {
	o.Path = &v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetPathSuffix() string {
	if o == nil || o.PathSuffix == nil {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetPathSuffixOk() (*string, bool) {
	if o == nil || o.PathSuffix == nil {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasPathSuffix() bool {
	if o != nil && o.PathSuffix != nil {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *SharingSmbCreate0) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetHome() bool {
	if o == nil || o.Home == nil {
		var ret bool
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetHomeOk() (*bool, bool) {
	if o == nil || o.Home == nil {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given bool and assigns it to the Home field.
func (o *SharingSmbCreate0) SetHome(v bool) {
	o.Home = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SharingSmbCreate0) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SharingSmbCreate0) SetComment(v string) {
	o.Comment = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetRo() bool {
	if o == nil || o.Ro == nil {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetRoOk() (*bool, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *SharingSmbCreate0) SetRo(v bool) {
	o.Ro = &v
}

// GetBrowsable returns the Browsable field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetBrowsable() bool {
	if o == nil || o.Browsable == nil {
		var ret bool
		return ret
	}
	return *o.Browsable
}

// GetBrowsableOk returns a tuple with the Browsable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetBrowsableOk() (*bool, bool) {
	if o == nil || o.Browsable == nil {
		return nil, false
	}
	return o.Browsable, true
}

// HasBrowsable returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasBrowsable() bool {
	if o != nil && o.Browsable != nil {
		return true
	}

	return false
}

// SetBrowsable gets a reference to the given bool and assigns it to the Browsable field.
func (o *SharingSmbCreate0) SetBrowsable(v bool) {
	o.Browsable = &v
}

// GetTimemachine returns the Timemachine field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetTimemachine() bool {
	if o == nil || o.Timemachine == nil {
		var ret bool
		return ret
	}
	return *o.Timemachine
}

// GetTimemachineOk returns a tuple with the Timemachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetTimemachineOk() (*bool, bool) {
	if o == nil || o.Timemachine == nil {
		return nil, false
	}
	return o.Timemachine, true
}

// HasTimemachine returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasTimemachine() bool {
	if o != nil && o.Timemachine != nil {
		return true
	}

	return false
}

// SetTimemachine gets a reference to the given bool and assigns it to the Timemachine field.
func (o *SharingSmbCreate0) SetTimemachine(v bool) {
	o.Timemachine = &v
}

// GetRecyclebin returns the Recyclebin field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetRecyclebin() bool {
	if o == nil || o.Recyclebin == nil {
		var ret bool
		return ret
	}
	return *o.Recyclebin
}

// GetRecyclebinOk returns a tuple with the Recyclebin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetRecyclebinOk() (*bool, bool) {
	if o == nil || o.Recyclebin == nil {
		return nil, false
	}
	return o.Recyclebin, true
}

// HasRecyclebin returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasRecyclebin() bool {
	if o != nil && o.Recyclebin != nil {
		return true
	}

	return false
}

// SetRecyclebin gets a reference to the given bool and assigns it to the Recyclebin field.
func (o *SharingSmbCreate0) SetRecyclebin(v bool) {
	o.Recyclebin = &v
}

// GetGuestok returns the Guestok field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetGuestok() bool {
	if o == nil || o.Guestok == nil {
		var ret bool
		return ret
	}
	return *o.Guestok
}

// GetGuestokOk returns a tuple with the Guestok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetGuestokOk() (*bool, bool) {
	if o == nil || o.Guestok == nil {
		return nil, false
	}
	return o.Guestok, true
}

// HasGuestok returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasGuestok() bool {
	if o != nil && o.Guestok != nil {
		return true
	}

	return false
}

// SetGuestok gets a reference to the given bool and assigns it to the Guestok field.
func (o *SharingSmbCreate0) SetGuestok(v bool) {
	o.Guestok = &v
}

// GetAbe returns the Abe field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetAbe() bool {
	if o == nil || o.Abe == nil {
		var ret bool
		return ret
	}
	return *o.Abe
}

// GetAbeOk returns a tuple with the Abe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetAbeOk() (*bool, bool) {
	if o == nil || o.Abe == nil {
		return nil, false
	}
	return o.Abe, true
}

// HasAbe returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasAbe() bool {
	if o != nil && o.Abe != nil {
		return true
	}

	return false
}

// SetAbe gets a reference to the given bool and assigns it to the Abe field.
func (o *SharingSmbCreate0) SetAbe(v bool) {
	o.Abe = &v
}

// GetHostsallow returns the Hostsallow field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetHostsallow() []interface{} {
	if o == nil || o.Hostsallow == nil {
		var ret []interface{}
		return ret
	}
	return *o.Hostsallow
}

// GetHostsallowOk returns a tuple with the Hostsallow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetHostsallowOk() (*[]interface{}, bool) {
	if o == nil || o.Hostsallow == nil {
		return nil, false
	}
	return o.Hostsallow, true
}

// HasHostsallow returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasHostsallow() bool {
	if o != nil && o.Hostsallow != nil {
		return true
	}

	return false
}

// SetHostsallow gets a reference to the given []interface{} and assigns it to the Hostsallow field.
func (o *SharingSmbCreate0) SetHostsallow(v []interface{}) {
	o.Hostsallow = &v
}

// GetHostsdeny returns the Hostsdeny field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetHostsdeny() []interface{} {
	if o == nil || o.Hostsdeny == nil {
		var ret []interface{}
		return ret
	}
	return *o.Hostsdeny
}

// GetHostsdenyOk returns a tuple with the Hostsdeny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetHostsdenyOk() (*[]interface{}, bool) {
	if o == nil || o.Hostsdeny == nil {
		return nil, false
	}
	return o.Hostsdeny, true
}

// HasHostsdeny returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasHostsdeny() bool {
	if o != nil && o.Hostsdeny != nil {
		return true
	}

	return false
}

// SetHostsdeny gets a reference to the given []interface{} and assigns it to the Hostsdeny field.
func (o *SharingSmbCreate0) SetHostsdeny(v []interface{}) {
	o.Hostsdeny = &v
}

// GetAaplNameMangling returns the AaplNameMangling field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetAaplNameMangling() bool {
	if o == nil || o.AaplNameMangling == nil {
		var ret bool
		return ret
	}
	return *o.AaplNameMangling
}

// GetAaplNameManglingOk returns a tuple with the AaplNameMangling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetAaplNameManglingOk() (*bool, bool) {
	if o == nil || o.AaplNameMangling == nil {
		return nil, false
	}
	return o.AaplNameMangling, true
}

// HasAaplNameMangling returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasAaplNameMangling() bool {
	if o != nil && o.AaplNameMangling != nil {
		return true
	}

	return false
}

// SetAaplNameMangling gets a reference to the given bool and assigns it to the AaplNameMangling field.
func (o *SharingSmbCreate0) SetAaplNameMangling(v bool) {
	o.AaplNameMangling = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetAcl() bool {
	if o == nil || o.Acl == nil {
		var ret bool
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetAclOk() (*bool, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given bool and assigns it to the Acl field.
func (o *SharingSmbCreate0) SetAcl(v bool) {
	o.Acl = &v
}

// GetDurablehandle returns the Durablehandle field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetDurablehandle() bool {
	if o == nil || o.Durablehandle == nil {
		var ret bool
		return ret
	}
	return *o.Durablehandle
}

// GetDurablehandleOk returns a tuple with the Durablehandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetDurablehandleOk() (*bool, bool) {
	if o == nil || o.Durablehandle == nil {
		return nil, false
	}
	return o.Durablehandle, true
}

// HasDurablehandle returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasDurablehandle() bool {
	if o != nil && o.Durablehandle != nil {
		return true
	}

	return false
}

// SetDurablehandle gets a reference to the given bool and assigns it to the Durablehandle field.
func (o *SharingSmbCreate0) SetDurablehandle(v bool) {
	o.Durablehandle = &v
}

// GetShadowcopy returns the Shadowcopy field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetShadowcopy() bool {
	if o == nil || o.Shadowcopy == nil {
		var ret bool
		return ret
	}
	return *o.Shadowcopy
}

// GetShadowcopyOk returns a tuple with the Shadowcopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetShadowcopyOk() (*bool, bool) {
	if o == nil || o.Shadowcopy == nil {
		return nil, false
	}
	return o.Shadowcopy, true
}

// HasShadowcopy returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasShadowcopy() bool {
	if o != nil && o.Shadowcopy != nil {
		return true
	}

	return false
}

// SetShadowcopy gets a reference to the given bool and assigns it to the Shadowcopy field.
func (o *SharingSmbCreate0) SetShadowcopy(v bool) {
	o.Shadowcopy = &v
}

// GetStreams returns the Streams field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetStreams() bool {
	if o == nil || o.Streams == nil {
		var ret bool
		return ret
	}
	return *o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetStreamsOk() (*bool, bool) {
	if o == nil || o.Streams == nil {
		return nil, false
	}
	return o.Streams, true
}

// HasStreams returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasStreams() bool {
	if o != nil && o.Streams != nil {
		return true
	}

	return false
}

// SetStreams gets a reference to the given bool and assigns it to the Streams field.
func (o *SharingSmbCreate0) SetStreams(v bool) {
	o.Streams = &v
}

// GetFsrvp returns the Fsrvp field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetFsrvp() bool {
	if o == nil || o.Fsrvp == nil {
		var ret bool
		return ret
	}
	return *o.Fsrvp
}

// GetFsrvpOk returns a tuple with the Fsrvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetFsrvpOk() (*bool, bool) {
	if o == nil || o.Fsrvp == nil {
		return nil, false
	}
	return o.Fsrvp, true
}

// HasFsrvp returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasFsrvp() bool {
	if o != nil && o.Fsrvp != nil {
		return true
	}

	return false
}

// SetFsrvp gets a reference to the given bool and assigns it to the Fsrvp field.
func (o *SharingSmbCreate0) SetFsrvp(v bool) {
	o.Fsrvp = &v
}

// GetAuxsmbconf returns the Auxsmbconf field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetAuxsmbconf() string {
	if o == nil || o.Auxsmbconf == nil {
		var ret string
		return ret
	}
	return *o.Auxsmbconf
}

// GetAuxsmbconfOk returns a tuple with the Auxsmbconf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetAuxsmbconfOk() (*string, bool) {
	if o == nil || o.Auxsmbconf == nil {
		return nil, false
	}
	return o.Auxsmbconf, true
}

// HasAuxsmbconf returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasAuxsmbconf() bool {
	if o != nil && o.Auxsmbconf != nil {
		return true
	}

	return false
}

// SetAuxsmbconf gets a reference to the given string and assigns it to the Auxsmbconf field.
func (o *SharingSmbCreate0) SetAuxsmbconf(v string) {
	o.Auxsmbconf = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SharingSmbCreate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingSmbCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SharingSmbCreate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SharingSmbCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SharingSmbCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PathSuffix != nil {
		toSerialize["path_suffix"] = o.PathSuffix
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
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Browsable != nil {
		toSerialize["browsable"] = o.Browsable
	}
	if o.Timemachine != nil {
		toSerialize["timemachine"] = o.Timemachine
	}
	if o.Recyclebin != nil {
		toSerialize["recyclebin"] = o.Recyclebin
	}
	if o.Guestok != nil {
		toSerialize["guestok"] = o.Guestok
	}
	if o.Abe != nil {
		toSerialize["abe"] = o.Abe
	}
	if o.Hostsallow != nil {
		toSerialize["hostsallow"] = o.Hostsallow
	}
	if o.Hostsdeny != nil {
		toSerialize["hostsdeny"] = o.Hostsdeny
	}
	if o.AaplNameMangling != nil {
		toSerialize["aapl_name_mangling"] = o.AaplNameMangling
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Durablehandle != nil {
		toSerialize["durablehandle"] = o.Durablehandle
	}
	if o.Shadowcopy != nil {
		toSerialize["shadowcopy"] = o.Shadowcopy
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	if o.Fsrvp != nil {
		toSerialize["fsrvp"] = o.Fsrvp
	}
	if o.Auxsmbconf != nil {
		toSerialize["auxsmbconf"] = o.Auxsmbconf
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSharingSmbCreate0 struct {
	value *SharingSmbCreate0
	isSet bool
}

func (v NullableSharingSmbCreate0) Get() *SharingSmbCreate0 {
	return v.value
}

func (v *NullableSharingSmbCreate0) Set(val *SharingSmbCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingSmbCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingSmbCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingSmbCreate0(val *SharingSmbCreate0) *NullableSharingSmbCreate0 {
	return &NullableSharingSmbCreate0{value: val, isSet: true}
}

func (v NullableSharingSmbCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingSmbCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


