package response

import "net/http"

type DataResponse struct {
	content interface{}
}

func NewDataResponse(content interface{}) *DataResponse {
	return &DataResponse{
		content: content,
	}
}

func (response *DataResponse) Write(writer http.ResponseWriter, r *http.Request) {
	jsonResponse := NewJsonResponse(response.content);
	jsonResponse.Write(writer, r);
}
