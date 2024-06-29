package redishelper_test

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleGetConfig() {
	config := make(map[string]string)

	log.Printf("-------- GetConfig() Test Start --------\n")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		goto end
	}
	defer c.Close()

	config, err = redishelper.GetConfig(c)
	if err != nil {
		goto end
	}

	log.Printf("Redis Config:\n")
	for k, v := range config {
		log.Printf("%v: %v\n", k, v)
	}
end:
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("-------- GetConfig() Test End --------\n\n")
	// Output:
}
