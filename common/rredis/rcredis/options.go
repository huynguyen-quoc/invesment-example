package rcredis

import (
	"time"
)

// Option is a functional parameter used to configure the clientImpl
type Option func(client *clientImpl)

// Addrs specified the cluster node addresses in []string, each addr is
// in the format of host:port
func Addrs(addrs []string) Option {
	return func(c *clientImpl) {
		c.options.Addrs = addrs
	}
}

// DialTimeout specifies the timeout for connecting (dialing) to a redis.
// default to defaultDialTimeout
func DialTimeout(timeout time.Duration) Option {
	return func(c *clientImpl) {
		c.options.DialTimeout = timeout
	}
}

// ReadTimeout specifies the timeout for reading a single command reply.
func ReadTimeout(timeout time.Duration) Option {
	return func(c *clientImpl) {
		c.options.ReadTimeout = timeout
	}
}

// WriteTimeout specifies the timeout for writing a single command.
func WriteTimeout(timeout time.Duration) Option {
	return func(c *clientImpl) {
		c.options.WriteTimeout = timeout
	}
}

// PoolTimeout specifies the timeout of waiting for connection if all connections are busy before returning an error.
func PoolTimeout(timeout time.Duration) Option {
	return func(c *clientImpl) {
		c.options.PoolTimeout = timeout
	}
}

// IdleTimeout specifies the IdleTimeout of the connections. Close
// connections after remaining idle for this duration. If the value is zero,
// then idle connections are not closed. Applications should set the timeout
// to a value less than the server's timeout.
// default to defaultIdleTimeout
func IdleTimeout(timeout time.Duration) Option {
	return func(c *clientImpl) {
		c.options.IdleTimeout = timeout
	}
}

// PoolSize sets the connection pool size
func PoolSize(poolSize int) Option {
	return func(c *clientImpl) {
		c.options.PoolSize = poolSize
	}
}

// MinIdleConnections specifies the mininum number idle connections maintained in the connection pool.
// Default is 0.
func MinIdleConnections(minIdleConnections int) Option {
	return func(c *clientImpl) {
		c.options.MinIdleConns = minIdleConnections
	}
}

// IdleCheckFrequency specifies the frequency for idle connections reaper.
// Check go-redis library to confirm.
func IdleCheckFrequency(idleCheckFrequency time.Duration) Option {
	return func(c *clientImpl) {
		c.options.IdleCheckFrequency = idleCheckFrequency
	}
}
