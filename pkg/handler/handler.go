package handler

import (
	"net/http"
	
	"github.com/maxpower89/gotroller/pkg/appstate"
	"github.com/maxpower89/gotroller/pkg/config"
	"github.com/maxpower89/gotroller/pkg/request"
	"github.com/maxpower89/gotroller/pkg/response"
)

type HandleFunc func(request *request.Request) response.Response

type Handler struct {
	Config      *config.Config
	HandlerFunc HandleFunc
	AppState    *appstate.AppState
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request.NewRequest(*h.Config,h.AppState, r);
	res := h.HandlerFunc(req);
	res.Write(w, r);
}
