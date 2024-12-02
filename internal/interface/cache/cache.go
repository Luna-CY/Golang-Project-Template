package cache

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"time"
)

type Cache interface {
	// Set sets the value for the given key with the specified timeout.
	// if key already exist and overwrite is false, do nothing, return nil
	Set(ctx context.Context, key string, value any, timeout time.Duration, overwrite bool) error

	// Get get the value stored at key.
	// parse the value into the dst interface.
	// if key not exist, return errors.ErrorRecordNotFound error
	// if dst is nil, only check the key existence.
	Get(ctx context.Context, key string, dst any) error

	// Ttl get the remaining time to live of the value stored at key.
	// if the key not exist or expired, return -1 and nil
	Ttl(ctx context.Context, key string) (time.Duration, error)

	// Expire sets the expiration time of the value stored at key.
	// if key not exist, return errors.ErrorRecordNotFound error
	Expire(ctx context.Context, key string, timeout time.Duration) error

	// Increment increments the value at key by the given delta.
	// delta allow negative values.
	Increment(ctx context.Context, key string, delta int64) (int64, error)

	// Delete deletes the specified keys.
	Delete(ctx context.Context, keys ...string) error
}
