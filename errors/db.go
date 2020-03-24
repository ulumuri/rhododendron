package errors

import (
	"github.com/ulumuri/rhododendron/meta"
)

const (
	CauseTypeCollectionNotFound meta.CauseType = "CollectionNotFound"
	CauseTypeDatabaseNotFound   meta.CauseType = "DatabaseNotFound"
)

func NewCollectionNotFound(src string, msg string) *StatusCauseError {
	return &StatusCauseError{
		StatusCauseError: meta.StatusCause{
			Type:    CauseTypeCollectionNotFound,
			Message: msg,
			Source:  src,
		}}
}

func NewDatabaseNotFound(src string, msg string) *StatusCauseError {
	return &StatusCauseError{
		StatusCauseError: meta.StatusCause{
			Type:    CauseTypeDatabaseNotFound,
			Message: msg,
			Source:  src,
		}}
}

func NewInvalidID(msg string) *StatusCauseError {
	return &StatusCauseError{
		StatusCauseError: meta.StatusCause{
			Type:    meta.CauseTypeFieldValueInvalid,
			Message: msg,
			Field:   "ID",
		}}
}

func NewIDNotFound(msg string) *StatusCauseError {
	return &StatusCauseError{
		StatusCauseError: meta.StatusCause{
			Type:    meta.CauseTypeFieldValueNotFound,
			Message: msg,
			Field:   "ID",
		}}
}

func NewUnknownCause(err error, src string) *StatusCauseError {
	return &StatusCauseError{
		StatusCauseError: meta.StatusCause{
			Type:   meta.CauseTypeUnknown,
			Error:  err,
			Source: src,
		}}
}
