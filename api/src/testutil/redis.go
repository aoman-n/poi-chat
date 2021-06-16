package testutil

import (
	"context"
	"testing"

	"github.com/laster18/poi/api/src/infra/redis"
)

func SetupRedis(t *testing.T) *redis.Client {
	t.Helper()
	r := redis.New()
	t.Cleanup(func() {
		r.FlushDB(context.Background())
	})

	return r
}
