package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"strings"
)

type Control struct {
	res http.ResponseWriter
	r   *http.Request
}

func (c *Control) InitC(res http.ResponseWriter, r *http.Request) {
	c.res = res
	c.r = r
}

func (c *Control) HtmlOut(data []byte) {
	c.res.Write(data)
}

type MyControl struct {
	Control
}

func (my *MyControl) Index() {
	b1 := tempjiexi()
	my.HtmlOut(b1)
}

type Server struct {
}

func returnBool(b bool) bool {
	return b
}

var ss = `{{if is_teacher_coming .}}Carefully!{{end}}`

var ss1 = `<html><body>jiangyibo</body></html>`

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	sp := strings.Split(r.URL.Path, "/")
	fmt.Println(sp[2])
	if sp[1] == "static" {
		http.ServeFile(w, r, path.Join("/root/work/gomysql", sp[2]))
	} else if sp[1] == "con" {
		p1 := reflect.New(LeiS["MyControl"])
		p2 := p1.MethodByName("InitC")
		val := make([]reflect.Value, 2)
		val[0] = reflect.ValueOf(w)
		val[1] = reflect.ValueOf(r)
		p2.Call(val)
		p3 := p1.MethodByName("Index")
		p3.Call(nil)
		fmt.Println("ooo")

	}

}

func add(a, b int) int {
	return a + b
}

var LeiS = make(map[string]reflect.Type, 10)

func tempjiexi() []byte {
	tfunc := make(map[string]interface{}, 10)
	var v1 bytes.Buffer
	tfunc["Add"] = add
	t1 := template.New("jiang")
	t1.Funcs(tfunc)
	i1, _ := ioutil.ReadFile("jiang.html")

	t2 := template.Must(t1.Parse(string(i1)))
	t2.ExecuteTemplate(&v1, "jiang", nil)
	fmt.Println(v1.Bytes())
	return v1.Bytes()

}

func main() {
	LeiS["MyControl"] = reflect.TypeOf(MyControl{})
	tempjiexi()

	a := Server{}
	fmt.Println("jiang")
	http.ListenAndServe(":8080", &a)

}
