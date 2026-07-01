package hd_redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type redisHook struct{}

func (h redisHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (h redisHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		fmt.Println("Executing:", cmd.Name())

		err := next(ctx, cmd)

		fmt.Println("Finished:", cmd.Name())

		return err
	}
}
func (h redisHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return next
}

func InitializeRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Username: "",
		Password: "",
	})
	rdb.AddHook(redisHook{})

	rdb.Ping(ctx)
}
