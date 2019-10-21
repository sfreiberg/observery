// Package observery implements the observery API in Go. See
// https://observery.com/apidocs/#introduction for more info.
package observery

import (
	"time"

	"github.com/go-playground/form"
)

var (
	encoder = form.NewEncoder()
	decoder = form.NewDecoder()
)

func init() {
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.ParseDuration(vals[0])
	}, time.Duration(0))
}

// PtrString takes a string and returns a pointer to the string
func PtrString(s string) *string {
	return &s
}

// PtrBool takes a bool and returns a pointer to the bool
func PtrBool(b bool) *bool {
	return &b
}

// PtrInt takes an int and returns a pointer to the int
func PtrInt(i int) *int {
	return &i
}
