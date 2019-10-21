package observery

import "context"

// Contact is the entry for managing observery contacts via the API.
type Contact struct {
	url    string
	client *Client
}

func newContact(url string, c *Client) *Contact {
	return &Contact{url: url, client: c}
}

// ContactAllResponse is the response when calling Contact.All.
type ContactAllResponse struct {
	Success  bool   `json:"success"`
	Reason   string `json:"reason"`
	Contacts []struct {
		ID                string `json:"id"`
		Type              string `json:"type"`
		Name              string `json:"name"`
		Verified          bool   `json:"verified"`
		Enabled           bool   `json:"enabled"`
		Email             string `json:"email,omitempty"`
		Format            string `json:"format,omitempty"`
		CheckMappingCount int    `json:"checkMappingCount"`
		Number            string `json:"number,omitempty"`
	} `json:"result"`
}

// ContactGetResponse is the response when calling Contact.Get.
type ContactGetResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Contact struct {
		ID                string   `json:"id"`
		Type              string   `json:"type"`
		Name              string   `json:"name"`
		Verified          bool     `json:"verified"`
		Enabled           bool     `json:"enabled"`
		Email             string   `json:"email"`
		Format            string   `json:"format"`
		CheckMappingCount int      `json:"checkMappingCount"`
		CheckIds          []string `json:"checkIds"`
		Checks            []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"checks"`
	} `json:"result"`
}

// ContactCreateRequest holds the values for creating a new contact.
type ContactCreateRequest struct {
	Type    string `schema:"type"`    // 'email' or 'sms'
	Name    string `schema:"name"`    // name for contact
	Email   string `schema:"email"`   // email for contact
	Number  string `schema:"number"`  // sms number for contact
	Enabled bool   `schema:"enabled"` // is contact enabled
	Format  string `schema:"format"`  // 'short' or 'long' - only when type = email
	Checks  string `schema:"checks"`  // comma-separated list of check ids to map to contact
}

// ContactCreateResponse is the response from the API when calling Contact.Create.
type ContactCreateResponse struct {
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

// ContactUpdateRequest holds the values for updating an existing contact.
type ContactUpdateRequest struct {
	ID      string  `schema:"-"`
	Name    *string `schema:"name"`
	Enabled *bool   `schema:"enabled"`
	Format  *string `schema:"format"`
	Checks  *string `schema:"checks"`
}

// ContactUpdateResponse holds the response from the API that is returned from Contact.Update.
type ContactUpdateResponse struct {
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

// ContactDeleteResponse holds the server response when calling Contact.Delete.
type ContactDeleteResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

// All returns all of the contacts.
func (c *Contact) All(ctx context.Context) (*ContactAllResponse, error) {
	resp := &ContactAllResponse{}
	err := c.client.get(ctx, c.url, nil, resp)
	return resp, err
}

// Get returns an invidual contact corresponding to the id.
func (c *Contact) Get(ctx context.Context, id string) (*ContactGetResponse, error) {
	resp := &ContactGetResponse{}
	err := c.client.get(ctx, c.url+"/"+id, nil, resp)
	return resp, err
}

// Create a new contact.
func (c *Contact) Create(ctx context.Context, req *ContactCreateRequest) (*ContactCreateResponse, error) {
	resp := &ContactCreateResponse{}
	err := c.client.post(ctx, c.url, req, resp)
	return resp, err
}

// Update an existing contact.
func (c *Contact) Update(ctx context.Context, req *ContactUpdateRequest) (*ContactUpdateResponse, error) {
	resp := &ContactUpdateResponse{}
	err := c.client.put(ctx, c.url+"/"+req.ID, req, resp)
	return resp, err
}

// Delete an existing contact.
func (c *Contact) Delete(ctx context.Context, id string) (*ContactDeleteResponse, error) {
	resp := &ContactDeleteResponse{}
	err := c.client.put(ctx, c.url+"/"+id, nil, resp)
	return resp, err
}
