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

// SystemAdvancedUpdate0 struct for SystemAdvancedUpdate0
type SystemAdvancedUpdate0 struct {
	Advancedmode *bool `json:"advancedmode,omitempty"`
	Autotune *bool `json:"autotune,omitempty"`
	BootScrub *int32 `json:"boot_scrub,omitempty"`
	Consolemenu *bool `json:"consolemenu,omitempty"`
	Consolemsg *bool `json:"consolemsg,omitempty"`
	Debugkernel *bool `json:"debugkernel,omitempty"`
	FqdnSyslog *bool `json:"fqdn_syslog,omitempty"`
	Motd *string `json:"motd,omitempty"`
	Powerdaemon *bool `json:"powerdaemon,omitempty"`
	Serialconsole *bool `json:"serialconsole,omitempty"`
	Serialport *string `json:"serialport,omitempty"`
	Serialspeed *string `json:"serialspeed,omitempty"`
	Swapondrive *int32 `json:"swapondrive,omitempty"`
	Overprovision NullableInt32 `json:"overprovision,omitempty"`
	Traceback *bool `json:"traceback,omitempty"`
	Uploadcrash *bool `json:"uploadcrash,omitempty"`
	Anonstats *bool `json:"anonstats,omitempty"`
	SedUser *string `json:"sed_user,omitempty"`
	SedPasswd *string `json:"sed_passwd,omitempty"`
	Sysloglevel *string `json:"sysloglevel,omitempty"`
	Syslogserver *string `json:"syslogserver,omitempty"`
	SyslogTransport *string `json:"syslog_transport,omitempty"`
	SyslogTlsCertificate NullableInt32 `json:"syslog_tls_certificate,omitempty"`
}

// NewSystemAdvancedUpdate0 instantiates a new SystemAdvancedUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemAdvancedUpdate0() *SystemAdvancedUpdate0 {
	this := SystemAdvancedUpdate0{}
	return &this
}

// NewSystemAdvancedUpdate0WithDefaults instantiates a new SystemAdvancedUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemAdvancedUpdate0WithDefaults() *SystemAdvancedUpdate0 {
	this := SystemAdvancedUpdate0{}
	return &this
}

// GetAdvancedmode returns the Advancedmode field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetAdvancedmode() bool {
	if o == nil || o.Advancedmode == nil {
		var ret bool
		return ret
	}
	return *o.Advancedmode
}

// GetAdvancedmodeOk returns a tuple with the Advancedmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetAdvancedmodeOk() (*bool, bool) {
	if o == nil || o.Advancedmode == nil {
		return nil, false
	}
	return o.Advancedmode, true
}

// HasAdvancedmode returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasAdvancedmode() bool {
	if o != nil && o.Advancedmode != nil {
		return true
	}

	return false
}

// SetAdvancedmode gets a reference to the given bool and assigns it to the Advancedmode field.
func (o *SystemAdvancedUpdate0) SetAdvancedmode(v bool) {
	o.Advancedmode = &v
}

// GetAutotune returns the Autotune field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetAutotune() bool {
	if o == nil || o.Autotune == nil {
		var ret bool
		return ret
	}
	return *o.Autotune
}

// GetAutotuneOk returns a tuple with the Autotune field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetAutotuneOk() (*bool, bool) {
	if o == nil || o.Autotune == nil {
		return nil, false
	}
	return o.Autotune, true
}

// HasAutotune returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasAutotune() bool {
	if o != nil && o.Autotune != nil {
		return true
	}

	return false
}

// SetAutotune gets a reference to the given bool and assigns it to the Autotune field.
func (o *SystemAdvancedUpdate0) SetAutotune(v bool) {
	o.Autotune = &v
}

// GetBootScrub returns the BootScrub field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetBootScrub() int32 {
	if o == nil || o.BootScrub == nil {
		var ret int32
		return ret
	}
	return *o.BootScrub
}

// GetBootScrubOk returns a tuple with the BootScrub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetBootScrubOk() (*int32, bool) {
	if o == nil || o.BootScrub == nil {
		return nil, false
	}
	return o.BootScrub, true
}

