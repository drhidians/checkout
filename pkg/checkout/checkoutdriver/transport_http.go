package checkoutdriver

import (
	"context"
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	api "github.com/drhidians/checkout/api/checkout/v1"
	"github.com/drhidians/checkout/pkg/checkout"
)

// RegisterHTTPHandlers mounts all of the service endpoints into a router.
func RegisterHTTPHandlers(endpoints Endpoints, router *mux.Router, options ...kithttp.ServerOption) {
	errorEncoder := kitxhttp.NewJSONProblemErrorResponseEncoder(appkithttp.NewDefaultProblemConverter())

	router.Methods(http.MethodPost).Path("").Handler(kithttp.NewServer(
		endpoints.AddProduct,
		decodeResetSkuPricingRulesHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeResetSkuPricingRulesHTTPResponse, errorEncoder),
		options...,
	))
}

func decodeResetSkuPricingRulesHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest api.Product
	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "decode request")
	}

	return AddProductRequest{
		Product:checkout.Product{
			Sku: apiRequest.Sku,
			Quantity: apiRequest.Quantity,
			Session: checkout.ClientInfo{
				Session: "1",
			}, // TODO: implement some kind of session in frontend and middleware for authorized user
		},
	}, nil
}

func encodeResetSkuPricingRulesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(nil, http.StatusNoContent))
}



