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
	"os"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

var (
	c             redis.Conn
	s             *miniredis.Miniredis
	sampleSegment = `[{"uuid": "mid.283ADF1E", "segment": "beeswax-1234"}]`
	sampleScore   = `[{"uuid": "mid.283ADF1F", "segment": 10.0}]`
)

func TestGetUserScore_notPresent(t *testing.T) {
	got := getUserScore(c, "doesnotexist")
	if want := 0.0; got != want {
		t.Errorf("TestGetUserScore_notPresent: got %v, want %v", got, want)
	}
}

func TestGetUserScore_badData(t *testing.T) {
	user, input := "mid.283ADF1D", "hi"
	s.Set(user, input)
	defer s.Del(user)
	got := getUserScore(c, user)
	if want := 0.0; got != want {
		t.Errorf("TestGetUserScore_notPresent: got %v, want %v", got, want)
	}
}

func TestGetUserScore_present(t *testing.T) {
	user, want := "mid.283ADF1E", 10.0
	s.Set(user, "10.0")
	defer s.Del(user)
	got := getUserScore(c, user)
	if got != want {
		t.Errorf("TestGetUserScore_present: got %v, want %v", got, want)
	}
}

func TestGetUserSegment_notPresent(t *testing.T) {
	user, want := "doesnotexist", ""
	if got := getScoredSegment(c, user); got != want {
		t.Errorf("TestGetUserSegment_present: got %v, want %v", got, want)
	}
}

func TestGetUserSegment_present(t *testing.T) {
	user, want := "mid.283ADF1F", "beeswax-1234"
	s.Set(user, want)
	defer s.Del(user)
	if got := getScoredSegment(c, user); got != want {
		t.Errorf("TestGetUserSegment_present: got %v, want %v", got, want)
	}
}

func TestNewRedis_badAddr(t *testing.T) {
	err := newRedis(":12345")
	if err == nil {
		t.Error("new redis bad address: want error, got none")
	}
}

func TestNewRedis(t *testing.T) {
	err := newRedis(s.Addr())
	if err != nil {
		t.Errorf("new redis: got error: %v", err)
	}
}

func TestMain(m *testing.M) {
	var err error
	s, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	c, err = redis.Dial("tcp", s.Addr())
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