// HasBootScrub returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasBootScrub() bool {
	if o != nil && o.BootScrub != nil {
		return true
	}

	return false
}

// SetBootScrub gets a reference to the given int32 and assigns it to the BootScrub field.
func (o *SystemAdvancedUpdate0) SetBootScrub(v int32) {
	o.BootScrub = &v
}

// GetConsolemenu returns the Consolemenu field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetConsolemenu() bool {
	if o == nil || o.Consolemenu == nil {
		var ret bool
		return ret
	}
	return *o.Consolemenu
}

// GetConsolemenuOk returns a tuple with the Consolemenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetConsolemenuOk() (*bool, bool) {
	if o == nil || o.Consolemenu == nil {
		return nil, false
	}
	return o.Consolemenu, true
}

// HasConsolemenu returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasConsolemenu() bool {
	if o != nil && o.Consolemenu != nil {
		return true
	}

	return false
}

// SetConsolemenu gets a reference to the given bool and assigns it to the Consolemenu field.
func (o *SystemAdvancedUpdate0) SetConsolemenu(v bool) {
	o.Consolemenu = &v
}

// GetConsolemsg returns the Consolemsg field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetConsolemsg() bool {
	if o == nil || o.Consolemsg == nil {
		var ret bool
		return ret
	}
	return *o.Consolemsg
}

// GetConsolemsgOk returns a tuple with the Consolemsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetConsolemsgOk() (*bool, bool) {
	if o == nil || o.Consolemsg == nil {
		return nil, false
	}
	return o.Consolemsg, true
}

// HasConsolemsg returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasConsolemsg() bool {
	if o != nil && o.Consolemsg != nil {
		return true
	}

	return false
}

// SetConsolemsg gets a reference to the given bool and assigns it to the Consolemsg field.
func (o *SystemAdvancedUpdate0) SetConsolemsg(v bool) {
	o.Consolemsg = &v
}

// GetDebugkernel returns the Debugkernel field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetDebugkernel() bool {
	if o == nil || o.Debugkernel == nil {
		var ret bool
		return ret
	}
	return *o.Debugkernel
}

// GetDebugkernelOk returns a tuple with the Debugkernel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetDebugkernelOk() (*bool, bool) {
	if o == nil || o.Debugkernel == nil {
		return nil, false
	}
	return o.Debugkernel, true
}

// HasDebugkernel returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasDebugkernel() bool {
	if o != nil && o.Debugkernel != nil {
		return true
	}

	return false
}

// SetDebugkernel gets a reference to the given bool and assigns it to the Debugkernel field.
func (o *SystemAdvancedUpdate0) SetDebugkernel(v bool) {
	o.Debugkernel = &v
}

// GetFqdnSyslog returns the FqdnSyslog field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetFqdnSyslog() bool {
	if o == nil || o.FqdnSyslog == nil {
		var ret bool
		return ret
	}
	return *o.FqdnSyslog
}

// GetFqdnSyslogOk returns a tuple with the FqdnSyslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetFqdnSyslogOk() (*bool, bool) {
	if o == nil || o.FqdnSyslog == nil {
		return nil, false
	}
	return o.FqdnSyslog, true
}

// HasFqdnSyslog returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasFqdnSyslog() bool {
	if o != nil && o.FqdnSyslog != nil {
		return true
	}

	return false
}

// SetFqdnSyslog gets a reference to the given bool and assigns it to the FqdnSyslog field.
func (o *SystemAdvancedUpdate0) SetFqdnSyslog(v bool) {
	o.FqdnSyslog = &v
}

// GetMotd returns the Motd field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetMotd() string {
	if o == nil || o.Motd == nil {
		var ret string
		return ret
	}
	return *o.Motd
}

// GetMotdOk returns a tuple with the Motd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetMotdOk() (*string, bool) {
	if o == nil || o.Motd == nil {
		return nil, false
	}
	return o.Motd, true
}

