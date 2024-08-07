package redishelper

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// GetConfig returns the Redis Config.
// Return a config map which key is parameter name and value is parameter value.
func GetConfig(c redis.Conn) (config map[string]string, err error) {
	config = make(map[string]string)
	l := 0

	items, err := redis.Strings(c.Do("CONFIG", "GET", "*"))
	if err != nil {
		goto end
	}

	l = len(items)
	if l == 0 || l%2 != 0 {
		err = fmt.Errorf("command 'CONFIG GET *' return error: %v", items)
		goto end
	}

	for i := 0; i < l; i += 2 {
		config[items[i]] = items[i+1]
	}

end:
	return config, err
}
