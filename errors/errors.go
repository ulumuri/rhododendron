package errors

import (
	"net/http"

	"github.com/ulumuri/rhododendron/meta"
)

type StatusError struct {
	ErrStatus meta.Status
}

func (e *StatusError) Error() string {
	return e.ErrStatus.Message
}

func (e *StatusError) Status() meta.Status {
	return e.ErrStatus
}

type StatusCauseError struct {
	StatusCauseError meta.StatusCause
}

func (e *StatusCauseError) Error() string {
	return e.StatusCauseError.Message
}

func NewSuccess(msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusSuccess,
		Message: msg,
		Code:    http.StatusOK,
	}}
}

func NewBadRequest(msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonBadRequest,
		Code:    http.StatusBadRequest,
	}}
}

func NewInvalidData(cause *StatusCauseError, msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInvalidData,
		Code:    http.StatusBadRequest,
		Details: &meta.StatusDetails{
			Cause: &cause.StatusCauseError,
		},
	}}
}

func NewInternalError(cause *StatusCauseError, msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInternalError,
		Details: &meta.StatusDetails{
			Cause: &cause.StatusCauseError,
		},
		Code: http.StatusInternalServerError,
	}}
}

func NewFailedConnection(cause *StatusCauseError, msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInternalError,
		Details: &meta.StatusDetails{
			Cause: &cause.StatusCauseError,
		},
		Code: http.StatusInternalServerError,
	}}
}

func NewUnknown(cause *StatusCauseError, msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonUnknown,
		Details: &meta.StatusDetails{
			Cause: &cause.StatusCauseError,
		},
		Code: 0,
	}}
}
