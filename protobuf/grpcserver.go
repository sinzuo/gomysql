package main

import (
	"fmt"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type HelloServer struct {
}

func (a *HelloServer) Hello(c context.Context, d *String) (*String, error) {

	return &String{Value: "jiangyibo"}, nil
}

func main() {
	p := grpc.NewServer()
	RegisterHelloServiceServer(p, &HelloServer{})

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	p.Serve(lis)

	fmt.Println("ok")
}
