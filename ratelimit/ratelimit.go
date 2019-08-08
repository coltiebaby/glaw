// The ratelimit package limits the number requests made toward the api.
// Based on the leaky bucket algorithm
package ratelimit

// ApiClient uses Limiter as the main way to ratelimit requests.
// It manages all the limiters inside
type Limiter interface {
    // Start rate limiting
	Activate()
    // Stop rate limiting
	Deactivate()
    // Add a new Limit into Limiter
	Add(int, Limit)
    // Take a token from a specific Limit
	Take(index int)
}

type Limit interface {
    // Fills the  bucket usually at a fixed rate
	Recharge()
    // Take a value from the bucket. Should wait if there are no values to get
	Take()
}
