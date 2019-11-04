/* 46.实现两个go轮流输出：A1B2C3.....Z26
 */

package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 100; i++ {

		select {
		case <-ch1:

			fmt.Println(i)
			ch2 <- 1

		}

	}
}

var s1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func f2() {
	for i := 0; i < len(s1); i++ {

		select {
		case <-ch2:
			fmt.Println(string(s1[i]))
			ch1 <- 1

		}

	}
}

var ch1 = make(chan int, 1)
var ch2 = make(chan int, 1)

func main() {
	go f1()
	go f2()
	//	time.Sleep(time.Second * 1)
	ch1 <- 1
	time.Sleep(time.Second * 10)

}
