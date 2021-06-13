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

// SshUpdate0 struct for SshUpdate0
type SshUpdate0 struct {
	Bindiface *[]string `json:"bindiface,omitempty"`
	Tcpport *int32 `json:"tcpport,omitempty"`
	Rootlogin *bool `json:"rootlogin,omitempty"`
	Passwordauth *bool `json:"passwordauth,omitempty"`
	Kerberosauth *bool `json:"kerberosauth,omitempty"`
	Tcpfwd *bool `json:"tcpfwd,omitempty"`
	Compression *bool `json:"compression,omitempty"`
	SftpLogLevel *string `json:"sftp_log_level,omitempty"`
	SftpLogFacility *string `json:"sftp_log_facility,omitempty"`
	WeakCiphers *[]string `json:"weak_ciphers,omitempty"`
	Options *string `json:"options,omitempty"`
}

// NewSshUpdate0 instantiates a new SshUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshUpdate0() *SshUpdate0 {
	this := SshUpdate0{}
	return &this
}

// NewSshUpdate0WithDefaults instantiates a new SshUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshUpdate0WithDefaults() *SshUpdate0 {
	this := SshUpdate0{}
	return &this
}

// GetBindiface returns the Bindiface field value if set, zero value otherwise.
func (o *SshUpdate0) GetBindiface() []string {
	if o == nil || o.Bindiface == nil {
		var ret []string
		return ret
	}
	return *o.Bindiface
}

// GetBindifaceOk returns a tuple with the Bindiface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetBindifaceOk() (*[]string, bool) {
	if o == nil || o.Bindiface == nil {
		return nil, false
	}
	return o.Bindiface, true
}

// HasBindiface returns a boolean if a field has been set.
func (o *SshUpdate0) HasBindiface() bool {
	if o != nil && o.Bindiface != nil {
		return true
	}

	return false
}

// SetBindiface gets a reference to the given []string and assigns it to the Bindiface field.
func (o *SshUpdate0) SetBindiface(v []string) {
	o.Bindiface = &v
}

// GetTcpport returns the Tcpport field value if set, zero value otherwise.
func (o *SshUpdate0) GetTcpport() int32 {
	if o == nil || o.Tcpport == nil {
		var ret int32
		return ret
	}
	return *o.Tcpport
}

// GetTcpportOk returns a tuple with the Tcpport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetTcpportOk() (*int32, bool) {
	if o == nil || o.Tcpport == nil {
		return nil, false
	}
	return o.Tcpport, true
}

// HasTcpport returns a boolean if a field has been set.
func (o *SshUpdate0) HasTcpport() bool {
	if o != nil && o.Tcpport != nil {
		return true
	}

	return false
}

// SetTcpport gets a reference to the given int32 and assigns it to the Tcpport field.
func (o *SshUpdate0) SetTcpport(v int32) {
	o.Tcpport = &v
}

// GetRootlogin returns the Rootlogin field value if set, zero value otherwise.
func (o *SshUpdate0) GetRootlogin() bool {
	if o == nil || o.Rootlogin == nil {
		var ret bool
		return ret
	}
	return *o.Rootlogin
}

// GetRootloginOk returns a tuple with the Rootlogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetRootloginOk() (*bool, bool) {
	if o == nil || o.Rootlogin == nil {
		return nil, false
	}
	return o.Rootlogin, true
}

// HasRootlogin returns a boolean if a field has been set.
func (o *SshUpdate0) HasRootlogin() bool {
	if o != nil && o.Rootlogin != nil {
		return true
	}

	return false
}

// SetRootlogin gets a reference to the given bool and assigns it to the Rootlogin field.
func (o *SshUpdate0) SetRootlogin(v bool) {
	o.Rootlogin = &v
}

// GetPasswordauth returns the Passwordauth field value if set, zero value otherwise.
func (o *SshUpdate0) GetPasswordauth() bool {
	if o == nil || o.Passwordauth == nil {
		var ret bool
		return ret
	}
	return *o.Passwordauth
}

// GetPasswordauthOk returns a tuple with the Passwordauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetPasswordauthOk() (*bool, bool) {
	if o == nil || o.Passwordauth == nil {
		return nil, false
	}
	return o.Passwordauth, true
}

// HasPasswordauth returns a boolean if a field has been set.
func (o *SshUpdate0) HasPasswordauth() bool {
	if o != nil && o.Passwordauth != nil {
		return true
	}

	return false
}

// SetPasswordauth gets a reference to the given bool and assigns it to the Passwordauth field.
func (o *SshUpdate0) SetPasswordauth(v bool) {
	o.Passwordauth = &v
}

// GetKerberosauth returns the Kerberosauth field value if set, zero value otherwise.
func (o *SshUpdate0) GetKerberosauth() bool {
	if o == nil || o.Kerberosauth == nil {
		var ret bool
		return ret
	}
	return *o.Kerberosauth
}

// GetKerberosauthOk returns a tuple with the Kerberosauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetKerberosauthOk() (*bool, bool) {
	if o == nil || o.Kerberosauth == nil {
		return nil, false
	}
	return o.Kerberosauth, true
}

// HasKerberosauth returns a boolean if a field has been set.
func (o *SshUpdate0) HasKerberosauth() bool {
	if o != nil && o.Kerberosauth != nil {
		return true
	}

	return false
}

// SetKerberosauth gets a reference to the given bool and assigns it to the Kerberosauth field.
func (o *SshUpdate0) SetKerberosauth(v bool) {
	o.Kerberosauth = &v
}

