package api

import (
	"encoding/json"
	"net/http"

	"github.com/ulumuri/rhododendron/status"
)

func RespondWithJsonStatus(w http.ResponseWriter, err error) {
	statusCode := err.(*status.StatusError).ErrStatus.Code
	response, err := json.Marshal(err)
	if err != nil {
		statusCode = http.StatusInternalServerError
		_, ok := err.(*json.UnsupportedTypeError)
		if ok {
			response = []byte(`"Error": "UnsupportedTypeError"`)
		} else {
			response = []byte(`"Error": "UnsupportedValueError"`)
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func RespondWithJson(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		apiErr := status.NewBadRequest("", err)
		RespondWithJsonStatus(w, apiErr)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
