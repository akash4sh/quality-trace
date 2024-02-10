/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type KafkaMessageHeader struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

// AssertKafkaMessageHeaderRequired checks if the required fields are not zero-ed
func AssertKafkaMessageHeaderRequired(obj KafkaMessageHeader) error {
	return nil
}

// AssertRecurseKafkaMessageHeaderRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of KafkaMessageHeader (e.g. [][]KafkaMessageHeader), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseKafkaMessageHeaderRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aKafkaMessageHeader, ok := obj.(KafkaMessageHeader)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertKafkaMessageHeaderRequired(aKafkaMessageHeader)
	})
}
