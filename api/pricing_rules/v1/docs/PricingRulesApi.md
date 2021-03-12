# \PricingRulesApi

All URIs are relative to *https://checkout.api/pricing-rules*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetCommonPricingRules**](PricingRulesApi.md#ResetCommonPricingRules) | **Post** /common-pricing-rules | reset common pricing rules
[**ResetPricingRules**](PricingRulesApi.md#ResetPricingRules) | **Post** /pricing-rules | reset pricing rules



## ResetCommonPricingRules

> ApplyNewPricingRuleResponse ResetCommonPricingRules(ctx).PricingRule(pricingRule).Execute()

reset common pricing rules

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pricingRule := []openapiclient.PricingRule{*openapiclient.NewPricingRule(int32(123), "Sku_example", int32(123))} // []PricingRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PricingRulesApi.ResetCommonPricingRules(context.Background()).PricingRule(pricingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesApi.ResetCommonPricingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetCommonPricingRules`: ApplyNewPricingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `PricingRulesApi.ResetCommonPricingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetCommonPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pricingRule** | [**[]PricingRule**](PricingRule.md) |  | 

### Return type

[**ApplyNewPricingRuleResponse**](ApplyNewPricingRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPricingRules

> ApplyNewPricingRuleResponse ResetPricingRules(ctx).PricingRule(pricingRule).Execute()

reset pricing rules

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pricingRule := []openapiclient.PricingRule{*openapiclient.NewPricingRule(int32(123), "Sku_example", int32(123))} // []PricingRule | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PricingRulesApi.ResetPricingRules(context.Background()).PricingRule(pricingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesApi.ResetPricingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetPricingRules`: ApplyNewPricingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `PricingRulesApi.ResetPricingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pricingRule** | [**[]PricingRule**](PricingRule.md) |  | 

### Return type

[**ApplyNewPricingRuleResponse**](ApplyNewPricingRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

