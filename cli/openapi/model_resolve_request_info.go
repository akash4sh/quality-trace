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

// checks if the ResolveRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResolveRequestInfo{}

// ResolveRequestInfo struct for ResolveRequestInfo
type ResolveRequestInfo struct {
	Expression *string         `json:"expression,omitempty"`
	Context    *ResolveContext `json:"context,omitempty"`
}

// NewResolveRequestInfo instantiates a new ResolveRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResolveRequestInfo() *ResolveRequestInfo {
	this := ResolveRequestInfo{}
	return &this
}

// NewResolveRequestInfoWithDefaults instantiates a new ResolveRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResolveRequestInfoWithDefaults() *ResolveRequestInfo {
	this := ResolveRequestInfo{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ResolveRequestInfo) GetExpression() string {
	if o == nil || isNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveRequestInfo) GetExpressionOk() (*string, bool) {
	if o == nil || isNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ResolveRequestInfo) HasExpression() bool {
	if o != nil && !isNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ResolveRequestInfo) SetExpression(v string) {
	o.Expression = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ResolveRequestInfo) GetContext() ResolveContext {
	if o == nil || isNil(o.Context) {
		var ret ResolveContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveRequestInfo) GetContextOk() (*ResolveContext, bool) {
	if o == nil || isNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ResolveRequestInfo) HasContext() bool {
	if o != nil && !isNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given ResolveContext and assigns it to the Context field.
func (o *ResolveRequestInfo) SetContext(v ResolveContext) {
	o.Context = &v
}

func (o ResolveRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResolveRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !isNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

type NullableResolveRequestInfo struct {
	value *ResolveRequestInfo
	isSet bool
}

func (v NullableResolveRequestInfo) Get() *ResolveRequestInfo {
	return v.value
}

func (v *NullableResolveRequestInfo) Set(val *ResolveRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResolveRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResolveRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResolveRequestInfo(val *ResolveRequestInfo) *NullableResolveRequestInfo {
	return &NullableResolveRequestInfo{value: val, isSet: true}
}

func (v NullableResolveRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResolveRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
