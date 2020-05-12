package response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	content interface{}
}

func NewJsonResponse(content interface{}) *JsonResponse {
	return &JsonResponse{
		content: content,
	}
}

func (response *JsonResponse) Write(writer http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(response.content)
	if (err == nil) {
		writer.WriteHeader(200);
		writer.Header().Set("content-type", "application/json");
		writer.Write(js);
	} else {
		NewErrorResponse(500, "Failed to encode json");
	}
}
