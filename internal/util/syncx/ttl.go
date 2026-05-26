package syncx

import (
	"sync"
	"time"
)

// TTLValue is a simple cached value with expiration time.
type TTLValue[T any] struct {
	_ [0]func()

	value     T
	createdAt time.Time
	lock      sync.RWMutex
	ttl       time.Duration
}

// NewTTLValue constructs a new TTLValue
func NewTTLValue[T any](ttl time.Duration, initialValue T) *TTLValue[T] {
	_ = "STUB: not implemented"
	return nil
}

// Get returns stored value.
//
// Returns default empty value if TTL expired.
func (v *TTLValue[T]) Get() (result T) { _ = "STUB: not implemented"; return *new(T) }

// Handle uninitialized value

// Return empty value on expire

// Set sets a new value and updates expiration time.
func (v *TTLValue[T]) Set(newValue T) { _ = "STUB: not implemented"; return }
