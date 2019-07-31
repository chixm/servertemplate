package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisConnections map[string]*redis.Pool

// Make connections to Redis from Config
func initializeRedis() {
	redisConnections = make(map[string]*redis.Pool)

	for _, r := range config.Redis {
		pool := makeConnectionPool(r.Host, r.Port, r.MaxIdle, r.MaxActive)
		redisConnections[r.Id] = pool

		// test connection
		_, err := pool.Get().Do("set", "test", "testvalue")
		if err != nil {
			panic(err)
		}
		logger.Println("Tested Redis Connection of " + r.Host)
		rep, err := redis.String(pool.Get().Do("get", "test"))
		if err != nil {
			panic(err)
		}
		if rep == "testvalue" {
			logger.Println("Tested Redis " + r.Host + " OK.")
		}
	}
}

func makeConnectionPool(host string, port int, maxIdle, maxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", host+":"+strconv.Itoa(port)) },
	}
}

func terminateRedis() {
	for _, c := range redisConnections {
		err := c.Close()
		if err != nil {
			logger.Println("Redis Close Error::" + err.Error())
		}
	}
}

// Use default setting and write to redis
func setRedisObject(key string, obj interface{}, second int) error {
	pool, ok := redisConnections[`default`]
	if !ok {
		panic(`No default redis pool found.`)
	}
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if _, err = pool.Get().Do("set", key, second, b); err != nil {
		return err
	}
	return nil
}

// get key data from redis
func getRedisObject(key string) ([]byte, error) {
	pool, ok := redisConnections[`default`]
	if !ok {
		panic(`No default redis pool found.`)
	}
	return redis.Bytes(pool.Get().Do("get", key))
}
