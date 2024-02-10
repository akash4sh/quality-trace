/*
Qualitytrace

OpenAPI definition for Qualitytrace endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DemoOpenTelemetryStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DemoOpenTelemetryStore{}

// DemoOpenTelemetryStore Represents the settings of the Open Telemetry Store demonstration.
type DemoOpenTelemetryStore struct {
	// Address of the root URL for the Frontend microservice on Open Telemetry Store.
	FrontendEndpoint *string `json:"frontendEndpoint,omitempty"`
	// Address of the root URL for the Product Catalog microservice on Open Telemetry Store.
	ProductCatalogEndpoint *string `json:"productCatalogEndpoint,omitempty"`
	// Address of the root URL for the Cart microservice on Open Telemetry Store.
	CartEndpoint *string `json:"cartEndpoint,omitempty"`
	// Address of the root URL for the Checkout microservice on Open Telemetry Store.
	CheckoutEndpoint *string `json:"checkoutEndpoint,omitempty"`
}

// NewDemoOpenTelemetryStore instantiates a new DemoOpenTelemetryStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDemoOpenTelemetryStore() *DemoOpenTelemetryStore {
	this := DemoOpenTelemetryStore{}
	return &this
}

// NewDemoOpenTelemetryStoreWithDefaults instantiates a new DemoOpenTelemetryStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDemoOpenTelemetryStoreWithDefaults() *DemoOpenTelemetryStore {
	this := DemoOpenTelemetryStore{}
	return &this
}

// GetFrontendEndpoint returns the FrontendEndpoint field value if set, zero value otherwise.
func (o *DemoOpenTelemetryStore) GetFrontendEndpoint() string {
	if o == nil || isNil(o.FrontendEndpoint) {
		var ret string
		return ret
	}
	return *o.FrontendEndpoint
}

// GetFrontendEndpointOk returns a tuple with the FrontendEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoOpenTelemetryStore) GetFrontendEndpointOk() (*string, bool) {
	if o == nil || isNil(o.FrontendEndpoint) {
		return nil, false
	}
	return o.FrontendEndpoint, true
}

// HasFrontendEndpoint returns a boolean if a field has been set.
func (o *DemoOpenTelemetryStore) HasFrontendEndpoint() bool {
	if o != nil && !isNil(o.FrontendEndpoint) {
		return true
	}

	return false
}

// SetFrontendEndpoint gets a reference to the given string and assigns it to the FrontendEndpoint field.
func (o *DemoOpenTelemetryStore) SetFrontendEndpoint(v string) {
	o.FrontendEndpoint = &v
}

// GetProductCatalogEndpoint returns the ProductCatalogEndpoint field value if set, zero value otherwise.
func (o *DemoOpenTelemetryStore) GetProductCatalogEndpoint() string {
	if o == nil || isNil(o.ProductCatalogEndpoint) {
		var ret string
		return ret
	}
	return *o.ProductCatalogEndpoint
}

// GetProductCatalogEndpointOk returns a tuple with the ProductCatalogEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoOpenTelemetryStore) GetProductCatalogEndpointOk() (*string, bool) {
	if o == nil || isNil(o.ProductCatalogEndpoint) {
		return nil, false
	}
	return o.ProductCatalogEndpoint, true
}

// HasProductCatalogEndpoint returns a boolean if a field has been set.
func (o *DemoOpenTelemetryStore) HasProductCatalogEndpoint() bool {
	if o != nil && !isNil(o.ProductCatalogEndpoint) {
		return true
	}

	return false
}

// SetProductCatalogEndpoint gets a reference to the given string and assigns it to the ProductCatalogEndpoint field.
func (o *DemoOpenTelemetryStore) SetProductCatalogEndpoint(v string) {
	o.ProductCatalogEndpoint = &v
}

// GetCartEndpoint returns the CartEndpoint field value if set, zero value otherwise.
func (o *DemoOpenTelemetryStore) GetCartEndpoint() string {
	if o == nil || isNil(o.CartEndpoint) {
		var ret string
		return ret
	}
	return *o.CartEndpoint
}

// GetCartEndpointOk returns a tuple with the CartEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoOpenTelemetryStore) GetCartEndpointOk() (*string, bool) {
	if o == nil || isNil(o.CartEndpoint) {
		return nil, false
	}
	return o.CartEndpoint, true
}

// HasCartEndpoint returns a boolean if a field has been set.
func (o *DemoOpenTelemetryStore) HasCartEndpoint() bool {
	if o != nil && !isNil(o.CartEndpoint) {
		return true
	}

	return false
}

// SetCartEndpoint gets a reference to the given string and assigns it to the CartEndpoint field.
func (o *DemoOpenTelemetryStore) SetCartEndpoint(v string) {
	o.CartEndpoint = &v
}

// GetCheckoutEndpoint returns the CheckoutEndpoint field value if set, zero value otherwise.
func (o *DemoOpenTelemetryStore) GetCheckoutEndpoint() string {
	if o == nil || isNil(o.CheckoutEndpoint) {
		var ret string
		return ret
	}
	return *o.CheckoutEndpoint
}

// GetCheckoutEndpointOk returns a tuple with the CheckoutEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemoOpenTelemetryStore) GetCheckoutEndpointOk() (*string, bool) {
	if o == nil || isNil(o.CheckoutEndpoint) {
		return nil, false
	}
	return o.CheckoutEndpoint, true
}

// HasCheckoutEndpoint returns a boolean if a field has been set.
func (o *DemoOpenTelemetryStore) HasCheckoutEndpoint() bool {
	if o != nil && !isNil(o.CheckoutEndpoint) {
		return true
	}

	return false
}

// SetCheckoutEndpoint gets a reference to the given string and assigns it to the CheckoutEndpoint field.
func (o *DemoOpenTelemetryStore) SetCheckoutEndpoint(v string) {
	o.CheckoutEndpoint = &v
}

func (o DemoOpenTelemetryStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DemoOpenTelemetryStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FrontendEndpoint) {
		toSerialize["frontendEndpoint"] = o.FrontendEndpoint
	}
	if !isNil(o.ProductCatalogEndpoint) {
		toSerialize["productCatalogEndpoint"] = o.ProductCatalogEndpoint
	}
	if !isNil(o.CartEndpoint) {
		toSerialize["cartEndpoint"] = o.CartEndpoint
	}
	if !isNil(o.CheckoutEndpoint) {
		toSerialize["checkoutEndpoint"] = o.CheckoutEndpoint
	}
	return toSerialize, nil
}

type NullableDemoOpenTelemetryStore struct {
	value *DemoOpenTelemetryStore
	isSet bool
}

func (v NullableDemoOpenTelemetryStore) Get() *DemoOpenTelemetryStore {
	return v.value
}

func (v *NullableDemoOpenTelemetryStore) Set(val *DemoOpenTelemetryStore) {
	v.value = val
	v.isSet = true
}

func (v NullableDemoOpenTelemetryStore) IsSet() bool {
	return v.isSet
}

func (v *NullableDemoOpenTelemetryStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDemoOpenTelemetryStore(val *DemoOpenTelemetryStore) *NullableDemoOpenTelemetryStore {
	return &NullableDemoOpenTelemetryStore{value: val, isSet: true}
}

func (v NullableDemoOpenTelemetryStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDemoOpenTelemetryStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
