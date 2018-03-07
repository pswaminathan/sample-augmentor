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
