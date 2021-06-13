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

// InterfaceUpdate1 struct for InterfaceUpdate1
type InterfaceUpdate1 struct {
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DisableOffloadCapabilities *bool `json:"disable_offload_capabilities,omitempty"`
	Ipv4Dhcp *bool `json:"ipv4_dhcp,omitempty"`
	Ipv6Auto *bool `json:"ipv6_auto,omitempty"`
	Aliases *[]map[string]interface{} `json:"aliases,omitempty"`
	FailoverCritical *bool `json:"failover_critical,omitempty"`
	FailoverGroup NullableInt32 `json:"failover_group,omitempty"`
	FailoverVhid NullableInt32 `json:"failover_vhid,omitempty"`
	FailoverAliases *[]map[string]interface{} `json:"failover_aliases,omitempty"`
	FailoverVirtualAliases *[]map[string]interface{} `json:"failover_virtual_aliases,omitempty"`
	BridgeMembers *[]interface{} `json:"bridge_members,omitempty"`
	LagProtocol *string `json:"lag_protocol,omitempty"`
	LagPorts *[]string `json:"lag_ports,omitempty"`
	VlanParentInterface *string `json:"vlan_parent_interface,omitempty"`
	VlanTag *int32 `json:"vlan_tag,omitempty"`
	VlanPcp NullableInt32 `json:"vlan_pcp,omitempty"`
	Mtu NullableInt32 `json:"mtu,omitempty"`
	Options *string `json:"options,omitempty"`
}

// NewInterfaceUpdate1 instantiates a new InterfaceUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceUpdate1() *InterfaceUpdate1 {
	this := InterfaceUpdate1{}
	return &this
}

// NewInterfaceUpdate1WithDefaults instantiates a new InterfaceUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceUpdate1WithDefaults() *InterfaceUpdate1 {
	this := InterfaceUpdate1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterfaceUpdate1) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceUpdate1) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceUpdate1) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InterfaceUpdate1) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InterfaceUpdate1) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InterfaceUpdate1) UnsetDescription() {
	o.Description.Unset()
}

// GetDisableOffloadCapabilities returns the DisableOffloadCapabilities field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetDisableOffloadCapabilities() bool {
	if o == nil || o.DisableOffloadCapabilities == nil {
		var ret bool
		return ret
	}
	return *o.DisableOffloadCapabilities
}

// GetDisableOffloadCapabilitiesOk returns a tuple with the DisableOffloadCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetDisableOffloadCapabilitiesOk() (*bool, bool) {
	if o == nil || o.DisableOffloadCapabilities == nil {
		return nil, false
	}
	return o.DisableOffloadCapabilities, true
}

// HasDisableOffloadCapabilities returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasDisableOffloadCapabilities() bool {
	if o != nil && o.DisableOffloadCapabilities != nil {
		return true
	}

	return false
}

// SetDisableOffloadCapabilities gets a reference to the given bool and assigns it to the DisableOffloadCapabilities field.
func (o *InterfaceUpdate1) SetDisableOffloadCapabilities(v bool) {
	o.DisableOffloadCapabilities = &v
}

// GetIpv4Dhcp returns the Ipv4Dhcp field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetIpv4Dhcp() bool {
	if o == nil || o.Ipv4Dhcp == nil {
		var ret bool
		return ret
	}
	return *o.Ipv4Dhcp
}

// GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetIpv4DhcpOk() (*bool, bool) {
	if o == nil || o.Ipv4Dhcp == nil {
		return nil, false
	}
	return o.Ipv4Dhcp, true
}

// HasIpv4Dhcp returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasIpv4Dhcp() bool {
	if o != nil && o.Ipv4Dhcp != nil {
		return true
	}

	return false
}

// SetIpv4Dhcp gets a reference to the given bool and assigns it to the Ipv4Dhcp field.
func (o *InterfaceUpdate1) SetIpv4Dhcp(v bool) {
	o.Ipv4Dhcp = &v
}

// GetIpv6Auto returns the Ipv6Auto field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetIpv6Auto() bool {
	if o == nil || o.Ipv6Auto == nil {
		var ret bool
		return ret
	}
	return *o.Ipv6Auto
}

// GetIpv6AutoOk returns a tuple with the Ipv6Auto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetIpv6AutoOk() (*bool, bool) {
	if o == nil || o.Ipv6Auto == nil {
		return nil, false
	}
	return o.Ipv6Auto, true
}

// HasIpv6Auto returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasIpv6Auto() bool {
	if o != nil && o.Ipv6Auto != nil {
		return true
	}

	return false
}

