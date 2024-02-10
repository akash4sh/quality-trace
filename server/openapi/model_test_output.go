/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TestOutput struct {
	Name string `json:"name,omitempty"`

	Selector string `json:"selector,omitempty"`

	SelectorParsed Selector `json:"selectorParsed,omitempty"`

	Value string `json:"value,omitempty"`
}

// AssertTestOutputRequired checks if the required fields are not zero-ed
func AssertTestOutputRequired(obj TestOutput) error {
	if err := AssertSelectorRequired(obj.SelectorParsed); err != nil {
		return err
	}
	return nil
}

// AssertRecurseTestOutputRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TestOutput (e.g. [][]TestOutput), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTestOutputRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTestOutput, ok := obj.(TestOutput)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTestOutputRequired(aTestOutput)
	})
}
