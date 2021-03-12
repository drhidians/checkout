# \CheckoutApi

All URIs are relative to *https://checkout.api/checkout*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Checkout**](CheckoutApi.md#Checkout) | **Post** /checkout | checkout



## Checkout

> Checkout(ctx).Product(product).Execute()

checkout

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
    product := []openapiclient.Product{*openapiclient.NewProduct("Sku_example", int32(123))} // []Product | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CheckoutApi.Checkout(context.Background()).Product(product).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CheckoutApi.Checkout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | [**[]Product**](Product.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

