package main

import "github.com/redis/go-redis/v9"

type RedisObj struct {
	*redis.Client
}

func NewRedis() RedisObj {
	client := redis.NewClient(&redis.Options{
		Addr:     "peifeng.site:6379", // redis地址
		Password: "ningzaichun",       // 密码
		DB:       0,                   // 使用默认数据库
	})
	return RedisObj{client}
}

var rds = NewRedis()
