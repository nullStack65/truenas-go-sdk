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

// CertificateCreate0CertExtensionsAuthorityKeyIdentifier struct for CertificateCreate0CertExtensionsAuthorityKeyIdentifier
type CertificateCreate0CertExtensionsAuthorityKeyIdentifier struct {
	AuthorityCertIssuer *bool `json:"authority_cert_issuer,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ExtensionCritical *bool `json:"extension_critical,omitempty"`
}

// NewCertificateCreate0CertExtensionsAuthorityKeyIdentifier instantiates a new CertificateCreate0CertExtensionsAuthorityKeyIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateCreate0CertExtensionsAuthorityKeyIdentifier() *CertificateCreate0CertExtensionsAuthorityKeyIdentifier {
	this := CertificateCreate0CertExtensionsAuthorityKeyIdentifier{}
	return &this
}

// NewCertificateCreate0CertExtensionsAuthorityKeyIdentifierWithDefaults instantiates a new CertificateCreate0CertExtensionsAuthorityKeyIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateCreate0CertExtensionsAuthorityKeyIdentifierWithDefaults() *CertificateCreate0CertExtensionsAuthorityKeyIdentifier {
	this := CertificateCreate0CertExtensionsAuthorityKeyIdentifier{}
	return &this
}

// GetAuthorityCertIssuer returns the AuthorityCertIssuer field value if set, zero value otherwise.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetAuthorityCertIssuer() bool {
	if o == nil || o.AuthorityCertIssuer == nil {
		var ret bool
		return ret
	}
	return *o.AuthorityCertIssuer
}

// GetAuthorityCertIssuerOk returns a tuple with the AuthorityCertIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetAuthorityCertIssuerOk() (*bool, bool) {
	if o == nil || o.AuthorityCertIssuer == nil {
		return nil, false
	}
	return o.AuthorityCertIssuer, true
}

// HasAuthorityCertIssuer returns a boolean if a field has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) HasAuthorityCertIssuer() bool {
	if o != nil && o.AuthorityCertIssuer != nil {
		return true
	}

	return false
}

// SetAuthorityCertIssuer gets a reference to the given bool and assigns it to the AuthorityCertIssuer field.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) SetAuthorityCertIssuer(v bool) {
	o.AuthorityCertIssuer = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExtensionCritical returns the ExtensionCritical field value if set, zero value otherwise.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetExtensionCritical() bool {
	if o == nil || o.ExtensionCritical == nil {
		var ret bool
		return ret
	}
	return *o.ExtensionCritical
}

// GetExtensionCriticalOk returns a tuple with the ExtensionCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) GetExtensionCriticalOk() (*bool, bool) {
	if o == nil || o.ExtensionCritical == nil {
		return nil, false
	}
	return o.ExtensionCritical, true
}

// HasExtensionCritical returns a boolean if a field has been set.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) HasExtensionCritical() bool {
	if o != nil && o.ExtensionCritical != nil {
		return true
	}

	return false
}

// SetExtensionCritical gets a reference to the given bool and assigns it to the ExtensionCritical field.
func (o *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) SetExtensionCritical(v bool) {
	o.ExtensionCritical = &v
}

func (o CertificateCreate0CertExtensionsAuthorityKeyIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorityCertIssuer != nil {
		toSerialize["authority_cert_issuer"] = o.AuthorityCertIssuer
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExtensionCritical != nil {
		toSerialize["extension_critical"] = o.ExtensionCritical
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier struct {
	value *CertificateCreate0CertExtensionsAuthorityKeyIdentifier
	isSet bool
}

func (v NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) Get() *CertificateCreate0CertExtensionsAuthorityKeyIdentifier {
	return v.value
}

func (v *NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) Set(val *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier(val *CertificateCreate0CertExtensionsAuthorityKeyIdentifier) *NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier {
	return &NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier{value: val, isSet: true}
}

func (v NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateCreate0CertExtensionsAuthorityKeyIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


