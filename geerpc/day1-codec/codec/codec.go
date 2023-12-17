package codec

import (
	"io"
)

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         error
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	      Write(*Header, any) error
}

type NewCodecFunc func (io.ReadWriteCloser) Codec 

type Type string

const (
	GobType Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc, 0)
	NewCodecFuncMap[GobType] = NewGobCodec
}