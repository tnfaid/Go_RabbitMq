package rediscli

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisHash(urlRedis string) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     urlRedis,
		Password: "",
		DB:       0,
	})
	defer client.Close()

	//set a value in the redis hash
	urutanString := "1"
	combinedString := fmt.Sprintf("IncentiveFulfillment:%s", urutanString)

	err := client.HMSet(ctx, combinedString, "id", "id1-isi", "value", "value1-isi").Err()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//get a value from the redis hash
	value, err := client.HGet(ctx, combinedString, "id").Result()
	if err == redis.Nil {
		fmt.Println("Key ga ada yang sama")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", value)
	}

	//Get all value from the Redis Hash
	hashValues, err := client.HGetAll(ctx, combinedString).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n\n\nAll Hash Values")
	for key, value := range hashValues {
		fmt.Printf("%s: %s\n", key, value)
	}

}
