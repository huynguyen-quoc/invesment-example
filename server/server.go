package server

import (
	"fmt"
	"runtime/debug"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
)

func Serve() {

}

func doServerPanicRecovery(serviceName string) {
	if r := recover(); r != nil {
		fmt.Printf("[%s] service got exception, failed with error [%s] [%s]", serviceName, r, string(debug.Stack()))
	}
}
