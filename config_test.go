package redishelper_test

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleGetRedisConfig() {
	config := make(map[string]string)

	log.Printf("-------- GetRedisConfig() Test Start --------\n")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		goto end
	}
	defer c.Close()

	config, err = redishelper.GetRedisConfig(c)
	if err != nil {
		goto end
	}

	log.Printf("Redis Config:\n")
	for k, v := range config {
		log.Printf("%v: %v\n", k, v)
	}
end:
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("-------- GetRedisConfig() Test End --------\n\n")
	// Output:
}
