/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PollingProfileList struct {
	Count int32 `json:"count,omitempty"`

	Items []PollingProfile `json:"items,omitempty"`
}

// AssertPollingProfileListRequired checks if the required fields are not zero-ed
func AssertPollingProfileListRequired(obj PollingProfileList) error {
	for _, el := range obj.Items {
		if err := AssertPollingProfileRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecursePollingProfileListRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PollingProfileList (e.g. [][]PollingProfileList), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePollingProfileListRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPollingProfileList, ok := obj.(PollingProfileList)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPollingProfileListRequired(aPollingProfileList)
	})
}
