package responses

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// NewErrorResponse creates a new error response.
func NewErrorResponse(message string, err error) ErrorResponse {
	errorResponse := ErrorResponse{
		Message: message,
	}
	if err != nil {
		errorResponse.Error = err.Error()
	}
	return errorResponse
}
