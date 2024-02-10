/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GrpcResponse struct {
	StatusCode int32 `json:"statusCode"`

	Metadata []GrpcHeader `json:"metadata,omitempty"`

	Body string `json:"body,omitempty"`
}

// AssertGrpcResponseRequired checks if the required fields are not zero-ed
func AssertGrpcResponseRequired(obj GrpcResponse) error {
	elements := map[string]interface{}{
		"statusCode": obj.StatusCode,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Metadata {
		if err := AssertGrpcHeaderRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseGrpcResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of GrpcResponse (e.g. [][]GrpcResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseGrpcResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aGrpcResponse, ok := obj.(GrpcResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertGrpcResponseRequired(aGrpcResponse)
	})
}
