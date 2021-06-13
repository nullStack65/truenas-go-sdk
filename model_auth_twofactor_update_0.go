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

// AuthTwofactorUpdate0 struct for AuthTwofactorUpdate0
type AuthTwofactorUpdate0 struct {
	Enabled *bool `json:"enabled,omitempty"`
	OtpDigits *int32 `json:"otp_digits,omitempty"`
	Window *int32 `json:"window,omitempty"`
	Interval *int32 `json:"interval,omitempty"`
	Services *AuthTwofactorUpdate0Services `json:"services,omitempty"`
}

// NewAuthTwofactorUpdate0 instantiates a new AuthTwofactorUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTwofactorUpdate0() *AuthTwofactorUpdate0 {
	this := AuthTwofactorUpdate0{}
	return &this
}

// NewAuthTwofactorUpdate0WithDefaults instantiates a new AuthTwofactorUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTwofactorUpdate0WithDefaults() *AuthTwofactorUpdate0 {
	this := AuthTwofactorUpdate0{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthTwofactorUpdate0) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOtpDigits returns the OtpDigits field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetOtpDigits() int32 {
	if o == nil || o.OtpDigits == nil {
		var ret int32
		return ret
	}
	return *o.OtpDigits
}

// GetOtpDigitsOk returns a tuple with the OtpDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetOtpDigitsOk() (*int32, bool) {
	if o == nil || o.OtpDigits == nil {
		return nil, false
	}
	return o.OtpDigits, true
}

// HasOtpDigits returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasOtpDigits() bool {
	if o != nil && o.OtpDigits != nil {
		return true
	}

	return false
}

// SetOtpDigits gets a reference to the given int32 and assigns it to the OtpDigits field.
func (o *AuthTwofactorUpdate0) SetOtpDigits(v int32) {
	o.OtpDigits = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetWindow() int32 {
	if o == nil || o.Window == nil {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetWindowOk() (*int32, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *AuthTwofactorUpdate0) SetWindow(v int32) {
	o.Window = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetInterval() int32 {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetIntervalOk() (*int32, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *AuthTwofactorUpdate0) SetInterval(v int32) {
	o.Interval = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AuthTwofactorUpdate0) GetServices() AuthTwofactorUpdate0Services {
	if o == nil || o.Services == nil {
		var ret AuthTwofactorUpdate0Services
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthTwofactorUpdate0) GetServicesOk() (*AuthTwofactorUpdate0Services, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AuthTwofactorUpdate0) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given AuthTwofactorUpdate0Services and assigns it to the Services field.
func (o *AuthTwofactorUpdate0) SetServices(v AuthTwofactorUpdate0Services) {
	o.Services = &v
}

func (o AuthTwofactorUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.OtpDigits != nil {
		toSerialize["otp_digits"] = o.OtpDigits
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableAuthTwofactorUpdate0 struct {
	value *AuthTwofactorUpdate0
	isSet bool
}

func (v NullableAuthTwofactorUpdate0) Get() *AuthTwofactorUpdate0 {
	return v.value
}

func (v *NullableAuthTwofactorUpdate0) Set(val *AuthTwofactorUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTwofactorUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTwofactorUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTwofactorUpdate0(val *AuthTwofactorUpdate0) *NullableAuthTwofactorUpdate0 {
	return &NullableAuthTwofactorUpdate0{value: val, isSet: true}
}

func (v NullableAuthTwofactorUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTwofactorUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


