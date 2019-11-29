package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Ser struct {
}

func (s *Ser) Request(a String, b *String) error {

	return nil

}

func main() {
	var s = String{Value: "jiang"}
	rpc.Register(Ser{})
	d := net.Listen("tcp", ":7800")

	fmt.Println(s.String())

}
