package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(time.Second * 5)

	tests := []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get success", http.MethodGet, "car", http.StatusOK, nil},
		{"create success", http.MethodPost, "car", http.StatusCreated, []byte(`{"make":"Toyota","model":"Camry","year":2022,"owner":"John Doe","in_garage":true}`)},
		{"get unknown endpoint", http.MethodGet, "unknown", http.StatusNotFound, nil},
		{"get invalid endpoint", http.MethodGet, "car/id", http.StatusNotFound, nil},
		{"unregistered route", http.MethodPut, "car", http.StatusMethodNotAllowed, nil},
		{"delete success", http.MethodDelete, "car/id", http.StatusNoContent, nil},
	}

	for i, tc := range tests {
		req, _ := request.NewMock(tc.method, "http://localhost:8097/"+tc.endpoint, bytes.NewBuffer(tc.body))
		c := http.Client{}

		resp, err := c.Do(req)
		if err != nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue 
		}

		if resp.StatusCode != tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}
}
