package limiter

import (
	"sync"
	"time"
)

// LeakyBucket rate limiter
type LeakyBucket struct {
	Rate     float64
	Capacity int

	amount float64
	last   time.Time
	mu     sync.Mutex
}

// NewLeakyBucket creates a LeakyBucket
func NewLeakyBucket(rate float64, capacity int) *LeakyBucket {
	return &LeakyBucket{Rate: rate, Capacity: capacity}
}

// Check checks if the rate limit is exceeded
func (b *LeakyBucket) Check() error {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	if !b.last.IsZero() {
		delta := float64(now.Sub(b.last)) * b.Rate
		if b.amount-delta > 0 {
			b.amount -= delta
		} else {
			b.amount = 0
		}
	}
	b.last = now
	if b.amount+1 > float64(b.Capacity) {
		return ErrLimitExceeded
	}
	b.amount++
	return nil
}
