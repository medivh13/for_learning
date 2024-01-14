package mock_limiter

import (
	"github.com/stretchr/testify/mock"
	limiter "for_learning/src/infra/limiter"
)

// Limiter adalah antarmuka untuk rate limiter
type Limiter interface {
	Allow() bool
}

// MockLimiter adalah implementasi mock dari Limiter
type MockLimiter struct {
	mock.Mock
}

func NewMockLimiter() *MockLimiter {
	return &MockLimiter{}
}

var _ limiter.RateLimiterInterface = &MockLimiter{}

// Allow adalah metode mock untuk memeriksa apakah rate limit diizinkan
func (m *MockLimiter) Allow() bool {
	args := m.Called()

	var data bool
	if n, ok := args.Get(0).(bool); ok {

		data = n
	}

	return data
}