// GetTcpfwd returns the Tcpfwd field value if set, zero value otherwise.
func (o *SshUpdate0) GetTcpfwd() bool {
	if o == nil || o.Tcpfwd == nil {
		var ret bool
		return ret
	}
	return *o.Tcpfwd
}

// GetTcpfwdOk returns a tuple with the Tcpfwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetTcpfwdOk() (*bool, bool) {
	if o == nil || o.Tcpfwd == nil {
		return nil, false
	}
	return o.Tcpfwd, true
}

// HasTcpfwd returns a boolean if a field has been set.
func (o *SshUpdate0) HasTcpfwd() bool {
	if o != nil && o.Tcpfwd != nil {
		return true
	}

	return false
}

// SetTcpfwd gets a reference to the given bool and assigns it to the Tcpfwd field.
func (o *SshUpdate0) SetTcpfwd(v bool) {
	o.Tcpfwd = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *SshUpdate0) GetCompression() bool {
	if o == nil || o.Compression == nil {
		var ret bool
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetCompressionOk() (*bool, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *SshUpdate0) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given bool and assigns it to the Compression field.
func (o *SshUpdate0) SetCompression(v bool) {
	o.Compression = &v
}

// GetSftpLogLevel returns the SftpLogLevel field value if set, zero value otherwise.
func (o *SshUpdate0) GetSftpLogLevel() string {
	if o == nil || o.SftpLogLevel == nil {
		var ret string
		return ret
	}
	return *o.SftpLogLevel
}

// GetSftpLogLevelOk returns a tuple with the SftpLogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetSftpLogLevelOk() (*string, bool) {
	if o == nil || o.SftpLogLevel == nil {
		return nil, false
	}
	return o.SftpLogLevel, true
}

// HasSftpLogLevel returns a boolean if a field has been set.
func (o *SshUpdate0) HasSftpLogLevel() bool {
	if o != nil && o.SftpLogLevel != nil {
		return true
	}

	return false
}

// SetSftpLogLevel gets a reference to the given string and assigns it to the SftpLogLevel field.
func (o *SshUpdate0) SetSftpLogLevel(v string) {
	o.SftpLogLevel = &v
}

// GetSftpLogFacility returns the SftpLogFacility field value if set, zero value otherwise.
func (o *SshUpdate0) GetSftpLogFacility() string {
	if o == nil || o.SftpLogFacility == nil {
		var ret string
		return ret
	}
	return *o.SftpLogFacility
}

// GetSftpLogFacilityOk returns a tuple with the SftpLogFacility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetSftpLogFacilityOk() (*string, bool) {
	if o == nil || o.SftpLogFacility == nil {
		return nil, false
	}
	return o.SftpLogFacility, true
}

// HasSftpLogFacility returns a boolean if a field has been set.
func (o *SshUpdate0) HasSftpLogFacility() bool {
	if o != nil && o.SftpLogFacility != nil {
		return true
	}

	return false
}

// SetSftpLogFacility gets a reference to the given string and assigns it to the SftpLogFacility field.
func (o *SshUpdate0) SetSftpLogFacility(v string) {
	o.SftpLogFacility = &v
}

// GetWeakCiphers returns the WeakCiphers field value if set, zero value otherwise.
func (o *SshUpdate0) GetWeakCiphers() []string {
	if o == nil || o.WeakCiphers == nil {
		var ret []string
		return ret
	}
	return *o.WeakCiphers
}

// GetWeakCiphersOk returns a tuple with the WeakCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetWeakCiphersOk() (*[]string, bool) {
	if o == nil || o.WeakCiphers == nil {
		return nil, false
	}
	return o.WeakCiphers, true
}

// HasWeakCiphers returns a boolean if a field has been set.
func (o *SshUpdate0) HasWeakCiphers() bool {
	if o != nil && o.WeakCiphers != nil {
		return true
	}

	return false
}

// SetWeakCiphers gets a reference to the given []string and assigns it to the WeakCiphers field.
func (o *SshUpdate0) SetWeakCiphers(v []string) {
	o.WeakCiphers = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SshUpdate0) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshUpdate0) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SshUpdate0) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *SshUpdate0) SetOptions(v string) {
	o.Options = &v
}

func (o SshUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bindiface != nil {
		toSerialize["bindiface"] = o.Bindiface
	}
	if o.Tcpport != nil {
		toSerialize["tcpport"] = o.Tcpport
	}
	if o.Rootlogin != nil {
		toSerialize["rootlogin"] = o.Rootlogin
	}
	if o.Passwordauth != nil {
		toSerialize["passwordauth"] = o.Passwordauth
	}
	if o.Kerberosauth != nil {
		toSerialize["kerberosauth"] = o.Kerberosauth
	}
	if o.Tcpfwd != nil {
		toSerialize["tcpfwd"] = o.Tcpfwd
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.SftpLogLevel != nil {
		toSerialize["sftp_log_level"] = o.SftpLogLevel
	}
	if o.SftpLogFacility != nil {
		toSerialize["sftp_log_facility"] = o.SftpLogFacility
	}
	if o.WeakCiphers != nil {
		toSerialize["weak_ciphers"] = o.WeakCiphers
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableSshUpdate0 struct {
	value *SshUpdate0
	isSet bool
}

func (v NullableSshUpdate0) Get() *SshUpdate0 {
	return v.value
}

func (v *NullableSshUpdate0) Set(val *SshUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSshUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSshUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshUpdate0(val *SshUpdate0) *NullableSshUpdate0 {
	return &NullableSshUpdate0{value: val, isSet: true}
}

func (v NullableSshUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


