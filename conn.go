package redishelper

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Dial connects to the Redis server and returns the connection.
func Dial(redisAddr, redisPassword string) (c redis.Conn, err error) {
	pongStr := ""

	if c, err = redis.Dial("tcp", redisAddr); err != nil {
		return nil, err
	}

	if len(redisPassword) != 0 {
		if _, err = c.Do("AUTH", redisPassword); err != nil {
			return nil, err
		}
	}

	if pongStr, err = redis.String(c.Do("PING")); err != nil {
		return nil, err
	}

	if pongStr != "PONG" {
		return nil, fmt.Errorf("redis PING != PONG(%v)", pongStr)
	}

	return c, nil
}
