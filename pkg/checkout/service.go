package checkout

import (
	"context"

	"github.com/drhidians/checkout/pkg/pricingrules"
)

// +kit:endpoint:errorStrategy=service

type Service interface {
	AddProduct (ctx context.Context, product Product) (totalPrice Price, err error)
	CalculateTotalPrice(context.Context,  Product) (Price, error)
}

type Price int32

type ClientInfo struct {
	Session string
}

type service struct {
	pricingRules pricingrules.Service
}

type Product struct {
	Sku string
	Quantity int32
	Session ClientInfo
}

func New(pricingRules pricingrules.Service) Service {
	return &service{
		pricingRules: pricingRules,
	}
}

func (s service) AddProduct (ctx context.Context, product Product) (Price, error){
	panic("implement me")
}

func (s service) CalculateTotalPrice(ctx context.Context, product Product) (Price, error){
	panic("implement me")
}
