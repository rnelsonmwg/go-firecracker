package firecracker

type TokenBucket struct {
	Size         int64 `json:"size"`
	OneTimeBurst int64 `json:"one_time_burst"`
	RefillTime   int64 `json:"refill_time"`
}

type RateLimiter struct {
	Bandwidth TokenBucket `json:"bandwidth"`
	Ops       TokenBucket `json:"ops"`
}
