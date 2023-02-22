/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SupportedClients the model 'SupportedClients'
type SupportedClients string

// List of SupportedClients
const (
	HTTP SupportedClients = "http"
	GRPC SupportedClients = "grpc"
)

// All allowed values of SupportedClients enum
var AllowedSupportedClientsEnumValues = []SupportedClients{
	"http",
	"grpc",
}

func (v *SupportedClients) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportedClients(value)
	for _, existing := range AllowedSupportedClientsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportedClients", value)
}

// NewSupportedClientsFromValue returns a pointer to a valid SupportedClients
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportedClientsFromValue(v string) (*SupportedClients, error) {
	ev := SupportedClients(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportedClients: valid values are %v", v, AllowedSupportedClientsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportedClients) IsValid() bool {
	for _, existing := range AllowedSupportedClientsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportedClients value
func (v SupportedClients) Ptr() *SupportedClients {
	return &v
}

type NullableSupportedClients struct {
	value *SupportedClients
	isSet bool
}

func (v NullableSupportedClients) Get() *SupportedClients {
	return v.value
}

func (v *NullableSupportedClients) Set(val *SupportedClients) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedClients) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedClients(val *SupportedClients) *NullableSupportedClients {
	return &NullableSupportedClients{value: val, isSet: true}
}

func (v NullableSupportedClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}