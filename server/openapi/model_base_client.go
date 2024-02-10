/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type BaseClient struct {
	Type SupportedClients `json:"type,omitempty"`

	Http HttpClientSettings `json:"http,omitempty"`

	Grpc GrpcClientSettings `json:"grpc,omitempty"`
}

// AssertBaseClientRequired checks if the required fields are not zero-ed
func AssertBaseClientRequired(obj BaseClient) error {
	if err := AssertHttpClientSettingsRequired(obj.Http); err != nil {
		return err
	}
	if err := AssertGrpcClientSettingsRequired(obj.Grpc); err != nil {
		return err
	}
	return nil
}

// AssertRecurseBaseClientRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of BaseClient (e.g. [][]BaseClient), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBaseClientRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBaseClient, ok := obj.(BaseClient)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBaseClientRequired(aBaseClient)
	})
}
