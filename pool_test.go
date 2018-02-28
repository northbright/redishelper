package redishelper_test

import (
	"log"
	"sync/atomic"

	"github.com/northbright/redishelper"
)

func ExampleNewRedisPool() {
	var (
		success            uint64
		fail               uint64
		total              uint64
		maxActive          = 1000 // try 10000(too many tcp connections.)
		maxIdle            = 100
		idleTimeout        = 60
		wait               = false
		limitedConcurrency = 2000 // try 2000(>maxActive, pool exhausted.)
		concurrency        = limitedConcurrency * 2
	)

	pool := redishelper.NewRedisPool(":6379", "", maxActive, maxIdle, idleTimeout, wait)
	defer pool.Close()

	sem := make(chan struct{}, limitedConcurrency)

	f := func(i int) {
		defer func() { <-sem }()

		atomic.AddUint64(&total, 1)

		conn := pool.Get()
		defer conn.Close()

		k := "pool_test"
		_, err := conn.Do("SET", k, i)
		if err != nil {
			atomic.AddUint64(&fail, 1)
			log.Printf("%v: SET error: %v", i, err)
			return
		}

		atomic.AddUint64(&success, 1)
		log.Printf("%v: SET ok", i)
	}

	for i := 0; i < concurrency; i++ {
		sem <- struct{}{}
		go f(i)
	}

	// After last goroutine is started,
	// there're still "concurrency" amount of goroutines running.
	// Make sure wait all goroutines to finish.
	for j := 0; j < cap(sem); j++ {
		sem <- struct{}{}
		log.Printf("----- j: %v", j)
	}

	totalFinal := atomic.LoadUint64(&total)
	successFinal := atomic.LoadUint64(&success)
	failFinal := atomic.LoadUint64(&fail)
	log.Printf("total: %v, success: %v, fail: %v", totalFinal, successFinal, failFinal)
	// Output:
}
