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

// CloudsyncSyncOnetime0 struct for CloudsyncSyncOnetime0
type CloudsyncSyncOnetime0 struct {
	Description *string `json:"description,omitempty"`
	Direction *string `json:"direction,omitempty"`
	TransferMode *string `json:"transfer_mode,omitempty"`
	Path *string `json:"path,omitempty"`
	Credentials *int32 `json:"credentials,omitempty"`
	Encryption *bool `json:"encryption,omitempty"`
	FilenameEncryption *bool `json:"filename_encryption,omitempty"`
	EncryptionPassword *string `json:"encryption_password,omitempty"`
	EncryptionSalt *string `json:"encryption_salt,omitempty"`
	Schedule *CloudsyncCreate0Schedule `json:"schedule,omitempty"`
	FollowSymlinks *bool `json:"follow_symlinks,omitempty"`
	Transfers NullableInt32 `json:"transfers,omitempty"`
	Bwlimit *[]map[string]interface{} `json:"bwlimit,omitempty"`
	Exclude *[]string `json:"exclude,omitempty"`
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	Snapshot *bool `json:"snapshot,omitempty"`
	PreScript *string `json:"pre_script,omitempty"`
	PostScript *string `json:"post_script,omitempty"`
	Args *string `json:"args,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCloudsyncSyncOnetime0 instantiates a new CloudsyncSyncOnetime0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudsyncSyncOnetime0() *CloudsyncSyncOnetime0 {
	this := CloudsyncSyncOnetime0{}
	return &this
}

// NewCloudsyncSyncOnetime0WithDefaults instantiates a new CloudsyncSyncOnetime0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudsyncSyncOnetime0WithDefaults() *CloudsyncSyncOnetime0 {
	this := CloudsyncSyncOnetime0{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudsyncSyncOnetime0) SetDescription(v string) {
	o.Description = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *CloudsyncSyncOnetime0) SetDirection(v string) {
	o.Direction = &v
}

// GetTransferMode returns the TransferMode field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetTransferMode() string {
	if o == nil || o.TransferMode == nil {
		var ret string
		return ret
	}
	return *o.TransferMode
}

// GetTransferModeOk returns a tuple with the TransferMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetTransferModeOk() (*string, bool) {
	if o == nil || o.TransferMode == nil {
		return nil, false
	}
	return o.TransferMode, true
}

// HasTransferMode returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasTransferMode() bool {
	if o != nil && o.TransferMode != nil {
		return true
	}

	return false
}

// SetTransferMode gets a reference to the given string and assigns it to the TransferMode field.
func (o *CloudsyncSyncOnetime0) SetTransferMode(v string) {
	o.TransferMode = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CloudsyncSyncOnetime0) SetPath(v string) {
	o.Path = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetCredentials() int32 {
	if o == nil || o.Credentials == nil {
		var ret int32
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetCredentialsOk() (*int32, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given int32 and assigns it to the Credentials field.
func (o *CloudsyncSyncOnetime0) SetCredentials(v int32) {
	o.Credentials = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *CloudsyncSyncOnetime0) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetFilenameEncryption returns the FilenameEncryption field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetFilenameEncryption() bool {
	if o == nil || o.FilenameEncryption == nil {
		var ret bool
		return ret
	}
	return *o.FilenameEncryption
}

// GetFilenameEncryptionOk returns a tuple with the FilenameEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetFilenameEncryptionOk() (*bool, bool) {
	if o == nil || o.FilenameEncryption == nil {
		return nil, false
	}
	return o.FilenameEncryption, true
}

// HasFilenameEncryption returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasFilenameEncryption() bool {
	if o != nil && o.FilenameEncryption != nil {
		return true
	}

	return false
}

// SetFilenameEncryption gets a reference to the given bool and assigns it to the FilenameEncryption field.
func (o *CloudsyncSyncOnetime0) SetFilenameEncryption(v bool) {
	o.FilenameEncryption = &v
}

// GetEncryptionPassword returns the EncryptionPassword field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetEncryptionPassword() string {
	if o == nil || o.EncryptionPassword == nil {
		var ret string
		return ret
	}
	return *o.EncryptionPassword
}

// GetEncryptionPasswordOk returns a tuple with the EncryptionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetEncryptionPasswordOk() (*string, bool) {
	if o == nil || o.EncryptionPassword == nil {
		return nil, false
	}
	return o.EncryptionPassword, true
}

// HasEncryptionPassword returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasEncryptionPassword() bool {
	if o != nil && o.EncryptionPassword != nil {
		return true
	}

	return false
}

// SetEncryptionPassword gets a reference to the given string and assigns it to the EncryptionPassword field.
func (o *CloudsyncSyncOnetime0) SetEncryptionPassword(v string) {
	o.EncryptionPassword = &v
}

// GetEncryptionSalt returns the EncryptionSalt field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetEncryptionSalt() string {
	if o == nil || o.EncryptionSalt == nil {
		var ret string
		return ret
	}
	return *o.EncryptionSalt
}

// GetEncryptionSaltOk returns a tuple with the EncryptionSalt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetEncryptionSaltOk() (*string, bool) {
	if o == nil || o.EncryptionSalt == nil {
		return nil, false
	}
	return o.EncryptionSalt, true
}

// HasEncryptionSalt returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasEncryptionSalt() bool {
	if o != nil && o.EncryptionSalt != nil {
		return true
	}

	return false
}

// SetEncryptionSalt gets a reference to the given string and assigns it to the EncryptionSalt field.
func (o *CloudsyncSyncOnetime0) SetEncryptionSalt(v string) {
	o.EncryptionSalt = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetSchedule() CloudsyncCreate0Schedule {
	if o == nil || o.Schedule == nil {
		var ret CloudsyncCreate0Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetScheduleOk() (*CloudsyncCreate0Schedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CloudsyncCreate0Schedule and assigns it to the Schedule field.
func (o *CloudsyncSyncOnetime0) SetSchedule(v CloudsyncCreate0Schedule) {
	o.Schedule = &v
}

// GetFollowSymlinks returns the FollowSymlinks field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetFollowSymlinks() bool {
	if o == nil || o.FollowSymlinks == nil {
		var ret bool
		return ret
	}
	return *o.FollowSymlinks
}

// GetFollowSymlinksOk returns a tuple with the FollowSymlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetFollowSymlinksOk() (*bool, bool) {
	if o == nil || o.FollowSymlinks == nil {
		return nil, false
	}
	return o.FollowSymlinks, true
}

// HasFollowSymlinks returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasFollowSymlinks() bool {
	if o != nil && o.FollowSymlinks != nil {
		return true
	}

	return false
}

// SetFollowSymlinks gets a reference to the given bool and assigns it to the FollowSymlinks field.
func (o *CloudsyncSyncOnetime0) SetFollowSymlinks(v bool) {
	o.FollowSymlinks = &v
}

// GetTransfers returns the Transfers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudsyncSyncOnetime0) GetTransfers() int32 {
	if o == nil || o.Transfers.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Transfers.Get()
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudsyncSyncOnetime0) GetTransfersOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Transfers.Get(), o.Transfers.IsSet()
}

// HasTransfers returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasTransfers() bool {
	if o != nil && o.Transfers.IsSet() {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given NullableInt32 and assigns it to the Transfers field.
func (o *CloudsyncSyncOnetime0) SetTransfers(v int32) {
	o.Transfers.Set(&v)
}
// SetTransfersNil sets the value for Transfers to be an explicit nil
func (o *CloudsyncSyncOnetime0) SetTransfersNil() {
	o.Transfers.Set(nil)
}

// UnsetTransfers ensures that no value is present for Transfers, not even an explicit nil
func (o *CloudsyncSyncOnetime0) UnsetTransfers() {
	o.Transfers.Unset()
}

// GetBwlimit returns the Bwlimit field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetBwlimit() []map[string]interface{} {
	if o == nil || o.Bwlimit == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Bwlimit
}

// GetBwlimitOk returns a tuple with the Bwlimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetBwlimitOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Bwlimit == nil {
		return nil, false
	}
	return o.Bwlimit, true
}

// HasBwlimit returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasBwlimit() bool {
	if o != nil && o.Bwlimit != nil {
		return true
	}

	return false
}

// SetBwlimit gets a reference to the given []map[string]interface{} and assigns it to the Bwlimit field.
func (o *CloudsyncSyncOnetime0) SetBwlimit(v []map[string]interface{}) {
	o.Bwlimit = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetExcludeOk() (*[]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *CloudsyncSyncOnetime0) SetExclude(v []string) {
	o.Exclude = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetAttributesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *CloudsyncSyncOnetime0) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetSnapshot() bool {
	if o == nil || o.Snapshot == nil {
		var ret bool
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetSnapshotOk() (*bool, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given bool and assigns it to the Snapshot field.
func (o *CloudsyncSyncOnetime0) SetSnapshot(v bool) {
	o.Snapshot = &v
}

// GetPreScript returns the PreScript field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetPreScript() string {
	if o == nil || o.PreScript == nil {
		var ret string
		return ret
	}
	return *o.PreScript
}

// GetPreScriptOk returns a tuple with the PreScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetPreScriptOk() (*string, bool) {
	if o == nil || o.PreScript == nil {
		return nil, false
	}
	return o.PreScript, true
}

// HasPreScript returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasPreScript() bool {
	if o != nil && o.PreScript != nil {
		return true
	}

	return false
}

// SetPreScript gets a reference to the given string and assigns it to the PreScript field.
func (o *CloudsyncSyncOnetime0) SetPreScript(v string) {
	o.PreScript = &v
}

// GetPostScript returns the PostScript field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetPostScript() string {
	if o == nil || o.PostScript == nil {
		var ret string
		return ret
	}
	return *o.PostScript
}

// GetPostScriptOk returns a tuple with the PostScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetPostScriptOk() (*string, bool) {
	if o == nil || o.PostScript == nil {
		return nil, false
	}
	return o.PostScript, true
}

// HasPostScript returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasPostScript() bool {
	if o != nil && o.PostScript != nil {
		return true
	}

	return false
}

// SetPostScript gets a reference to the given string and assigns it to the PostScript field.
func (o *CloudsyncSyncOnetime0) SetPostScript(v string) {
	o.PostScript = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetArgs() string {
	if o == nil || o.Args == nil {
		var ret string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetArgsOk() (*string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given string and assigns it to the Args field.
func (o *CloudsyncSyncOnetime0) SetArgs(v string) {
	o.Args = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudsyncSyncOnetime0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudsyncSyncOnetime0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudsyncSyncOnetime0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudsyncSyncOnetime0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CloudsyncSyncOnetime0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.TransferMode != nil {
		toSerialize["transfer_mode"] = o.TransferMode
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Encryption != nil {
		toSerialize["encryption"] = o.Encryption
	}
	if o.FilenameEncryption != nil {
		toSerialize["filename_encryption"] = o.FilenameEncryption
	}
	if o.EncryptionPassword != nil {
		toSerialize["encryption_password"] = o.EncryptionPassword
	}
	if o.EncryptionSalt != nil {
		toSerialize["encryption_salt"] = o.EncryptionSalt
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.FollowSymlinks != nil {
		toSerialize["follow_symlinks"] = o.FollowSymlinks
	}
	if o.Transfers.IsSet() {
		toSerialize["transfers"] = o.Transfers.Get()
	}
	if o.Bwlimit != nil {
		toSerialize["bwlimit"] = o.Bwlimit
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Snapshot != nil {
		toSerialize["snapshot"] = o.Snapshot
	}
	if o.PreScript != nil {
		toSerialize["pre_script"] = o.PreScript
	}
	if o.PostScript != nil {
		toSerialize["post_script"] = o.PostScript
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableCloudsyncSyncOnetime0 struct {
	value *CloudsyncSyncOnetime0
	isSet bool
}

func (v NullableCloudsyncSyncOnetime0) Get() *CloudsyncSyncOnetime0 {
	return v.value
}

func (v *NullableCloudsyncSyncOnetime0) Set(val *CloudsyncSyncOnetime0) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudsyncSyncOnetime0) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudsyncSyncOnetime0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudsyncSyncOnetime0(val *CloudsyncSyncOnetime0) *NullableCloudsyncSyncOnetime0 {
	return &NullableCloudsyncSyncOnetime0{value: val, isSet: true}
}

func (v NullableCloudsyncSyncOnetime0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudsyncSyncOnetime0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


