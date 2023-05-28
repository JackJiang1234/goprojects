package main

import (
	"context"
	"demo/pb"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SimpleRPC(ctx context.Context, in *pb.HelloRequest)(*pb.HelloResponse, error){
	log.Println("client call simpleRPC...")
	log.Println(in)

	return &pb.HelloResponse{ Reply: "hello " + in.Name}, nil
}

func (s *server) ServerSideStreamingRPC(in *pb.HelloRequest, stream pb.Greeter_ServerSideStreamingRPCServer) error {
	log.Println("client call ServerSideStreamingRPC...")
	words := []string{
		"ni hao",
		"hello",
		"good morning",
		"yyds",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + " " + in.Name,
		}
		if err := stream.SendMsg(data); err != nil {
			return err
		}
	}

	return nil
}

// server.go
func (s *server) ClientSideStreamingRPC(stream pb.Greeter_ClientSideStreamingRPCServer) error {
	log.Println("client call ClientSideStreamingRPC...")
	reply := "Hello: "
	for {
        // 从流中接收客户端发送的数据
		res, err := stream.Recv()
        // 判断流是否关闭
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{Reply: reply})
		}

		if err != nil {
			return err
		}

		reply += ", " + res.GetName()
	}
}

// server.go
func (s *server) BidrectionalStreamingRPC(stream pb.Greeter_BidrectionalStreamingRPCServer) error {
	log.Println("client call BidrectionalStreamingRPC...")
	for {
        // 接收来自客户端的数据
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := "Hello " + res.GetName()
        // 向客户端发送数据
		if err := stream.SendMsg(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
		log.Println("send " + reply)
	}
}



func main() {
	listen, err := net.Listen("tcp", ":5678")
	if err != nil {
		log.Fatal(err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("grpc server starts running...")
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}