// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/http.proto

package envoy_data_tap_v2alpha

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

// Validate checks the field values on HttpBufferedTrace with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpBufferedTrace) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpBufferedTraceValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpBufferedTraceValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpBufferedTraceValidationError is the validation error returned by
// HttpBufferedTrace.Validate if the designated constraints aren't met.
type HttpBufferedTraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpBufferedTraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpBufferedTraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpBufferedTraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpBufferedTraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpBufferedTraceValidationError) ErrorName() string {
	return "HttpBufferedTraceValidationError"
}

// Error satisfies the builtin error interface
func (e HttpBufferedTraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpBufferedTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpBufferedTraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpBufferedTraceValidationError{}

// Validate checks the field values on HttpStreamedTraceSegment with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpStreamedTraceSegment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TraceId

	switch m.MessagePiece.(type) {

	case *HttpStreamedTraceSegment_RequestHeaders:

		if v, ok := interface{}(m.GetRequestHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "RequestHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpStreamedTraceSegment_RequestBodyChunk:

		if v, ok := interface{}(m.GetRequestBodyChunk()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "RequestBodyChunk",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpStreamedTraceSegment_RequestTrailers:

		if v, ok := interface{}(m.GetRequestTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "RequestTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpStreamedTraceSegment_ResponseHeaders:

		if v, ok := interface{}(m.GetResponseHeaders()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "ResponseHeaders",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpStreamedTraceSegment_ResponseBodyChunk:

		if v, ok := interface{}(m.GetResponseBodyChunk()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "ResponseBodyChunk",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpStreamedTraceSegment_ResponseTrailers:

		if v, ok := interface{}(m.GetResponseTrailers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpStreamedTraceSegmentValidationError{
					field:  "ResponseTrailers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpStreamedTraceSegmentValidationError is the validation error returned by
// HttpStreamedTraceSegment.Validate if the designated constraints aren't met.
type HttpStreamedTraceSegmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpStreamedTraceSegmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpStreamedTraceSegmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpStreamedTraceSegmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpStreamedTraceSegmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpStreamedTraceSegmentValidationError) ErrorName() string {
	return "HttpStreamedTraceSegmentValidationError"
}

// Error satisfies the builtin error interface
func (e HttpStreamedTraceSegmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpStreamedTraceSegment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpStreamedTraceSegmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpStreamedTraceSegmentValidationError{}

// Validate checks the field values on HttpBufferedTrace_Message with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpBufferedTrace_Message) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpBufferedTrace_MessageValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpBufferedTrace_MessageValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTrailers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpBufferedTrace_MessageValidationError{
					field:  fmt.Sprintf("Trailers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpBufferedTrace_MessageValidationError is the validation error returned by
// HttpBufferedTrace_Message.Validate if the designated constraints aren't met.
type HttpBufferedTrace_MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpBufferedTrace_MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpBufferedTrace_MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpBufferedTrace_MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpBufferedTrace_MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpBufferedTrace_MessageValidationError) ErrorName() string {
	return "HttpBufferedTrace_MessageValidationError"
}

// Error satisfies the builtin error interface
func (e HttpBufferedTrace_MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpBufferedTrace_Message.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpBufferedTrace_MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpBufferedTrace_MessageValidationError{}
