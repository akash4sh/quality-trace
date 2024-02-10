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

// checks if the AssertionResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssertionResults{}

// AssertionResults struct for AssertionResults
type AssertionResults struct {
	AllPassed *bool                          `json:"allPassed,omitempty"`
	Results   []AssertionResultsResultsInner `json:"results,omitempty"`
}

// NewAssertionResults instantiates a new AssertionResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertionResults() *AssertionResults {
	this := AssertionResults{}
	return &this
}

// NewAssertionResultsWithDefaults instantiates a new AssertionResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertionResultsWithDefaults() *AssertionResults {
	this := AssertionResults{}
	return &this
}

// GetAllPassed returns the AllPassed field value if set, zero value otherwise.
func (o *AssertionResults) GetAllPassed() bool {
	if o == nil || isNil(o.AllPassed) {
		var ret bool
		return ret
	}
	return *o.AllPassed
}

// GetAllPassedOk returns a tuple with the AllPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertionResults) GetAllPassedOk() (*bool, bool) {
	if o == nil || isNil(o.AllPassed) {
		return nil, false
	}
	return o.AllPassed, true
}

// HasAllPassed returns a boolean if a field has been set.
func (o *AssertionResults) HasAllPassed() bool {
	if o != nil && !isNil(o.AllPassed) {
		return true
	}

	return false
}

// SetAllPassed gets a reference to the given bool and assigns it to the AllPassed field.
func (o *AssertionResults) SetAllPassed(v bool) {
	o.AllPassed = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AssertionResults) GetResults() []AssertionResultsResultsInner {
	if o == nil || isNil(o.Results) {
		var ret []AssertionResultsResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertionResults) GetResultsOk() ([]AssertionResultsResultsInner, bool) {
	if o == nil || isNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AssertionResults) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []AssertionResultsResultsInner and assigns it to the Results field.
func (o *AssertionResults) SetResults(v []AssertionResultsResultsInner) {
	o.Results = v
}

func (o AssertionResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssertionResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AllPassed) {
		toSerialize["allPassed"] = o.AllPassed
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAssertionResults struct {
	value *AssertionResults
	isSet bool
}

func (v NullableAssertionResults) Get() *AssertionResults {
	return v.value
}

func (v *NullableAssertionResults) Set(val *AssertionResults) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertionResults) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertionResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertionResults(val *AssertionResults) *NullableAssertionResults {
	return &NullableAssertionResults{value: val, isSet: true}
}

func (v NullableAssertionResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertionResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
