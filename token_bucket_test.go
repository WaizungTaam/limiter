package limiter

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	c := NewTokenBucket(0.02, 50)
	pass := 0
	start := time.Now()
	for i := 0; i < 100; i++ {
		if err := c.Check(); err != nil {
			pass++
		}
	}
	end := time.Now()
	if end.Sub(start) <= 1*time.Second && pass != 50 {
		t.Errorf("wrong #pass: %v", pass)
	}
}
