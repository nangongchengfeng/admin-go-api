package redis

import (
	"admin-go-api/common/config"
	"context"
)
import "github.com/go-redis/redis/v8"

/**
 * @Author: 南宫乘风
 * @Description: redis 配置初始化
 * @File:  redis.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 15:25
 */

var (
	RedisDb *redis.Client
)

// 初始化连接
func SetupRdisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password, // no password set
		DB:       9,                            // use default DB
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
