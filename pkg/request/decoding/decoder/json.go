package decoder

import (
	"encoding/json"
	"io"
	"reflect"
	
	"github.com/maxpower89/gotroller/pkg/config"
)

type JsonDecoder struct {
	Config *config.Config
}

func NewJsonDecoder(config *config.Config) *JsonDecoder {
	return &JsonDecoder{Config: config};
}

func (*JsonDecoder) Decode(closer io.ReadCloser, target interface{}) (interface{}, error) {
	if (reflect.TypeOf(target).Kind() == reflect.Struct) {
		target = &target;
	}
	decoder := json.NewDecoder(closer);
	err := decoder.Decode(target);
	return target, err;
}
