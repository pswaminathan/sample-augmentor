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
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoggingHandler(t *testing.T) {
	out := bytes.NewBuffer(nil)
	logger := log.New(&logWriter{out}, " | ", 0)
	handler := http.HandlerFunc(noOpHandler)
	finalHandler := loggingHandler(logger, handler)
	rr, req := httptest.NewRecorder(), httptest.NewRequest("GET", "/augment", nil)
	finalHandler.ServeHTTP(rr, req)
	logOutput := out.Bytes()
	shouldContain := []string{
		" | 200 | ",
		" | GET",
		"/augment |",
	}
	for _, want := range shouldContain {
		if !bytes.Contains(logOutput, []byte(want)) {
			t.Errorf("bad log line. got %v, want %v contained", string(logOutput), want)
		}
	}
}

func TestLoggingResponseWriter(t *testing.T) {
	rr := httptest.NewRecorder()
	lrw := &loggingResponseWriter{ResponseWriter: rr}
	want := http.StatusOK
	lrw.WriteHeader(want)
	if got := lrw.statusCode; got != want {
		t.Errorf("bad stored status code: got %v, want %v", got, want)
	}
	if got := rr.Result().StatusCode; got != want {
		t.Errorf("bad returned status code: got %v, want %v", got, want)
	}
}

// This test can in theory fail due to a race condition.
// We grab time.Now() and compare it to the time.Now() calculation in
// w.Write(). These two are run as close together as possible at the moment,
// so this will only trigger in rare situations where they cross a second
// boundary. The right way to do it here is to mock the call to time.Now()
// to return a consistent time. We can't just compare to a more granular boundary
// like minute, because there still is a race condition where the two calls
// could cross that boundary as well. For now, just take that into consideration
// and if the test fails, run it again. It should be a rare occurrence.
func TestLogWriter(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	w := &logWriter{buf}
	msg := []byte("test")
	now := time.Now().Format(time.RFC3339)
	n, err := w.Write(msg)
	want := append([]byte(now), msg...)
	if err != nil {
		t.Errorf("log write error: %v", err)
	}
	if want := len(now) + len(msg); n != want {
		t.Errorf("incorrect length: got %v, want %v", n, want)
	}
	if got := buf.Bytes(); !bytes.Equal(got, want) {
		t.Errorf("log write bytes: got %v, want %v", got, want)
	}
}

func TestLogWriter_error(t *testing.T) {
	w := &logWriter{errorWriter{}}
	msg := []byte("test")
	n, err := w.Write(msg)
	if err == nil {
		t.Error("want error, got none")
	}
	if got, want := n, 0; got != want {
		t.Errorf("error bad bytes: got %v, want %v", got, want)
	}
}

type errorWriter struct{}

func (w errorWriter) Write(p []byte) (n int, err error) {
	err = errors.New("oops")
	return
}

func noOpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