// HasMotd returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasMotd() bool {
	if o != nil && o.Motd != nil {
		return true
	}

	return false
}

// SetMotd gets a reference to the given string and assigns it to the Motd field.
func (o *SystemAdvancedUpdate0) SetMotd(v string) {
	o.Motd = &v
}

// GetPowerdaemon returns the Powerdaemon field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetPowerdaemon() bool {
	if o == nil || o.Powerdaemon == nil {
		var ret bool
		return ret
	}
	return *o.Powerdaemon
}

// GetPowerdaemonOk returns a tuple with the Powerdaemon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetPowerdaemonOk() (*bool, bool) {
	if o == nil || o.Powerdaemon == nil {
		return nil, false
	}
	return o.Powerdaemon, true
}

// HasPowerdaemon returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasPowerdaemon() bool {
	if o != nil && o.Powerdaemon != nil {
		return true
	}

	return false
}

// SetPowerdaemon gets a reference to the given bool and assigns it to the Powerdaemon field.
func (o *SystemAdvancedUpdate0) SetPowerdaemon(v bool) {
	o.Powerdaemon = &v
}

// GetSerialconsole returns the Serialconsole field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSerialconsole() bool {
	if o == nil || o.Serialconsole == nil {
		var ret bool
		return ret
	}
	return *o.Serialconsole
}

// GetSerialconsoleOk returns a tuple with the Serialconsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSerialconsoleOk() (*bool, bool) {
	if o == nil || o.Serialconsole == nil {
		return nil, false
	}
	return o.Serialconsole, true
}

// HasSerialconsole returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSerialconsole() bool {
	if o != nil && o.Serialconsole != nil {
		return true
	}

	return false
}

// SetSerialconsole gets a reference to the given bool and assigns it to the Serialconsole field.
func (o *SystemAdvancedUpdate0) SetSerialconsole(v bool) {
	o.Serialconsole = &v
}

// GetSerialport returns the Serialport field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSerialport() string {
	if o == nil || o.Serialport == nil {
		var ret string
		return ret
	}
	return *o.Serialport
}

// GetSerialportOk returns a tuple with the Serialport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSerialportOk() (*string, bool) {
	if o == nil || o.Serialport == nil {
		return nil, false
	}
	return o.Serialport, true
}

// HasSerialport returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSerialport() bool {
	if o != nil && o.Serialport != nil {
		return true
	}

	return false
}

// SetSerialport gets a reference to the given string and assigns it to the Serialport field.
func (o *SystemAdvancedUpdate0) SetSerialport(v string) {
	o.Serialport = &v
}

// GetSerialspeed returns the Serialspeed field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSerialspeed() string {
	if o == nil || o.Serialspeed == nil {
		var ret string
		return ret
	}
	return *o.Serialspeed
}

// GetSerialspeedOk returns a tuple with the Serialspeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSerialspeedOk() (*string, bool) {
	if o == nil || o.Serialspeed == nil {
		return nil, false
	}
	return o.Serialspeed, true
}

// HasSerialspeed returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSerialspeed() bool {
	if o != nil && o.Serialspeed != nil {
		return true
	}

	return false
}

// SetSerialspeed gets a reference to the given string and assigns it to the Serialspeed field.
func (o *SystemAdvancedUpdate0) SetSerialspeed(v string) {
	o.Serialspeed = &v
}

// GetSwapondrive returns the Swapondrive field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSwapondrive() int32 {
	if o == nil || o.Swapondrive == nil {
		var ret int32
		return ret
	}
	return *o.Swapondrive
}

// GetSwapondriveOk returns a tuple with the Swapondrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSwapondriveOk() (*int32, bool) {
	if o == nil || o.Swapondrive == nil {
		return nil, false
	}
	return o.Swapondrive, true
}

// HasSwapondrive returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSwapondrive() bool {
	if o != nil && o.Swapondrive != nil {
		return true
	}

	return false
}

// SetSwapondrive gets a reference to the given int32 and assigns it to the Swapondrive field.
func (o *SystemAdvancedUpdate0) SetSwapondrive(v int32) {
	o.Swapondrive = &v
}

