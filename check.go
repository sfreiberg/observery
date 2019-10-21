package observery

import "context"

// Check is the entry for managing observery checks via the API.
type Check struct {
	url    string
	client *Client
}

func newCheck(url string, c *Client) *Check {
	return &Check{url: url, client: c}
}

// CheckAllResponse is the response when calling Check.All.
type CheckAllResponse struct {
	Success bool `json:"success"`
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
}

// CheckGetResponse is the response when calling Check.Get.
type CheckGetResponse struct {
	Success bool `json:"success"`
	Result  struct {
		ID                     string `json:"id"`
		Name                   string `json:"name"`
		Type                   string `json:"type"`
		State                  string `json:"state"`
		Since                  string `json:"since"`
		OutageID               string `json:"outageId"`
		URL                    string `json:"url"`
		Active                 bool   `json:"active"`
		Interval               int    `json:"interval"`
		EmailNotificationDelay int    `json:"emailNotificationDelay"`
		SmsNotificationDelay   int    `json:"smsNotificationDelay"`
		InMaintenance          bool   `json:"inMaintenance"`
		MaintenanceModeActive  bool   `json:"maintenanceModeActive"`
		MaintenanceSchedules   []struct {
			Days     string `json:"days"`
			Start    string `json:"start"`
			Stop     string `json:"stop"`
			Timezone string `json:"timezone"`
		} `json:"maintenanceSchedules"`
		ContactIds []string `json:"contactIds"`
		Contacts   []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"contacts"`
	} `json:"result"`
}

// CheckCreateRequest holds the values for creating a new check.
type CheckCreateRequest struct {
	// http, ping, ssh, ftp, pop, smtp, imap or cert
	Type string `schema:"type"`

	// name for check
	Name string `schema:"name"`

	// is check active
	Active bool `schema:"active"`

	// interval (in minutes) this check is ran
	Interval int `schema:"interval"`

	// comma-separated list of checks ids to map to this check
	Contacts string `schema:"contacts"`

	// fields for http type
	URL string `schema:"url"`

	// username for http or ftp, optional
	Username string `schema:"username"`

	// password for http or ftp, optional
	Password string `schema:"password"`

	// post data to send for http, optional
	SendData string `schema:"sendData"`

	// headers to send for http, optional
	// TODO: what format?
	HTTPHeaders string `schema:"httpHeaders"`

	// host to check, required for ping, ssh, ftp, pop, smtp, imap and cert types
	Host string `schema:"host"`

	// port to check, optional for ssh, ftp, pop, smtp, imap and cert types
	Port int `schema:"port"`

	// whether to use the secure version of the protocol for ftp, pop, smtp and imap, optional
	Secure bool `schema:"secure"`

	// number of days until cert expiration that should result in down status in cert type, required
	CertExpirationDays int `schema:"certExpirationDays"`
}

// CheckCreateResponse is the response from the API when calling Check.Create.
type CheckCreateResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Reasons []struct {
		Field string `json:"field"`
		Error string `json:"error"`
	} `json:"reasons"`
	Result struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"result"`
}

// CheckUpdateRequest holds the values for updating an existing check.
type CheckUpdateRequest struct {
	// ID of the check to be updated
	ID string

	// name for check
	Name *string `schema:"name"`

	// is check active
	Active *bool `schema:"active"`

	// interval (in minutes) this check is ran
	Interval *int `schema:"interval"`

	// comma-separated list of checks ids to map to this check
	Contacts *string `schema:"contacts"`

	// fields for http type
	URL *string `schema:"url"`

	// username for http or ftp, optional
	Username *string `schema:"username"`

	// password for http or ftp, optional
	Password *string `schema:"password"`

	// post data to send for http, optional
	SendData *string `schema:"sendData"`

	// headers to send for http, optional
	// TODO: what format?
	HTTPHeaders *string `schema:"httpHeaders"`

	// host to check, required for ping, ssh, ftp, pop, smtp, imap and cert types
	Host *string `schema:"host"`

	// port to check, optional for ssh, ftp, pop, smtp, imap and cert types
	Port *int `schema:"port"`

	// whether to use the secure version of the protocol for ftp, pop, smtp and imap, optional
	Secure *bool `schema:"secure"`

	// number of days until cert expiration that should result in down status in cert type, required
	CertExpirationDays *int `schema:"certExpirationDays"`
}

// CheckUpdateResponse holds the response from the API that is returned from Check.Update.
type CheckUpdateResponse struct {
	Success bool `json:"success"`
	Result  struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"result"`
}

// CheckDeleteResponse holds the server response when calling Check.Delete.
type CheckDeleteResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

// All returns all of the checks.
func (c *Check) All(ctx context.Context) (*CheckAllResponse, error) {
	resp := &CheckAllResponse{}
	err := c.client.get(ctx, c.url, nil, resp)
	return resp, err
}

// Get returns an invidual check corresponding to the id.
func (c *Check) Get(ctx context.Context, id string) (*CheckGetResponse, error) {
	resp := &CheckGetResponse{}
	err := c.client.get(ctx, c.url+"/"+id, nil, resp)
	return resp, err
}

// Create a new check.
func (c *Check) Create(ctx context.Context, req *CheckCreateRequest) (*CheckCreateResponse, error) {
	resp := &CheckCreateResponse{}
	err := c.client.post(ctx, c.url, req, resp)
	return resp, err
}

// Update an existing check.
func (c *Check) Update(ctx context.Context, req *CheckUpdateRequest) (*CheckUpdateResponse, error) {
	resp := &CheckUpdateResponse{}
	err := c.client.put(ctx, c.url+req.ID, req, resp)
	return resp, err
}

// Delete an existing check.
func (c *Check) Delete(ctx context.Context, id string) (*CheckDeleteResponse, error) {
	resp := &CheckDeleteResponse{}
	err := c.client.delete(ctx, c.url+"/"+id, nil, resp)
	return resp, err
}
