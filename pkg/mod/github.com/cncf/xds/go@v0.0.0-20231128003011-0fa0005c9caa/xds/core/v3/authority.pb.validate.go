// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: xds/core/v3/authority.proto

package v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Authority with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Authority) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return AuthorityValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// AuthorityValidationError is the validation error returned by
// Authority.Validate if the designated constraints aren't met.
type AuthorityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorityValidationError) ErrorName() string { return "AuthorityValidationError" }

// Error satisfies the builtin error interface
func (e AuthorityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthority.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorityValidationError{}
