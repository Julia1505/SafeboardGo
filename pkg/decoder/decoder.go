package decoder

import (
	"errors"
	"io"
)

var (
	ISO_1 = "ISO8859_1"
)

var (
	TypeNotExist = errors.New("This type doesn't exist")
)

type Decoder interface {
	GetType() string
	Decode(file io.Reader, out chan<- string)
}

func NewDecoder(typeDecoder string) (Decoder, error) {
	switch typeDecoder {
	case ISO_1:
		return NewISO1Decoder(), nil
	default:
		return nil, TypeNotExist
	}
}
