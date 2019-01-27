package service

import (
	"context"
	"github.com/amku0xb/go-starter/config"
	"github.com/amku0xb/go-starter/constant"
	"github.com/amku0xb/go-starter/logger"
	"github.com/go-redis/redis"
)

func InitRedisClient(config *config.AppConfig) *redis.Client {
	log := logger.GetLogger(context.Background())

	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    config.RedisMaster,
		SentinelAddrs: config.RedisSentinelServers,
		Password:      config.RedisPassword,
		PoolSize:      config.RedisPoolSize,
		DB:            config.RedisDB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Logger.WithError(err).Error(constant.InitRedisFail)
		return nil
	}
	log.Logger.Info(constant.InitRedisSuccess)
	return client
}
