package limiter

import (
	"sync"
	"time"
)

// TokenBucket rate limiter
type TokenBucket struct {
	Rate     float64
	Capacity int

	tokens float64
	last   time.Time
	mu     sync.Mutex
}

// NewTokenBucket creates a TokenBucket
func NewTokenBucket(rate float64, capacity int) *TokenBucket {
	return &TokenBucket{Rate: rate, Capacity: capacity, tokens: float64(capacity)}
}

// Check checks if the rate limit is exceeded
func (t *TokenBucket) Check() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	now := time.Now()
	if !t.last.IsZero() {
		delta := float64(now.Sub(t.last)) * t.Rate
		if t.tokens+delta < float64(t.Capacity) {
			t.tokens += delta
		} else {
			t.tokens = float64(t.Capacity)
		}
	}
	t.last = now
	if t.tokens-1 < 0 {
		return ErrLimitExceeded
	}
	t.tokens--
	return nil
}
