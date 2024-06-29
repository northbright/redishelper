package redishelper_test

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleGetRedisConn() {
	var err error
	var c redis.Conn
	redisAddr := ":6379"
	redisPassword := ""

	if c, err = redishelper.GetRedisConn(redisAddr, redisPassword); err != nil {
		log.Printf("GetRedisConn() error: %v", err)
		return
	}
	defer c.Close()

	log.Printf("Get Redis connection successfully.\n")

	// Output:
}
