package util

import (
	"admin-go-api/common/constant"
	"admin-go-api/pkg/redis"
	"context"
	"log"
	"time"
)

/**
 * @Author: 南宫乘风
 * @Description: redis 存取 验证码
 * @File:  redisStore.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 17:32
 */

var ctx = context.Background()

type RedisStore struct{}

// Set 存取验证码
func (r RedisStore) Set(id string, value string) {
	key := constant.LOGIN_CODE + id
	err := redis.RedisDb.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		log.Panicln(err.Error())
	}
}

// Get 获取验证码
func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + id
	val, err := redis.RedisDb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// Verify 校检验证码
func (r RedisStore) Verify(id string, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