// GetOverprovision returns the Overprovision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemAdvancedUpdate0) GetOverprovision() int32 {
	if o == nil || o.Overprovision.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Overprovision.Get()
}

// GetOverprovisionOk returns a tuple with the Overprovision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemAdvancedUpdate0) GetOverprovisionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Overprovision.Get(), o.Overprovision.IsSet()
}

// HasOverprovision returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasOverprovision() bool {
	if o != nil && o.Overprovision.IsSet() {
		return true
	}

	return false
}

// SetOverprovision gets a reference to the given NullableInt32 and assigns it to the Overprovision field.
func (o *SystemAdvancedUpdate0) SetOverprovision(v int32) {
	o.Overprovision.Set(&v)
}
// SetOverprovisionNil sets the value for Overprovision to be an explicit nil
func (o *SystemAdvancedUpdate0) SetOverprovisionNil() {
	o.Overprovision.Set(nil)
}

// UnsetOverprovision ensures that no value is present for Overprovision, not even an explicit nil
func (o *SystemAdvancedUpdate0) UnsetOverprovision() {
	o.Overprovision.Unset()
}

// GetTraceback returns the Traceback field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetTraceback() bool {
	if o == nil || o.Traceback == nil {
		var ret bool
		return ret
	}
	return *o.Traceback
}

// GetTracebackOk returns a tuple with the Traceback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetTracebackOk() (*bool, bool) {
	if o == nil || o.Traceback == nil {
		return nil, false
	}
	return o.Traceback, true
}

// HasTraceback returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasTraceback() bool {
	if o != nil && o.Traceback != nil {
		return true
	}

	return false
}

// SetTraceback gets a reference to the given bool and assigns it to the Traceback field.
func (o *SystemAdvancedUpdate0) SetTraceback(v bool) {
	o.Traceback = &v
}

// GetUploadcrash returns the Uploadcrash field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetUploadcrash() bool {
	if o == nil || o.Uploadcrash == nil {
		var ret bool
		return ret
	}
	return *o.Uploadcrash
}

// GetUploadcrashOk returns a tuple with the Uploadcrash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetUploadcrashOk() (*bool, bool) {
	if o == nil || o.Uploadcrash == nil {
		return nil, false
	}
	return o.Uploadcrash, true
}

// HasUploadcrash returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasUploadcrash() bool {
	if o != nil && o.Uploadcrash != nil {
		return true
	}

	return false
}

// SetUploadcrash gets a reference to the given bool and assigns it to the Uploadcrash field.
func (o *SystemAdvancedUpdate0) SetUploadcrash(v bool) {
	o.Uploadcrash = &v
}

// GetAnonstats returns the Anonstats field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetAnonstats() bool {
	if o == nil || o.Anonstats == nil {
		var ret bool
		return ret
	}
	return *o.Anonstats
}

// GetAnonstatsOk returns a tuple with the Anonstats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetAnonstatsOk() (*bool, bool) {
	if o == nil || o.Anonstats == nil {
		return nil, false
	}
	return o.Anonstats, true
}

// HasAnonstats returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasAnonstats() bool {
	if o != nil && o.Anonstats != nil {
		return true
	}

	return false
}

// SetAnonstats gets a reference to the given bool and assigns it to the Anonstats field.
func (o *SystemAdvancedUpdate0) SetAnonstats(v bool) {
	o.Anonstats = &v
}

// GetSedUser returns the SedUser field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSedUser() string {
	if o == nil || o.SedUser == nil {
		var ret string
		return ret
	}
	return *o.SedUser
}

// GetSedUserOk returns a tuple with the SedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSedUserOk() (*string, bool) {
	if o == nil || o.SedUser == nil {
		return nil, false
	}
	return o.SedUser, true
}

// HasSedUser returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSedUser() bool {
	if o != nil && o.SedUser != nil {
		return true
	}

	return false
}

