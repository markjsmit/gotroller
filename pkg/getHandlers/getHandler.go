package getHandlers

import "net/http"

type GetHandler func(key string, r *http.Request)(string,error)
