package pricingrules

import (
	"context"
	"sort"
	"sync"
)

// InMemoryStore keeps pricingRules in the memory.
// Use it in tests or for development/demo purposes.
type InMemoryStore struct {
	pricingRules     []PricingRule
	commonPricingRules []CommonPricingRule
	once sync.Once
	mu        sync.RWMutex
}

// NewInMemoryStore returns a new in-memory pricingRule store.
func NewInMemoryStore() *InMemoryStore {
	store := &InMemoryStore{}

	return store
}


// Reset re-sets a new pricingRules.
func (s *InMemoryStore) Reset(_ context.Context, pricingRules []PricingRule) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pricingRules = pricingRules

	return nil
}

// Reset re-sets a new pricingRules.
func (s *InMemoryStore) ResetCommonRules(_ context.Context, commonRules []CommonPricingRule) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.commonPricingRules = commonRules

	return nil
}

// GetAll returns all pricingRules.
func (s *InMemoryStore) GetCommonRules(_ context.Context) ([]CommonPricingRule, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pricingRules := make([]CommonPricingRule, 0,  len(s.commonPricingRules))

	for _,pr := range s.commonPricingRules{
			pricingRules = append(pricingRules, pr)
	}

	sort.Slice(pricingRules, func(i, j int) bool {
		return pricingRules[i].Weight > pricingRules[j].Weight
	})


	return pricingRules, nil
}

func (s *InMemoryStore) GetRulesBySkus(_ context.Context, skus []int64) ([]PricingRule, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	pricingRules := make([]PricingRule, 0,  len(s.pricingRules))

	for _, pr := range s.pricingRules{
			pricingRules = append(pricingRules, pr)
	}

	sort.Slice(pricingRules, func(i, j int) bool {
		return pricingRules[i].Weight > pricingRules[j].Weight
	})


	return pricingRules, nil
}