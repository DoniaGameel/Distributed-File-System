// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/network/connection_balance/dlb/v3alpha/dlb.proto

package v3alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Dlb with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Dlb) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dlb with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DlbMultiError, or nil if none found.
func (m *Dlb) ValidateAll() error {
	return m.validate(true)
}

func (m *Dlb) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for MaxRetries

	// no validation rules for FallbackPolicy

	if len(errors) > 0 {
		return DlbMultiError(errors)
	}

	return nil
}

// DlbMultiError is an error wrapping multiple validation errors returned by
// Dlb.ValidateAll() if the designated constraints aren't met.
type DlbMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DlbMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DlbMultiError) AllErrors() []error { return m }

// DlbValidationError is the validation error returned by Dlb.Validate if the
// designated constraints aren't met.
type DlbValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DlbValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DlbValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DlbValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DlbValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DlbValidationError) ErrorName() string { return "DlbValidationError" }

// Error satisfies the builtin error interface
func (e DlbValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDlb.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DlbValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DlbValidationError{}
