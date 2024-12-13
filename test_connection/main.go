package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options(
		Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
	))

	if err := rdb.Ping(ctx).Err(); err != nil (
		log.Fatal(err)
	)

	fmt.Printf("success connection to redis")
}