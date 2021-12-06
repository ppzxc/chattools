package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ppzxc/chattools/common"
	"github.com/ppzxc/chattools/types"
)

type Adapter interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key ...string) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HSet(ctx context.Context, key string, value ...interface{}) error
	HDel(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key string, field string) error

	PubSubNumSub(ctx context.Context, key ...string) (map[string]int64, error)
	Subscribe(ctx context.Context, key ...string) (*redis.PubSub, error)
	PSubscribe(ctx context.Context, key ...string) (*redis.PubSub, error)
	Publish(ctx context.Context, key string, message interface{}) error
}

type redisCache struct {
	rdb *redis.Client
}

func (r redisCache) HExists(ctx context.Context, key string, field string) error {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	exist, err := r.rdb.HExists(redisCtx, key, field).Result()
	cancel()
	if err != nil {
		return err
	}

	if !exist {
		return types.ErrNoExistsKeys
	}
	return nil
}

func (r redisCache) Exists(ctx context.Context, key ...string) error {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	count, err := r.rdb.Exists(redisCtx, key...).Result()
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

func (r redisCache) PubSubNumSub(ctx context.Context, key ...string) (map[string]int64, error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	stringIntMapCmd := r.rdb.PubSubNumSub(redisCtx, key...)
	cancel()

	if stringIntMapCmd.Err() != nil {
		return nil, stringIntMapCmd.Err()
	}
	result, err := stringIntMapCmd.Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r redisCache) PSubscribe(ctx context.Context, key ...string) (*redis.PubSub, error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	pub := r.rdb.PSubscribe(redisCtx, key...)
	cancel()

	redisCtx2, cancel2 := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err := pub.Ping(redisCtx2)
	cancel2()
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func (r redisCache) Subscribe(ctx context.Context, key ...string) (*redis.PubSub, error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	pub := r.rdb.Subscribe(redisCtx, key...)
	cancel()

	redisCtx2, cancel2 := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err := pub.Ping(redisCtx2)
	cancel2()
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func (r redisCache) Publish(ctx context.Context, key string, message interface{}) error {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	cmd := r.rdb.Publish(redisCtx, key, message)
	cancel()
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}

func (r redisCache) HDel(ctx context.Context, key string, fields ...string) (err error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err = r.rdb.HDel(redisCtx, key, fields...).Err()
	cancel()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r redisCache) Del(ctx context.Context, key string) (err error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err = r.rdb.Del(redisCtx, key).Err()
	cancel()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (r redisCache) Get(ctx context.Context, key string) (interface{}, error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	value, err := r.rdb.Get(redisCtx, key).Result()
	cancel()
	if err == redis.Nil {
		return nil, redis.Nil
	} else if err != nil {
		return nil, err
	} else {
		return []byte(value), nil
	}
}

func (r redisCache) Set(ctx context.Context, key string, value interface{}) error {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err := r.rdb.Set(redisCtx, key, value, 0).Err()
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func (r redisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	value, err := r.rdb.HGetAll(redisCtx, key).Result()
	cancel()
	if err == redis.Nil {
		return nil, redis.Nil
	} else if err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func (r redisCache) HSet(ctx context.Context, key string, values ...interface{}) error {
	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	err := r.rdb.HSet(redisCtx, key, values...).Err()
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func NewRedisCache(ctx context.Context, address string, password string, db int, isReset bool) Adapter {
	r := &redisCache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password, // no password set
			DB:       db,       // use default DB
		}),
	}

	redisCtx, cancel := context.WithTimeout(ctx, common.RedisCmdTimeOut)
	cmd := r.rdb.Ping(redisCtx)
	cancel()
	if err := cmd.Err(); err != nil {
		panic(err)
	}

	if isReset {
		redisCtx2, cancel2 := context.WithTimeout(ctx, common.RedisCmdTimeOut)
		r.rdb.FlushDB(redisCtx2)
		cancel2()
	}

	return r
}
