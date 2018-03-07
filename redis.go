package main

import (
	"github.com/garyburd/redigo/redis"
)

var (
	conn redis.Conn
)

// newRedis dials a Redis connection, stores it in the global conn,
// and returns any error encountered.
func newRedis(addr string) error {
	var err error
	conn, err = redis.Dial("tcp", addr)
	return err
}

// getUserScore looks up a user ID in Redis and returns a float
// score and an error.
func getUserScore(c redis.Conn, uuid string) (float64, error) {
	val, err := redis.Float64(c.Do("GET", uuid))
	if err != nil {
		return 0.0, err
	}
	return val, nil
}

// getScoredSegment looks up a user ID in Redis and returns a
// segment string.
func getScoredSegment(c redis.Conn, uuid string) string {
	val, _ := redis.String(c.Do("GET", uuid))
	return val
}
