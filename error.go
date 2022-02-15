package limiter

import "fmt"

var (
	// ErrLimitExceeded indicates rate limit exceeded
	ErrLimitExceeded = fmt.Errorf("limit exceeded")
)
