package observery

import (
	"context"
	"time"
)

// ListChecksResponse is the response when calling Client.ListChecks.
type ListChecksResponse struct {
	// Success will be false in the event of a failure.
	Success bool

	// Reason will contain a message about why the request failed.
	Reason string

	// Checks is a list of all checks.
	Checks []struct {
		// ID of the check.
		ID string

		// Name of the check.
		Name string

		// Active
		Active bool

		// Type will be one of: http, ping, ssh, ftp, pop, smtp, imap or cert.
		Type string

		// State is the current state of the check. Possible states are:
		// up, down or waiting.
		State string

		// Since holds the time of the last state change.
		Since time.Time

		// URL is the url to check for type http.
		URL string

		// Host holds the host for ping, ssh, ftp, pop, smtp, imap and cert.
		Host string
	}
}

// GetCheckResponse is the response when calling Client.GetCheck.
type GetCheckResponse struct {
	// Success will be false in the event of a failure.
	Success bool `json:"success"`

	// Reason will contain a message about why the request failed.
	Reason string `json:"Reason"`

	// Check hold the requested check.
	Check struct {
		// ID of the check.
		ID string `json:"id"`

		// Name of the check.
		Name string `json:"name"`

		// Type will be one of: http, ping, ssh, ftp, pop, smtp, imap or cert.
		Type string `json:"type"`

		// State is the current state of the check. Possible states are:
		// up, down or waiting.
		State string `json:"state"`

		// Since holds the time of the last state change.
		Since string `json:"since"`

		// OutageID is the outage id if the check is currently down.
		OutageID *string `json:"outageId"`

		// URL to check if Check.Type is 'http'.
		URL *string `json:"url"`

		// Active is true when the check is being execute.
		Active bool `json:"active"`

		// Interval is how often the check in minutes.
		Interval int `json:"interval"`

		// EmailNotificationDelay is how long in minutes observery will wait to
		// notify of an outage.
		EmailNotificationDelay int `json:"emailNotificationDelay"`

		// SmsNotificationDelay is how long in minutes observery will wait to
		// notify of an outage.
		SmsNotificationDelay int `json:"smsNotificationDelay"`

		// InMaintenance returns true if the check is currently in
		// a maintenance window.
		InMaintenance bool `json:"inMaintenance"`

		// MaintenanceModeActive return true if maintenance mode is currently
		// active.
		MaintenanceModeActive bool `json:"maintenanceModeActive"`

		// MaintenanceSchedules configured for this check.
		MaintenanceSchedules []struct {
			// Days is a comma-seperated list of days.
			Days string `json:"days"`

			// Start is when the maintenance schedule starts.
			Start time.Time `json:"start"`

			// Stop is when the maintenance schedule ends.
			Stop time.Time `json:"stop"`

			// Timezone for this maintenance schedule in continent/city format.
			Timezone string `json:"timezone"`
		} `json:"maintenanceSchedules"`

		// Contacts that are mapped to this is check.
		Contacts []struct {
			// ID of the contact.
			ID string `json:"id"`

			// Name of the contact.
			Name string `json:"name"`
		} `json:"contacts"`
	} `json:"result"`
}

// CreateCheckRequest holds the values for creating a new check.
type CreateCheckRequest struct {
	// Type is the type of check to create. Will be one of:
	// http, ping, ssh, ftp, pop, smtp, imap or cert.
	Type string `form:"type"`

	// Name of the check.
	Name string `form:"name"`

	// Active is true if the check is active.
	Active bool `form:"active"`

	// Interval (in minutes) this check should run.
	Interval int `form:"interval"`

	// Contacts is a comma-separated list of checks ids to map to this check.
	Contacts string `form:"contacts"`

	// URL holds the URL to check for Type 'http'. Mandatory for 'http' checks.
	URL *string `form:"url"`

	// Username for http or ftp, optional.
	Username *string `form:"username"`

	// Password for http or ftp, optional
	Password *string `form:"password"`

	// SendData is the data to send for an 'http' check, optional.
	SendData *string `form:"sendData"`

	// HTTPHeaders is an optional field for 'http' checks.
	// Headers need to be formatted as 'key: value'. One heaer per line.
	HTTPHeaders *string `form:"httpHeaders"`

	// Host to check, required for ping, ssh, ftp, pop, smtp, imap and cert types.
	Host *string `form:"host"`

	// Port to check, optional for ssh, ftp, pop, smtp, imap and cert types.
	Port *int `form:"port"`

	// Secure tells the check to use the secure version of the protocol for:
	// ftp, pop, smtp and imap, optional.
	Secure *bool `form:"secure"`

	// CertExpirationDays is the number of days until cert expiration that
	// should result in down status in cert type, required.
	CertExpirationDays *int `form:"certExpirationDays"`
}

