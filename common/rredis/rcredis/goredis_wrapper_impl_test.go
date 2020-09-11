package rcredis

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/go-redis/redis/v8"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewGoRedisCluster(t *testing.T) {
	Convey("New Go Redis Cluster", t, func() {
		mock, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		Convey("should returns object goRedisWrapper", func() {
			redisConfig := &goredis.ClusterOptions{
				Addrs:        []string{mock.Addr()},
				ReadTimeout:  1 * time.Second,
				WriteTimeout: 1 * time.Second,
				PoolSize:     6500,
				PoolTimeout:  30 * time.Second}

			redis := newGoRedisClusterWrapperImpl(redisConfig)
			So(redis, ShouldBeNil)
		})
	})
}
