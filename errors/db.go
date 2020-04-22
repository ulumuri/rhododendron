package errors

import (
	"fmt"
	"net/http"

	"github.com/ulumuri/rhododendron/meta"
)

const PostMaxSize int = 255

func NewFailedConnection(msg string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInternalError,
		Code:    http.StatusInternalServerError,
		Error:   err.Error(),
	}}
}

func NewCollectionNotFound(name string) *StatusError {
	msg := fmt.Sprintf("Couldn't find a collection with a given name: %s", name)
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInternalError,
		Details: meta.StatusDetailCollectionNotFound,
		Code:    http.StatusInternalServerError,
	}}
}

func NewInvalidID(msg string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInvalidData,
		Details: meta.StatusDetailFieldValueInvalid,
		Code:    http.StatusBadRequest,
		Error:   err.Error(),
	}}
}

func NewIDNotFound(msg string, err error) *StatusError {
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInvalidData,
		Details: meta.StatusDetailFieldValueNotFound,
		Code:    http.StatusNotFound,
		Error:   err.Error(),
	}}
}

func NewPostMaxSizeExceeded() *StatusError {
	msg := fmt.Sprintf("Post message is too long, it should be less than %d characters.", PostMaxSize)
	return &StatusError{meta.Status{
		Status:  meta.StatusFailure,
		Message: msg,
		Reason:  meta.StatusReasonInvalidData,
		Details: meta.StatusDetailPostMaxSizeExceeded,
		Code:    http.StatusBadRequest,
	}}
}
