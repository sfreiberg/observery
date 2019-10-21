package observery

import (
	"net/http"
)

func ExampleWebhookHandler() {
	callback := func(w *Webhook, e error) {
		//Do something
	}
	http.HandleFunc("/observery", WebhookHandler(callback))
}
