package codec

import "io"

type Header struct {
	ServiceMehtod string //service and method name
	Seq           uint64 //seq num
	Error         string
}

// 编解码接口设计，内容是调用服务与方法名称，序列号， 请求参数
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType Type = "appliction/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}


