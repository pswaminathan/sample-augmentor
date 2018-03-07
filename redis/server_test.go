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
	c, err = redis.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
