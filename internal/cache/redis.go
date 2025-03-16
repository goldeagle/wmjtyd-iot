package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisDB *redis.Client //创建redis客户端实例

func InitDatabaseRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("Redis.Host") + ":" + viper.GetString("Redis.Port"),
		Password:     viper.GetString("Redis.Password"),
		DB:           viper.GetInt("Redis.Database"),
		ReadTimeout:  2 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		PoolTimeout:  2 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		PoolSize:     1000,
	})
	var ctx = context.Background()
	_, err := RedisDB.Ping(ctx).Result()

	if err != nil {
		panic("连接 Redis 失败：" + err.Error())
	} else {
		log.Println("Redis connected")
	}
}
