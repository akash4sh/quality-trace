/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SelectorFilter struct {
	Property string `json:"property"`

	Operator string `json:"operator"`

	Value string `json:"value"`
}

// AssertSelectorFilterRequired checks if the required fields are not zero-ed
func AssertSelectorFilterRequired(obj SelectorFilter) error {
	elements := map[string]interface{}{
		"property": obj.Property,
		"operator": obj.Operator,
		"value":    obj.Value,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseSelectorFilterRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SelectorFilter (e.g. [][]SelectorFilter), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSelectorFilterRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSelectorFilter, ok := obj.(SelectorFilter)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSelectorFilterRequired(aSelectorFilter)
	})
}
