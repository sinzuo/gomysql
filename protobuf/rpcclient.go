package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	d1, _ := rpc.Dial("tcp", ":8700")
	var a String
	var b String
	d1.Call("Ser.Request", a, &b)

	fmt.Println(b)

	fmt.Println("ok")

}
