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

	if s1 == "0x2" {
		fmt.Println("{"+s1, ","+s2, ","+s3[0:4]+",0x"+s3[4:6]+"},")
	} else if s1 == "0x3" {
		fmt.Println("{"+s1, ","+s2, ","+s3[0:4]+",0x"+s3[4:6], ",0x"+s3[6:8]+"},")
	} else if s1 == "0x4" {
		fmt.Println("{"+s1, ","+s2, ","+s3[0:4]+",0x"+s3[4:6], ",0x"+s3[6:8]+",0x"+s3[8:10]+"},")
	} else if s1 == "0x6" {
		fmt.Println("{"+s1, ","+s2, ","+s3[0:4]+",0x"+s3[4:6], ",0x"+s3[6:8]+",0x"+s3[8:10]+",0x"+s3[10:12]+",0x"+s3[12:14]+"},")
	} else if s1 == "0x7" {
		fmt.Println("{"+s1, ","+s2, ","+s3[0:4]+",0x"+s3[4:6], ",0x"+s3[6:8]+",0x"+s3[8:10]+",0x"+s3[10:12]+",0x"+s3[12:14]+",0x"+s3[14:16]+"},")
	} else {

		fmt.Println("{"+s1, ","+s2, ","+s3+"},")
	}
}

func zhuanhuaS(s1 string, s2 string, s3 string) {
	if s1 == "0x1" {

		fmt.Println(s1, s2, s3)

	} else if s1 == "0x2" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])

	} else if s1 == "0x3" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])

	} else if s1 == "0x4" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])

	} else if s1 == "0x5" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])

	} else if s1 == "0x6" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		offset = strconv.FormatInt(Value+5, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[12:14])

	} else if s1 == "0x7" {

		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		offset = strconv.FormatInt(Value+5, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[12:14])
		offset = strconv.FormatInt(Value+6, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[14:16])

	} else {
		fmt.Println(s1, s2, s3)
	}

}

func zhuanhuaT(s1 string, s2 string, s3 string) {
	if s1 == "0x1" {

		fmt.Println(s1, s2, s3)

	} else if s1 == "0x2" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		fmt.Println("\n")
	} else if s1 == "0x3" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		fmt.Println("\n")
	} else if s1 == "0x4" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		fmt.Println("\n")
	} else if s1 == "0x5" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		fmt.Println("\n")
	} else if s1 == "0x6" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		offset = strconv.FormatInt(Value+5, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[12:14])
		fmt.Println("\n")
	} else if s1 == "0x7" {
		fmt.Println("\n"+s1, s2, s3)
		Value, _ := strconv.ParseInt(s2, 0, 16)
		fmt.Println("0x1", s2, s3[0:4])
		offset := strconv.FormatInt(Value+1, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[4:6])
		offset = strconv.FormatInt(Value+2, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[6:8])
		offset = strconv.FormatInt(Value+3, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[8:10])
		offset = strconv.FormatInt(Value+4, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[10:12])
		offset = strconv.FormatInt(Value+5, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[12:14])
		offset = strconv.FormatInt(Value+6, 16)
		fmt.Println("0x1", "0x"+offset, "0x"+s3[14:16])
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
		// if i > 1 {
		// 	break
		// }
		ts1 := strings.TrimRight(arr[1], ",")
		ts2 := "0x" + strings.TrimRight(arr[3], ",")
		ts3 := arr[5]
		//fmt.Println(ts1, "0x"+ts2, ts3)
		zhuanhua(ts1, ts2, ts3)

	}
	fo.Close()
	//	os.Remove(os.Args[1])
	//os.Rename(os.Args[1]+"temp", os.Args[1])

}
