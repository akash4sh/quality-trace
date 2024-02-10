/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MissingVariable struct {
	TestId string `json:"testId,omitempty"`

	Variables []Variable `json:"variables,omitempty"`
}

// AssertMissingVariableRequired checks if the required fields are not zero-ed
func AssertMissingVariableRequired(obj MissingVariable) error {
	for _, el := range obj.Variables {
		if err := AssertVariableRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseMissingVariableRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of MissingVariable (e.g. [][]MissingVariable), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseMissingVariableRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aMissingVariable, ok := obj.(MissingVariable)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertMissingVariableRequired(aMissingVariable)
	})
}
