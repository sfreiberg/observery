package observery

import (
	"net/http"
	"time"
)

// Webhook is a struct for holding information from observery webhooks.
// https://observery.com/account/_/integration/webhook
type Webhook struct {
	// CheckID is the unique identifier of the check.
	CheckID string `form:"checkId"`

	// CheckName is user supplied name of the check.
	CheckName string `form:"checkName"`

	// CheckType is the check type and is one of the following:
	// * http
	// * ping
	// * ssh
	// * ftp
	// * pop
	// * smtp
	// * imap
	// * cert
	CheckType string `form:"checkType"`

	// State indicates whether the check was up or dowm.
	State string `form:"state"`

	// HTTPStatusCode holds the status code if the type is http.
	HTTPStatusCode int `form:"httpStatusCode"`

	// ResponseTime holds the duration of the last response.
	ResponseTime time.Duration `form:"responseTime"`

	// TimedOut indicates whether the check timed out waiting for a response.
	TimedOut bool `form:"timedOut"`

	// Details holds any additional details.
	Details string `form:"details"`
}

// Decode takes an http.Request and decodes it directly into a Webhook.
func (w *Webhook) Decode(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Add unit so it can be parsed into a time.Duration type.
	r.Form.Set("responseTime", r.Form.Get("responseTime")+"ms")

	return decoder.Decode(w, r.Form)
}

// WebhookHandler takes a function that will be called whenever the handler
// is called by the observery.com webhook. The func f will be called in a
// goroutine so it doesn't tie up the observery caller in the event that f is a
// long running task.
func WebhookHandler(f func(*Webhook, error)) func(w http.ResponseWriter, r *http.Request) {
	var (
		hook *Webhook
		err  error
	)
	callback := func(w http.ResponseWriter, r *http.Request) {
		err = hook.Decode(r)
	}
	go f(hook, err)
	return callback
}
