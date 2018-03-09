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
	"os"
	"strings"
	"testing"

	"github.com/garyburd/redigo/redis"
)

var (
	addr    string
	c       redis.Conn
	data    = `[{"uuid": "mid.34582634", "segment": "beeswax-123"}]`
	badData = `[{"uuid": "mid.34582635", "segment": 123}]`
)

func TestLoad_good(t *testing.T) {
	userID, want := "mid.34582634", "beeswax-123"
	r := strings.NewReader(data)
	err := Load(r, c)
	if err != nil {
		t.Errorf("load error: %v", err)
	}
	got, err := redis.String(c.Do("GET", userID))
	if err != nil {
		t.Errorf("redis get error: %v", err)
	}
	if got != want {
		t.Errorf("redis get: got %v, want %v", got, want)
	}
}

func TestLoad_bad(t *testing.T) {
	r := strings.NewReader(badData)
	err := Load(r, c)
	if err == nil {
		t.Error("load bad data: want error, got none")
	}
}

func TestMain(m *testing.M) {
	var err error
	addr, err = Start()
	if err != nil {
		panic(err)
	}
	defer server.Close()
	c, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	os.Exit(m.Run())
}
