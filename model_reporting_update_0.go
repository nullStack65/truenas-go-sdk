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

// ReportingUpdate0 struct for ReportingUpdate0
type ReportingUpdate0 struct {
	CpuInPercentage *bool `json:"cpu_in_percentage,omitempty"`
	Graphite *string `json:"graphite,omitempty"`
	GraphiteSeparateinstances *bool `json:"graphite_separateinstances,omitempty"`
	GraphAge *int32 `json:"graph_age,omitempty"`
	GraphPoints *int32 `json:"graph_points,omitempty"`
	ConfirmRrdDestroy *bool `json:"confirm_rrd_destroy,omitempty"`
}

// NewReportingUpdate0 instantiates a new ReportingUpdate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportingUpdate0() *ReportingUpdate0 {
	this := ReportingUpdate0{}
	return &this
}

// NewReportingUpdate0WithDefaults instantiates a new ReportingUpdate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportingUpdate0WithDefaults() *ReportingUpdate0 {
	this := ReportingUpdate0{}
	return &this
}

// GetCpuInPercentage returns the CpuInPercentage field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetCpuInPercentage() bool {
	if o == nil || o.CpuInPercentage == nil {
		var ret bool
		return ret
	}
	return *o.CpuInPercentage
}

// GetCpuInPercentageOk returns a tuple with the CpuInPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetCpuInPercentageOk() (*bool, bool) {
	if o == nil || o.CpuInPercentage == nil {
		return nil, false
	}
	return o.CpuInPercentage, true
}

// HasCpuInPercentage returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasCpuInPercentage() bool {
	if o != nil && o.CpuInPercentage != nil {
		return true
	}

	return false
}

// SetCpuInPercentage gets a reference to the given bool and assigns it to the CpuInPercentage field.
func (o *ReportingUpdate0) SetCpuInPercentage(v bool) {
	o.CpuInPercentage = &v
}

// GetGraphite returns the Graphite field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetGraphite() string {
	if o == nil || o.Graphite == nil {
		var ret string
		return ret
	}
	return *o.Graphite
}

// GetGraphiteOk returns a tuple with the Graphite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetGraphiteOk() (*string, bool) {
	if o == nil || o.Graphite == nil {
		return nil, false
	}
	return o.Graphite, true
}

// HasGraphite returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasGraphite() bool {
	if o != nil && o.Graphite != nil {
		return true
	}

	return false
}

// SetGraphite gets a reference to the given string and assigns it to the Graphite field.
func (o *ReportingUpdate0) SetGraphite(v string) {
	o.Graphite = &v
}

// GetGraphiteSeparateinstances returns the GraphiteSeparateinstances field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetGraphiteSeparateinstances() bool {
	if o == nil || o.GraphiteSeparateinstances == nil {
		var ret bool
		return ret
	}
	return *o.GraphiteSeparateinstances
}

// GetGraphiteSeparateinstancesOk returns a tuple with the GraphiteSeparateinstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetGraphiteSeparateinstancesOk() (*bool, bool) {
	if o == nil || o.GraphiteSeparateinstances == nil {
		return nil, false
	}
	return o.GraphiteSeparateinstances, true
}

// HasGraphiteSeparateinstances returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasGraphiteSeparateinstances() bool {
	if o != nil && o.GraphiteSeparateinstances != nil {
		return true
	}

	return false
}

// SetGraphiteSeparateinstances gets a reference to the given bool and assigns it to the GraphiteSeparateinstances field.
func (o *ReportingUpdate0) SetGraphiteSeparateinstances(v bool) {
	o.GraphiteSeparateinstances = &v
}

// GetGraphAge returns the GraphAge field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetGraphAge() int32 {
	if o == nil || o.GraphAge == nil {
		var ret int32
		return ret
	}
	return *o.GraphAge
}

// GetGraphAgeOk returns a tuple with the GraphAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetGraphAgeOk() (*int32, bool) {
	if o == nil || o.GraphAge == nil {
		return nil, false
	}
	return o.GraphAge, true
}

// HasGraphAge returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasGraphAge() bool {
	if o != nil && o.GraphAge != nil {
		return true
	}

	return false
}

// SetGraphAge gets a reference to the given int32 and assigns it to the GraphAge field.
func (o *ReportingUpdate0) SetGraphAge(v int32) {
	o.GraphAge = &v
}

// GetGraphPoints returns the GraphPoints field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetGraphPoints() int32 {
	if o == nil || o.GraphPoints == nil {
		var ret int32
		return ret
	}
	return *o.GraphPoints
}

// GetGraphPointsOk returns a tuple with the GraphPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetGraphPointsOk() (*int32, bool) {
	if o == nil || o.GraphPoints == nil {
		return nil, false
	}
	return o.GraphPoints, true
}

// HasGraphPoints returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasGraphPoints() bool {
	if o != nil && o.GraphPoints != nil {
		return true
	}

	return false
}

// SetGraphPoints gets a reference to the given int32 and assigns it to the GraphPoints field.
func (o *ReportingUpdate0) SetGraphPoints(v int32) {
	o.GraphPoints = &v
}

// GetConfirmRrdDestroy returns the ConfirmRrdDestroy field value if set, zero value otherwise.
func (o *ReportingUpdate0) GetConfirmRrdDestroy() bool {
	if o == nil || o.ConfirmRrdDestroy == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmRrdDestroy
}

// GetConfirmRrdDestroyOk returns a tuple with the ConfirmRrdDestroy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportingUpdate0) GetConfirmRrdDestroyOk() (*bool, bool) {
	if o == nil || o.ConfirmRrdDestroy == nil {
		return nil, false
	}
	return o.ConfirmRrdDestroy, true
}

// HasConfirmRrdDestroy returns a boolean if a field has been set.
func (o *ReportingUpdate0) HasConfirmRrdDestroy() bool {
	if o != nil && o.ConfirmRrdDestroy != nil {
		return true
	}

	return false
}

// SetConfirmRrdDestroy gets a reference to the given bool and assigns it to the ConfirmRrdDestroy field.
func (o *ReportingUpdate0) SetConfirmRrdDestroy(v bool) {
	o.ConfirmRrdDestroy = &v
}

func (o ReportingUpdate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuInPercentage != nil {
		toSerialize["cpu_in_percentage"] = o.CpuInPercentage
	}
	if o.Graphite != nil {
		toSerialize["graphite"] = o.Graphite
	}
	if o.GraphiteSeparateinstances != nil {
		toSerialize["graphite_separateinstances"] = o.GraphiteSeparateinstances
	}
	if o.GraphAge != nil {
		toSerialize["graph_age"] = o.GraphAge
	}
	if o.GraphPoints != nil {
		toSerialize["graph_points"] = o.GraphPoints
	}
	if o.ConfirmRrdDestroy != nil {
		toSerialize["confirm_rrd_destroy"] = o.ConfirmRrdDestroy
	}
	return json.Marshal(toSerialize)
}

type NullableReportingUpdate0 struct {
	value *ReportingUpdate0
	isSet bool
}

func (v NullableReportingUpdate0) Get() *ReportingUpdate0 {
	return v.value
}

func (v *NullableReportingUpdate0) Set(val *ReportingUpdate0) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingUpdate0) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingUpdate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingUpdate0(val *ReportingUpdate0) *NullableReportingUpdate0 {
	return &NullableReportingUpdate0{value: val, isSet: true}
}

func (v NullableReportingUpdate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingUpdate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


