package rcredis

import (
	"context"
	"fmt"
	goredis "github.com/go-redis/redis/v8"
	"huynguyen-quoc/investment/common/rredis/rredisapi"
)

const (
	defaultHostAndPort = "localhost:6379"
)

type clientImpl struct {
	redisClient goRedisWrapper
	options     *goredis.ClusterOptions
}


func NewClusterClient(ctx context.Context, options ...Option) (rredisapi.CacheEngine, error) {
	c := &clientImpl{
		options: &goredis.ClusterOptions{
			Addrs:       []string{defaultHostAndPort},
			DialTimeout: defaultDialTimeout,
		},
	}

	for _, opt := range options {
		opt(c)
	}

	c.redisClient = newGoRedisClusterWrapperImpl(c.options)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		return c, nil
	}
}

// Do implements Client interface
func (c clientImpl) Do(ctx context.Context, cmdName string, args ...interface{}) (interface{}, error) {
	allArgs := rredisapi.NewArgs(cmdName).Add(args...)
	cmd := goredis.NewCmd(ctx, allArgs.Value()...)

	return c.redisClient.Process(ctx, cmd)
}

func (c clientImpl) Publish(ctx context.Context, channelName string, value interface{}) (interface{}, error) {
	return c.redisClient.Publish(ctx, channelName, value)
}

func (c clientImpl) ShutDown(ctx context.Context) {
	if _, ok := ctx.Deadline(); !ok {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, defaultShutdownTimeout)
		ctx = ctxWithTimeout
		defer cancel()
	}
	done := make(chan struct{})
	go func() {
		if err := c.redisClient.Close(); err != nil {
			fmt.Printf("Failed to close redisClient with error: [%s]", err)
		}
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Gracefully shutdown")
	case <-ctx.Done():
		fmt.Printf("ShutDown ctx Done with: [%s]", ctx.Err())
	}
}