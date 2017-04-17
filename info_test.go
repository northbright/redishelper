package redishelper_test

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/northbright/redishelper"
)

func ExampleGetInfo() {
	info := make(map[string]map[string]string)

	log.Printf("-------- GetInfo() Test Start --------\n")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		goto end
	}
	defer c.Close()

	info, err = redishelper.GetInfo(c)
	if err != nil {
		goto end
	}

	log.Printf("INFO:\n")
	for k, v := range info {
		log.Printf("# %v\n", k)
		for name, value := range v {
			log.Printf("%v:%v\n", name, value)
		}
		log.Printf("\n")
	}

end:
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	log.Printf("-------- GetInfo() Test End --------\n\n")
	// Output:
}
