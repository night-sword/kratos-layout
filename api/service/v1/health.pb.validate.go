// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/v1/health.proto

package v1

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

// Validate checks the field values on HealthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthRequestMultiError, or
// nil if none found.
func (m *HealthRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Service

	if len(errors) > 0 {
		return HealthRequestMultiError(errors)
	}

	return nil
}

// HealthRequestMultiError is an error wrapping multiple validation errors
// returned by HealthRequest.ValidateAll() if the designated constraints
// aren't met.
type HealthRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthRequestMultiError) AllErrors() []error { return m }

// HealthRequestValidationError is the validation error returned by
// HealthRequest.Validate if the designated constraints aren't met.
type HealthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthRequestValidationError) ErrorName() string { return "HealthRequestValidationError" }

// Error satisfies the builtin error interface
func (e HealthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthRequestValidationError{}

// Validate checks the field values on HealthReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthReplyMultiError, or
// nil if none found.
func (m *HealthReply) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	if len(errors) > 0 {
		return HealthReplyMultiError(errors)
	}

	return nil
}

// HealthReplyMultiError is an error wrapping multiple validation errors
// returned by HealthReply.ValidateAll() if the designated constraints aren't met.
type HealthReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthReplyMultiError) AllErrors() []error { return m }

// HealthReplyValidationError is the validation error returned by
// HealthReply.Validate if the designated constraints aren't met.
type HealthReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthReplyValidationError) ErrorName() string { return "HealthReplyValidationError" }

// Error satisfies the builtin error interface
func (e HealthReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthReplyValidationError{}
