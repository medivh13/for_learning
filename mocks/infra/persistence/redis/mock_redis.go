package mock_redis

import (
	"context"
	redisServ "for_learning/src/infra/persistence/redis/service"
	"time"

	"github.com/stretchr/testify/mock"
)

// MockRedis adalah implementasi mock dari redis service
type MockRedis struct {
	mock.Mock
}

func NewMockRedis() *MockRedis {
	return &MockRedis{}
}

var _ redisServ.ServRedisInt = &MockRedis{}

func (m *MockRedis) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := m.Called(ctx, key, value, ttl)

	var err error
	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockRedis) GetData(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)

	var (
		data string
		err  error
	)

	if n, ok := args.Get(0).(string); ok {

		data = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return data, err
}
