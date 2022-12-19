package rediscache

import (
	"context"
	"demo-fiber-mysql-gorm/config"

	"github.com/go-redis/redis/v9"
)

var RedisCaching RedisCache

type (
	RedisCache struct {
		RedisClient *redis.Client   `json:"redis_client"`
		EnvPrefix   string          `json:"env_prefix"`
		Contx       context.Context `json:"contx"`
		MakeCache   MakeCache       `json:"make_cache"`
	}

	MakeCache struct {
		Key    string      `json:"key" validate:"required"`
		Data   interface{} `json:"data" validate:"required"`
		Expire string      `json:"expire"`
	}
)

func InitRedisCache() {

	RedisCaching.EnvPrefix = "LOCAL_"
	RedisCaching.Contx = context.Background()
	RedisCaching.RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Env.REDIS_ADDRESS,
		Password: config.Env.REDIS_PASSWORD,
		DB:       config.Env.REDIS_DB_NUM,
	})

}
