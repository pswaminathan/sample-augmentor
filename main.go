package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pswaminathan/sample-augmentor/redis"
)

var (
	port       int
	sampleFile string
)

func main() {
	flag.Parse()
	addr, err := redis.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Redis addr: %v", addr)
	err = newRedis(addr)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(sampleFile)
	if err != nil {
		log.Fatal(err)
	}
	err = redis.Load(f, conn)
	if err != nil {
		log.Fatal(err)
	}

	server.Addr = fmt.Sprintf(":%d", port)
	server.ListenAndServe()
}

func init() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&sampleFile, "samplefile", "", "REQUIRED: Filename of sample data")
}
