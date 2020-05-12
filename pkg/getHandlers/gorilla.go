package getHandlers

import (
	"errors"
	"net/http"
	
	"github.com/gorilla/mux"
)

func GorillaGetHandler(key string, r *http.Request) (string, error) {
	params := mux.Vars(r);
	result, ok := params[key];
	if (!ok) {
		return "", errors.New("Url param not found")
	}
	return result, nil;
}
