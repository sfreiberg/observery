package observery

import (
	"fmt"
	"net/http"
	"time"
)

// Webhook is a struct for holding information from observery webhooks.
// https://observery.com/account/_/integration/webhook
type Webhook struct {
	// ID is the unique identifier of the check.
	ID int `form:"checkId"`

	// Name is user supplied name of the check.
	Name string `form:"checkName"`

	// Type is the check type and is one of the following:
	// * http
	// * ping
	// * ssh
	// * ftp
	// * pop
	// * smtp
	// * imap
	// * cert
	Type string `form:"checkType"`

	// Status indicates whether the check was up or dowm.
	Status string `form:"status"`

	// Code holds the status code if the type is http.
	Code int `form:"code"`

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

	// Change from yes or no to true or false so we can properly decode
	// to a bool.
	if r.Form.Get("timedOut") == "yes" {
		r.Form.Set("timedOut", "true")
	} else {
		r.Form.Set("timedOut", "false")
	}

	fmt.Printf("%+v\n", r.Form)

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