// CreateCheckResponse is the response from the API when calling Client.CreateCheck.
type CreateCheckResponse struct {
	// Success returns true if the update was successful, false otherwise.
	Success bool `json:"success"`

	// Reason is a human readable message about the status of the request.
	Reason string `json:"reason"`

	// Reasons is a slice of Field/Error messages that explain why the create
	// was unsuccessful.
	Reasons []struct {
		// Field is the name of the field that was incorrect.
		Field string `json:"field"`

		// Error is the actual error message that caused the failure.
		Error string `json:"error"`
	} `json:"reasons"`

	// Result contains information about the request.
	Result struct {
		// ID of the newly created check.
		ID string `json:"id"`

		// Message from the server about the success or failure of the request.
		Message string `json:"message"`
	} `json:"result"`
}

// UpdateCheckRequest holds the values for updating an existing check.
type UpdateCheckRequest struct {
	// ID of the check to be updated
	ID string

	// Name of the check.
	Name *string `form:"name"`

	// is check active.
	Active *bool `form:"active"`

	// Interval (in minutes) this check is ran.
	Interval *int `form:"interval"`

	// Contacts comma-separated list of checks ids to map to this check.
	Contacts *string `form:"contacts"`

	// fields for http type

	// URL of website to check.
	URL *string `form:"url"`

	// Username for http or ftp, optional
	Username *string `form:"username"`

	// Password for http or ftp, optional
	Password *string `form:"password"`

	// SendData is post data to send for http, optional
	SendData *string `form:"sendData"`

	// HTTPHeaders to send for http, optional.
	// Headers need to be formatted as 'key: value'. One heaer per line.
	HTTPHeaders *string `form:"httpHeaders"`

	// Host to check, required for ping, ssh, ftp, pop, smtp, imap and cert
	// types.
	Host *string `form:"host"`

	// Port to check, optional for ssh, ftp, pop, smtp, imap and cert types
	Port *int `form:"port"`

	// whether to use the secure version of the protocol for ftp, pop, smtp
	// and imap, optional.
	Secure *bool `form:"secure"`

	// CertExpirationDays is the number of days until cert expiration that
	// should result in down status in cert type, required.
	CertExpirationDays *int `form:"certExpirationDays"`
}

// UpdateCheckResponse holds the response from the API that is returned from
// Client.UpdateCheck.
type UpdateCheckResponse struct {
	// Success returns true if the update was successful, false otherwise.
	Success bool `json:"success"`

	// Result contains information about the update.
	Result struct {
		// ID of the check that was updated.
		ID string `json:"id"`

		// Message from the server about the success or failure of the update.
		Message string `json:"message"`
	} `json:"result"`
}

// DeleteCheckResponse holds the server response when calling Client.DeleteCheck.
type DeleteCheckResponse struct {
	// Success returns true if the update was successful, false otherwise.
	Success bool `json:"success"`

	// Result is a message from the server about the requested action.
	Result string `json:"result"`
}

// ListChecks returns all of the checks.
func (c *Client) ListChecks(ctx context.Context) (*ListChecksResponse, error) {
	url := api + "/check"
	s := &struct {
		Success bool   `json:"success"`
		Reason  string `json:"reason"`
		Checks  []struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Active bool   `json:"active"`
			Type   string `json:"type"`
			State  string `json:"state"`
			Since  string `json:"since,omitempty"`
			URL    string `json:"url,omitempty"`
			Host   string `json:"host,omitempty"`
		} `json:"result"`
	}{}

	type Check struct {
		ID     string
		Name   string
		Active bool
		Type   string
		State  string
		Since  time.Time
		URL    string
		Host   string
	}

	err := c.get(ctx, url, nil, s)

	resp := &ListChecksResponse{Success: s.Success, Reason: s.Reason}
	for _, check := range s.Checks {
		newCheck := Check{
			ID:     check.ID,
			Name:   check.Name,
			Active: check.Active,
			Type:   check.Type,
			State:  check.State,
			URL:    check.URL,
			Host:   check.Host,
		}
		if check.Since != "" {
			since, err := time.Parse("2006-01-02T15:04:05", check.Since)
			if err != nil {
				return nil, err
			}
			newCheck.Since = since
		}
		resp.Checks = append(resp.Checks, newCheck)
	}

	return resp, err
}

// GetCheck returns an invidual check corresponding to the id.
func (c *Client) GetCheck(ctx context.Context, id string) (*GetCheckResponse, error) {
	url := api + "/check/" + id
	resp := &GetCheckResponse{}
	err := c.get(ctx, url, nil, resp)
	return resp, err
}

// CreateCheck a new check.
func (c *Client) CreateCheck(ctx context.Context, req *CreateCheckRequest) (*CreateCheckResponse, error) {
	url := api + "/check"
	resp := &CreateCheckResponse{}
	err := c.post(ctx, url, req, resp)
	return resp, err
}

// UpdateCheck an existing check.
func (c *Client) UpdateCheck(ctx context.Context, req *UpdateCheckRequest) (*UpdateCheckResponse, error) {
	url := api + "/check/" + req.ID
	resp := &UpdateCheckResponse{}
	err := c.put(ctx, url, req, resp)
	return resp, err
}

// DeleteCheck deletes an existing check.
func (c *Client) DeleteCheck(ctx context.Context, id string) (*DeleteCheckResponse, error) {
	url := api + "/check/" + id
	resp := &DeleteCheckResponse{}
	err := c.delete(ctx, url, nil, resp)
	return resp, err
}
