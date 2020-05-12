package response

import (
	"net/http"
)

type ErrorResponse struct {
	errorMessage string
	statusCode   int
}

func NewErrorResponse(statusCode int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		statusCode:   statusCode,
		errorMessage: errorMessage,
	}
}

func (errorResponse *ErrorResponse) Write(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(errorResponse.statusCode);
	writer.Header().Set("content-type", "application/json");
	writer.Write([]byte(errorResponse.errorMessage));
}
