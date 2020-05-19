package status

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

func NewSuccess(msg string) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusSuccess,
		Message: msg,
		Code:    http.StatusOK,
	}}
}

func NewBadRequest(msg string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonBadRequest,
		Code:    http.StatusBadRequest,
		Error:   err.Error(),
	}}
}

func NewUnknown(msg string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonUnknown,
		Code:    http.StatusInternalServerError,
		Error:   err.Error(),
	}}
}
