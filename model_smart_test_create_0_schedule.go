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

// SmartTestCreate0Schedule struct for SmartTestCreate0Schedule
type SmartTestCreate0Schedule struct {
	Hour *string `json:"hour,omitempty"`
	Dom *string `json:"dom,omitempty"`
	Month *string `json:"month,omitempty"`
	Dow *string `json:"dow,omitempty"`
}

// NewSmartTestCreate0Schedule instantiates a new SmartTestCreate0Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartTestCreate0Schedule() *SmartTestCreate0Schedule {
	this := SmartTestCreate0Schedule{}
	return &this
}

// NewSmartTestCreate0ScheduleWithDefaults instantiates a new SmartTestCreate0Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartTestCreate0ScheduleWithDefaults() *SmartTestCreate0Schedule {
	this := SmartTestCreate0Schedule{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *SmartTestCreate0Schedule) GetHour() string {
	if o == nil || o.Hour == nil {
		var ret string
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestCreate0Schedule) GetHourOk() (*string, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *SmartTestCreate0Schedule) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given string and assigns it to the Hour field.
func (o *SmartTestCreate0Schedule) SetHour(v string) {
	o.Hour = &v
}

// GetDom returns the Dom field value if set, zero value otherwise.
func (o *SmartTestCreate0Schedule) GetDom() string {
	if o == nil || o.Dom == nil {
		var ret string
		return ret
	}
	return *o.Dom
}

// GetDomOk returns a tuple with the Dom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestCreate0Schedule) GetDomOk() (*string, bool) {
	if o == nil || o.Dom == nil {
		return nil, false
	}
	return o.Dom, true
}

// HasDom returns a boolean if a field has been set.
func (o *SmartTestCreate0Schedule) HasDom() bool {
	if o != nil && o.Dom != nil {
		return true
	}

	return false
}

// SetDom gets a reference to the given string and assigns it to the Dom field.
func (o *SmartTestCreate0Schedule) SetDom(v string) {
	o.Dom = &v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *SmartTestCreate0Schedule) GetMonth() string {
	if o == nil || o.Month == nil {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestCreate0Schedule) GetMonthOk() (*string, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *SmartTestCreate0Schedule) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *SmartTestCreate0Schedule) SetMonth(v string) {
	o.Month = &v
}

// GetDow returns the Dow field value if set, zero value otherwise.
func (o *SmartTestCreate0Schedule) GetDow() string {
	if o == nil || o.Dow == nil {
		var ret string
		return ret
	}
	return *o.Dow
}

// GetDowOk returns a tuple with the Dow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartTestCreate0Schedule) GetDowOk() (*string, bool) {
	if o == nil || o.Dow == nil {
		return nil, false
	}
	return o.Dow, true
}

// HasDow returns a boolean if a field has been set.
func (o *SmartTestCreate0Schedule) HasDow() bool {
	if o != nil && o.Dow != nil {
		return true
	}

	return false
}

// SetDow gets a reference to the given string and assigns it to the Dow field.
func (o *SmartTestCreate0Schedule) SetDow(v string) {
	o.Dow = &v
}

func (o SmartTestCreate0Schedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.Dom != nil {
		toSerialize["dom"] = o.Dom
	}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Dow != nil {
		toSerialize["dow"] = o.Dow
	}
	return json.Marshal(toSerialize)
}

type NullableSmartTestCreate0Schedule struct {
	value *SmartTestCreate0Schedule
	isSet bool
}

func (v NullableSmartTestCreate0Schedule) Get() *SmartTestCreate0Schedule {
	return v.value
}

func (v *NullableSmartTestCreate0Schedule) Set(val *SmartTestCreate0Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartTestCreate0Schedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartTestCreate0Schedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartTestCreate0Schedule(val *SmartTestCreate0Schedule) *NullableSmartTestCreate0Schedule {
	return &NullableSmartTestCreate0Schedule{value: val, isSet: true}
}

func (v NullableSmartTestCreate0Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartTestCreate0Schedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


