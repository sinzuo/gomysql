package main

import (
	"log"
	"reflect"
)

type A struct {
	Age  int
	Name string
}

func (a A) Da() {
	log.Println("ok")
}

func (a A) DaReal(b interface{}) {
	v1 := reflect.ValueOf(b)
	v2 := v1.MethodByName("Wori")
	v2.Call(nil)
	log.Println("okmmm")
}

type B struct {
	A
}

func (a B) Da() {
	log.Println("B")
}

func (a B) Wori() {
	log.Println("Bbbbb")
}

func main() {
	var p1 = B{}
	p1.Da()
	p1.DaReal(p1)
	log.Println("jiang")
}