// SetIpv6Auto gets a reference to the given bool and assigns it to the Ipv6Auto field.
func (o *InterfaceUpdate1) SetIpv6Auto(v bool) {
	o.Ipv6Auto = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetAliases() []map[string]interface{} {
	if o == nil || o.Aliases == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetAliasesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Aliases == nil {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasAliases() bool {
	if o != nil && o.Aliases != nil {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []map[string]interface{} and assigns it to the Aliases field.
func (o *InterfaceUpdate1) SetAliases(v []map[string]interface{}) {
	o.Aliases = &v
}

// GetFailoverCritical returns the FailoverCritical field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetFailoverCritical() bool {
	if o == nil || o.FailoverCritical == nil {
		var ret bool
		return ret
	}
	return *o.FailoverCritical
}

// GetFailoverCriticalOk returns a tuple with the FailoverCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetFailoverCriticalOk() (*bool, bool) {
	if o == nil || o.FailoverCritical == nil {
		return nil, false
	}
	return o.FailoverCritical, true
}

// HasFailoverCritical returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasFailoverCritical() bool {
	if o != nil && o.FailoverCritical != nil {
		return true
	}

	return false
}

// SetFailoverCritical gets a reference to the given bool and assigns it to the FailoverCritical field.
func (o *InterfaceUpdate1) SetFailoverCritical(v bool) {
	o.FailoverCritical = &v
}

// GetFailoverGroup returns the FailoverGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceUpdate1) GetFailoverGroup() int32 {
	if o == nil || o.FailoverGroup.Get() == nil {
		var ret int32
		return ret
	}
	return *o.FailoverGroup.Get()
}

// GetFailoverGroupOk returns a tuple with the FailoverGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceUpdate1) GetFailoverGroupOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailoverGroup.Get(), o.FailoverGroup.IsSet()
}

// HasFailoverGroup returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasFailoverGroup() bool {
	if o != nil && o.FailoverGroup.IsSet() {
		return true
	}

	return false
}

// SetFailoverGroup gets a reference to the given NullableInt32 and assigns it to the FailoverGroup field.
func (o *InterfaceUpdate1) SetFailoverGroup(v int32) {
	o.FailoverGroup.Set(&v)
}
// SetFailoverGroupNil sets the value for FailoverGroup to be an explicit nil
func (o *InterfaceUpdate1) SetFailoverGroupNil() {
	o.FailoverGroup.Set(nil)
}

// UnsetFailoverGroup ensures that no value is present for FailoverGroup, not even an explicit nil
func (o *InterfaceUpdate1) UnsetFailoverGroup() {
	o.FailoverGroup.Unset()
}

// GetFailoverVhid returns the FailoverVhid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceUpdate1) GetFailoverVhid() int32 {
	if o == nil || o.FailoverVhid.Get() == nil {
		var ret int32
		return ret
	}
	return *o.FailoverVhid.Get()
}

// GetFailoverVhidOk returns a tuple with the FailoverVhid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceUpdate1) GetFailoverVhidOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailoverVhid.Get(), o.FailoverVhid.IsSet()
}

// HasFailoverVhid returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasFailoverVhid() bool {
	if o != nil && o.FailoverVhid.IsSet() {
		return true
	}

	return false
}

// SetFailoverVhid gets a reference to the given NullableInt32 and assigns it to the FailoverVhid field.
func (o *InterfaceUpdate1) SetFailoverVhid(v int32) {
	o.FailoverVhid.Set(&v)
}
// SetFailoverVhidNil sets the value for FailoverVhid to be an explicit nil
func (o *InterfaceUpdate1) SetFailoverVhidNil() {
	o.FailoverVhid.Set(nil)
}

// UnsetFailoverVhid ensures that no value is present for FailoverVhid, not even an explicit nil
func (o *InterfaceUpdate1) UnsetFailoverVhid() {
	o.FailoverVhid.Unset()
}

// GetFailoverAliases returns the FailoverAliases field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetFailoverAliases() []map[string]interface{} {
	if o == nil || o.FailoverAliases == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.FailoverAliases
}

// GetFailoverAliasesOk returns a tuple with the FailoverAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetFailoverAliasesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.FailoverAliases == nil {
		return nil, false
	}
	return o.FailoverAliases, true
}

// HasFailoverAliases returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasFailoverAliases() bool {
	if o != nil && o.FailoverAliases != nil {
		return true
	}

	return false
}

// SetFailoverAliases gets a reference to the given []map[string]interface{} and assigns it to the FailoverAliases field.
func (o *InterfaceUpdate1) SetFailoverAliases(v []map[string]interface{}) {
	o.FailoverAliases = &v
}

