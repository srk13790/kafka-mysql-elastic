package database

import (
    "context"
    "github.com/redis/go-redis/v9"
)

func InsertData(data string) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-11553.c305.ap-south-1-1.ec2.redns.redis-cloud.com:11553", // Redis server address
		Password: "FPBSzOyu4qtLvRmNRMTpSlxeAiYPFBb8",               // No password set
		DB:       0,                // Use default DB
	})

	err := rdb.Set(ctx, "crawledData", data, 0).Err()
	if err != nil {
		panic(err)
	}
}