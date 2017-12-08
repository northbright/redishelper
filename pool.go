package redishelper

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// NewRedisPool creates a new pool.
func NewRedisPool(redisAddr, redisPassword string, maxIdle, maxActive, idleTimeout int) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddr)
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
