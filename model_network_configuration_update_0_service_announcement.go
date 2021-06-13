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

// NetworkConfigurationUpdate0ServiceAnnouncement struct for NetworkConfigurationUpdate0ServiceAnnouncement
type NetworkConfigurationUpdate0ServiceAnnouncement struct {
	Netbios *bool `json:"netbios,omitempty"`
	Mdns *bool `json:"mdns,omitempty"`
	Wsd *bool `json:"wsd,omitempty"`
}

// NewNetworkConfigurationUpdate0ServiceAnnouncement instantiates a new NetworkConfigurationUpdate0ServiceAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfigurationUpdate0ServiceAnnouncement() *NetworkConfigurationUpdate0ServiceAnnouncement {
	this := NetworkConfigurationUpdate0ServiceAnnouncement{}
	return &this
}

// NewNetworkConfigurationUpdate0ServiceAnnouncementWithDefaults instantiates a new NetworkConfigurationUpdate0ServiceAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigurationUpdate0ServiceAnnouncementWithDefaults() *NetworkConfigurationUpdate0ServiceAnnouncement {
	this := NetworkConfigurationUpdate0ServiceAnnouncement{}
	return &this
}

// GetNetbios returns the Netbios field value if set, zero value otherwise.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetNetbios() bool {
	if o == nil || o.Netbios == nil {
		var ret bool
		return ret
	}
	return *o.Netbios
}

// GetNetbiosOk returns a tuple with the Netbios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetNetbiosOk() (*bool, bool) {
	if o == nil || o.Netbios == nil {
		return nil, false
	}
	return o.Netbios, true
}

// HasNetbios returns a boolean if a field has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) HasNetbios() bool {
	if o != nil && o.Netbios != nil {
		return true
	}

	return false
}

// SetNetbios gets a reference to the given bool and assigns it to the Netbios field.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) SetNetbios(v bool) {
	o.Netbios = &v
}

// GetMdns returns the Mdns field value if set, zero value otherwise.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetMdns() bool {
	if o == nil || o.Mdns == nil {
		var ret bool
		return ret
	}
	return *o.Mdns
}

// GetMdnsOk returns a tuple with the Mdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetMdnsOk() (*bool, bool) {
	if o == nil || o.Mdns == nil {
		return nil, false
	}
	return o.Mdns, true
}

// HasMdns returns a boolean if a field has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) HasMdns() bool {
	if o != nil && o.Mdns != nil {
		return true
	}

	return false
}

// SetMdns gets a reference to the given bool and assigns it to the Mdns field.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) SetMdns(v bool) {
	o.Mdns = &v
}

// GetWsd returns the Wsd field value if set, zero value otherwise.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetWsd() bool {
	if o == nil || o.Wsd == nil {
		var ret bool
		return ret
	}
	return *o.Wsd
}

// GetWsdOk returns a tuple with the Wsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) GetWsdOk() (*bool, bool) {
	if o == nil || o.Wsd == nil {
		return nil, false
	}
	return o.Wsd, true
}

// HasWsd returns a boolean if a field has been set.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) HasWsd() bool {
	if o != nil && o.Wsd != nil {
		return true
	}

	return false
}

// SetWsd gets a reference to the given bool and assigns it to the Wsd field.
func (o *NetworkConfigurationUpdate0ServiceAnnouncement) SetWsd(v bool) {
	o.Wsd = &v
}

func (o NetworkConfigurationUpdate0ServiceAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Netbios != nil {
		toSerialize["netbios"] = o.Netbios
	}
	if o.Mdns != nil {
		toSerialize["mdns"] = o.Mdns
	}
	if o.Wsd != nil {
		toSerialize["wsd"] = o.Wsd
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkConfigurationUpdate0ServiceAnnouncement struct {
	value *NetworkConfigurationUpdate0ServiceAnnouncement
	isSet bool
}

func (v NullableNetworkConfigurationUpdate0ServiceAnnouncement) Get() *NetworkConfigurationUpdate0ServiceAnnouncement {
	return v.value
}

func (v *NullableNetworkConfigurationUpdate0ServiceAnnouncement) Set(val *NetworkConfigurationUpdate0ServiceAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfigurationUpdate0ServiceAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfigurationUpdate0ServiceAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfigurationUpdate0ServiceAnnouncement(val *NetworkConfigurationUpdate0ServiceAnnouncement) *NullableNetworkConfigurationUpdate0ServiceAnnouncement {
	return &NullableNetworkConfigurationUpdate0ServiceAnnouncement{value: val, isSet: true}
}

func (v NullableNetworkConfigurationUpdate0ServiceAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfigurationUpdate0ServiceAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


