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
