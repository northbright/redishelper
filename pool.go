package redishelper

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// NewRedisPool creates a new pool.
// redisServer: redis server address.
// redisPassword: redis password. Leave it empty if there's no password.
// maxActive: Maximum number of connections allocated by the pool at a given time. When zero, there is no limit on the number of connections in the pool.
// maxIdle: Maximum number of idle connections in the pool.
// idleTimeout: Close connections after remaining idle for this duration(in second).
// If the value is zero, the idle connections are not closed.
// Applications should set the timeout to a value less than the server's timeout.
// wait: If Wait is true and the pool is at the MaxActive limit, then Get() waits for a connection to be returned to the pool before returning.
func NewRedisPool(redisServer, redisPassword string, maxActive, maxIdle, idleTimeout int, wait bool) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Wait:        wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer)
			if err != nil {
				return nil, err
			}

			if len(redisPassword) != 0 {
				if _, err = c.Do("AUTH", redisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
}
