package rredisapi

import "context"
// UnsubscribeFunc tells a connection to cancel all it's subscriptions
type UnsubscribeFunc func()

// SubscribeMessage message notification from the pubsub channel
type SubscribeMessage struct {
	// The originating channel.
	Channel string

	// The message data.
	Data []byte
}

type ReplyPair struct {
	Value interface{}
	Err   error
}

type SubscribeResponse struct {
	// ResultChan returns either a SubscribeMessage or an error
	ResultChan  <-chan interface{}
	Unsubscribe UnsubscribeFunc
}

// CacheEngine is an interface for cache engine (e.g. in-memory cache or Redis Cache)
type CacheEngine interface {
	Doer
	Publisher

	ShutDown(ctx context.Context)
}

// Doer interface defines something send single redis commands
type Doer interface {
	// Do sends a redis command to a read and write enabled node
	Do(ctx context.Context, cmdName string, args ...interface{}) (interface{}, error)
}

// Publisher interface defines osmething that can perform redis publish
type Publisher interface {
	// Publish publishes to a Redis channel and returns an error
	Publish(ctx context.Context, channelName string, value interface{}) (interface{}, error)
}
