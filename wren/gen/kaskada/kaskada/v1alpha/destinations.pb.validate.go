// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kaskada/kaskada/v1alpha/destinations.proto

package kaskadav1alpha

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

// Validate checks the field values on ObjectStoreDestination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ObjectStoreDestination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ObjectStoreDestination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ObjectStoreDestinationMultiError, or nil if none found.
func (m *ObjectStoreDestination) ValidateAll() error {
	return m.validate(true)
}

func (m *ObjectStoreDestination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FileType

	// no validation rules for OutputPrefixUri

	if len(errors) > 0 {
		return ObjectStoreDestinationMultiError(errors)
	}

	return nil
}

// ObjectStoreDestinationMultiError is an error wrapping multiple validation
// errors returned by ObjectStoreDestination.ValidateAll() if the designated
// constraints aren't met.
type ObjectStoreDestinationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ObjectStoreDestinationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ObjectStoreDestinationMultiError) AllErrors() []error { return m }

// ObjectStoreDestinationValidationError is the validation error returned by
// ObjectStoreDestination.Validate if the designated constraints aren't met.
type ObjectStoreDestinationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ObjectStoreDestinationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ObjectStoreDestinationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ObjectStoreDestinationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ObjectStoreDestinationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ObjectStoreDestinationValidationError) ErrorName() string {
	return "ObjectStoreDestinationValidationError"
}

// Error satisfies the builtin error interface
func (e ObjectStoreDestinationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sObjectStoreDestination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ObjectStoreDestinationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ObjectStoreDestinationValidationError{}

// Validate checks the field values on RedisDestination with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RedisDestination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedisDestination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RedisDestinationMultiError, or nil if none found.
func (m *RedisDestination) ValidateAll() error {
	return m.validate(true)
}

func (m *RedisDestination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HostName

	// no validation rules for Port

	// no validation rules for UseTls

	// no validation rules for DatabaseNumber

	// no validation rules for Password

	// no validation rules for TlsCert

	// no validation rules for TlsKey

	// no validation rules for TlsCaCert

	// no validation rules for InsecureSkipVerify

	if len(errors) > 0 {
		return RedisDestinationMultiError(errors)
	}

	return nil
}

// RedisDestinationMultiError is an error wrapping multiple validation errors
// returned by RedisDestination.ValidateAll() if the designated constraints
// aren't met.
type RedisDestinationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedisDestinationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedisDestinationMultiError) AllErrors() []error { return m }

// RedisDestinationValidationError is the validation error returned by
// RedisDestination.Validate if the designated constraints aren't met.
type RedisDestinationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisDestinationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisDestinationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisDestinationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisDestinationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisDestinationValidationError) ErrorName() string { return "RedisDestinationValidationError" }

// Error satisfies the builtin error interface
func (e RedisDestinationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisDestination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisDestinationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisDestinationValidationError{}

// Validate checks the field values on PulsarDestination with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PulsarDestination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PulsarDestination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PulsarDestinationMultiError, or nil if none found.
func (m *PulsarDestination) ValidateAll() error {
	return m.validate(true)
}

func (m *PulsarDestination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Tenant

	// no validation rules for Namespace

	// no validation rules for TopicName

	// no validation rules for BrokerServiceUrl

	if len(errors) > 0 {
		return PulsarDestinationMultiError(errors)
	}

	return nil
}

// PulsarDestinationMultiError is an error wrapping multiple validation errors
// returned by PulsarDestination.ValidateAll() if the designated constraints
// aren't met.
type PulsarDestinationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PulsarDestinationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PulsarDestinationMultiError) AllErrors() []error { return m }

// PulsarDestinationValidationError is the validation error returned by
// PulsarDestination.Validate if the designated constraints aren't met.
type PulsarDestinationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PulsarDestinationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PulsarDestinationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PulsarDestinationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PulsarDestinationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PulsarDestinationValidationError) ErrorName() string {
	return "PulsarDestinationValidationError"
}

// Error satisfies the builtin error interface
func (e PulsarDestinationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPulsarDestination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PulsarDestinationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PulsarDestinationValidationError{}
