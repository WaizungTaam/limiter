package limiter

import (
	"sync"
	"time"
)

// Counter rate limiter
type Counter struct {
	Limit  int
	Window time.Duration

	count int
	last  time.Time
	mu    sync.Mutex
}

// NewCounter creates a Counter
func NewCounter(limit int, window time.Duration) *Counter {
	return &Counter{Limit: limit, Window: window}
}

// Check checks if the rate limit is exceeded
func (c *Counter) Check() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	if c.last.IsZero() || now.Sub(c.last) >= c.Window {
		c.count = 0
		c.last = now
	}
	if c.count < c.Limit {
		c.count++
		return nil
	}
	return ErrLimitExceeded
}
