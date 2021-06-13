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

// JailCreate0 struct for JailCreate0
type JailCreate0 struct {
	Release *string `json:"release,omitempty"`
	Template *string `json:"template,omitempty"`
	Pkglist *[]string `json:"pkglist,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Basejail *bool `json:"basejail,omitempty"`
	Empty *bool `json:"empty,omitempty"`
	Short *bool `json:"short,omitempty"`
	Props *[]interface{} `json:"props,omitempty"`
	Https *bool `json:"https,omitempty"`
}

// NewJailCreate0 instantiates a new JailCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJailCreate0() *JailCreate0 {
	this := JailCreate0{}
	return &this
}

// NewJailCreate0WithDefaults instantiates a new JailCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJailCreate0WithDefaults() *JailCreate0 {
	this := JailCreate0{}
	return &this
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *JailCreate0) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *JailCreate0) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *JailCreate0) SetRelease(v string) {
	o.Release = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *JailCreate0) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *JailCreate0) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *JailCreate0) SetTemplate(v string) {
	o.Template = &v
}

// GetPkglist returns the Pkglist field value if set, zero value otherwise.
func (o *JailCreate0) GetPkglist() []string {
	if o == nil || o.Pkglist == nil {
		var ret []string
		return ret
	}
	return *o.Pkglist
}

// GetPkglistOk returns a tuple with the Pkglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetPkglistOk() (*[]string, bool) {
	if o == nil || o.Pkglist == nil {
		return nil, false
	}
	return o.Pkglist, true
}

// HasPkglist returns a boolean if a field has been set.
func (o *JailCreate0) HasPkglist() bool {
	if o != nil && o.Pkglist != nil {
		return true
	}

	return false
}

// SetPkglist gets a reference to the given []string and assigns it to the Pkglist field.
func (o *JailCreate0) SetPkglist(v []string) {
	o.Pkglist = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *JailCreate0) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *JailCreate0) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *JailCreate0) SetUuid(v string) {
	o.Uuid = &v
}

// GetBasejail returns the Basejail field value if set, zero value otherwise.
func (o *JailCreate0) GetBasejail() bool {
	if o == nil || o.Basejail == nil {
		var ret bool
		return ret
	}
	return *o.Basejail
}

// GetBasejailOk returns a tuple with the Basejail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetBasejailOk() (*bool, bool) {
	if o == nil || o.Basejail == nil {
		return nil, false
	}
	return o.Basejail, true
}

// HasBasejail returns a boolean if a field has been set.
func (o *JailCreate0) HasBasejail() bool {
	if o != nil && o.Basejail != nil {
		return true
	}

	return false
}

// SetBasejail gets a reference to the given bool and assigns it to the Basejail field.
func (o *JailCreate0) SetBasejail(v bool) {
	o.Basejail = &v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *JailCreate0) GetEmpty() bool {
	if o == nil || o.Empty == nil {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetEmptyOk() (*bool, bool) {
	if o == nil || o.Empty == nil {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *JailCreate0) HasEmpty() bool {
	if o != nil && o.Empty != nil {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *JailCreate0) SetEmpty(v bool) {
	o.Empty = &v
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *JailCreate0) GetShort() bool {
	if o == nil || o.Short == nil {
		var ret bool
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetShortOk() (*bool, bool) {
	if o == nil || o.Short == nil {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *JailCreate0) HasShort() bool {
	if o != nil && o.Short != nil {
		return true
	}

	return false
}

// SetShort gets a reference to the given bool and assigns it to the Short field.
func (o *JailCreate0) SetShort(v bool) {
	o.Short = &v
}

// GetProps returns the Props field value if set, zero value otherwise.
func (o *JailCreate0) GetProps() []interface{} {
	if o == nil || o.Props == nil {
		var ret []interface{}
		return ret
	}
	return *o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetPropsOk() (*[]interface{}, bool) {
	if o == nil || o.Props == nil {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *JailCreate0) HasProps() bool {
	if o != nil && o.Props != nil {
		return true
	}

	return false
}

// SetProps gets a reference to the given []interface{} and assigns it to the Props field.
func (o *JailCreate0) SetProps(v []interface{}) {
	o.Props = &v
}

// GetHttps returns the Https field value if set, zero value otherwise.
func (o *JailCreate0) GetHttps() bool {
	if o == nil || o.Https == nil {
		var ret bool
		return ret
	}
	return *o.Https
}

// GetHttpsOk returns a tuple with the Https field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailCreate0) GetHttpsOk() (*bool, bool) {
	if o == nil || o.Https == nil {
		return nil, false
	}
	return o.Https, true
}

// HasHttps returns a boolean if a field has been set.
func (o *JailCreate0) HasHttps() bool {
	if o != nil && o.Https != nil {
		return true
	}

	return false
}

// SetHttps gets a reference to the given bool and assigns it to the Https field.
func (o *JailCreate0) SetHttps(v bool) {
	o.Https = &v
}

func (o JailCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Release != nil {
		toSerialize["release"] = o.Release
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Pkglist != nil {
		toSerialize["pkglist"] = o.Pkglist
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Basejail != nil {
		toSerialize["basejail"] = o.Basejail
	}
	if o.Empty != nil {
		toSerialize["empty"] = o.Empty
	}
	if o.Short != nil {
		toSerialize["short"] = o.Short
	}
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	if o.Https != nil {
		toSerialize["https"] = o.Https
	}
	return json.Marshal(toSerialize)
}

type NullableJailCreate0 struct {
	value *JailCreate0
	isSet bool
}

func (v NullableJailCreate0) Get() *JailCreate0 {
	return v.value
}

func (v *NullableJailCreate0) Set(val *JailCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableJailCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableJailCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJailCreate0(val *JailCreate0) *NullableJailCreate0 {
	return &NullableJailCreate0{value: val, isSet: true}
}

func (v NullableJailCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJailCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


