package redisDB

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/seew0/rceIdx/domain/models"
)

type RedisStore struct {
	store *redis.Client
}

var ctx = context.Background()

func NewRedisStore(addr string, password string) *RedisStore {
	store := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // uses default db
	})

	return &RedisStore{
		store: store,
	}
}

func (r *RedisStore) Get(uuid string) (models.RespStore, error) {
	val, err := r.store.Get(ctx, uuid).Result()
	if err == redis.Nil {
		return models.RespStore{}, fmt.Errorf("entry does not exist")
	} else if err != nil {
		return models.RespStore{}, err
	}

	var resp models.RespStore
	if err := json.Unmarshal([]byte(val), &resp); err != nil {
		return models.RespStore{}, err
	}

	return resp, nil
}

func (r *RedisStore) Store(uuid string, resp models.RespStore) error {
	if err := r.store.Set(ctx, uuid, resp, 0).Err(); err != nil {
		return err
	}
	
	return nil
}
