package worker

// Response is worker function call result
type Response struct {
	// Error is error
	Error string `json:"error,omitempty"`

	// Result is execution result
	Result interface{} `json:"result,omitempty"`
}

// JSON returns value as JSON string
func (r Response) JSON() string { _ = "STUB: not implemented"; return "" }

// Return manual JSON in case of error

// NewErrorResponse returns a new response with error
func NewErrorResponse(err error) Response { _ = "STUB: not implemented"; return *new(Response) }

// NewResponse is Response constructor
func NewResponse(result interface{}, err error) Response {
	_ = "STUB: not implemented"
	return *new(Response)
}
