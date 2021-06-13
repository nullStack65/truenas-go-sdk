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

// IscsiInitiatorUpdate1 struct for IscsiInitiatorUpdate1
type IscsiInitiatorUpdate1 struct {
	Initiators *[]interface{} `json:"initiators,omitempty"`
	AuthNetwork *[]string `json:"auth_network,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewIscsiInitiatorUpdate1 instantiates a new IscsiInitiatorUpdate1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIscsiInitiatorUpdate1() *IscsiInitiatorUpdate1 {
	this := IscsiInitiatorUpdate1{}
	return &this
}

// NewIscsiInitiatorUpdate1WithDefaults instantiates a new IscsiInitiatorUpdate1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIscsiInitiatorUpdate1WithDefaults() *IscsiInitiatorUpdate1 {
	this := IscsiInitiatorUpdate1{}
	return &this
}

// GetInitiators returns the Initiators field value if set, zero value otherwise.
func (o *IscsiInitiatorUpdate1) GetInitiators() []interface{} {
	if o == nil || o.Initiators == nil {
		var ret []interface{}
		return ret
	}
	return *o.Initiators
}

// GetInitiatorsOk returns a tuple with the Initiators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiInitiatorUpdate1) GetInitiatorsOk() (*[]interface{}, bool) {
	if o == nil || o.Initiators == nil {
		return nil, false
	}
	return o.Initiators, true
}

// HasInitiators returns a boolean if a field has been set.
func (o *IscsiInitiatorUpdate1) HasInitiators() bool {
	if o != nil && o.Initiators != nil {
		return true
	}

	return false
}

// SetInitiators gets a reference to the given []interface{} and assigns it to the Initiators field.
func (o *IscsiInitiatorUpdate1) SetInitiators(v []interface{}) {
	o.Initiators = &v
}

// GetAuthNetwork returns the AuthNetwork field value if set, zero value otherwise.
func (o *IscsiInitiatorUpdate1) GetAuthNetwork() []string {
	if o == nil || o.AuthNetwork == nil {
		var ret []string
		return ret
	}
	return *o.AuthNetwork
}

// GetAuthNetworkOk returns a tuple with the AuthNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiInitiatorUpdate1) GetAuthNetworkOk() (*[]string, bool) {
	if o == nil || o.AuthNetwork == nil {
		return nil, false
	}
	return o.AuthNetwork, true
}

// HasAuthNetwork returns a boolean if a field has been set.
func (o *IscsiInitiatorUpdate1) HasAuthNetwork() bool {
	if o != nil && o.AuthNetwork != nil {
		return true
	}

	return false
}

// SetAuthNetwork gets a reference to the given []string and assigns it to the AuthNetwork field.
func (o *IscsiInitiatorUpdate1) SetAuthNetwork(v []string) {
	o.AuthNetwork = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IscsiInitiatorUpdate1) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IscsiInitiatorUpdate1) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IscsiInitiatorUpdate1) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *IscsiInitiatorUpdate1) SetComment(v string) {
	o.Comment = &v
}

func (o IscsiInitiatorUpdate1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Initiators != nil {
		toSerialize["initiators"] = o.Initiators
	}
	if o.AuthNetwork != nil {
		toSerialize["auth_network"] = o.AuthNetwork
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableIscsiInitiatorUpdate1 struct {
	value *IscsiInitiatorUpdate1
	isSet bool
}

func (v NullableIscsiInitiatorUpdate1) Get() *IscsiInitiatorUpdate1 {
	return v.value
}

func (v *NullableIscsiInitiatorUpdate1) Set(val *IscsiInitiatorUpdate1) {
	v.value = val
	v.isSet = true
}

func (v NullableIscsiInitiatorUpdate1) IsSet() bool {
	return v.isSet
}

func (v *NullableIscsiInitiatorUpdate1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIscsiInitiatorUpdate1(val *IscsiInitiatorUpdate1) *NullableIscsiInitiatorUpdate1 {
	return &NullableIscsiInitiatorUpdate1{value: val, isSet: true}
}

func (v NullableIscsiInitiatorUpdate1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIscsiInitiatorUpdate1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


