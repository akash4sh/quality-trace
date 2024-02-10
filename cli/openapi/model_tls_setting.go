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

// checks if the TLSSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TLSSetting{}

// TLSSetting struct for TLSSetting
type TLSSetting struct {
	CAFile     *string `json:"cAFile,omitempty"`
	CertFile   *string `json:"certFile,omitempty"`
	KeyFile    *string `json:"keyFile,omitempty"`
	MinVersion *string `json:"minVersion,omitempty"`
	MaxVersion *string `json:"maxVersion,omitempty"`
}

// NewTLSSetting instantiates a new TLSSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSSetting() *TLSSetting {
	this := TLSSetting{}
	return &this
}

// NewTLSSettingWithDefaults instantiates a new TLSSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSSettingWithDefaults() *TLSSetting {
	this := TLSSetting{}
	return &this
}

// GetCAFile returns the CAFile field value if set, zero value otherwise.
func (o *TLSSetting) GetCAFile() string {
	if o == nil || isNil(o.CAFile) {
		var ret string
		return ret
	}
	return *o.CAFile
}

// GetCAFileOk returns a tuple with the CAFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSetting) GetCAFileOk() (*string, bool) {
	if o == nil || isNil(o.CAFile) {
		return nil, false
	}
	return o.CAFile, true
}

// HasCAFile returns a boolean if a field has been set.
func (o *TLSSetting) HasCAFile() bool {
	if o != nil && !isNil(o.CAFile) {
		return true
	}

	return false
}

// SetCAFile gets a reference to the given string and assigns it to the CAFile field.
func (o *TLSSetting) SetCAFile(v string) {
	o.CAFile = &v
}

// GetCertFile returns the CertFile field value if set, zero value otherwise.
func (o *TLSSetting) GetCertFile() string {
	if o == nil || isNil(o.CertFile) {
		var ret string
		return ret
	}
	return *o.CertFile
}

// GetCertFileOk returns a tuple with the CertFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSetting) GetCertFileOk() (*string, bool) {
	if o == nil || isNil(o.CertFile) {
		return nil, false
	}
	return o.CertFile, true
}

// HasCertFile returns a boolean if a field has been set.
func (o *TLSSetting) HasCertFile() bool {
	if o != nil && !isNil(o.CertFile) {
		return true
	}

	return false
}

// SetCertFile gets a reference to the given string and assigns it to the CertFile field.
func (o *TLSSetting) SetCertFile(v string) {
	o.CertFile = &v
}

// GetKeyFile returns the KeyFile field value if set, zero value otherwise.
func (o *TLSSetting) GetKeyFile() string {
	if o == nil || isNil(o.KeyFile) {
		var ret string
		return ret
	}
	return *o.KeyFile
}

// GetKeyFileOk returns a tuple with the KeyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSetting) GetKeyFileOk() (*string, bool) {
	if o == nil || isNil(o.KeyFile) {
		return nil, false
	}
	return o.KeyFile, true
}

// HasKeyFile returns a boolean if a field has been set.
func (o *TLSSetting) HasKeyFile() bool {
	if o != nil && !isNil(o.KeyFile) {
		return true
	}

	return false
}

// SetKeyFile gets a reference to the given string and assigns it to the KeyFile field.
func (o *TLSSetting) SetKeyFile(v string) {
	o.KeyFile = &v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *TLSSetting) GetMinVersion() string {
	if o == nil || isNil(o.MinVersion) {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSetting) GetMinVersionOk() (*string, bool) {
	if o == nil || isNil(o.MinVersion) {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *TLSSetting) HasMinVersion() bool {
	if o != nil && !isNil(o.MinVersion) {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *TLSSetting) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetMaxVersion returns the MaxVersion field value if set, zero value otherwise.
func (o *TLSSetting) GetMaxVersion() string {
	if o == nil || isNil(o.MaxVersion) {
		var ret string
		return ret
	}
	return *o.MaxVersion
}

// GetMaxVersionOk returns a tuple with the MaxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSetting) GetMaxVersionOk() (*string, bool) {
	if o == nil || isNil(o.MaxVersion) {
		return nil, false
	}
	return o.MaxVersion, true
}

// HasMaxVersion returns a boolean if a field has been set.
func (o *TLSSetting) HasMaxVersion() bool {
	if o != nil && !isNil(o.MaxVersion) {
		return true
	}

	return false
}

// SetMaxVersion gets a reference to the given string and assigns it to the MaxVersion field.
func (o *TLSSetting) SetMaxVersion(v string) {
	o.MaxVersion = &v
}

func (o TLSSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TLSSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CAFile) {
		toSerialize["cAFile"] = o.CAFile
	}
	if !isNil(o.CertFile) {
		toSerialize["certFile"] = o.CertFile
	}
	if !isNil(o.KeyFile) {
		toSerialize["keyFile"] = o.KeyFile
	}
	if !isNil(o.MinVersion) {
		toSerialize["minVersion"] = o.MinVersion
	}
	if !isNil(o.MaxVersion) {
		toSerialize["maxVersion"] = o.MaxVersion
	}
	return toSerialize, nil
}

type NullableTLSSetting struct {
	value *TLSSetting
	isSet bool
}

func (v NullableTLSSetting) Get() *TLSSetting {
	return v.value
}

func (v *NullableTLSSetting) Set(val *TLSSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableTLSSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableTLSSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTLSSetting(val *TLSSetting) *NullableTLSSetting {
	return &NullableTLSSetting{value: val, isSet: true}
}

func (v NullableTLSSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTLSSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
