package redishelper_test

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleGetRedisConn() {
	var err error
	var c redis.Conn
	redisAddr := ":6379"
	redisPassword := ""

	log.Printf("-------- GetRedisConn() Test Start --------\n")

	if c, err = redishelper.GetRedisConn(redisAddr, redisPassword); err != nil {
		goto end
	}
	defer c.Close()

	log.Printf("Get Redis connection successfully.\n")

end:
	if err != nil {
		fmt.Printf("GetRedisConn() error: %v\n", err)
	}

	log.Printf("-------- GetRedisConn() Test End --------\n\n")
	// Output:
}
