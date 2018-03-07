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

package redis

import (
	"encoding/json"
	"io"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

var (
	server *miniredis.Miniredis
)

// Start starts a Miniredis server and returns its address as a string and an error.
// If an error occurs in starting, the address will be empty.
func Start() (string, error) {
	var err error
	server, err = miniredis.Run()
	if err != nil {
		return "", err
	}
	return server.Addr(), nil
}

type segment struct {
	UUID string `json:"uuid"`
	ID   string `json:"segment"`
}

// Load loads sample data into the Miniredis server.
// `data` should be an io,Reader of a JSON object in the form of
// [{"uuid": "x", "segment": y}]
func Load(data io.Reader, c redis.Conn) error {
	d := json.NewDecoder(data)
	var segs []segment
	err := d.Decode(&segs)
	if err != nil {
		return err
	}

	c.Send("MULTI")
	for _, seg := range segs {
		c.Send("SET", seg.UUID, seg.ID)
	}
	_, err = c.Do("EXEC")
	return err
}