// SetSedUser gets a reference to the given string and assigns it to the SedUser field.
func (o *SystemAdvancedUpdate0) SetSedUser(v string) {
	o.SedUser = &v
}

// GetSedPasswd returns the SedPasswd field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSedPasswd() string {
	if o == nil || o.SedPasswd == nil {
		var ret string
		return ret
	}
	return *o.SedPasswd
}

// GetSedPasswdOk returns a tuple with the SedPasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSedPasswdOk() (*string, bool) {
	if o == nil || o.SedPasswd == nil {
		return nil, false
	}
	return o.SedPasswd, true
}

// HasSedPasswd returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSedPasswd() bool {
	if o != nil && o.SedPasswd != nil {
		return true
	}

	return false
}

// SetSedPasswd gets a reference to the given string and assigns it to the SedPasswd field.
func (o *SystemAdvancedUpdate0) SetSedPasswd(v string) {
	o.SedPasswd = &v
}

// GetSysloglevel returns the Sysloglevel field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSysloglevel() string {
	if o == nil || o.Sysloglevel == nil {
		var ret string
		return ret
	}
	return *o.Sysloglevel
}

// GetSysloglevelOk returns a tuple with the Sysloglevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSysloglevelOk() (*string, bool) {
	if o == nil || o.Sysloglevel == nil {
		return nil, false
	}
	return o.Sysloglevel, true
}

// HasSysloglevel returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSysloglevel() bool {
	if o != nil && o.Sysloglevel != nil {
		return true
	}

	return false
}

// SetSysloglevel gets a reference to the given string and assigns it to the Sysloglevel field.
func (o *SystemAdvancedUpdate0) SetSysloglevel(v string) {
	o.Sysloglevel = &v
}

// GetSyslogserver returns the Syslogserver field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSyslogserver() string {
	if o == nil || o.Syslogserver == nil {
		var ret string
		return ret
	}
	return *o.Syslogserver
}

// GetSyslogserverOk returns a tuple with the Syslogserver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSyslogserverOk() (*string, bool) {
	if o == nil || o.Syslogserver == nil {
		return nil, false
	}
	return o.Syslogserver, true
}

// HasSyslogserver returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSyslogserver() bool {
	if o != nil && o.Syslogserver != nil {
		return true
	}

	return false
}

// SetSyslogserver gets a reference to the given string and assigns it to the Syslogserver field.
func (o *SystemAdvancedUpdate0) SetSyslogserver(v string) {
	o.Syslogserver = &v
}

// GetSyslogTransport returns the SyslogTransport field value if set, zero value otherwise.
func (o *SystemAdvancedUpdate0) GetSyslogTransport() string {
	if o == nil || o.SyslogTransport == nil {
		var ret string
		return ret
	}
	return *o.SyslogTransport
}

// GetSyslogTransportOk returns a tuple with the SyslogTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAdvancedUpdate0) GetSyslogTransportOk() (*string, bool) {
	if o == nil || o.SyslogTransport == nil {
		return nil, false
	}
	return o.SyslogTransport, true
}

// HasSyslogTransport returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSyslogTransport() bool {
	if o != nil && o.SyslogTransport != nil {
		return true
	}

	return false
}

// SetSyslogTransport gets a reference to the given string and assigns it to the SyslogTransport field.
func (o *SystemAdvancedUpdate0) SetSyslogTransport(v string) {
	o.SyslogTransport = &v
}

// GetSyslogTlsCertificate returns the SyslogTlsCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemAdvancedUpdate0) GetSyslogTlsCertificate() int32 {
	if o == nil || o.SyslogTlsCertificate.Get() == nil {
		var ret int32
		return ret
	}
	return *o.SyslogTlsCertificate.Get()
}

// GetSyslogTlsCertificateOk returns a tuple with the SyslogTlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemAdvancedUpdate0) GetSyslogTlsCertificateOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SyslogTlsCertificate.Get(), o.SyslogTlsCertificate.IsSet()
}

// HasSyslogTlsCertificate returns a boolean if a field has been set.
func (o *SystemAdvancedUpdate0) HasSyslogTlsCertificate() bool {
	if o != nil && o.SyslogTlsCertificate.IsSet() {
		return true
	}

	return false
}

