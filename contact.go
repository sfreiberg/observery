package observery

import "context"

// ListContactsResponse is the response when calling Client.ListContacts.
type ListContactsResponse struct {
	// Success returns false if there was an error. A failure message will be
	// stored in Reason.
	Success bool `json:"success"`

	// Reason will be populated if Success is false.
	Reason string `json:"reason"`

	// Contacts holds
	Contacts []struct {
		// ID of the contact.
		ID string `json:"id"`

		// Type will be 'email' or 'sms'.
		Type string `json:"type"`

		// Name is the friendly name of the contact.
		Name string `json:"name"`

		// Verified returns true if the contact has been verified.
		Verified bool `json:"verified"`

		// Enabled returns true if the contact is enabled. If true the contact
		// will receive updates.
		Enabled bool `json:"enabled"`

		// Email holds the email address of the contact if Type is 'email'.
		Email *string `json:"email,omitempty"`

		// Format will be either 'short' or 'long'. Only applicable to Type
		// 'email'.
		Format *string `json:"format,omitempty"`

		// CheckMappingCount is the number of checks mapped to this Contact.
		CheckMappingCount int `json:"checkMappingCount"`

		// Number is the telephone number used for 'sms' messages.
		Number *string `json:"number,omitempty"`
	} `json:"result"`
}

// GetContactResponse is the response when calling Client.GetContact.
type GetContactResponse struct {
	// Success returns false if there was an error. A failure message will be
	// stored in Reason.
	Success bool `json:"success"`

	// Reason will be populated if Success is false.
	Reason string `json:"reason"`

	// Contact
	Contact struct {
		// ID of the contact.
		ID string `json:"id"`

		// Type of contact. Will be 'email' or 'sms'.
		Type string `json:"type"`

		// Name of the contact.
		Name string `json:"name"`

		// Verified is true if the contact has been verified.
		Verified bool `json:"verified"`

		// Enabled returns true if the contact is enabled. If true the contact
		// will receive updates.
		Enabled bool `json:"enabled"`

		// Email holds the email address of the contact if Type is 'email'.
		Email *string `json:"email"`

		// Format will be either 'short' or 'long'. Only applicable to Type
		// 'email'.
		Format *string `json:"format"`

		// Number is the telephone number used for 'sms' messages.
		Number *string `json:"number"`

		// CheckMappingCount is the number of checks mapped to this Contact.
		CheckMappingCount int `json:"checkMappingCount"`

		// Checks contains all of the checks mapped to this Contact.
		Checks []struct {

			// ID of the Check.
			ID string `json:"id"`

			// Name of the check.
			Name string `json:"name"`

			// Type of check, 'email' or 'sms'
			Type string `json:"type"`
		} `json:"checks"`
	} `json:"result"`
}

// CreateContactRequest holds the values for creating a new contact.
type CreateContactRequest struct {
	// Type holds the type of contact. Must be email or sms.
	Type string `form:"type"`

	// Name for the contact.
	Name string `form:"name"`

	// Email address for the contact. Required for type email.
	Email string `form:"email"`

	// Number is the phone number for sending SMS messages. Must be in the
	// following format: +{country code}{phone number}. For example:
	//   +18885551234
	Number string `form:"number"`

	// Enabled determines whether the contact should receive messages.
	Enabled bool `form:"enabled"`

	// Format specifies what size of message the contact should receive.
	// Only applies to type email and must be either 'short' or 'long'.
	Format string `form:"format"`

	// Checks is a comma separated list of check ids that should be
	// contacted when the check changes state.
	Checks string `form:"checks"`
}

// CreateContactResponse is the response from the API when calling Client.CreateContact.
type CreateContactResponse struct {
	// Success returns false if the Contact couldn't be created.
	Success bool `json:"success"`

	// Reason will hold a message about why the Contact couldn't be created.
	Reason string `json:"reason"`

	// Reasons is a slice of error messages.
	Reasons []struct {
		// Field contains the name of the field that contained an error.
		Field string `json:"field"`

		// Error is the message that explains why the field was invalid.
		Error string `json:"error"`
	} `json:"reasons"`

	// Result will be populated when the contact was successfully created.
	Result *struct {
		// ID of the newly created contact.
		ID string `json:"id"`

		// Message holds the success string.
		Message string `json:"message"`
	} `json:"result"`
}

// UpdateContactRequest holds the values for updating an existing contact.
// Only updated when populated.
type UpdateContactRequest struct {
	// ID of the Contact to update.
	ID string `form:"-"`

	// Name of the contact.
	Name *string `form:"name"`

	// Enabled determines whether the contact should receive messages.
	Enabled *bool `form:"enabled"`

	// Format specifies what size of message the contact should receive.
	// Only applies to type email and must be either 'short' or 'long'.
	Format *string `form:"format"`

	// Checks is a comma separated list of check ids that should be
	// contacted when the check changes state.
	Checks *string `form:"checks"`
}

// UpdateContactResponse holds the response from the API that is returned from Contact.Update.
type UpdateContactResponse struct {
	// Success returns false if the Contact couldn't be updated.
	Success bool `json:"success"`

	// Reason will hold a message about why the Contact couldn't be updated.
	Reason string `json:"reason"`

	// Reasons is a slice of error messages.
	Reasons []struct {
		// Field contains the name of the field that contained an error.
		Field string `json:"field"`

		// Error is the message that explains why the field was invalid.
		Error string `json:"error"`
	} `json:"reasons"`

	// Result will be populated when the contact was successfully updated.
	Result *struct {
		// ID of the updated contact.
		ID string `json:"id"`

		// Message holds the success string.
		Message string `json:"message"`
	} `json:"result"`
}

// DeleteContactResponse holds the server response when calling Contact.Delete.
type DeleteContactResponse struct {
	// Success returns false if the contact couldn't be deleted. The reason
	// will be held in DeleteContactResponse.Result.
	Success bool `json:"success"`

	// Result holds a message explaining the DeleteContactResponse.Success
	// response.
	Result string `json:"result"`
}

// ListContacts returns all of the contacts.
func (c *Client) ListContacts(ctx context.Context) (*ListContactsResponse, error) {
	url := api + "/contact"
	resp := &ListContactsResponse{}
	err := c.get(ctx, url, nil, resp)
	return resp, err
}

// GetContact returns an invidual contact corresponding to the id.
func (c *Client) GetContact(ctx context.Context, id string) (*GetContactResponse, error) {
	url := api + "/contact/" + id
	resp := &GetContactResponse{}
	err := c.get(ctx, url, nil, resp)
	return resp, err
}

// CreateContact creates a new contact. New contacts must be verified in
// the front-end.
func (c *Client) CreateContact(ctx context.Context, req *CreateContactRequest) (*CreateContactResponse, error) {
	url := api + "/contact"
	resp := &CreateContactResponse{}
	err := c.post(ctx, url, req, resp)
	return resp, err
}

// UpdateContact updates an existing contact.
func (c *Client) UpdateContact(ctx context.Context, req *UpdateContactRequest) (*UpdateContactResponse, error) {
	url := api + "/contact/" + req.ID
	resp := &UpdateContactResponse{}
	err := c.put(ctx, url, req, resp)
	return resp, err
}

// DeleteContact deletes an existing contact.
func (c *Client) DeleteContact(ctx context.Context, id string) (*DeleteContactResponse, error) {
	url := api + "/contact/" + id
	resp := &DeleteContactResponse{}
	err := c.delete(ctx, url, nil, resp)
	return resp, err
}
