package decoderRegistry

import (
	"errors"
	"fmt"
	"strings"
	
	"github.com/maxpower89/gotroller/pkg/config"
	decoder2 "github.com/maxpower89/gotroller/pkg/request/decoding/decoder"
)

type DecoderRegistry struct {
	config   *config.Config
	decoders map[string]decoder2.Decoder
}

func NewDecoderRegistry(config *config.Config) *DecoderRegistry {
	
	reg := &DecoderRegistry{config: config, decoders: map[string]decoder2.Decoder{}}
	reg.Register("application/json", decoder2.NewJsonDecoder(config));
	return reg;
}

func (registry *DecoderRegistry) Register(mime string, decoder decoder2.Decoder) {
	registry.decoders[mime] = decoder;
}

func (registry *DecoderRegistry) GetForMime(mime string) (decoder2.Decoder, error) {
	if (len(mime) < 1) {
		mime = registry.config.RequestConfig.DefaultFormat;
	}
	
	mime = strings.ToLower(mime);
	decoder, ok := registry.decoders[mime];
	if (!ok) {
		if (mime == strings.ToLower(registry.config.RequestConfig.DefaultFormat)) {
			panic("Default decoder doesn't exist");
		}
		defaultDecoder, _ := registry.GetForMime("");
		return defaultDecoder, errors.New(fmt.Sprint("Decoder for ", mime, " not found."))
	}
	
	return decoder, nil;
}
