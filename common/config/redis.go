package rredis

import (
	"context"
)

// RedisConfig contains configuration for redis
type RedisConfig struct {
	Host                string `json:"host"`
	SlaveHost           string `json:"slaveHost"`
	Port                int    `json:"port"`
	MaxConnInSec        int    `json:"maxConn"`
	MaxIdleInSec        int    `json:"maxIdle"`
	ConnectTimeoutInSec int    `json:"connectTimeoutInSec"` // In Second
	ReadTimeoutInSec    int    `json:"readTimeout"`         // In Second
	WriteTimeoutInSec   int    `json:"writeTimeout"`        // In Second
	Db                  int    `json:"db"`
}

// RedisClusterConfig defines the connection parameters used when connecting to a redis cluster server.
type RedisClusterConfig struct {
	Addr               string `json:"addr"`
	PoolSize           int    `json:"poolSize"`
	ReadTimeoutInSec   int    `json:"readTimeoutInSec"`
	WriteTimeoutInSec  int    `json:"writeTimeoutInSec"`
	IdleTimeoutInSec   int    `json:"idleTimeoutInSec"`
	ReadOnlyFromSlaves bool   `json:"readOnlyFromSlaves"`
}

func (rcc *RedisClusterConfig) NewRedisClusterClient(ctx context.Context)  {

}

func (rc *RedisConfig) NewRedisClient()  {

}

