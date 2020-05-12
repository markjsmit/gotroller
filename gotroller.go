package gotroller

import (
	"github.com/maxpower89/gotroller/pkg/appstate"
	"github.com/maxpower89/gotroller/pkg/config"
	"github.com/maxpower89/gotroller/pkg/getHandlers"
	"github.com/maxpower89/gotroller/pkg/handler"
	"github.com/maxpower89/gotroller/pkg/request/decoding/decoderRegistry"
)

type Gotroller struct {
	config   *config.Config
	appState *appstate.AppState
}

func NewGotroller(config *config.Config) *Gotroller {
	return &Gotroller{
		config: config,
		appState: &appstate.AppState{
			DecoderRegistry: decoderRegistry.NewDecoderRegistry(config),
		},
	};
}

func (gc *Gotroller) SetAdditionalGetHandler(additionalGethandler getHandlers.GetHandler) *Gotroller {
	gc.appState.AdditionalGetHandler = additionalGethandler;
	return gc;
}

func (gc *Gotroller) CreateHandler(handlerFunc handler.HandleFunc) *handler.Handler {
	return &handler.Handler{Config: gc.config, AppState: gc.appState, HandlerFunc: handlerFunc};
}
