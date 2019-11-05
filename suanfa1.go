// suanfa1
package main

import (
	"fmt"
)

func bijiao(a byte, b byte) {
	if a < b {
		c := a
		a = b
		b = c
	}

}

func paixu(in []byte) {
	if len(in)-1 >= 1 {
		for i := 0; i < len(in)-1; i++ {
			bijiao(in[0], in[i+1])
		}
	}

}

func main() {
	a := []byte{10, 23, 1, 88, 13, 15}
	paixu(a)
	fmt.Println("Hello World!", a)
}
