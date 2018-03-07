package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	conn redis.Conn
)

func newRedis(addr string) error {
	var err error
	conn, err = redis.Dial("tcp", addr)
	return err
}

func getUserScore(c redis.Conn, uuid string) (float64, error) {
	val, err := redis.Float64(c.Do("GET", uuid))
	if err != nil {
		return 0.0, err
	}
	return val, nil
}

func getScoredSegment(c redis.Conn, uuid string) string {
	val, err := redis.String(c.Do("GET", uuid))
	if err != nil {
		log.Println("Redis error on get UUID:", err)
	}
	return val
}
