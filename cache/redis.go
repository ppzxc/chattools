package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Adapter interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Del(key string) error
}

type redisCache struct {
	rdb *redis.Client
}

func (r redisCache) Del(key string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	err = r.rdb.Del(ctx, key).Err()
	cancel()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r redisCache) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	value, err := r.rdb.Get(ctx, key).Result()
	cancel()
	if err == redis.Nil {
		return nil, redis.Nil
	} else if err != nil {
		return nil, err
	} else {
		return []byte(value), nil
	}
}

func (r redisCache) Set(key string, value interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	err := r.rdb.Set(ctx, key, value, 0).Err()
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisCache(address string, password string, db int, isReset bool) Adapter {
	r := &redisCache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password, // no password set
			DB:       db,                              // use default DB
		}),
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := r.rdb.Ping(ctx)
	cancel()
	if err := cmd.Err(); err != nil {
		panic(err)
	}

	if isReset {
		ctx, cancel = context.WithCancel(context.Background())
		r.rdb.FlushDB(ctx)
		cancel()
	}

	return r
}
