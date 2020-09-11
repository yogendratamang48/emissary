// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v4alpha/fault.proto

package envoy_extensions_filters_http_fault_v4alpha

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _fault_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on FaultAbort with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FaultAbort) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPercentage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultAbortValidationError{
				field:  "Percentage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.ErrorType.(type) {

	case *FaultAbort_HttpStatus:

		if val := m.GetHttpStatus(); val < 200 || val >= 600 {
			return FaultAbortValidationError{
				field:  "HttpStatus",
				reason: "value must be inside range [200, 600)",
			}
		}

	case *FaultAbort_GrpcStatus:
		// no validation rules for GrpcStatus

	case *FaultAbort_HeaderAbort_:

		if v, ok := interface{}(m.GetHeaderAbort()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultAbortValidationError{
					field:  "HeaderAbort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return FaultAbortValidationError{
			field:  "ErrorType",
			reason: "value is required",
		}

	}

	return nil
}

// FaultAbortValidationError is the validation error returned by
// FaultAbort.Validate if the designated constraints aren't met.
type FaultAbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultAbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultAbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultAbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultAbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultAbortValidationError) ErrorName() string { return "FaultAbortValidationError" }

// Error satisfies the builtin error interface
func (e FaultAbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultAbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultAbortValidationError{}

// Validate checks the field values on HTTPFault with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HTTPFault) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "Delay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAbort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "Abort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UpstreamCluster

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HTTPFaultValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetMaxActiveFaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "MaxActiveFaults",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponseRateLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HTTPFaultValidationError{
				field:  "ResponseRateLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DelayPercentRuntime

	// no validation rules for AbortPercentRuntime

	// no validation rules for DelayDurationRuntime

	// no validation rules for AbortHttpStatusRuntime

	// no validation rules for MaxActiveFaultsRuntime

	// no validation rules for ResponseRateLimitPercentRuntime

	// no validation rules for AbortGrpcStatusRuntime

	return nil
}

// HTTPFaultValidationError is the validation error returned by
// HTTPFault.Validate if the designated constraints aren't met.
type HTTPFaultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HTTPFaultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HTTPFaultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HTTPFaultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HTTPFaultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HTTPFaultValidationError) ErrorName() string { return "HTTPFaultValidationError" }

// Error satisfies the builtin error interface
func (e HTTPFaultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHTTPFault.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HTTPFaultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HTTPFaultValidationError{}

// Validate checks the field values on FaultAbort_HeaderAbort with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FaultAbort_HeaderAbort) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// FaultAbort_HeaderAbortValidationError is the validation error returned by
// FaultAbort_HeaderAbort.Validate if the designated constraints aren't met.
type FaultAbort_HeaderAbortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultAbort_HeaderAbortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultAbort_HeaderAbortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultAbort_HeaderAbortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultAbort_HeaderAbortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultAbort_HeaderAbortValidationError) ErrorName() string {
	return "FaultAbort_HeaderAbortValidationError"
}

// Error satisfies the builtin error interface
func (e FaultAbort_HeaderAbortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultAbort_HeaderAbort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultAbort_HeaderAbortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultAbort_HeaderAbortValidationError{}
