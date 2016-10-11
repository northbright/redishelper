package redishelper

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// GetRedisConfig() returns the Redis Config.
//
//     Return:
//         config: key: parameter name, value: parameter value.
func GetRedisConfig(c redis.Conn) (config map[string]string, err error) {
	config = make(map[string]string)
	l := 0

	items, err := redis.Strings(c.Do("CONFIG", "GET", "*"))
	if err != nil {
		goto end
	}

	l = len(items)
	if l == 0 || l%2 != 0 {
		errors.New(fmt.Sprintf("Command \"CONFIG GET *\" return error: %v\n", items))
		goto end
	}

	for i := 0; i < l; i += 2 {
		config[items[i]] = items[i+1]
	}

end:
	if err != nil {
		return make(map[string]string), err
	}

	return config, nil
}