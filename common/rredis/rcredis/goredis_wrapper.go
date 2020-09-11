package rcredis
import (
	"context"
	goredis "github.com/go-redis/redis/v8"
	"huynguyen-quoc/investment/common/rredis/rredisapi"
)

//go:generate mockery -name goRedisWrapper  -case=underscore -testonly
type goRedisWrapper interface {
	Process(ctx context.Context, cmd goredis.Cmder) (interface{}, error)
	Pipeline() goredis.Pipeliner
	Close() error
	Name() string
	getResultFromCommands(cmd []goredis.Cmder) ([]rredisapi.ReplyPair, error)
	Publish(ctx context.Context, channel string, message interface{}) (interface{}, error)
	Subscribe(ctx context.Context, bufferSize int, channels ...string) (response *rredisapi.SubscribeResponse, err error)
}
//go:generate mockery -name clusterClient  -case=underscore -testonly
type clusterClient interface {
	Process(ctx context.Context, cmd goredis.Cmder) error
	Pipeline() goredis.Pipeliner
	Close() error
	Publish(ctx context.Context, channel string, message interface{}) *goredis.IntCmd
	Subscribe(ctx context.Context, channels ...string) *goredis.PubSub
}
