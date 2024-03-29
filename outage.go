package observery

import (
	"context"
	"time"
)

// ListOutagesResponse contains the server response when requesting multiple
// outages.
type ListOutagesResponse struct {
	// Success returns true if the request was successful and false otherwise.
	Success bool

	// Result is a message from the server about the requested action.
	Reason string

	// Outage is a slice of outages.
	Outages []struct {
		// ID of the outage.
		ID string

		// CheckID of the check that the outage belongs to.
		CheckID string

		// CheckName is the friendly name of the check that the outage belongs
		// to.
		CheckName string

		// Ongoing returns true it the outage is ongoing.
		Ongoing bool

		// Start of when the outage began.
		Start time.Time

		// Stop is the date/time when the outage concluded.
		Stop time.Time

		// Duration of the outage.
		Duration time.Duration
	}
}

// GetOutageResponse contains the server response when requesting an individual
// outage.
type GetOutageResponse struct {
	// Success returns true if the request was successful and false otherwise.
	Success bool

	// Result is a message from the server about the requested action.
	Reason string

	// Outage contains the information about the requested outage.
	Outage struct {
		// ID of the outage.
		ID string

		// CheckID of the check that the outage belongs to.
		CheckID string

		// CheckName is the friendly name of the check that the outage belongs
		// to.
		CheckName string

		// Ongoing returns true it the outage is ongoing.
		Ongoing bool

		// Start of when the outage began.
		Start time.Time

		// Stop is the date/time when the outage concluded.
		Stop time.Time

		// Duration of the outage.
		Duration time.Duration

		// ResponseTime is how long it took for end point to respond when it
		// came back online.
		ResponseTime time.Duration

		// Details is a human readable message about what caused the outage.
		Details string
	}
}

// ListOutages returns the 100 most recent outages.
func (c *Client) ListOutages(ctx context.Context) (*ListOutagesResponse, error) {
	// This is used internally because the response contains an anonymous
	// struct. I could have pulled this but getting all outages has different
	// fields than getting a single outage. So more code here but better
	// and more obvious user experience.
	type Outage struct {
		ID        string
		CheckID   string
		CheckName string
		Ongoing   bool
		Start     time.Time
		Stop      time.Time
		Duration  time.Duration
	}

	s := &struct {
		Success bool   `json:"success"`
		Reason  string `json:"reason"`
		Outages []struct {
			ID        string `json:"id"`
			CheckID   string `json:"checkId"`
			CheckName string `json:"checkName"`
			Ongoing   bool   `json:"ongoing"`
			Start     string `json:"start"`
			Stop      string `json:"stop,omitempty"`
			Duration  int    `json:"duration"`
		} `json:"result"`
	}{}
	url := api + "/outage"
	if err := c.get(ctx, url, nil, s); err != nil {
		return nil, err
	}

	resp := &ListOutagesResponse{
		Success: s.Success,
		Reason:  s.Reason,
	}

	for _, o := range s.Outages {
		outage := Outage{
			ID:        o.ID,
			CheckID:   o.CheckID,
			CheckName: o.CheckName,
			Ongoing:   o.Ongoing,
			Duration:  time.Duration(o.Duration) * time.Millisecond,
		}

		start, err := time.Parse("2006-01-02T15:04:05", o.Start)
		if err != nil {
			return nil, err
		}
		outage.Start = start

		if o.Stop != "" {
			stop, err := time.Parse("2006-01-02T15:04:05", o.Stop)
			if err != nil {
				return nil, err
			}
			outage.Stop = stop
		}

		resp.Outages = append(resp.Outages, outage)
	}

	return resp, nil
}

// GetOutage returns an invidual outage corresponding to the id.
func (c *Client) GetOutage(ctx context.Context, id string) (*GetOutageResponse, error) {
	s := &struct {
		Success bool   `json:"success"`
		Reason  string `json:"reason"`
		Outage  struct {
			ID           string `json:"id"`
			CheckID      string `json:"checkId"`
			CheckName    string `json:"checkName"`
			Ongoing      bool   `json:"ongoing"`
			Start        string `json:"start"`
			Stop         string `json:"stop,omitempty"`
			Duration     int    `json:"duration"`
			ResponseTime int    `json:"responseTime"`
			Details      string `json:"details"`
		} `json:"result"`
	}{}
	url := api + "/outage/" + id
	if err := c.get(ctx, url, nil, s); err != nil {
		return nil, err
	}

	resp := &GetOutageResponse{
		Success: s.Success,
		Reason:  s.Reason,
	}

	resp.Outage.ID = s.Outage.ID
	resp.Outage.CheckID = s.Outage.CheckID
	resp.Outage.CheckName = s.Outage.CheckName
	resp.Outage.Ongoing = s.Outage.Ongoing
	resp.Outage.Duration = time.Duration(s.Outage.Duration) * time.Millisecond
	resp.Outage.ResponseTime = time.Duration(s.Outage.Duration) * time.Millisecond
	resp.Outage.Details = s.Outage.Details

	start, err := time.Parse("2006-01-02T15:04:05", s.Outage.Start)
	if err != nil {
		return nil, err
	}
	resp.Outage.Start = start

	if s.Outage.Stop != "" {
		stop, err := time.Parse("2006-01-02T15:04:05", s.Outage.Stop)
		if err != nil {
			return nil, err
		}
		resp.Outage.Stop = stop
	}

	return resp, nil
}
