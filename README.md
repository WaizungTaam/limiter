# Go Limiter

Rate limiters.
- Counter
- Sliding window
- Leaky bucket
- Token bucket

## Usage
```go
package main

import (
    "fmt"
    "time"

    "github.com/waizungtaam/limiter"
)

func main() {
    b := limiter.NewTokenBucket(0.2, 50)
    for i := 0; i < 100; i++ {
        limiter.Wait(b, 20*time.MilliSecond)
        fmt.Println("Hello")
    }
}
```

## Author
waizungtaam

## License
MIT
