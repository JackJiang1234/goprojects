package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

type FramePayLoad []byte

type StreamFrameCodec interface {
	Encode(io.Writer, FramePayLoad) error
	Decode(io.Reader) (FramePayLoad, error)
}

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct{}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, payload FramePayLoad) error {
	var f = payload
	var totalLen int32 = int32(len(payload)) + 4

	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := w.Write([]byte(f))
	if err != nil {
		return err
	}

	if n != len(payload) {
		return ErrShortWrite
	}
	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayLoad, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, totalLen-4)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}

	return FramePayLoad(buf), nil
}
