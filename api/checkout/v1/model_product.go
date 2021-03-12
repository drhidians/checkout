/*
 * Checkout API
 *
 * The Checkout API manages checkout items 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Product struct for Product
type Product struct {
	Sku string `json:"sku"`
	Quantity int32 `json:"quantity"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(sku string, quantity int32, ) *Product {
	this := Product{}
	this.Sku = sku
	this.Quantity = quantity
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetSku returns the Sku field value
func (o *Product) GetSku() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *Product) GetSkuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *Product) SetSku(v string) {
	o.Sku = v
}

// GetQuantity returns the Quantity field value
func (o *Product) GetQuantity() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Product) GetQuantityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Product) SetQuantity(v int32) {
	o.Quantity = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sku"] = o.Sku
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

