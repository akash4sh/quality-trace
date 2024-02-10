/*
Qualitytrace

OpenAPI definition for Qualitytrace endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AwsXRay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsXRay{}

// AwsXRay struct for AwsXRay
type AwsXRay struct {
	Region          *string `json:"region,omitempty"`
	AccessKeyId     *string `json:"accessKeyId,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	SessionToken    *string `json:"sessionToken,omitempty"`
	UseDefaultAuth  *bool   `json:"useDefaultAuth,omitempty"`
}

// NewAwsXRay instantiates a new AwsXRay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsXRay() *AwsXRay {
	this := AwsXRay{}
	return &this
}

// NewAwsXRayWithDefaults instantiates a new AwsXRay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsXRayWithDefaults() *AwsXRay {
	this := AwsXRay{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AwsXRay) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsXRay) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AwsXRay) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AwsXRay) SetRegion(v string) {
	o.Region = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AwsXRay) GetAccessKeyId() string {
	if o == nil || isNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsXRay) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || isNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AwsXRay) HasAccessKeyId() bool {
	if o != nil && !isNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AwsXRay) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *AwsXRay) GetSecretAccessKey() string {
	if o == nil || isNil(o.SecretAccessKey) {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsXRay) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || isNil(o.SecretAccessKey) {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *AwsXRay) HasSecretAccessKey() bool {
	if o != nil && !isNil(o.SecretAccessKey) {
		return true
	}

	return false
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *AwsXRay) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *AwsXRay) GetSessionToken() string {
	if o == nil || isNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsXRay) GetSessionTokenOk() (*string, bool) {
	if o == nil || isNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *AwsXRay) HasSessionToken() bool {
	if o != nil && !isNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *AwsXRay) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetUseDefaultAuth returns the UseDefaultAuth field value if set, zero value otherwise.
func (o *AwsXRay) GetUseDefaultAuth() bool {
	if o == nil || isNil(o.UseDefaultAuth) {
		var ret bool
		return ret
	}
	return *o.UseDefaultAuth
}

// GetUseDefaultAuthOk returns a tuple with the UseDefaultAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsXRay) GetUseDefaultAuthOk() (*bool, bool) {
	if o == nil || isNil(o.UseDefaultAuth) {
		return nil, false
	}
	return o.UseDefaultAuth, true
}

// HasUseDefaultAuth returns a boolean if a field has been set.
func (o *AwsXRay) HasUseDefaultAuth() bool {
	if o != nil && !isNil(o.UseDefaultAuth) {
		return true
	}

	return false
}

// SetUseDefaultAuth gets a reference to the given bool and assigns it to the UseDefaultAuth field.
func (o *AwsXRay) SetUseDefaultAuth(v bool) {
	o.UseDefaultAuth = &v
}

func (o AwsXRay) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsXRay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.AccessKeyId) {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if !isNil(o.SecretAccessKey) {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if !isNil(o.SessionToken) {
		toSerialize["sessionToken"] = o.SessionToken
	}
	if !isNil(o.UseDefaultAuth) {
		toSerialize["useDefaultAuth"] = o.UseDefaultAuth
	}
	return toSerialize, nil
}

type NullableAwsXRay struct {
	value *AwsXRay
	isSet bool
}

func (v NullableAwsXRay) Get() *AwsXRay {
	return v.value
}

func (v *NullableAwsXRay) Set(val *AwsXRay) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsXRay) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsXRay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsXRay(val *AwsXRay) *NullableAwsXRay {
	return &NullableAwsXRay{value: val, isSet: true}
}

func (v NullableAwsXRay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsXRay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
