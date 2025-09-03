package utils

import (
	"sync"
	"time"
)

// RateLimiter structure to manage rate limits
type RateLimiter struct {
	mu       sync.Mutex
	store    map[string]int       // Stores the count of requests per phone number
	expiry   map[string]time.Time // Stores the expiry time for each phone number
	maxRequests int               // Maximum requests allowed
	timeWindow time.Duration      // Time window for limiting requests
}

// NewRateLimiter initializes the rate limiter
func NewRateLimiter(maxRequests int, timeWindow time.Duration) *RateLimiter {
	return &RateLimiter{
		store:      make(map[string]int),
		expiry:     make(map[string]time.Time),
		maxRequests: maxRequests,
		timeWindow:  timeWindow,
	}
}

// IsAllowed checks if a phone number can make a request
func (rl *RateLimiter) IsAllowed(phone string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Check if the rate limit has expired
	if time.Now().After(rl.expiry[phone]) {
		// Reset the count and expiry if the time window has passed
		rl.store[phone] = 0
		rl.expiry[phone] = time.Now().Add(rl.timeWindow)
	}

	// Check if the phone number has exceeded the rate limit
	if rl.store[phone] >= rl.maxRequests {
		return false
	}

	// Increment the request count and allow the request
	rl.store[phone]++
	return true
}