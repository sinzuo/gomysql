//ding shi renwu zhixing
package main

import (
	"log"
	"os"
	"syscall"
	"testing"
)

func init() {

	f1, _ := os.OpenFile("jiang.txt", syscall.O_CREAT|syscall.O_RDWR, os.ModePerm)
	log.SetOutput(f1)
	log.Println("ok11")
}

func TestOk(t *testing.T) {

	log.Println("ok")
}

func BenchmarkXinming(t *testing.B) {
	for i := 0; i < t.N; i++ {
		log.Println("xingneng ok")
	}
}
