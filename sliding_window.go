package limiter

import (
	"sync"
	"time"
)

// SlidingWindow rate limiter
type SlidingWindow struct {
	Limit  int
	Window time.Duration
	Step   time.Duration

	count int
	start time.Time
	queue []int
	mu    sync.Mutex
}

// NewSlidingWindow creates a SlidingWindow
func NewSlidingWindow(limit int, window time.Duration, step time.Duration) *SlidingWindow {
	return &SlidingWindow{Limit: limit, Window: window, Step: step}
}

// Check checks if the rate limit is exceeded
func (s *SlidingWindow) Check() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if s.start.IsZero() || now.Sub(s.start) >= s.Window {
		s.count = 1
		s.queue = []int{1}
		s.start = now
		return nil
	}
	if now.Sub(s.start) < time.Duration(len(s.queue))*s.Step {
		if s.count >= s.Limit {
			return ErrLimitExceeded
		}
		s.count++
		s.queue[len(s.queue)-1]++
		return nil
	}
	s.start = s.start.Add(s.Step)
	s.count -= s.queue[0]
	s.queue = s.queue[1:]
	if s.count >= s.Limit {
		return ErrLimitExceeded
	}
	s.count++
	s.queue = append(s.queue, 1)
	return nil
}
