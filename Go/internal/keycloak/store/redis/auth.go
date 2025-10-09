package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// AuthStore defines the interface for auth state management
type AuthStore interface {
	SetState(ctx context.Context, state string) error
	GetState(ctx context.Context, state string) (string, error)
	DeleteState(ctx context.Context, state string) error
}

type authRedisManager struct {
	client *redis.Client
}

func NewAuthRedisManager(client *redis.Client) AuthStore {
	return &authRedisManager{client: client}
}

// Implement AuthStore methods with stub logic for now
func (a *authRedisManager) SetState(ctx context.Context, state string) error {
	// Store the state in Redis with a 10 minute expiration
	err := a.client.Set(ctx, state, state, 10*60*1e9).Err() // 10 minutes in nanoseconds
	if err != nil {
		return err
	}
	return nil
}

func (a *authRedisManager) GetState(ctx context.Context, state string) (string, error) {
	// Use the state as the key to get the value from Redis
	val, err := a.client.Get(ctx, state).Result()
	if err == redis.Nil {
		return "", nil // state not found
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (a *authRedisManager) DeleteState(ctx context.Context, state string) error {
	err := a.client.Del(ctx, state).Err()
	if err != nil {
		return err
	}
	return nil
}
