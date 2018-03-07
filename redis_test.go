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
	got, err := getUserScore(c, "doesnotexist")
	if err == nil {
		t.Error("TestGetUserScore_notPresent: want error, got none")
	}
	if want := 0.0; got != want {
		t.Errorf("TestGetUserScore_notPresent: got %v, want %v", got, want)
	}
}

func TestGetUserScore_badData(t *testing.T) {
	user, input := "mid.283ADF1D", "hi"
	s.Set(user, input)
	defer s.Del(user)
	got, err := getUserScore(c, user)
	if err == nil {
		t.Error("TestGetUserScore_notPresent: want error, got none")
	}
	if want := 0.0; got != want {
		t.Errorf("TestGetUserScore_notPresent: got %v, want %v", got, want)
	}
}

func TestGetUserScore_present(t *testing.T) {
	user, want := "mid.283ADF1E", 10.0
	s.Set(user, "10.0")
	defer s.Del(user)
	got, err := getUserScore(c, user)
	if err != nil {
		t.Errorf("TestGetUserScore_present: expected no error, got %v", err)
	}
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
