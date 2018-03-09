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
	"beeswax/augment"
	"beeswax/openrtb"
	"bytes"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
)

var (
	server  *http.Server
	bufPool *sync.Pool
)

func init() {
	server = &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	finalHandler := http.HandlerFunc(augmentHandler)
	http.Handle("/augment", loggingHandler(augLogger, finalHandler))
	bufPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// getUserID retrieves the user ID from a given bid request
func getUserID(br *openrtb.BidRequest) string {
	usr := br.GetUser()
	if usr == nil {
		return ""
	}
	ue := usr.GetExt()
	if ue == nil {
		return ""
	}
	return ue.GetUserId()
}

// augmentHandler reads the request, which should be a POST
// with a protobuf augmentor request in the body, looks up the
// user ID to retrieve a segment to enrich, and writes
// the response back.
func augmentHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request's body into a bytes.Buffer
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufPool.Put(buf)
	}()
	n, _ := buf.ReadFrom(r.Body)
	if n == 0 {
		w.Header().Set("X-Pass-Reason", "No body")
		w.WriteHeader(204)
		return
	}

	// The body should unarmshal into an AugmentorRequest
	augReq := new(augment.AugmentorRequest)
	err := proto.Unmarshal(buf.Bytes(), augReq)
	if err != nil {
		w.Header().Set("X-Pass-Reason", "Protobuf unmarshal error")
		w.WriteHeader(204)
		return
	}

	// Get the UUID from the request and retrieve the segment from Redis
	uuid := getUserID(augReq.GetBidRequest())
	id := getScoredSegment(conn, uuid)
	if id == "" {
		w.Header().Set("X-Pass-Reason", "No segment to augment")
		w.WriteHeader(204)
		return
	}

	// Create an Augmentor Response and write it back
	resp, respSeg := augment.AugmentorResponse{}, &augment.AugmentorResponse_Segment{}
	respSeg.Id = &id
	resp.Segments = []*augment.AugmentorResponse_Segment{respSeg}
	d, err := resp.Marshal()
	if err != nil {
		log.Printf("Protobuf marshal error: %v", err)
	}
	w.Write(d)
}
