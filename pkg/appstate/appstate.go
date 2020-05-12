package appstate

import (
	"github.com/maxpower89/gotroller/pkg/getHandlers"
	"github.com/maxpower89/gotroller/pkg/request/decoding/decoderRegistry"
)

type AppState struct {
	DecoderRegistry      *decoderRegistry.DecoderRegistry
	AdditionalGetHandler getHandlers.GetHandler
}
