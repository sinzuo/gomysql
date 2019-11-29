package main

import (
	"fmt"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	d, _ := grpc.Dial(":8081", grpc.WithInsecure())
	a := NewHelloServiceClient(d)

	b, _ := a.Hello(context.Background(), &String{Value: "jiang"})
	fmt.Println(b)

	fmt.Println("ok")
}
