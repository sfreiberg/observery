package observery

import (
	"context"
	"os"
	"testing"
)

func newClient(t *testing.T) *Client {
	var (
		username = os.Getenv("OBSERVERY_USERNAME")
		password = os.Getenv("OBSERVERY_PASSWORD")
	)

	if username == "" || password == "" {
		t.Fatalf("You must set the OBSERVERY_USERNAME and OBSERVERY_PASSWORD environment variables")
	}

	return NewClient(username, password)
}

func TestOutages(t *testing.T) {
	var (
		ctx    = context.Background()
		client = newClient(t)
	)

	if resp, err := client.ListOutages(ctx); err != nil {
		t.Fatalf("Error getting outages: %s\n", err)
	} else if !resp.Success {
		t.Fatalf("Unable to get outages: %s\n", resp.Reason)
	}
}

func TestCheck(t *testing.T) {
	var (
		ctx    = context.Background()
		client = newClient(t)
	)

	// Create a check
	createCheckReq := &CreateCheckRequest{
		Type:     "http",
		Name:     "Test Check",
		Active:   true,
		Interval: 1,
		URL:      PtrString("http://example.com"),
	}
	createCheckResp, err := client.CreateCheck(ctx, createCheckReq)
	if err != nil {
		t.Fatalf("Error creating check: %s\n", err)
	}
	if !createCheckResp.Success {
		for _, reason := range createCheckResp.Reasons {
			t.Errorf("%s: %s\n", reason.Field, reason.Error)
		}
		t.Fatalf(
			"Unable to create check: %+v\n",
			createCheckResp.Reason,
		)
	}

	listChecksResp, err := client.ListChecks(ctx)
	if err != nil {
		t.Fatalf("Error getting checks: %s\n", err)
	}
	if !listChecksResp.Success {
		t.Fatalf("Unable to get checks: %s\n", listChecksResp.Reason)
	}
	if len(listChecksResp.Checks) == 0 {
		t.Fatal("Couldn't find any checks")
	}

	getCheckResp, err := client.GetCheck(ctx, createCheckResp.Result.ID)
	if err != nil {
		t.Fatalf("Error getting check: %s\n", err)
	}
	if !getCheckResp.Success {
		t.Fatalf("Unable to get check: %s\n", getCheckResp.Reason)
	}

	updateCheckReq := &UpdateCheckRequest{
		ID:       createCheckResp.Result.ID,
		Name:     PtrString("Check #2"),
		Active:   PtrBool(false),
		Interval: PtrInt(2),
		URL:      PtrString("http://www.example.com"),
	}
	updateCheckResp, err := client.UpdateCheck(ctx, updateCheckReq)
	if err != nil {
		t.Fatalf("Error updating check: %s\n", err)
	}
	if !updateCheckResp.Success {
		t.Fatalf("Unable to update check: %s\n", updateCheckResp.Result.Message)
	}

	deleteCheckResp, err := client.DeleteCheck(ctx, createCheckResp.Result.ID)
	if err != nil {
		t.Fatalf("Error deleting check: %s\n", err)
	}
	if !deleteCheckResp.Success {
		t.Fatalf("Unable to delete check: %s\n", deleteCheckResp.Result)
	}
}

func TestContact(t *testing.T) {
	var (
		ctx    = context.Background()
		client = newClient(t)
	)

	// Create a contact
	contact := &CreateContactRequest{
		Type:    "email",
		Name:    "Test Email",
		Email:   "me@example.com",
		Enabled: true,
		Format:  "long",
	}
	ccr, err := client.CreateContact(ctx, contact)
	if err != nil {
		t.Fatalf("Unable to create contact: %s\n", err)
	}
	if !ccr.Success {
		t.Fatalf("Unable to create contact: %+v\n", ccr)
	}

	// Verify we got at least one contact
	contacts, err := client.ListContacts(ctx)
	if err != nil {
		t.Fatalf("Unable to get contacts: %s\n", err)
	}
	if len(contacts.Contacts) == 0 {
		t.Fatal("Expected at least one contact but found none.\n")
	}

	gcr, err := client.GetContact(ctx, ccr.Result.ID)
	if err != nil {
		t.Fatalf("Unable to get contact: %s\n", err)
	}
	if !gcr.Success {
		t.Fatalf("Unable to get contact: %s\n", gcr.Reason)
	}

	if contact.Type != gcr.Contact.Type {
		t.Fatalf("Type doesn't match. Got %v expected %v\n", contact.Type, gcr.Contact.Type)
	}

	if contact.Name != gcr.Contact.Name {
		t.Fatalf("Name doesn't match. Got %s expected %s.\n", contact.Name, gcr.Contact.Name)
	}

	if contact.Email != *gcr.Contact.Email {
		t.Fatalf("Email doesn't match. Got %q expected %q.\n", contact.Email, *gcr.Contact.Email)
	}

	if contact.Enabled != gcr.Contact.Enabled {
		t.Fatalf("Enabled doesn't match. Got %v expected %v.\n", contact.Enabled, gcr.Contact.Enabled)
	}

	if contact.Format != *gcr.Contact.Format {
		t.Fatalf("Format doesn't match. Got %s expected %s.\n", contact.Format, *gcr.Contact.Format)
	}

	contactUpdate := &UpdateContactRequest{
		ID:      ccr.Result.ID,
		Name:    PtrString("Test #2"),
		Enabled: PtrBool(false),
		Format:  PtrString("short"),
	}
	ucr, err := client.UpdateContact(ctx, contactUpdate)
	if err != nil {
		t.Fatalf("Unable to update contact: %s\n", err)
	}
	if !ucr.Success {
		t.Fatalf("Unable to update contact; %+v\n", ucr)
	}

	gcr, err = client.GetContact(ctx, ccr.Result.ID)
	if err != nil {
		t.Fatalf("Unable to get contact: %s\n", err)
	}
	if !gcr.Success {
		t.Fatalf("Unable to get contact: %s\n", gcr.Reason)
	}

	if gcr.Contact.Name != "Test #2" {
		t.Fatalf("Error doesn't match. Got %s expected %s\n", gcr.Contact.Name, "Test #2")
	}

	if gcr.Contact.Enabled {
		t.Fatalf("Error doesn't match. Got %v expected %v\n", gcr.Contact.Enabled, false)
	}

	if *gcr.Contact.Format != "short" {
		t.Fatalf("Error doesn't match. Got %s expected %s\n", *gcr.Contact.Format, "short")
	}

	deleteResp, err := client.DeleteContact(ctx, gcr.Contact.ID)
	if err != nil {
		t.Fatalf("Error deleting contact: %s\n", err)
	}
	if !deleteResp.Success {
		t.Fatalf("Unable to delete contact: %s\n", deleteResp.Result)
	}
}
