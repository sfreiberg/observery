package observery

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
)

const (
	api = "https://api.observery.com/api/v1"
)

// Client is the main entry point into the observery API and its endpoints.
type Client struct {
	username string
	password string
	client   *http.Client
	Contact  *Contact
	Outage   *Outage
	Check    *Check
}

// NewClient creates a new client with appropriate API keys.
func NewClient(username, password string) *Client {
	c := &Client{
		username: username,
		password: password,
		client:   &http.Client{},
	}
	c.Contact = newContact(api+"/contact", c)
	c.Outage = newOutage(api+"/outage", c)
	c.Check = newCheck(api+"/check", c)
	return c
}

func (c *Client) get(ctx context.Context, url string, input, output interface{}) error {
	return c.exec(ctx, url, "GET", input, output)
}

func (c *Client) post(ctx context.Context, u string, input, output interface{}) error {
	return c.exec(ctx, u, "POST", input, output)
}

func (c *Client) put(ctx context.Context, url string, input, output interface{}) error {
	return c.exec(ctx, url, "PUT", input, output)
}

func (c *Client) delete(ctx context.Context, url string, input, output interface{}) error {
	return c.exec(ctx, url, "DELETE", input, output)
}

func (c *Client) exec(ctx context.Context, u, method string, input, output interface{}) error {
	var body io.Reader
	if input != nil {
		dst := url.Values{}
		e := schema.NewEncoder()
		e.Encode(input, dst)
		body = strings.NewReader(dst.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return err
	}

	if input != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.SetBasicAuth(c.username, c.password)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(output)
}
