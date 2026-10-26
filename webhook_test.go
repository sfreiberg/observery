package observery

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func newWebhookRequest(form url.Values) *http.Request {
	r := httptest.NewRequest("POST", "/observery", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r
}

func TestWebhookDecode(t *testing.T) {
	r := newWebhookRequest(url.Values{
		"checkId":        {"abc"},
		"checkName":      {"Homepage"},
		"checkType":      {"http"},
		"state":          {"down"},
		"httpStatusCode": {"502"},
		"responseTime":   {"250"},
		"timedOut":       {"true"},
		"details":        {"Bad Gateway"},
	})

	var hook Webhook
	if err := hook.Decode(r); err != nil {
		t.Fatal(err)
	}

	want := Webhook{
		CheckID:        "abc",
		CheckName:      "Homepage",
		CheckType:      "http",
		State:          "down",
		HTTPStatusCode: 502,
		ResponseTime:   250 * time.Millisecond,
		TimedOut:       true,
		Details:        "Bad Gateway",
	}
	if hook != want {
		t.Errorf("Decode() = %+v, want %+v", hook, want)
	}
}

func TestWebhookDecodeWithoutResponseTime(t *testing.T) {
	r := newWebhookRequest(url.Values{"checkId": {"abc"}, "state": {"up"}})

	var hook Webhook
	if err := hook.Decode(r); err != nil {
		t.Fatalf("Decode() without responseTime: %v", err)
	}
	if hook.CheckID != "abc" || hook.ResponseTime != 0 {
		t.Errorf("Decode() = %+v", hook)
	}
}

func TestWebhookHandler(t *testing.T) {
	type call struct {
		hook *Webhook
		err  error
	}
	calls := make(chan call, 2)
	handler := WebhookHandler(func(w *Webhook, err error) {
		calls <- call{w, err}
	})

	for _, id := range []string{"first", "second"} {
		rec := httptest.NewRecorder()
		handler(rec, newWebhookRequest(url.Values{"checkId": {id}, "responseTime": {"10"}}))
		if rec.Code != http.StatusOK {
			t.Errorf("status = %d, want 200", rec.Code)
		}

		select {
		case c := <-calls:
			if c.err != nil {
				t.Fatalf("callback error: %v", c.err)
			}
			if c.hook == nil || c.hook.CheckID != id {
				t.Fatalf("callback got %+v, want CheckID %q", c.hook, id)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("callback was not called for the request")
		}
	}
}
