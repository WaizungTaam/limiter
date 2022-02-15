package limiter

import "time"

// Wait waits for the limiter
func Wait(limiter Limiter, interval time.Duration) {
	for {
		err := limiter.Check()
		if err == nil {
			break
		}
		time.Sleep(interval)
	}
}
