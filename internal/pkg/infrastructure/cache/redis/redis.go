package redis

import (
	"context"
	"time"

	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/cache"
	"github.com/cold-runner/Hylark/internal/pkg/instance"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func NewCache(opt *instance.RedisConfig) cache.Cache {
	client, err := instance.NewRedis(opt)
	if err != nil {
		panic(errors.Errorf("init redis failed! err: %v", err))
	}
	return &redisClient{client: client}
}

func (r *redisClient) Set(c context.Context, key string, value interface{}) error {
	if err := r.client.Set(c, key, value, redis.KeepTTL).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) SetExpiration(c context.Context, key string, value interface{}, expiration time.Duration) error {
	result, err := r.client.SetNX(c, key, value, expiration).Result()
	if err != nil {
		return err
	}
	if !result {
		return errors.WithMessagef(err, "设置失败，key: %v, val: %v", key, value)
	}
	return nil
}

func (r *redisClient) SetHash(c context.Context, hashName string, key string, value interface{}) error {
	result, err := r.client.HSet(c, hashName, key, value).Result()
	if err != nil {
		return err
	}
	if result != 1 {
		return errors.WithMessagef(err, "设置HSET失败，haseName: %v, key: %v, val: %v", hashName, key, value)
	}
	return nil
}

func (r *redisClient) SetHashExpiration(c context.Context, hashName string, key string, value interface{}, expiration time.Duration) error {
	_, err := r.client.HSet(c, hashName, key, value).Result()
	if err != nil {
		return err
	}
	return r.Expire(c, hashName, expiration)
}

func (r *redisClient) SetHashMulti(c context.Context, hashName string, kvPair map[string]interface{}) error {
	result, err := r.client.HSet(c, hashName, kvPair).Result()
	if err != nil {
		return err
	}
	if result != int64(len(kvPair)) {
		return errors.WithMessagef(err, "批量设置HSET失败，haseName: %v, kvPair: %v", hashName, kvPair)
	}
	return nil
}

func (r *redisClient) Get(c context.Context, key string) (interface{}, error) {
	result, err := r.client.Get(c, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		return nil, nil
	case err != nil:
		return nil, err
	}

	return result, nil
}

func (r *redisClient) GetDel(c context.Context, key string) (interface{}, error) {
	result, err := r.client.GetDel(c, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *redisClient) Del(c context.Context, key string) error {
	result, err := r.client.Del(c, key).Result()
	if err != nil {
		return err
	}
	if result == int64(1) {
		return nil
	}
	return errors.Errorf("redis删除键错误 err: %v", err)
}

func (r *redisClient) IsExpired(c context.Context, key string) (bool, error) {
	result, err := r.client.TTL(c, key).Result()
	if err != nil {
		return false, err
	}
	if result <= 0 {
		return true, nil
	}
	return false, nil
}

func (r *redisClient) Expire(c context.Context, key string, expiration time.Duration) error {
	success, err := r.client.Expire(c, key, expiration).Result()
	if err != nil {
		return err
	}
	if !success {
		return errors.New("设置超时时间失败")
	}
	return nil
}
