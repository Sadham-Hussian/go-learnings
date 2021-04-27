package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// Product struct
type Product struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrize int    `json:"product_prize"`
}

func main() {
	fmt.Println("Connecting Golang and Redis")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(redisClient.Context()).Result()
	fmt.Println(pong, err)

	// Setting key value
	err = redisClient.Set(redisClient.Context(), "name", "jack", 0).Err()

	if err != nil {
		fmt.Println(err)
	}

	// Getting value
	val, err := redisClient.Get(redisClient.Context(), "name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)

	// storing composite data
	product, err := json.Marshal(Product{
		ProductID:    1234,
		ProductName:  "pen",
		ProductPrize: 10,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = redisClient.Set(redisClient.Context(), "1234", product, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err = redisClient.Get(redisClient.Context(), "1234").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
