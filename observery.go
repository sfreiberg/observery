// Package observery implements the observery API in Go. See
// https://observery.com/apidocs/#introduction for more info.
package observery

import (
	"time"

	"github.com/go-playground/form/v4"
)

var (
	encoder = form.NewEncoder()
	decoder = form.NewDecoder()
)

func init() {
	decoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		return time.ParseDuration(vals[0])
	}, time.Duration(0))
}

// PtrString takes a string and returns a pointer to the string
//
//go:fix inline
func PtrString(s string) *string {
	return new(s)
}

// PtrBool takes a bool and returns a pointer to the bool
//
//go:fix inline
func PtrBool(b bool) *bool {
	return new(b)
}

// PtrInt takes an int and returns a pointer to the int
//
//go:fix inline
func PtrInt(i int) *int {
	return new(i)
}
