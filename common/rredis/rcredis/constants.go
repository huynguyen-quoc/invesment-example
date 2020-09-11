package rcredis

import "time"

const (
	// default timeout
	defaultShutdownTimeout = 5 * time.Second
	defaultDialTimeout     = 5 * time.Second
)
