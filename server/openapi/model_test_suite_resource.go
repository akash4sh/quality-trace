/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TestSuiteResource - Represents a TestSuite structured into the Resources format.
type TestSuiteResource struct {

	// Represents the type of this resource. It should always be set as 'TestSuite'.
	Type string `json:"type,omitempty"`

	Spec TestSuite `json:"spec,omitempty"`
}

// AssertTestSuiteResourceRequired checks if the required fields are not zero-ed
func AssertTestSuiteResourceRequired(obj TestSuiteResource) error {
	if err := AssertTestSuiteRequired(obj.Spec); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTestSuiteResourceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestSuiteResource (e.g. [][]TestSuiteResource), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestSuiteResourceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestSuiteResource, ok := obj.(TestSuiteResource)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestSuiteResourceRequired(aTestSuiteResource)
	})
}
