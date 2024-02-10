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

// checks if the DemoList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoList{}

// DemoList struct for DemoList
type DemoList struct {
	Count *int32 `json:"count,omitempty"`
	Items []Demo `json:"items,omitempty"`
}

// NewDemoList instantiates a new DemoList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoList() *DemoList {
	this := DemoList{}
	return &this
}

// NewDemoListWithDefaults instantiates a new DemoList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoListWithDefaults() *DemoList {
	this := DemoList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DemoList) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoList) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DemoList) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DemoList) SetCount(v int32) {
	o.Count = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DemoList) GetItems() []Demo {
	if o == nil || isNil(o.Items) {
		var ret []Demo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoList) GetItemsOk() ([]Demo, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DemoList) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Demo and assigns it to the Items field.
func (o *DemoList) SetItems(v []Demo) {
	o.Items = v
}

func (o DemoList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableDemoList struct {
	value *DemoList
	isSet bool
}

func (v NullableDemoList) Get() *DemoList {
	return v.value
}

func (v *NullableDemoList) Set(val *DemoList) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoList) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoList(val *DemoList) *NullableDemoList {
	return &NullableDemoList{value: val, isSet: true}
}

func (v NullableDemoList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
