package limiter

// Limiter rate limiter
type Limiter interface {
	Check() error
}
