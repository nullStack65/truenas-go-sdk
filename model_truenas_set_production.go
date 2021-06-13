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

// TruenasSetProduction struct for TruenasSetProduction
type TruenasSetProduction struct {
	Production *bool `json:"production,omitempty"`
	AttachDebug *bool `json:"attach_debug,omitempty"`
}

// NewTruenasSetProduction instantiates a new TruenasSetProduction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTruenasSetProduction() *TruenasSetProduction {
	this := TruenasSetProduction{}
	var attachDebug bool = false
	this.AttachDebug = &attachDebug
	return &this
}

// NewTruenasSetProductionWithDefaults instantiates a new TruenasSetProduction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTruenasSetProductionWithDefaults() *TruenasSetProduction {
	this := TruenasSetProduction{}
	var attachDebug bool = false
	this.AttachDebug = &attachDebug
	return &this
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *TruenasSetProduction) GetProduction() bool {
	if o == nil || o.Production == nil {
		var ret bool
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruenasSetProduction) GetProductionOk() (*bool, bool) {
	if o == nil || o.Production == nil {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *TruenasSetProduction) HasProduction() bool {
	if o != nil && o.Production != nil {
		return true
	}

	return false
}

// SetProduction gets a reference to the given bool and assigns it to the Production field.
func (o *TruenasSetProduction) SetProduction(v bool) {
	o.Production = &v
}

// GetAttachDebug returns the AttachDebug field value if set, zero value otherwise.
func (o *TruenasSetProduction) GetAttachDebug() bool {
	if o == nil || o.AttachDebug == nil {
		var ret bool
		return ret
	}
	return *o.AttachDebug
}

// GetAttachDebugOk returns a tuple with the AttachDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruenasSetProduction) GetAttachDebugOk() (*bool, bool) {
	if o == nil || o.AttachDebug == nil {
		return nil, false
	}
	return o.AttachDebug, true
}

// HasAttachDebug returns a boolean if a field has been set.
func (o *TruenasSetProduction) HasAttachDebug() bool {
	if o != nil && o.AttachDebug != nil {
		return true
	}

	return false
}

// SetAttachDebug gets a reference to the given bool and assigns it to the AttachDebug field.
func (o *TruenasSetProduction) SetAttachDebug(v bool) {
	o.AttachDebug = &v
}

func (o TruenasSetProduction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Production != nil {
		toSerialize["production"] = o.Production
	}
	if o.AttachDebug != nil {
		toSerialize["attach_debug"] = o.AttachDebug
	}
	return json.Marshal(toSerialize)
}

type NullableTruenasSetProduction struct {
	value *TruenasSetProduction
	isSet bool
}

func (v NullableTruenasSetProduction) Get() *TruenasSetProduction {
	return v.value
}

func (v *NullableTruenasSetProduction) Set(val *TruenasSetProduction) {
	v.value = val
	v.isSet = true
}

func (v NullableTruenasSetProduction) IsSet() bool {
	return v.isSet
}

func (v *NullableTruenasSetProduction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTruenasSetProduction(val *TruenasSetProduction) *NullableTruenasSetProduction {
	return &NullableTruenasSetProduction{value: val, isSet: true}
}

func (v NullableTruenasSetProduction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTruenasSetProduction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


