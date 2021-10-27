package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/types"
)

type Adapter interface {
	MGet(keys ...string) ([]interface{}, error)
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	HGetAll(key string) (map[string]string, error)
	HSet(key string, value ...interface{}) error
	Del(key string) error
	HDel(key string, fields ...string) error

	Exists(key ...string) error

	Subscribe(ctx context.Context, key string) (*redis.PubSub, error)
	Publish(ctx context.Context, key string, message interface{}) error
}

type redisCache struct {
	rdb *redis.Client
}

func (r redisCache) Exists(key ...string) error {
	cCtx, cancel := context.WithCancel(context.Background())
	count, err := r.rdb.Exists(cCtx, key...).Result()
	cancel()
	if err != nil {
		return err
	}

	if len(key) > 0 && len(key) == int(count) {
		return nil
	} else {
		return types.ErrNoExistsKeys
	}
}

func (r redisCache) Subscribe(ctx context.Context, key string) (*redis.PubSub, error) {
	cCtx, cancel := context.WithCancel(ctx)
	pub := r.rdb.Subscribe(cCtx, key)
	cancel()

	cCtx, cancel = context.WithCancel(ctx)
	err := pub.Ping(cCtx)
	cancel()
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func (r redisCache) Publish(ctx context.Context, key string, message interface{}) error {
	cCtx, cancel := context.WithCancel(ctx)
	cmd := r.rdb.Publish(cCtx, key, message)
	cancel()
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

func (r redisCache) HDel(key string, fields ...string) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	err = r.rdb.HDel(ctx, key, fields...).Err()
	cancel()
	if err != nil {
		return err
	} else {
		return nil
	}
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

func (r redisCache) MGet(keys ...string) ([]interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := r.rdb.MGet(ctx, keys...)
	cancel()
	fmt.Println(cmd)
	if cmd.Err() == redis.Nil {
		return nil, redis.Nil
	} else if cmd.Err() != nil {
		return nil, cmd.Err()
	} else {
		return cmd.Val(), nil
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

func (r redisCache) HGetAll(key string) (map[string]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	value, err := r.rdb.HGetAll(ctx, key).Result()
	cancel()
	if err == redis.Nil {
		return nil, redis.Nil
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (r redisCache) HSet(key string, values ...interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	err := r.rdb.HSet(ctx, key, values...).Err()
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
			DB:       db,       // use default DB
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