// GetFailoverVirtualAliases returns the FailoverVirtualAliases field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetFailoverVirtualAliases() []map[string]interface{} {
	if o == nil || o.FailoverVirtualAliases == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.FailoverVirtualAliases
}

// GetFailoverVirtualAliasesOk returns a tuple with the FailoverVirtualAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetFailoverVirtualAliasesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.FailoverVirtualAliases == nil {
		return nil, false
	}
	return o.FailoverVirtualAliases, true
}

// HasFailoverVirtualAliases returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasFailoverVirtualAliases() bool {
	if o != nil && o.FailoverVirtualAliases != nil {
		return true
	}

	return false
}

// SetFailoverVirtualAliases gets a reference to the given []map[string]interface{} and assigns it to the FailoverVirtualAliases field.
func (o *InterfaceUpdate1) SetFailoverVirtualAliases(v []map[string]interface{}) {
	o.FailoverVirtualAliases = &v
}

// GetBridgeMembers returns the BridgeMembers field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetBridgeMembers() []interface{} {
	if o == nil || o.BridgeMembers == nil {
		var ret []interface{}
		return ret
	}
	return *o.BridgeMembers
}

// GetBridgeMembersOk returns a tuple with the BridgeMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetBridgeMembersOk() (*[]interface{}, bool) {
	if o == nil || o.BridgeMembers == nil {
		return nil, false
	}
	return o.BridgeMembers, true
}

// HasBridgeMembers returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasBridgeMembers() bool {
	if o != nil && o.BridgeMembers != nil {
		return true
	}

	return false
}

// SetBridgeMembers gets a reference to the given []interface{} and assigns it to the BridgeMembers field.
func (o *InterfaceUpdate1) SetBridgeMembers(v []interface{}) {
	o.BridgeMembers = &v
}

// GetLagProtocol returns the LagProtocol field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetLagProtocol() string {
	if o == nil || o.LagProtocol == nil {
		var ret string
		return ret
	}
	return *o.LagProtocol
}

// GetLagProtocolOk returns a tuple with the LagProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetLagProtocolOk() (*string, bool) {
	if o == nil || o.LagProtocol == nil {
		return nil, false
	}
	return o.LagProtocol, true
}

// HasLagProtocol returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasLagProtocol() bool {
	if o != nil && o.LagProtocol != nil {
		return true
	}

	return false
}

// SetLagProtocol gets a reference to the given string and assigns it to the LagProtocol field.
func (o *InterfaceUpdate1) SetLagProtocol(v string) {
	o.LagProtocol = &v
}

// GetLagPorts returns the LagPorts field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetLagPorts() []string {
	if o == nil || o.LagPorts == nil {
		var ret []string
		return ret
	}
	return *o.LagPorts
}

// GetLagPortsOk returns a tuple with the LagPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetLagPortsOk() (*[]string, bool) {
	if o == nil || o.LagPorts == nil {
		return nil, false
	}
	return o.LagPorts, true
}

// HasLagPorts returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasLagPorts() bool {
	if o != nil && o.LagPorts != nil {
		return true
	}

	return false
}

// SetLagPorts gets a reference to the given []string and assigns it to the LagPorts field.
func (o *InterfaceUpdate1) SetLagPorts(v []string) {
	o.LagPorts = &v
}

// GetVlanParentInterface returns the VlanParentInterface field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetVlanParentInterface() string {
	if o == nil || o.VlanParentInterface == nil {
		var ret string
		return ret
	}
	return *o.VlanParentInterface
}

// GetVlanParentInterfaceOk returns a tuple with the VlanParentInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetVlanParentInterfaceOk() (*string, bool) {
	if o == nil || o.VlanParentInterface == nil {
		return nil, false
	}
	return o.VlanParentInterface, true
}

// HasVlanParentInterface returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasVlanParentInterface() bool {
	if o != nil && o.VlanParentInterface != nil {
		return true
	}

	return false
}

// SetVlanParentInterface gets a reference to the given string and assigns it to the VlanParentInterface field.
func (o *InterfaceUpdate1) SetVlanParentInterface(v string) {
	o.VlanParentInterface = &v
}

// GetVlanTag returns the VlanTag field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetVlanTag() int32 {
	if o == nil || o.VlanTag == nil {
		var ret int32
		return ret
	}
	return *o.VlanTag
}

// GetVlanTagOk returns a tuple with the VlanTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetVlanTagOk() (*int32, bool) {
	if o == nil || o.VlanTag == nil {
		return nil, false
	}
	return o.VlanTag, true
}

// HasVlanTag returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasVlanTag() bool {
	if o != nil && o.VlanTag != nil {
		return true
	}

	return false
}

