package request

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/gorilla/Schema"
	"github.com/maxpower89/gotroller/pkg/appstate"
	"github.com/maxpower89/gotroller/pkg/config"
)

type Request struct {
	config   config.Config
	appState *appstate.AppState
	Original *http.Request
}

func NewRequest(config config.Config, appState *appstate.AppState, original *http.Request) *Request {
	return &Request{config: config, Original: original, appState: appState}
}

func (r Request) Get(key string) (string, error) {
	keys, ok := r.Original.URL.Query()[key];
	if !ok || len(keys[0]) < 1 {
		if (r.appState.AdditionalGetHandler != nil) {
			return r.appState.AdditionalGetHandler(key, r.Original);
		}
		return "", errors.New(fmt.Sprint("Url Param ", key, " Is missing"));
	}
	return string(keys[0]), nil;
}

func (r Request) GetInt(key string, fallback int) (int) {
	str, error := r.Get(key);
	if (error != nil) {
		return fallback;
	}
	i, err := strconv.Atoi(str);
	if (err != nil) {
		return fallback;
	}
	return i;
}

func (r Request) Post(key string) (string, error) {
	error := r.Original.ParseForm();
	if (error != nil) {
		return "", errors.New(fmt.Sprint("Form request is invalid"));
	}
	keys, ok := r.Original.Form[key];
	if !ok || len(keys[0]) < 1 {
		return "", errors.New(fmt.Sprint("Form params ", key, " Is missing"));
	}
	return string(keys[0]), nil;
}

func (r Request) Decode(target interface{}) (interface{}, error) {
	mime := r.Original.Header.Get("content-type");
	dec, error := r.appState.DecoderRegistry.GetForMime(mime);
	if (error != nil) {
		return target, error;
	}
	
	result, err := dec.Decode(r.Original.Body, target);
	return result, err;
}

func (r Request) DecodeQuery(target interface{}) error {
	decoder := schema.NewDecoder();
	err := decoder.Decode(target, r.Original.URL.Query());
	return err;
}
