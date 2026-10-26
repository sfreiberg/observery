package observery

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

// roundTripFunc lets tests answer API requests without the network.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func newStubClient(body string) *Client {
	c := NewClient("user", "pass")
	c.client.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{"Content-Type": {"application/json"}},
			Body:       io.NopCloser(bytes.NewBufferString(body)),
			Request:    r,
		}, nil
	})
	return c
}

func TestGetOutage(t *testing.T) {
	c := newStubClient(`{
		"success": true,
		"result": {
			"id": "o1",
			"checkId": "c1",
			"checkName": "Homepage",
			"ongoing": false,
			"start": "2024-01-02T03:04:05",
			"stop": "2024-01-02T03:14:05",
			"duration": 600000,
			"responseTime": 1500,
			"details": "timeout"
		}
	}`)

	resp, err := c.GetOutage(context.Background(), "o1")
	if err != nil {
		t.Fatal(err)
	}

	o := resp.Outage
	if o.Duration != 10*time.Minute {
		t.Errorf("Duration = %v, want 10m", o.Duration)
	}
	if o.ResponseTime != 1500*time.Millisecond {
		t.Errorf("ResponseTime = %v, want 1.5s", o.ResponseTime)
	}
	if want := time.Date(2024, 1, 2, 3, 14, 5, 0, time.UTC); !o.Stop.Equal(want) {
		t.Errorf("Stop = %v, want %v", o.Stop, want)
	}
	if o.ID != "o1" || o.CheckID != "c1" || o.Details != "timeout" {
		t.Errorf("Outage = %+v", o)
	}
}
