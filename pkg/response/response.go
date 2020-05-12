package response

import "net/http"


type Response interface {
	Write(http.ResponseWriter, *http.Request)
}
