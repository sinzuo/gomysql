package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 12
	b := reflect.ValueOf(&a)
	fmt.Println(b.Elem().Interface())
	fmt.Println(b.Elem().Type())
	b.Elem().SetInt(16)

	fmt.Println(a)

}
