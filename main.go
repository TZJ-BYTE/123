package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "meter@aly",      // 密码，如果没有设置密码则为空
		DB:       0,                // 默认数据库
	})

	// 使用上下文
	ctx := context.Background()

	// 循环存入 dt1 到 dt100
	for i := 7; i <= 100; i++ {
		key := fmt.Sprintf("dt%d", i)
		err := rdb.HSet(ctx, "type", key, "PressureMeter").Err()
		if err != nil {
			fmt.Printf("Error setting %s: %v\n", key, err)
			return
		}
	}

	fmt.Println("Data inserted successfully")
}
