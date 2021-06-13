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

// IpmiUpdate1 struct for IpmiUpdate1
type IpmiUpdate1 struct {
	Ipaddress *string `json:"ipaddress,omitempty"`
	Netmask *string `json:"netmask,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Password *string `json:"password,omitempty"`
	Dhcp *bool `json:"dhcp,omitempty"`
	Vlan NullableInt32 `json:"vlan,omitempty"`
}

// NewIpmiUpdate1 instantiates a new IpmiUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpmiUpdate1() *IpmiUpdate1 {
	this := IpmiUpdate1{}
	return &this
}

// NewIpmiUpdate1WithDefaults instantiates a new IpmiUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpmiUpdate1WithDefaults() *IpmiUpdate1 {
	this := IpmiUpdate1{}
	return &this
}

// GetIpaddress returns the Ipaddress field value if set, zero value otherwise.
func (o *IpmiUpdate1) GetIpaddress() string {
	if o == nil || o.Ipaddress == nil {
		var ret string
		return ret
	}
	return *o.Ipaddress
}

// GetIpaddressOk returns a tuple with the Ipaddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmiUpdate1) GetIpaddressOk() (*string, bool) {
	if o == nil || o.Ipaddress == nil {
		return nil, false
	}
	return o.Ipaddress, true
}

// HasIpaddress returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasIpaddress() bool {
	if o != nil && o.Ipaddress != nil {
		return true
	}

	return false
}

// SetIpaddress gets a reference to the given string and assigns it to the Ipaddress field.
func (o *IpmiUpdate1) SetIpaddress(v string) {
	o.Ipaddress = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *IpmiUpdate1) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmiUpdate1) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *IpmiUpdate1) SetNetmask(v string) {
	o.Netmask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *IpmiUpdate1) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmiUpdate1) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *IpmiUpdate1) SetGateway(v string) {
	o.Gateway = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IpmiUpdate1) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmiUpdate1) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IpmiUpdate1) SetPassword(v string) {
	o.Password = &v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *IpmiUpdate1) GetDhcp() bool {
	if o == nil || o.Dhcp == nil {
		var ret bool
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpmiUpdate1) GetDhcpOk() (*bool, bool) {
	if o == nil || o.Dhcp == nil {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given bool and assigns it to the Dhcp field.
func (o *IpmiUpdate1) SetDhcp(v bool) {
	o.Dhcp = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpmiUpdate1) GetVlan() int32 {
	if o == nil || o.Vlan.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpmiUpdate1) GetVlanOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *IpmiUpdate1) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableInt32 and assigns it to the Vlan field.
func (o *IpmiUpdate1) SetVlan(v int32) {
	o.Vlan.Set(&v)
}
// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *IpmiUpdate1) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *IpmiUpdate1) UnsetVlan() {
	o.Vlan.Unset()
}

func (o IpmiUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ipaddress != nil {
		toSerialize["ipaddress"] = o.Ipaddress
	}
	if o.Netmask != nil {
		toSerialize["netmask"] = o.Netmask
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIpmiUpdate1 struct {
	value *IpmiUpdate1
	isSet bool
}

func (v NullableIpmiUpdate1) Get() *IpmiUpdate1 {
	return v.value
}

func (v *NullableIpmiUpdate1) Set(val *IpmiUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableIpmiUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableIpmiUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpmiUpdate1(val *IpmiUpdate1) *NullableIpmiUpdate1 {
	return &NullableIpmiUpdate1{value: val, isSet: true}
}

func (v NullableIpmiUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpmiUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


