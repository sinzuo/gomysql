package main

import "fmt"

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		//msg.Name = "qq"
		fmt.Print(msg)
	}
}

func main() {
	var a student
	a.Name = "string"
	zhoujielun(a)
}
