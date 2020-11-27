//
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handle(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("jiangyibo"))

}

func zhuanhua(s1 string, s2 string, s3 string) {
	if s1 == "0x1" || s1 == "0x2" || s1 == "0x3" || s1 == "0x4" {
		fmt.Println(s1, s2, s3)

	} else if s1 == "0x5" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])

	} else if s1 == "0x6" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x2", "0x"+offset, "0x"+s3[10:14])

	} else if s1 == "0x7" {

		Value, _ := strconv.ParseInt(s2, 0, 16)

		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x3", "0x"+offset, "0x"+s3[10:16])

	} else if s1 == "0x8" {

		Value, _ := strconv.ParseInt(s2, 0, 16)

		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])

	} else if s1 == "0x9" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[18:20])

	} else if s1 == "0xA" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x2", "0x"+offset, "0x"+s3[18:22])

	} else if s1 == "0xC" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[18:26])

	} else if s1 == "0xD" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[18:26])
		offset = strconv.FormatInt(Value+12, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[26:28])

	} else {
		fmt.Println(s1, s2, s3)
	}

}

func zhuanhuaTest(s1 string, s2 string, s3 string) {
	if s1 == "0x1" || s1 == "0x2" || s1 == "0x3" || s1 == "0x4" {
		fmt.Println(s1, s2, s3)

	} else if s1 == "0x5" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		fmt.Println("\n")
	} else if s1 == "0x6" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x2", "0x"+offset, "0x"+s3[10:14])
		fmt.Println("\n")
	} else if s1 == "0x8" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)

		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:14])
		fmt.Println("\n")
	} else if s1 == "0x9" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[18:20])
		fmt.Println("\n")
	} else if s1 == "0xA" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x2", "0x"+offset, "0x"+s3[18:22])
		fmt.Println("\n")
	} else if s1 == "0xC" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[18:26])
		fmt.Println("\n")
	} else if s1 == "0xD" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x4", s2, s3[0:10])
		offset := strconv.FormatInt(Value+4, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[10:18])
		offset = strconv.FormatInt(Value+8, 16)
		fmt.Println("0x4", "0x"+offset, "0x"+s3[18:26])
		offset = strconv.FormatInt(Value+12, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[26:28])
		fmt.Println("\n")
	} else {
		fmt.Println(s1, s2, s3)
	}

}

//var s1 = http.Handle("bobo", handle)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input file name")
		return
	}

	f1, e := os.Open(os.Args[1])
	if e != nil {
		fmt.Println("input file error")
		return
	}
	defer f1.Close()
	fo, _ := os.OpenFile(os.Args[1]+"temp", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	defer fo.Close()
	f2 := bufio.NewReader(f1)
	i := 0
	for {
		b1, _, e := f2.ReadLine()
		if e != nil {
			break
		}
		s := string(b1)
		arr := strings.Fields(s)
		i++
		// if i > 10 {
		// 	break
		// }
		ts1 := strings.TrimRight(arr[1], ",")
		ts2 := strings.TrimRight(arr[3], ",")
		ts3 := arr[5]
		zhuanhua(ts1, ts2, ts3)

	}
	fo.Close()
	//	os.Remove(os.Args[1])
	//os.Rename(os.Args[1]+"temp", os.Args[1])

}