// SetVlanTag gets a reference to the given int32 and assigns it to the VlanTag field.
func (o *InterfaceUpdate1) SetVlanTag(v int32) {
	o.VlanTag = &v
}

// GetVlanPcp returns the VlanPcp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceUpdate1) GetVlanPcp() int32 {
	if o == nil || o.VlanPcp.Get() == nil {
		var ret int32
		return ret
	}
	return *o.VlanPcp.Get()
}

// GetVlanPcpOk returns a tuple with the VlanPcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceUpdate1) GetVlanPcpOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanPcp.Get(), o.VlanPcp.IsSet()
}

// HasVlanPcp returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasVlanPcp() bool {
	if o != nil && o.VlanPcp.IsSet() {
		return true
	}

	return false
}

// SetVlanPcp gets a reference to the given NullableInt32 and assigns it to the VlanPcp field.
func (o *InterfaceUpdate1) SetVlanPcp(v int32) {
	o.VlanPcp.Set(&v)
}
// SetVlanPcpNil sets the value for VlanPcp to be an explicit nil
func (o *InterfaceUpdate1) SetVlanPcpNil() {
	o.VlanPcp.Set(nil)
}

// UnsetVlanPcp ensures that no value is present for VlanPcp, not even an explicit nil
func (o *InterfaceUpdate1) UnsetVlanPcp() {
	o.VlanPcp.Unset()
}

// GetMtu returns the Mtu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceUpdate1) GetMtu() int32 {
	if o == nil || o.Mtu.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Mtu.Get()
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceUpdate1) GetMtuOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mtu.Get(), o.Mtu.IsSet()
}

// HasMtu returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasMtu() bool {
	if o != nil && o.Mtu.IsSet() {
		return true
	}

	return false
}

// SetMtu gets a reference to the given NullableInt32 and assigns it to the Mtu field.
func (o *InterfaceUpdate1) SetMtu(v int32) {
	o.Mtu.Set(&v)
}
// SetMtuNil sets the value for Mtu to be an explicit nil
func (o *InterfaceUpdate1) SetMtuNil() {
	o.Mtu.Set(nil)
}

// UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
func (o *InterfaceUpdate1) UnsetMtu() {
	o.Mtu.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InterfaceUpdate1) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceUpdate1) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InterfaceUpdate1) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *InterfaceUpdate1) SetOptions(v string) {
	o.Options = &v
}

func (o InterfaceUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisableOffloadCapabilities != nil {
		toSerialize["disable_offload_capabilities"] = o.DisableOffloadCapabilities
	}
	if o.Ipv4Dhcp != nil {
		toSerialize["ipv4_dhcp"] = o.Ipv4Dhcp
	}
	if o.Ipv6Auto != nil {
		toSerialize["ipv6_auto"] = o.Ipv6Auto
	}
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	if o.FailoverCritical != nil {
		toSerialize["failover_critical"] = o.FailoverCritical
	}
	if o.FailoverGroup.IsSet() {
		toSerialize["failover_group"] = o.FailoverGroup.Get()
	}
	if o.FailoverVhid.IsSet() {
		toSerialize["failover_vhid"] = o.FailoverVhid.Get()
	}
	if o.FailoverAliases != nil {
		toSerialize["failover_aliases"] = o.FailoverAliases
	}
	if o.FailoverVirtualAliases != nil {
		toSerialize["failover_virtual_aliases"] = o.FailoverVirtualAliases
	}
	if o.BridgeMembers != nil {
		toSerialize["bridge_members"] = o.BridgeMembers
	}
	if o.LagProtocol != nil {
		toSerialize["lag_protocol"] = o.LagProtocol
	}
	if o.LagPorts != nil {
		toSerialize["lag_ports"] = o.LagPorts
	}
	if o.VlanParentInterface != nil {
		toSerialize["vlan_parent_interface"] = o.VlanParentInterface
	}
	if o.VlanTag != nil {
		toSerialize["vlan_tag"] = o.VlanTag
	}
	if o.VlanPcp.IsSet() {
		toSerialize["vlan_pcp"] = o.VlanPcp.Get()
	}
	if o.Mtu.IsSet() {
		toSerialize["mtu"] = o.Mtu.Get()
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInterfaceUpdate1 struct {
	value *InterfaceUpdate1
	isSet bool
}

func (v NullableInterfaceUpdate1) Get() *InterfaceUpdate1 {
	return v.value
}

func (v *NullableInterfaceUpdate1) Set(val *InterfaceUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceUpdate1(val *InterfaceUpdate1) *NullableInterfaceUpdate1 {
	return &NullableInterfaceUpdate1{value: val, isSet: true}
}

func (v NullableInterfaceUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


