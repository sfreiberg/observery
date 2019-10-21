package observery

import "context"

// Outage is the entry for managing observery outages via the API.
type Outage struct {
	url    string
	client *Client
}

func newOutage(url string, c *Client) *Outage {
	return &Outage{url: url, client: c}
}

// OutageRecentResponse is the response when calling Outage.Recent.
type OutageRecentResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Outages []struct {
		ID               string `json:"id"`
		CheckID          string `json:"checkId"`
		CheckName        string `json:"checkName"`
		Ongoing          bool   `json:"ongoing"`
		Start            string `json:"start"`
		Duration         int64  `json:"duration"`
		DurationFriendly string `json:"durationFriendly"`
		Stop             string `json:"stop,omitempty"`
	} `json:"result"`
}

// OutageGetResponse is the response when calling Check.Get.
type OutageGetResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Outage  struct {
		ID               string `json:"id"`
		CheckID          string `json:"checkId"`
		CheckName        string `json:"checkName"`
		Ongoing          bool   `json:"ongoing"`
		Start            string `json:"start"`
		Stop             string `json:"stop,omitempty"`
		Duration         int64  `json:"duration"`
		DurationFriendly string `json:"durationFriendly"`
		ResponseTime     int    `json:"responseTime"`
		Details          string `json:"details"`
	} `json:"result"`
}

// Recent returns the 100 most recent outages.
func (o *Outage) Recent(ctx context.Context) (*OutageRecentResponse, error) {
	resp := &OutageRecentResponse{}
	err := o.client.get(ctx, o.url, nil, resp)
	return resp, err
}

// Get returns an invidual outage corresponding to the id.
func (o *Outage) Get(ctx context.Context, id string) (*OutageGetResponse, error) {
	resp := &OutageGetResponse{}
	err := o.client.get(ctx, o.url+"/"+id, nil, resp)
	return resp, err
}
