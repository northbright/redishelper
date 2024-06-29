package redishelper_test

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleDial() {
	var err error
	var c redis.Conn
	redisAddr := ":6379"
	redisPassword := ""

	if c, err = redishelper.Dial(redisAddr, redisPassword); err != nil {
		log.Printf("Dial() error: %v", err)
		return
	}
	defer c.Close()

	log.Printf("Dial() successfully.")

	// Output:
}
