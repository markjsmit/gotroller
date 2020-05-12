package response

import "net/http"

type PlainResponse struct {
	content string
}

func NewPlainResponse(content string) *PlainResponse {
	return &PlainResponse{
		content: content,
	}
}

func (response *PlainResponse) Write(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200);
	writer.Header().Set("content-type", "plain/text");
	writer.Write([]byte(response.content));
}
