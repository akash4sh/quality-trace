/*
 * Qualitytrace
 *
 * OpenAPI definition for Qualitytrace endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// ResourceApiApiService is a service that implements the logic for the ResourceApiApiServicer
// This service should implement the business logic for every endpoint for the ResourceApiApi API.
// Include any external packages or services that will be required by this service.
type ResourceApiApiService struct {
}

// NewResourceApiApiService creates a default api service
func NewResourceApiApiService() ResourceApiApiServicer {
	return &ResourceApiApiService{}
}

// CreateDemo - Create a Demonstration setting
func (s *ResourceApiApiService) CreateDemo(ctx context.Context, demo Demo) (ImplResponse, error) {
	// TODO - update CreateDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Demo{}) or use other options such as http.Ok ...
	//return Response(201, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateDemo method not implemented")
}

// CreateLinter - Create an Linter
func (s *ResourceApiApiService) CreateLinter(ctx context.Context, linterResource LinterResource) (ImplResponse, error) {
	// TODO - update CreateLinter with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, LinterResource{}) or use other options such as http.Ok ...
	//return Response(201, LinterResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateLinter method not implemented")
}

// CreateTest - Create new test
func (s *ResourceApiApiService) CreateTest(ctx context.Context, test Test) (ImplResponse, error) {
	// TODO - update CreateTest with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Test{}) or use other options such as http.Ok ...
	//return Response(200, Test{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateTest method not implemented")
}

// CreateTestSuite - Create new TestSuite
func (s *ResourceApiApiService) CreateTestSuite(ctx context.Context, testSuiteResource TestSuiteResource) (ImplResponse, error) {
	// TODO - update CreateTestSuite with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, TestSuiteResource{}) or use other options such as http.Ok ...
	//return Response(201, TestSuiteResource{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateTestSuite method not implemented")
}

// CreateVariableSet - Create a VariableSet
func (s *ResourceApiApiService) CreateVariableSet(ctx context.Context, variableSetResource VariableSetResource) (ImplResponse, error) {
	// TODO - update CreateVariableSet with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, VariableSetResource{}) or use other options such as http.Ok ...
	//return Response(201, VariableSetResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateVariableSet method not implemented")
}

// DeleteDataStore - Delete a Data Store
func (s *ResourceApiApiService) DeleteDataStore(ctx context.Context, dataStoreId string) (ImplResponse, error) {
	// TODO - update DeleteDataStore with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteDataStore method not implemented")
}

// DeleteDemo - Delete a Demonstration setting
func (s *ResourceApiApiService) DeleteDemo(ctx context.Context, demoId string) (ImplResponse, error) {
	// TODO - update DeleteDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteDemo method not implemented")
}

// DeleteLinter - Delete an Linter
func (s *ResourceApiApiService) DeleteLinter(ctx context.Context, linterId string) (ImplResponse, error) {
	// TODO - update DeleteLinter with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteLinter method not implemented")
}

// DeleteTest - delete a test
func (s *ResourceApiApiService) DeleteTest(ctx context.Context, testId string) (ImplResponse, error) {
	// TODO - update DeleteTest with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteTest method not implemented")
}

// DeleteTestSuite - delete a TestSuite
func (s *ResourceApiApiService) DeleteTestSuite(ctx context.Context, testSuiteId string) (ImplResponse, error) {
	// TODO - update DeleteTestSuite with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteTestSuite method not implemented")
}

// DeleteVariableSet - Delete a variable set
func (s *ResourceApiApiService) DeleteVariableSet(ctx context.Context, variableSetId string) (ImplResponse, error) {
	// TODO - update DeleteVariableSet with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteVariableSet method not implemented")
}

// GetConfiguration - Get Qualitytrace configuration
func (s *ResourceApiApiService) GetConfiguration(ctx context.Context, configId string) (ImplResponse, error) {
	// TODO - update GetConfiguration with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigurationResource{}) or use other options such as http.Ok ...
	//return Response(200, ConfigurationResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetConfiguration method not implemented")
}

// GetDataStore - Get a Data Store
func (s *ResourceApiApiService) GetDataStore(ctx context.Context, dataStoreId string) (ImplResponse, error) {
	// TODO - update GetDataStore with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DataStoreResource{}) or use other options such as http.Ok ...
	//return Response(200, DataStoreResource{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDataStore method not implemented")
}

// GetDemo - Get Demonstration setting
func (s *ResourceApiApiService) GetDemo(ctx context.Context, demoId string) (ImplResponse, error) {
	// TODO - update GetDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Demo{}) or use other options such as http.Ok ...
	//return Response(200, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDemo method not implemented")
}

// GetLinter - Get a specific Linter
func (s *ResourceApiApiService) GetLinter(ctx context.Context, linterId string) (ImplResponse, error) {
	// TODO - update GetLinter with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LinterResource{}) or use other options such as http.Ok ...
	//return Response(200, LinterResource{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetLinter method not implemented")
}

// GetPollingProfile - Get Polling Profile
func (s *ResourceApiApiService) GetPollingProfile(ctx context.Context, pollingProfileId string) (ImplResponse, error) {
	// TODO - update GetPollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollingProfile{}) or use other options such as http.Ok ...
	//return Response(200, PollingProfile{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPollingProfile method not implemented")
}

// GetTestSuite - get TestSuite
func (s *ResourceApiApiService) GetTestSuite(ctx context.Context, testSuiteId string) (ImplResponse, error) {
	// TODO - update GetTestSuite with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TestSuiteResource{}) or use other options such as http.Ok ...
	//return Response(200, TestSuiteResource{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetTestSuite method not implemented")
}

// GetTestSuites - Get testsuites
func (s *ResourceApiApiService) GetTestSuites(ctx context.Context, take int32, skip int32, query string, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update GetTestSuites with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TestSuiteResourceList{}) or use other options such as http.Ok ...
	//return Response(200, TestSuiteResourceList{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetTestSuites method not implemented")
}

// GetTests - Get tests
func (s *ResourceApiApiService) GetTests(ctx context.Context, take int32, skip int32, query string, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update GetTests with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TestResourceList{}) or use other options such as http.Ok ...
	//return Response(200, TestResourceList{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetTests method not implemented")
}

// GetVariableSet - Get a specific VariableSet
func (s *ResourceApiApiService) GetVariableSet(ctx context.Context, variableSetId string) (ImplResponse, error) {
	// TODO - update GetVariableSet with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VariableSetResource{}) or use other options such as http.Ok ...
	//return Response(200, VariableSetResource{}), nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetVariableSet method not implemented")
}

// ListConfiguration - List Qualitytrace configuration
func (s *ResourceApiApiService) ListConfiguration(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListConfiguration with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigurationResourceList{}) or use other options such as http.Ok ...
	//return Response(200, ConfigurationResourceList{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListConfiguration method not implemented")
}

// ListDataStore - List Data Store
func (s *ResourceApiApiService) ListDataStore(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListDataStore with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DataStoreList{}) or use other options such as http.Ok ...
	//return Response(200, DataStoreList{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListDataStore method not implemented")
}

// ListDemos - List Demonstrations
func (s *ResourceApiApiService) ListDemos(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListDemos with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DemoList{}) or use other options such as http.Ok ...
	//return Response(200, DemoList{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListDemos method not implemented")
}

// ListLinters - List Linters
func (s *ResourceApiApiService) ListLinters(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListLinters with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LinterResourceList{}) or use other options such as http.Ok ...
	//return Response(200, LinterResourceList{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListLinters method not implemented")
}

// ListPollingProfile - List Polling Profile Configuration
func (s *ResourceApiApiService) ListPollingProfile(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListPollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollingProfileList{}) or use other options such as http.Ok ...
	//return Response(200, PollingProfileList{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListPollingProfile method not implemented")
}

// ListVariableSets - List VariableSets
func (s *ResourceApiApiService) ListVariableSets(ctx context.Context, take int32, skip int32, sortBy string, sortDirection string) (ImplResponse, error) {
	// TODO - update ListVariableSets with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VariableSetResourceList{}) or use other options such as http.Ok ...
	//return Response(200, VariableSetResourceList{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListVariableSets method not implemented")
}

// TestsTestIdGet - get test
func (s *ResourceApiApiService) TestsTestIdGet(ctx context.Context, testId string) (ImplResponse, error) {
	// TODO - update TestsTestIdGet with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TestResource{}) or use other options such as http.Ok ...
	//return Response(200, TestResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("TestsTestIdGet method not implemented")
}

// UpdateConfiguration - Update Qualitytrace configuration
func (s *ResourceApiApiService) UpdateConfiguration(ctx context.Context, configId string, configurationResource ConfigurationResource) (ImplResponse, error) {
	// TODO - update UpdateConfiguration with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ConfigurationResource{}) or use other options such as http.Ok ...
	//return Response(200, ConfigurationResource{}), nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateConfiguration method not implemented")
}

// UpdateDataStore - Update a Data Store
func (s *ResourceApiApiService) UpdateDataStore(ctx context.Context, dataStoreId string, dataStore DataStore) (ImplResponse, error) {
	// TODO - update UpdateDataStore with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateDataStore method not implemented")
}

// UpdateDemo - Update a Demonstration setting
func (s *ResourceApiApiService) UpdateDemo(ctx context.Context, demoId string, demo Demo) (ImplResponse, error) {
	// TODO - update UpdateDemo with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Demo{}) or use other options such as http.Ok ...
	//return Response(200, Demo{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateDemo method not implemented")
}

// UpdateLinter - Update a Linter
func (s *ResourceApiApiService) UpdateLinter(ctx context.Context, linterId string, linterResource LinterResource) (ImplResponse, error) {
	// TODO - update UpdateLinter with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, LinterResource{}) or use other options such as http.Ok ...
	//return Response(200, LinterResource{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateLinter method not implemented")
}

// UpdatePollingProfile - Update a Polling Profile
func (s *ResourceApiApiService) UpdatePollingProfile(ctx context.Context, pollingProfileId string, pollingProfile PollingProfile) (ImplResponse, error) {
	// TODO - update UpdatePollingProfile with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollingProfile{}) or use other options such as http.Ok ...
	//return Response(200, PollingProfile{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePollingProfile method not implemented")
}

// UpdateTest - update test
func (s *ResourceApiApiService) UpdateTest(ctx context.Context, testId string, test Test) (ImplResponse, error) {
	// TODO - update UpdateTest with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateTest method not implemented")
}

// UpdateTestSuite - update TestSuite
func (s *ResourceApiApiService) UpdateTestSuite(ctx context.Context, testSuiteId string, testSuiteResource TestSuiteResource) (ImplResponse, error) {
	// TODO - update UpdateTestSuite with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TestSuiteResource{}) or use other options such as http.Ok ...
	//return Response(200, TestSuiteResource{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateTestSuite method not implemented")
}

// UpdateVariableSet - Update a VariableSet
func (s *ResourceApiApiService) UpdateVariableSet(ctx context.Context, variableSetId string, variableSetResource VariableSetResource) (ImplResponse, error) {
	// TODO - update UpdateVariableSet with the required logic for this service method.
	// Add api_resource_api_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, VariableSetResource{}) or use other options such as http.Ok ...
	//return Response(200, VariableSetResource{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateVariableSet method not implemented")
}
