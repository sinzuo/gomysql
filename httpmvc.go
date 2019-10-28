package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type A struct {
}

func (a *A) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	fmt.Println(req.URL.Path)
	if req.URL.Path == "/jiang" {
		http.ServeFile(res, req, "/root/work/gomysql/test.sql")
	} else {
		e1 := make([]reflect.Value, 2)
		t1 := ControlTable["MyControl"]
		fmt.Println(t1)
		v1 := reflect.New(reflect.TypeOf(MyControl{}))
		fmt.Println(v1.NumMethod())
		v2 := v1.MethodByName("Init")
		fmt.Println(v2.Type().Name())
		e1[0] = reflect.ValueOf(res)
		e1[1] = reflect.ValueOf(req)

		v2.Call(e1)

		v3 := v1.MethodByName("Index")

		v3.Call(nil)

		// res.Header().Add("", "text/html")
		// res.Write([]byte("jiang"))
	}

	fmt.Println("ok")
}

var ControlTable = make(map[string]reflect.Type, 10)

func main() {
	ControlTable["MyControl"] = reflect.TypeOf(&MyControl{})
	ControlTable["SecControl"] = reflect.TypeOf(&SecControl{})
	p1 := reflect.New(reflect.TypeOf(&MyControl{}).Elem())
	//p1 := reflect.ValueOf(&MyControl{})
	p2 := p1.MethodByName("Init")
	//fmt.Println(p2.Type().NumIn())
	if p2.IsValid() {
		e1 := make([]reflect.Value, 1)
		e1[0] = reflect.ValueOf(p1)
		fmt.Println("fun ok")
		//		p2.Call(nil)
	} else {
		fmt.Println("fun ri")
	}

	http.ListenAndServe(":8888", &A{})

}

type IControl interface {
	Init(res http.ResponseWriter, req *http.Request)
	ReturnJson(in interface{})
	ReturnHtml(in interface{})
	ReturnText(in interface{})
}

type Control struct {
	res http.ResponseWriter
	req *http.Request
}

func (A *Control) Init(res http.ResponseWriter, req *http.Request) {
	A.res = res
	A.req = req
	fmt.Println("ok" + req.URL.Path)

}

func (A *Control) ReturnJson(in interface{}) {
	A.res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	A.res.Write([]byte("ok"))
}

func (A *Control) ReturnHtml(in interface{}) {
	fmt.Println("jiangbbbbbbbb")
	A.res.Header().Set("Content-Type", "text/html;charset=UTF-8")
	A.res.Write([]byte("ok"))
}

func (A *Control) ReturnText(in interface{}) {
	A.res.Header().Set("Content-Type", "text/html;charset=UTF-8")
	A.res.Write([]byte("ok"))
}

type MyControl struct {
	Control
}

func (A *MyControl) Index() {
	fmt.Println("rrrrrr")
	A.ReturnHtml("jiang")
}

type SecControl struct {
	Control
}

func (A *SecControl) Index() {
	A.ReturnHtml("jiang")
}

type Abc struct {
}

func (A *Abc) Index() {
	fmt.Println("jiangyibo")
}
