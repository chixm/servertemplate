package server

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var redisConnections map[string]*redis.Pool

// Make connections to Redis from Config
func InitializeRedis() {
	redisConnections = make(map[string]*redis.Pool)

	for _, r := range Config.Redis {
		pool := makeConnectionPool(r.Host, r.Port, r.MaxIdle, r.MaxActive)
		redisConnections[r.Id] = pool

		// test connection
		_, err := pool.Get().Do("set", "test", "testvalue")
		if err != nil {
			panic(err)
		}
		Logger.Println("Tested Redis Connection of " + r.Host)
		rep, err := redis.String(pool.Get().Do("get", "test"))
		if err != nil {
			panic(err)
		}
		if rep == "testvalue" {
			Logger.Println("Tested Redis " + r.Host + " OK.")
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

func TerminateRedis() {
	for _, c := range redisConnections {
		err := c.Close()
		if err != nil {
			Logger.Println("Redis Close Error::" + err.Error())
		}
	}
}
