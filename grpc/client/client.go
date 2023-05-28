package main

import (
	"context"
	"demo/pb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SimpleRPC(c pb.GreeterClient) {
	ctx := context.Background()
	reply, err := c.SimpleRPC(ctx, &pb.HelloRequest{Name: "simplePRC"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetReply())
}

func serverSideStreamRPC(c pb.GreeterClient) {
	ctx := context.Background()
    // 调用服务端方法
	stream, err := c.ServerSideStreamingRPC(ctx, &pb.HelloRequest{Name: "gRPC!"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
		res, err := stream.Recv()
        // 判断流是否关闭
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("got reply: %q\n", res.GetReply())
	}
}

// client.go
func clientSideStreamRPC(c pb.GreeterClient) {
	ctx := context.Background()
    // 调用服务端方法
	stream, err := c.ClientSideStreamingRPC(ctx)
	if err != nil {
		log.Fatal(err)
	}

	names := []string{"a1", "a2", "a3", "a4"}
	for _, name := range names {
        // 向流中不断写入数据
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatal(err)
		}
	}

    // 向服务端发送关闭流的信号，并接收数据
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Reply)
}

func bidStreamRPC(c pb.GreeterClient) {
    // 调用服务端方法
	stream, err := c.BidrectionalStreamingRPC(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	names := []string{"a1", "a2", "a3", "a4"}
	for _, name := range names {
        // 向服务端发送数据
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatal(err)
		}

        // 从客户端接收数据
		reply := new(pb.HelloResponse)
		err = stream.RecvMsg(reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(reply.GetReply())
	}

    // 关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
}



func main() {
	dial, err := grpc.Dial("127.0.0.1:5678", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dial.Close()

	conn := pb.NewGreeterClient(dial)
	//SimpleRPC(conn)
	//serverSideStreamRPC(conn)
	//clientSideStreamRPC(conn)
	bidStreamRPC(conn);
}
