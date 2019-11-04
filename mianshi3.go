package main

import (
	"fmt"
	"sync"
	"time"
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

func runA() {

	fmt.Println("ok")

	time.Sleep(time.Second * 5)
	wait.Done()

}

var wait sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		go runA()
		wait.Add(1)
	}

	wait.Wait()

	fmt.Println("run ok")
}
