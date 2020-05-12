package decoder

import "io"

type Decoder interface {
	Decode(closer io.ReadCloser, target interface{}) (interface{}, error)
}
