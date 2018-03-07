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
	http.HandleFunc("/augment", augmentHandler)
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
