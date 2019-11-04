package main

import (
	"fmt"
	"sync"
)

type A interface{}

var a *A = nil

var std sync.Mutex

func InitA() *A {

	if a == nil {
		std.Lock()
		defer std.Unlock()
		a = new(A)
	}
	return a
}

var pk sync.Once

func InitB() *A {
	pk.Do(func() {
		if a == nil {
			a = new(A)
		}
	})
	return a
}

func main() {
	b := InitB()
	fmt.Println(b)
}
