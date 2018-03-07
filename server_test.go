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
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/gogo/protobuf/proto"
)

func TestAugment_noBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/augment", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(augmentHandler)
	handler.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNoContent; got != want {
		t.Errorf("TestAugment_noBody: handler returned wrong status code: got %v, want %v", got, want)
	}
	if got, want := rr.Header().Get("X-Pass-Reason"), "No body"; got != want {
		t.Errorf("pass reason incorrect: got %v, want %v", got, want)
	}
}

func TestAugment_badBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/augment", strings.NewReader("not a bid request"))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(augmentHandler)
	handler.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNoContent; got != want {
		t.Errorf("TestAugment_noBody: handler returned wrong status code: got %v, want %v", got, want)
	}
	if got, want := rr.Header().Get("X-Pass-Reason"), "Protobuf unmarshal error"; got != want {
		t.Errorf("pass reason incorrect: got %v, want %v", got, want)
	}
}

func TestAugment_noSegment(t *testing.T) {
	conn, _ = redis.Dial("tcp", s.Addr())
	defer conn.Close()
	f, err := os.Open("testdata/sample_proto_nosegment")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	req := httptest.NewRequest("POST", "/augment", f)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(augmentHandler)
	handler.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusNoContent; got != want {
		t.Errorf("no segment handler returned wrong status code: got %v, want %v", got, want)
	}
	if got, want := rr.Header().Get("X-Pass-Reason"), "No segment to augment"; got != want {
		t.Errorf("no segment pass reason incorrect: got %v, want %v", got, want)
	}
}

func TestAugment_segment(t *testing.T) {
	conn, _ = redis.Dial("tcp", s.Addr())
	defer conn.Close()
	userID, wantSeg := "mid.A4B841F6-C109-41F0-9738-A366E222C2CC", "beeswax-1234"
	conn.Do("SET", userID, wantSeg)
	f, err := os.Open("testdata/sample_proto_segment")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	req := httptest.NewRequest("POST", "/augment", f)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(augmentHandler)
	handler.ServeHTTP(rr, req)

	if got, want := rr.Code, http.StatusOK; got != want {
		t.Errorf("no segment handler returned wrong status code: got %v, want %v", got, want)
	}
	if got, want := rr.Header().Get("X-Pass-Reason"), ""; got != want {
		t.Errorf("no segment pass reason incorrect: got %q, want %q", got, want)
	}

	var resp augment.AugmentorResponse
	if err := proto.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("unmarshal response error: %v", err)
	}
	segs := resp.GetSegments()
	if segs == nil || len(segs) == 0 || segs[0] == nil {
		t.Errorf("response segments is bad: %v", segs)
	}
	if gotSeg := *(segs[0]).Id; gotSeg != wantSeg {
		t.Errorf("bad response segment: got %v, want %v", gotSeg, wantSeg)
	}
}

func TestBufferPool(t *testing.T) {
	buf := bufPool.Get()
	want := new(bytes.Buffer)
	switch got := buf.(type) {
	case *bytes.Buffer:
	default:
		t.Errorf("Buffer pool type: got %T, want %T", got, want)
	}
}

func TestGetUserID_nils(t *testing.T) {
	tests := []*openrtb.BidRequest{
		&openrtb.BidRequest{},
		&openrtb.BidRequest{User: &openrtb.BidRequest_User{}},
	}
	for _, br := range tests {
		u := getUserID(br)
		if u != "" {
			t.Errorf("getUserID with a nil field returned a user ID: %v", u)
		}
	}
}

func TestGetUserID_UserID(t *testing.T) {
	tests := []string{
		"",
		"mid.283ADF1E-96A9-49DA-A5AE-B5AA211C1159",
	}
	for _, want := range tests {
		br := getBRWithUserID(want)
		got := getUserID(br)
		if got != want {
			t.Errorf("TestGetUserID: got %v, want %v", got, want)
		}
	}
}

func getBRWithUserID(userID string) *openrtb.BidRequest {
	return &openrtb.BidRequest{
		User: &openrtb.BidRequest_User{
			Ext: &openrtb.UserExtensions{
				UserId: &userID,
			},
		},
	}
}
