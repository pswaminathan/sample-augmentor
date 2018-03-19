// Copyright 2018 BeeswaxIO Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// getUserScore looks up a user ID in Redis and returns a float score
func getUserScore(c redis.Conn, uuid string) float64 {
	val, _ := redis.Float64(c.Do("GET", uuid))
	return val
}

// getScoredSegment looks up a user ID in Redis and returns a segment string.
func getScoredSegment(c redis.Conn, uuid string) string {
	val, _ := redis.String(c.Do("GET", uuid))
	return val
}
