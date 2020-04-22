package api

import (
	"encoding/json"
	"net/http"

	"github.com/ulumuri/rhododendron/errors"
)

func RespondWithJsonStatus(w http.ResponseWriter, payload *errors.StatusError) {
	response, err := json.Marshal(payload)
	if err != nil {
		payload.ErrStatus.Code = http.StatusInternalServerError
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			response = []byte(`"Error": "UnsupportedTypeError"`)
		} else {
			response = []byte(`"Error": "UnsupportedValueError"`)
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(payload.ErrStatus.Code)
	w.Write(response)
}

func RespondWithJson(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		apiErr := errors.NewBadRequest("", err)
		RespondWithJsonStatus(w, apiErr)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
