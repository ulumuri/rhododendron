package meta

const (
	StatusSuccess string = "success"
	StatusFailure string = "failure"
)

type Status struct {
	Status  string         `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
	Reason  StatusReason   `json:"reason,omitempty"`
	Details *StatusDetails `json:"details,omitempty"`
	Code    int            `json:"code,omitempty"`
}

type StatusReason string

const (
	StatusReasonUnknown       StatusReason = "Unknown"
	StatusReasonInvalidData   StatusReason = "InvalidData"
	StatusReasonBadRequest    StatusReason = "BadRequest"
	StatusReasonInternalError StatusReason = "InternalError"
)

type StatusDetails struct {
	Cause             *StatusCause `json:"cause,omitempty"`
	RetryAfterSeconds int32        `json:"retryAfterSeconds,omitempty"`
}

type StatusCause struct {
	Type    CauseType `json:"cause,omitempty"`
	Message string    `json:"message,omitempty"`
	Field   string    `json:"field,omitempty"`
	Error   error     `json:"error,omitempty"`
	Source  string    `json:"source,omitempty"`
}

type CauseType string

const (
	CauseTypeFieldValueNotFound CauseType = "FieldValueNotFound"
	CauseTypeFieldValueInvalid  CauseType = "FieldValueInvalid"
	CauseTypeUnknown            CauseType = "Unknown"
)
