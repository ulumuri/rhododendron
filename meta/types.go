package meta

const (
	StatusSuccess string = "success"
	StatusFailure string = "failure"
)

type Status struct {
	Status  string       `json:"status,omitempty"`
	Message string       `json:"message,omitempty"`
	Reason  StatusReason `json:"reason,omitempty"`
	Details StatusDetail `json:"details,omitempty"`
	Field   string       `json:"field,omitempty"`
	Code    int          `json:"code,omitempty"`
	Error   string       `json:"error,omitempty"`
}

type StatusReason string

const (
	StatusReasonUnknown       StatusReason = "StatusDetailUnknown"
	StatusReasonInvalidData   StatusReason = "InvalidData"
	StatusReasonBadRequest    StatusReason = "BadRequest"
	StatusReasonInternalError StatusReason = "InternalError"
)

type StatusDetail string

const (
	// Database specific details
	StatusDetailCollectionNotFound  StatusDetail = "CollectionNotFound"
	StatusDetailPostMaxSizeExceeded StatusDetail = "PostMaxSizeExceeded"

	// Other details
	StatusDetailFieldValueNotFound StatusDetail = "FieldValueNotFound"
	StatusDetailFieldValueInvalid  StatusDetail = "FieldValueInvalid"
)
