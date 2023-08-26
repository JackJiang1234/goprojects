package main

import (
	"fmt"
	"net"

	"github.com/jackjiang/tcp-server-demo1/frame"
	"github.com/jackjiang/tcp-server-demo1/packet"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	fmt.Println("server start ok(on *.8888)")

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()

	for {
		framePayLoad, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn: frame decode error:", err)
		}

		ackFramePayload, err := handlePacket(framePayLoad)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
		}

		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
	}
}

func handlePacket(framePayload []byte)(ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}

	switch p := p.(type) {
	case *packet.Submit:
		submit := p
		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := &packet.SubmitAck{
			ID: submit.ID,
			Result: 0,
		}

		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		} else {
			return ackFramePayload, nil
		}
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}