package firecracker

type TokenBucket struct {
	// Size of the bucket
	Size int64 `json:"size,omitempty"`
	// OneTimeBurst is the initial size of the bucket, may be bigger than overall Size
	OneTimeBurst int64 `json:"one_time_burst,omitempty"`
	// RefillTime - the amount of milliseconds for the bucket to refill
	RefillTime int64 `json:"refill_time,omitempty"`
}

type RateLimiter struct {
	// Bytes is a Bytes per Second limiter
	Bytes TokenBucket `json:"bandwidth,omitempty"`
	// Operations is a operations per second limiter
	Operations TokenBucket `json:"ops,omitempty"`
}
