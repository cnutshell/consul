// Code generated by protoc-gen-validate
// source: envoy/service/auth/v2alpha/external_auth.proto
// DO NOT EDIT!!!

package v2alpha

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckRequestValidationError{
				Field:  "Attributes",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CheckRequestValidationError{}

// Validate checks the field values on DeniedHttpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeniedHttpResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStatus() == nil {
		return DeniedHttpResponseValidationError{
			Field:  "Status",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeniedHttpResponseValidationError{
				Field:  "Status",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeniedHttpResponseValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for Body

	return nil
}

// DeniedHttpResponseValidationError is the validation error returned by
// DeniedHttpResponse.Validate if the designated constraints aren't met.
type DeniedHttpResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e DeniedHttpResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeniedHttpResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = DeniedHttpResponseValidationError{}

// Validate checks the field values on OkHttpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OkHttpResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OkHttpResponseValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// OkHttpResponseValidationError is the validation error returned by
// OkHttpResponse.Validate if the designated constraints aren't met.
type OkHttpResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e OkHttpResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOkHttpResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = OkHttpResponseValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CheckResponseValidationError{
				Field:  "Status",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.HttpResponse.(type) {

	case *CheckResponse_DeniedResponse:

		if v, ok := interface{}(m.GetDeniedResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					Field:  "DeniedResponse",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *CheckResponse_OkResponse:

		if v, ok := interface{}(m.GetOkResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					Field:  "OkResponse",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CheckResponseValidationError{}