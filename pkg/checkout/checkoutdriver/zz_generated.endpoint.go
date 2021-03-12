// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package checkoutdriver

import (
	"context"
	"errors"
	"github.com/drhidians/checkout/pkg/checkout"
	"github.com/go-kit/kit/endpoint"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
)

// endpointError identifies an error that should be returned as an endpoint error.
type endpointError interface {
	EndpointError() bool
}

// serviceError identifies an error that should be returned as a service error.
type serviceError interface {
	ServiceError() bool
}

// Endpoints collects all of the endpoints that compose the underlying service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	AddProduct          endpoint.Endpoint
	CalculateTotalPrice endpoint.Endpoint
}

// MakeEndpoints returns a(n) Endpoints struct where each endpoint invokes
// the corresponding method on the provided service.
func MakeEndpoints(service checkout.Service, middleware ...endpoint.Middleware) Endpoints {
	mw := kitxendpoint.Combine(middleware...)

	return Endpoints{
		AddProduct:          kitxendpoint.OperationNameMiddleware("checkout.AddProduct")(mw(MakeAddProductEndpoint(service))),
		CalculateTotalPrice: kitxendpoint.OperationNameMiddleware("checkout.CalculateTotalPrice")(mw(MakeCalculateTotalPriceEndpoint(service))),
	}
}

// AddProductRequest is a request struct for AddProduct endpoint.
type AddProductRequest struct {
	Product checkout.Product
}

// AddProductResponse is a response struct for AddProduct endpoint.
type AddProductResponse struct {
	TotalPrice checkout.Price
	Err        error
}

func (r AddProductResponse) Failed() error {
	return r.Err
}

// MakeAddProductEndpoint returns an endpoint for the matching method of the underlying service.
func MakeAddProductEndpoint(service checkout.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddProductRequest)

		totalPrice, err := service.AddProduct(ctx, req.Product)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return AddProductResponse{
					Err:        err,
					TotalPrice: totalPrice,
				}, nil
			}

			return AddProductResponse{
				Err:        err,
				TotalPrice: totalPrice,
			}, err
		}

		return AddProductResponse{TotalPrice: totalPrice}, nil
	}
}

// CalculateTotalPriceRequest is a request struct for CalculateTotalPrice endpoint.
type CalculateTotalPriceRequest struct {
	P0 checkout.Product
}

// CalculateTotalPriceResponse is a response struct for CalculateTotalPrice endpoint.
type CalculateTotalPriceResponse struct {
	R0  checkout.Price
	Err error
}

func (r CalculateTotalPriceResponse) Failed() error {
	return r.Err
}

// MakeCalculateTotalPriceEndpoint returns an endpoint for the matching method of the underlying service.
func MakeCalculateTotalPriceEndpoint(service checkout.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculateTotalPriceRequest)

		r0, err := service.CalculateTotalPrice(ctx, req.P0)

		if err != nil {
			if serviceErr := serviceError(nil); errors.As(err, &serviceErr) && serviceErr.ServiceError() {
				return CalculateTotalPriceResponse{
					Err: err,
					R0:  r0,
				}, nil
			}

			return CalculateTotalPriceResponse{
				Err: err,
				R0:  r0,
			}, err
		}

		return CalculateTotalPriceResponse{R0: r0}, nil
	}
}
