package builder

import "github.com/stretchr/testify/mock"

// FakeBuilder implements a mockable Builder interface
type FakeBuilder struct {
	mock.Mock
}

var _ Builder = &FakeBuilder{}

// Build implement Builder
func (b *FakeBuilder) Build(bo BuildOpts) *Result {
	o := b.Called(bo)
	return o.Get(0).(*Result)
}
