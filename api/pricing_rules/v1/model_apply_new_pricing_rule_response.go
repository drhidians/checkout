/*
 * Pricing Rules API
 *
 * The Pricing Rules API manages a pricing rules of items 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApplyNewPricingRuleResponse struct for ApplyNewPricingRuleResponse
type ApplyNewPricingRuleResponse struct {
	Count int32 `json:"count"`
}

// NewApplyNewPricingRuleResponse instantiates a new ApplyNewPricingRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyNewPricingRuleResponse(count int32, ) *ApplyNewPricingRuleResponse {
	this := ApplyNewPricingRuleResponse{}
	this.Count = count
	return &this
}

// NewApplyNewPricingRuleResponseWithDefaults instantiates a new ApplyNewPricingRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyNewPricingRuleResponseWithDefaults() *ApplyNewPricingRuleResponse {
	this := ApplyNewPricingRuleResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *ApplyNewPricingRuleResponse) GetCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ApplyNewPricingRuleResponse) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ApplyNewPricingRuleResponse) SetCount(v int32) {
	o.Count = v
}

func (o ApplyNewPricingRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableApplyNewPricingRuleResponse struct {
	value *ApplyNewPricingRuleResponse
	isSet bool
}

func (v NullableApplyNewPricingRuleResponse) Get() *ApplyNewPricingRuleResponse {
	return v.value
}

func (v *NullableApplyNewPricingRuleResponse) Set(val *ApplyNewPricingRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyNewPricingRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyNewPricingRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyNewPricingRuleResponse(val *ApplyNewPricingRuleResponse) *NullableApplyNewPricingRuleResponse {
	return &NullableApplyNewPricingRuleResponse{value: val, isSet: true}
}

func (v NullableApplyNewPricingRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyNewPricingRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


