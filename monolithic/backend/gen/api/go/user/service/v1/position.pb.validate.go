// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/service/v1/position.proto

package servicev1

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

// Validate checks the field values on Position with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Position) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Position with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PositionMultiError, or nil
// if none found.
func (m *Position) ValidateAll() error {
	return m.validate(true)
}

func (m *Position) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.OrderNo != nil {
		// no validation rules for OrderNo
	}

	if m.Code != nil {
		// no validation rules for Code
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Remark != nil {
		// no validation rules for Remark
	}

	if m.CreateTime != nil {
		// no validation rules for CreateTime
	}

	if m.UpdateTime != nil {
		// no validation rules for UpdateTime
	}

	if m.DeleteTime != nil {
		// no validation rules for DeleteTime
	}

	if len(errors) > 0 {
		return PositionMultiError(errors)
	}

	return nil
}

// PositionMultiError is an error wrapping multiple validation errors returned
// by Position.ValidateAll() if the designated constraints aren't met.
type PositionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionMultiError) AllErrors() []error { return m }

// PositionValidationError is the validation error returned by
// Position.Validate if the designated constraints aren't met.
type PositionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionValidationError) ErrorName() string { return "PositionValidationError" }

// Error satisfies the builtin error interface
func (e PositionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPosition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionValidationError{}

// Validate checks the field values on ListPositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPositionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPositionResponseMultiError, or nil if none found.
func (m *ListPositionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPositionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPositionResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPositionResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPositionResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListPositionResponseMultiError(errors)
	}

	return nil
}

// ListPositionResponseMultiError is an error wrapping multiple validation
// errors returned by ListPositionResponse.ValidateAll() if the designated
// constraints aren't met.
type ListPositionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPositionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPositionResponseMultiError) AllErrors() []error { return m }

// ListPositionResponseValidationError is the validation error returned by
// ListPositionResponse.Validate if the designated constraints aren't met.
type ListPositionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPositionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPositionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPositionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPositionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPositionResponseValidationError) ErrorName() string {
	return "ListPositionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPositionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPositionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPositionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPositionResponseValidationError{}

// Validate checks the field values on GetPositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPositionRequestMultiError, or nil if none found.
func (m *GetPositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetPositionRequestMultiError(errors)
	}

	return nil
}

// GetPositionRequestMultiError is an error wrapping multiple validation errors
// returned by GetPositionRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPositionRequestMultiError) AllErrors() []error { return m }

// GetPositionRequestValidationError is the validation error returned by
// GetPositionRequest.Validate if the designated constraints aren't met.
type GetPositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPositionRequestValidationError) ErrorName() string {
	return "GetPositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPositionRequestValidationError{}

// Validate checks the field values on CreatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePositionRequestMultiError, or nil if none found.
func (m *CreatePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePositionRequestValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreatePositionRequestMultiError(errors)
	}

	return nil
}

// CreatePositionRequestMultiError is an error wrapping multiple validation
// errors returned by CreatePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type CreatePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePositionRequestMultiError) AllErrors() []error { return m }

// CreatePositionRequestValidationError is the validation error returned by
// CreatePositionRequest.Validate if the designated constraints aren't met.
type CreatePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePositionRequestValidationError) ErrorName() string {
	return "CreatePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePositionRequestValidationError{}

// Validate checks the field values on UpdatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePositionRequestMultiError, or nil if none found.
func (m *UpdatePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePositionRequestValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return UpdatePositionRequestMultiError(errors)
	}

	return nil
}

// UpdatePositionRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePositionRequestMultiError) AllErrors() []error { return m }

// UpdatePositionRequestValidationError is the validation error returned by
// UpdatePositionRequest.Validate if the designated constraints aren't met.
type UpdatePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePositionRequestValidationError) ErrorName() string {
	return "UpdatePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePositionRequestValidationError{}

// Validate checks the field values on DeletePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePositionRequestMultiError, or nil if none found.
func (m *DeletePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeletePositionRequestMultiError(errors)
	}

	return nil
}

// DeletePositionRequestMultiError is an error wrapping multiple validation
// errors returned by DeletePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type DeletePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePositionRequestMultiError) AllErrors() []error { return m }

// DeletePositionRequestValidationError is the validation error returned by
// DeletePositionRequest.Validate if the designated constraints aren't met.
type DeletePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePositionRequestValidationError) ErrorName() string {
	return "DeletePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePositionRequestValidationError{}
