package rcredis

import (
	"context"
	"huynguyen-quoc/investment/common/rredis/rredisapi"
	"strings"

	goredis "github.com/go-redis/redis/v8"
)

func newGoRedisClusterWrapperImpl(options *goredis.ClusterOptions) goRedisWrapper {
	wrapper := &goRedisWrapperImpl{
		client:    goredis.NewClusterClient(options),
		addrs:     options.Addrs,
		closeChan: make(chan struct{}),
	}
	return wrapper
}

type goRedisWrapperImpl struct {
	client    clusterClient
	addrs     []string
	closeChan chan struct{}
}

func (w *goRedisWrapperImpl) Process(ctx context.Context, cmd goredis.Cmder) (interface{}, error) {
	_ = w.client.Process(ctx, cmd)
	reply, err := cmd.(*goredis.Cmd).Result()
	if err == goredis.Nil {
		err = nil
	}
	return reply, err
}

func (w *goRedisWrapperImpl) Pipeline() goredis.Pipeliner {
	return w.client.Pipeline()
}

func (w *goRedisWrapperImpl) Close() error {
	if w.closeChan != nil {
		close(w.closeChan) // stop ticker
	}
	return w.client.Close()
}

func (w *goRedisWrapperImpl) Name() string {
	return strings.Join(w.addrs, ",")
}

// getResultFromCommands implements goRedisWrapper
func (w *goRedisWrapperImpl) getResultFromCommands(cmds []goredis.Cmder) ([]rredisapi.ReplyPair, error) {
	results := make([]rredisapi.ReplyPair, len(cmds))
	var err error
	for idx, cmd := range cmds {
		results[idx].Value, results[idx].Err = cmd.(*goredis.Cmd).Result()
		if results[idx].Err == goredis.Nil {
			results[idx].Err = nil
			continue
		}
		if err == nil && results[idx].Err != nil {
			err = results[idx].Err
		}
	}

	return results, err
}

// Publish implements Publisher
func (w *goRedisWrapperImpl) Publish(ctx context.Context, channelName string, value interface{}) (interface{}, error) {
	cmd := w.client.Publish(ctx, channelName, value)
	return cmd.Result()
}

// Subscribe implements Subscriber
func (w *goRedisWrapperImpl) Subscribe(ctx context.Context, bufferSize int, channels ...string) (*rredisapi.SubscribeResponse, error) {
	sub := w.client.Subscribe(ctx, channels...)

	ch := sub.Channel()
	resultChan := make(chan interface{}, bufferSize)
	go func() {
		for msg := range ch {
			resultChan <- &rredisapi.SubscribeMessage{
				Channel: msg.Channel,
				Data:    []byte(msg.Payload),
			}
		}

		// close channel
		close(resultChan)
	}()

	return &rredisapi.SubscribeResponse{
		ResultChan: resultChan,
		Unsubscribe: func() {
			_ = sub.Unsubscribe(ctx, channels...)
		},
	}, nil
}
