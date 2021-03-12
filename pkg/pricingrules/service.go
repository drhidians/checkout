package pricingrules

import "context"

// +kit:endpoint:errorStrategy=service

type Service interface {
	ApplyNewPricingRules(_ context.Context, rules []PricingRule) error
	ApplyNewCommonPricingRules(ctx context.Context, rules []CommonPricingRule) error
}

type Rule int
const(
	// sku specific rules
	MultiPriceDiscount Rule = iota

	// common rules
	OverAllDiscount

	// user(personalized) specific rules e.g. premium user...

	// region specific rules ...

	// time specific rules ...

	// store specific rules ....
)

type BaseRule struct {
	Rule Rule
	Weight int32 // priority of pricing rule

	// Period, Start date, End date etc.
}

type PricingRule struct {
	BaseRule
	Sku string
}

type CommonPricingRule struct {
	BaseRule
}

type service struct {
	store Store
}

func New(store Store) Service{
	return &service{
		store: store,
	}
}

func (pr service) ApplyNewPricingRules(ctx context.Context, rules []PricingRule) error {
	return pr.store.Reset(ctx, rules)
}

func (pr service) ApplyNewCommonPricingRules(ctx context.Context, rules []CommonPricingRule) error {
	return pr.store.ResetCommonRules(ctx, rules)
}

type Store interface {
	 GetRulesBySkus(ctx context.Context, skus []int64) ([]PricingRule, error)
	 GetCommonRules(context.Context) ([]CommonPricingRule, error)
	 Reset(context.Context, []PricingRule) error
	 ResetCommonRules(_ context.Context, commonRules []CommonPricingRule) error
}