// SetSyslogTlsCertificate gets a reference to the given NullableInt32 and assigns it to the SyslogTlsCertificate field.
func (o *SystemAdvancedUpdate0) SetSyslogTlsCertificate(v int32) {
	o.SyslogTlsCertificate.Set(&v)
}
// SetSyslogTlsCertificateNil sets the value for SyslogTlsCertificate to be an explicit nil
func (o *SystemAdvancedUpdate0) SetSyslogTlsCertificateNil() {
	o.SyslogTlsCertificate.Set(nil)
}

// UnsetSyslogTlsCertificate ensures that no value is present for SyslogTlsCertificate, not even an explicit nil
func (o *SystemAdvancedUpdate0) UnsetSyslogTlsCertificate() {
	o.SyslogTlsCertificate.Unset()
}

func (o SystemAdvancedUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Advancedmode != nil {
		toSerialize["advancedmode"] = o.Advancedmode
	}
	if o.Autotune != nil {
		toSerialize["autotune"] = o.Autotune
	}
	if o.BootScrub != nil {
		toSerialize["boot_scrub"] = o.BootScrub
	}
	if o.Consolemenu != nil {
		toSerialize["consolemenu"] = o.Consolemenu
	}
	if o.Consolemsg != nil {
		toSerialize["consolemsg"] = o.Consolemsg
	}
	if o.Debugkernel != nil {
		toSerialize["debugkernel"] = o.Debugkernel
	}
	if o.FqdnSyslog != nil {
		toSerialize["fqdn_syslog"] = o.FqdnSyslog
	}
	if o.Motd != nil {
		toSerialize["motd"] = o.Motd
	}
	if o.Powerdaemon != nil {
		toSerialize["powerdaemon"] = o.Powerdaemon
	}
	if o.Serialconsole != nil {
		toSerialize["serialconsole"] = o.Serialconsole
	}
	if o.Serialport != nil {
		toSerialize["serialport"] = o.Serialport
	}
	if o.Serialspeed != nil {
		toSerialize["serialspeed"] = o.Serialspeed
	}
	if o.Swapondrive != nil {
		toSerialize["swapondrive"] = o.Swapondrive
	}
	if o.Overprovision.IsSet() {
		toSerialize["overprovision"] = o.Overprovision.Get()
	}
	if o.Traceback != nil {
		toSerialize["traceback"] = o.Traceback
	}
	if o.Uploadcrash != nil {
		toSerialize["uploadcrash"] = o.Uploadcrash
	}
	if o.Anonstats != nil {
		toSerialize["anonstats"] = o.Anonstats
	}
	if o.SedUser != nil {
		toSerialize["sed_user"] = o.SedUser
	}
	if o.SedPasswd != nil {
		toSerialize["sed_passwd"] = o.SedPasswd
	}
	if o.Sysloglevel != nil {
		toSerialize["sysloglevel"] = o.Sysloglevel
	}
	if o.Syslogserver != nil {
		toSerialize["syslogserver"] = o.Syslogserver
	}
	if o.SyslogTransport != nil {
		toSerialize["syslog_transport"] = o.SyslogTransport
	}
	if o.SyslogTlsCertificate.IsSet() {
		toSerialize["syslog_tls_certificate"] = o.SyslogTlsCertificate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSystemAdvancedUpdate0 struct {
	value *SystemAdvancedUpdate0
	isSet bool
}

func (v NullableSystemAdvancedUpdate0) Get() *SystemAdvancedUpdate0 {
	return v.value
}

func (v *NullableSystemAdvancedUpdate0) Set(val *SystemAdvancedUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemAdvancedUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemAdvancedUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemAdvancedUpdate0(val *SystemAdvancedUpdate0) *NullableSystemAdvancedUpdate0 {
	return &NullableSystemAdvancedUpdate0{value: val, isSet: true}
}

func (v NullableSystemAdvancedUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemAdvancedUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


