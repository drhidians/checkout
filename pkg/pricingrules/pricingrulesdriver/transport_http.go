package pricingrulesdriver

import (
	"context"
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	appkithttp "github.com/sagikazarmark/appkit/transport/http"
	kitxhttp "github.com/sagikazarmark/kitx/transport/http"

	api "github.com/drhidians/checkout/api/pricing_rules/v1"
	"github.com/drhidians/checkout/pkg/pricingrules"
)

type contextKey int

const (
	ContextKeyBaseURL contextKey = iota
)

// RegisterHTTPHandlers mounts all of the service endpoints into a router.
func RegisterHTTPHandlers(endpoints Endpoints, router *mux.Router, options ...kithttp.ServerOption) {
	errorEncoder := kitxhttp.NewJSONProblemErrorResponseEncoder(appkithttp.NewDefaultProblemConverter())

	router.Methods(http.MethodPost).Path("/reset-sku-rules").Handler(kithttp.NewServer(
		endpoints.ApplyNewPricingRules,
		decodeResetSkuPricingRulesHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeResetSkuPricingRulesHTTPResponse, errorEncoder),
		options...,
	))

	router.Methods(http.MethodPost).Path("/reset-common-rules").Handler(kithttp.NewServer(
		endpoints.ApplyNewPricingRules,
		decodeResetCommonPricingRulesHTTPRequest,
		kitxhttp.ErrorResponseEncoder(encodeResetCommonPricingRulesHTTPResponse, errorEncoder),
		options...,
	))
}

func decodeResetSkuPricingRulesHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest []api.PricingRule
	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "decode request")
	}

	var rules = make([]pricingrules.PricingRule, 0, len(apiRequest))

	for _, r := range apiRequest {
		rules = append(rules, pricingrules.PricingRule{
			BaseRule: pricingrules.BaseRule{
				Rule:   pricingrules.Rule(r.Rule),
				Weight: r.Weight,
			},
			Sku: r.Sku,
		})
	}

	return ApplyNewPricingRulesRequest{
		Rules: rules,
	}, nil
}

func encodeResetSkuPricingRulesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ApplyNewPricingRulesRequest)

	apiResponse := marshalPricingRulesHTTP(ctx, resp.Rules)

	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(apiResponse, http.StatusCreated))
}


func decodeResetCommonPricingRulesHTTPRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var apiRequest []api.CommonPricingRule
	err := json.NewDecoder(r.Body).Decode(&apiRequest)
	if err != nil {
		return nil, errors.Wrap(err, "decode request")
	}

	var rules = make([]pricingrules.CommonPricingRule, 0, len(apiRequest))

	for _, r := range apiRequest {
		rules = append(rules, pricingrules.CommonPricingRule{
			BaseRule: pricingrules.BaseRule{
				Rule:   pricingrules.Rule(r.Rule),
				Weight: r.Weight,
			},
		})
	}

	return ApplyNewCommonPricingRulesRequest{
		Rules: rules,
	}, nil
}

func encodeResetCommonPricingRulesHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	resp := response.(ApplyNewPricingRulesRequest)

	apiResponse := marshalPricingRulesHTTP(ctx, resp.Rules)

	return kitxhttp.JSONResponseEncoder(ctx, w, kitxhttp.WithStatusCode(apiResponse, http.StatusCreated))
}

func marshalPricingRulesHTTP(_ context.Context, rules []pricingrules.PricingRule) api.ApplyNewPricingRuleResponse {
	return api.ApplyNewPricingRuleResponse{
		Count: int32(len(rules)),
	}
}
