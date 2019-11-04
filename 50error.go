目录
初级篇
开大括号不能放在单独的一行
未使用的变量
未使用的Imports
简式的变量声明仅可以在函数内部使用
使用简式声明重复声明变量
偶然的变量隐藏Accidental Variable Shadowing
不使用显式类型，无法使用“nil”来初始化变量
使用“nil” Slices and Maps
Map的容量
字符串不会为“nil”
Array函数的参数
在Slice和Array使用“range”语句时的出现的不希望得到的值
Slices和Arrays是一维的
访问不存在的Map Keys
Strings无法修改
String和Byte Slice之间的转换
String和索引操作
字符串不总是UTF8文本
字符串的长度
在多行的Slice、Array和Map语句中遗漏逗号
log.Fatal和log.Panic不仅仅是Log
内建的数据结构操作不是同步的
String在“range”语句中的迭代值
对Map使用“for range”语句迭代
"switch"声明中的失效行为
自增和自减
按位NOT操作
操作优先级的差异
未导出的结构体不会被编码
有活动的Goroutines下的应用退出
向无缓存的Channel发送消息，只要目标接收者准备好就会立即返回
向已关闭的Channel发送会引起Panic
使用"nil" Channels
传值方法的接收者无法修改原有的值
进阶篇
关闭HTTP的响应
关闭HTTP的连接
比较Structs, Arrays, Slices, and Maps
从Panic中恢复
在Slice, Array, and Map "range"语句中更新引用元素的值
在Slice中"隐藏"数据
Slice的数据“毁坏”
"走味的"Slices
类型声明和方法
从"for switch"和"for select"代码块中跳出
"for"声明中的迭代变量和闭包
Defer函数调用参数的求值
被Defer的函数调用执行
失败的类型断言
阻塞的Goroutine和资源泄露
高级篇
使用指针接收方法的值的实例
更新Map的值
"nil" Interfaces和"nil" Interfaces的值
栈和堆变量
GOMAXPROCS, 并发, 和并行
读写操作的重排顺序
优先调度
进阶篇
关闭HTTP的响应
level: intermediate
当你使用标准http库发起请求时，你得到一个http的响应变量。如果你不读取响应主体，你依旧需要关闭它。注意对于空的响应你也一定要这么做。对于新的Go开发者而言，这个很容易就会忘掉。

一些新的Go开发者确实尝试关闭响应主体，但他们在错误的地方做。

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    resp, err := http.Get("https://api.ipify.org?format=json")
    defer resp.Body.Close()//not ok
    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}
这段代码对于成功的请求没问题，但如果http的请求失败， resp变量可能会是 nil，这将导致一个runtime panic。

最常见的关闭响应主体的方法是在http响应的错误检查后调用 defer。

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    resp, err := http.Get("https://api.ipify.org?format=json")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()//ok, most of the time :-)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}
大多数情况下，当你的http响应失败时， resp变量将为 nil，而 err变量将是 non-nil。然而，当你得到一个重定向的错误时，两个变量都将是 non-nil。这意味着你最后依然会内存泄露。

通过在http响应错误处理中添加一个关闭 non-nil响应主体的的调用来修复这个问题。另一个方法是使用一个 defer调用来关闭所有失败和成功的请求的响应主体。

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    resp, err := http.Get("https://api.ipify.org?format=json")
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}
resp.Body.Close()的原始实现也会读取并丢弃剩余的响应主体数据。这确保了http的链接在keepalive http连接行为开启的情况下，可以被另一个请求复用。最新的http客户端的行为是不同的。现在读取并丢弃剩余的响应数据是你的职责。如果你不这么做，http的连接可能会关闭，而无法被重用。这个小技巧应该会写在Go 1.5的文档中。

如果http连接的重用对你的应用很重要，你可能需要在响应处理逻辑的后面添加像下面的代码：

_, err = io.Copy(ioutil.Discard, resp.Body)  
如果你不立即读取整个响应将是必要的，这可能在你处理json API响应时会发生：

json.NewDecoder(resp.Body).Decode(&data)
关闭HTTP的连接
level: intermediate
一些HTTP服务器保持会保持一段时间的网络连接（根据HTTP 1.1的说明和服务器端的“keep-alive”配置）。默认情况下，标准http库只在目标HTTP服务器要求关闭时才会关闭网络连接。这意味着你的应用在某些条件下消耗完sockets/file的描述符。

你可以通过设置请求变量中的 Close域的值为 true，来让http库在请求完成时关闭连接。

另一个选项是添加一个 Connection的请求头，并设置为 close。目标HTTP服务器应该也会响应一个 Connection: close的头。当http库看到这个响应头时，它也将会关闭连接。

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    req, err := http.NewRequest("GET","http://golang.org",nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Close = true
    //or do this:
    //req.Header.Add("Connection", "close")

    resp, err := http.DefaultClient.Do(req)
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(len(string(body)))
}
你也可以取消http的全局连接复用。你将需要为此创建一个自定义的http传输配置。

package main

import (  
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {  
    tr := &http.Transport{DisableKeepAlives: true}
    client := &http.Client{Transport: tr}

    resp, err := client.Get("http://golang.org")
    if resp != nil {
        defer resp.Body.Close()
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp.StatusCode)

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(len(string(body)))
}
如果你向同一个HTTP服务器发送大量的请求，那么把保持网络连接的打开是没问题的。然而，如果你的应用在短时间内向大量不同的HTTP服务器发送一两个请求，那么在引用收到响应后立刻关闭网络连接是一个好主意。增加打开文件的限制数可能也是个好主意。当然，正确的选择源自于应用。

比较Structs, Arrays, Slices, and Maps
level: intermediate
如果结构体中的各个元素都可以用你可以使用等号来比较的话，那就可以使用相号, ==，来比较结构体变量。

package main

import "fmt"

type data struct {  
    num int
    fp float32
    complex complex64
    str string
    char rune
    yes bool
    events <-chan string
    handler interface{}
    ref *byte
    raw [10]byte
}

func main() {  
    v1 := data{}
    v2 := data{}
    fmt.Println("v1 == v2:",v1 == v2) //prints: v1 == v2: true
}
如果结构体中的元素无法比较，那使用等号将导致编译错误。注意数组仅在它们的数据元素可比较的情况下才可以比较。

package main

import "fmt"

type data struct {  
    num int                //ok
    checks [10]func() bool //not comparable
    doit func() bool       //not comparable
    m map[string] string   //not comparable
    bytes []byte           //not comparable
}

func main() {  
    v1 := data{}
    v2 := data{}
    fmt.Println("v1 == v2:",v1 == v2)
}
Go确实提供了一些助手函数，用于比较那些无法使用等号比较的变量。

最常用的方法是使用 reflect包中的 DeepEqual()函数。

package main

import (  
    "fmt"
    "reflect"
)

type data struct {  
    num int                //ok
    checks [10]func() bool //not comparable
    doit func() bool       //not comparable
    m map[string] string   //not comparable
    bytes []byte           //not comparable
}

func main() {  
    v1 := data{}
    v2 := data{}
    fmt.Println("v1 == v2:",reflect.DeepEqual(v1,v2)) //prints: v1 == v2: true

    m1 := map[string]string{"one": "a","two": "b"}
    m2 := map[string]string{"two": "b", "one": "a"}
    fmt.Println("m1 == m2:",reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true

    s1 := []int{1, 2, 3}
    s2 := []int{1, 2, 3}
    fmt.Println("s1 == s2:",reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true
}
除了很慢（这个可能会也可能不会影响你的应用）， DeepEqual()也有其他自身的技巧。

package main

import (  
    "fmt"
    "reflect"
)

func main() {  
    var b1 []byte = nil
    b2 := []byte{}
    fmt.Println("b1 == b2:",reflect.DeepEqual(b1, b2)) //prints: b1 == b2: false
}
DeepEqual()不会认为空的slice与“nil”的slice相等。这个行为与你使用 bytes.Equal()函数的行为不同。 bytes.Equal()认为“nil”和空的slice是相等的。

package main

import (  
    "fmt"
    "bytes"
)

func main() {  
    var b1 []byte = nil
    b2 := []byte{}
    fmt.Println("b1 == b2:",bytes.Equal(b1, b2)) //prints: b1 == b2: true
}
DeepEqual()在比较slice时并不总是完美的。

package main

import (  
    "fmt"
    "reflect"
    "encoding/json"
)

func main() {  
    var str string = "one"
    var in interface{} = "one"
    fmt.Println("str == in:",str == in,reflect.DeepEqual(str, in))
    //prints: str == in: true true

    v1 := []string{"one","two"}
    v2 := []interface{}{"one","two"}
    fmt.Println("v1 == v2:",reflect.DeepEqual(v1, v2))
    //prints: v1 == v2: false (not ok)

    data := map[string]interface{}{
        "code": 200,
        "value": []string{"one","two"},
    }
    encoded, _ := json.Marshal(data)
    var decoded map[string]interface{}
    json.Unmarshal(encoded, &decoded)
    fmt.Println("data == decoded:",reflect.DeepEqual(data, decoded))
    //prints: data == decoded: false (not ok)
}
如果你的byte slice（或者字符串）中包含文字数据，而当你要不区分大小写形式的值时（在使用 ==， bytes.Equal()，或者 bytes.Compare()），你可能会尝试使用“bytes”和“string”包中的 ToUpper()或者 ToLower()函数。对于英语文本，这么做是没问题的，但对于许多其他的语言来说就不行了。这时应该使用 strings.EqualFold()和 bytes.EqualFold()。

如果你的byte slice中包含需要验证用户数据的隐私信息（比如，加密哈希、tokens等），不要使用 reflect.DeepEqual()、 bytes.Equal()，或者 bytes.Compare()，因为这些函数将会让你的应用易于被定时攻击。为了避免泄露时间信息，使用 'crypto/subtle'包中的函数（即， subtle.ConstantTimeCompare()）。

从Panic中恢复
level: intermediate
recover()函数可以用于获取/拦截panic。仅当在一个defer函数中被完成时，调用 recover()将会完成这个小技巧。

Incorrect:

ackage main

import "fmt"

func main() {  
    recover() //doesn't do anything
    panic("not good")
    recover() //won't be executed :)
    fmt.Println("ok")
}
Works:

package main

import "fmt"

func main() {  
    defer func() {
        fmt.Println("recovered:",recover())
    }()

    panic("not good")
}
recover()的调用仅当它在defer函数中被直接调用时才有效。

Fails:

package main

import "fmt"

func doRecover() {  
    fmt.Println("recovered =>",recover()) //prints: recovered => <nil>
}

func main() {  
    defer func() {
        doRecover() //panic is not recovered
    }()

    panic("not good")
}
在Slice, Array, and Map "range"语句中更新引用元素的值
level: intermediate
在“range”语句中生成的数据的值是真实集合元素的拷贝。它们不是原有元素的引用。这意味着更新这些值将不会修改原来的数据。同时也意味着使用这些值的地址将不会得到原有数据的指针。

package main

import "fmt"

func main() {  
    data := []int{1,2,3}
    for _,v := range data {
        v *= 10 //original item is not changed
    }

    fmt.Println("data:",data) //prints data: [1 2 3]
}
如果你需要更新原有集合中的数据，使用索引操作符来获得数据。

package main

import "fmt"

func main() {  
    data := []int{1,2,3}
    for i,_ := range data {
        data[i] *= 10
    }

    fmt.Println("data:",data) //prints data: [10 20 30]
}
如果你的集合保存的是指针，那规则会稍有不同。如果要更新原有记录指向的数据，你依然需要使用索引操作，但你可以使用 for range语句中的第二个值来更新存储在目标位置的数据。

package main

import "fmt"

func main() {  
    data := []*struct{num int} {{1},{2},{3}}

    for _,v := range data {
        v.num *= 10
    }

    fmt.Println(data[0],data[1],data[2]) //prints &{10} &{20} &{30}
}
在Slice中"隐藏"数据
level: intermediate
当你重新划分一个slice时，新的slice将引用原有slice的数组。如果你忘了这个行为的话，在你的应用分配大量临时的slice用于创建新的slice来引用原有数据的一小部分时，会导致难以预期的内存使用。

package main

import "fmt"

func get() []byte {  
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    return raw[:3]
}

func main() {  
    data := get()
    fmt.Println(len(data),cap(data),&data[0]) //prints: 3 10000 <byte_addr_x>
}
为了避免这个陷阱，你需要从临时的slice中拷贝数据（而不是重新划分slice）。

package main

import "fmt"

func get() []byte {  
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    res := make([]byte,3)
    copy(res,raw[:3])
    return res
}

func main() {  
    data := get()
    fmt.Println(len(data),cap(data),&data[0]) //prints: 3 3 <byte_addr_y>
}
Slice的数据“毁坏”
level: intermediate
比如说你需要重新一个路径（在slice中保存）。你通过修改第一个文件夹的名字，然后把名字合并来创建新的路劲，来重新划分指向各个文件夹的路径。

package main

import (  
    "fmt"
    "bytes"
)

func main() {  
    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex]
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

    dir1 = append(dir1,"suffix"...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB (not ok)

    fmt.Println("new path =>",string(path))
}
结果与你想的不一样。与"AAAAsuffix/BBBBBBBBB"相反，你将会得到"AAAAsuffix/uffixBBBB"。这个情况的发生是因为两个文件夹的slice都潜在的引用了同一个原始的路径slice。这意味着原始路径也被修改了。根据你的应用，这也许会是个问题。

通过分配新的slice并拷贝需要的数据，你可以修复这个问题。另一个选择是使用完整的slice表达式。

package main

import (  
    "fmt"
    "bytes"
)

func main() {  
    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/')
    dir1 := path[:sepIndex:sepIndex] //full slice expression
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

    dir1 = append(dir1,"suffix"...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB (ok now)

    fmt.Println("new path =>",string(path))
}
完整的slice表达式中的额外参数可以控制新的slice的容量。现在在那个slice后添加元素将会触发一个新的buffer分配，而不是覆盖第二个slice中的数据。

"走味的"Slices
level: intermediate
多个slice可以引用同一个数据。比如，当你从一个已有的slice创建一个新的slice时，这就会发生。如果你的应用功能需要这种行为，那么你将需要关注下“走味的”slice。

在某些情况下，在一个slice中添加新的数据，在原有数组无法保持更多新的数据时，将导致分配一个新的数组。而现在其他的slice还指向老的数组（和老的数据）。

import "fmt"

func main() {  
    s1 := []int{1,2,3}
    fmt.Println(len(s1),cap(s1),s1) //prints 3 3 [1 2 3]

    s2 := s1[1:]
    fmt.Println(len(s2),cap(s2),s2) //prints 2 2 [2 3]

    for i := range s2 { s2[i] += 20 }

    //still referencing the same array
    fmt.Println(s1) //prints [1 22 23]
    fmt.Println(s2) //prints [22 23]

    s2 = append(s2,4)

    for i := range s2 { s2[i] += 10 }

    //s1 is now "stale"
    fmt.Println(s1) //prints [1 22 23]
    fmt.Println(s2) //prints [32 33 14]
}
类型声明和方法
level: intermediate
当你通过把一个现有（非interface）的类型定义为一个新的类型时，新的类型不会继承现有类型的方法。

Fails:

package main

import "sync"

type myMutex sync.Mutex

func main() {  
    var mtx myMutex
    mtx.Lock() //error
    mtx.Unlock() //error  
}
Compile Errors:

/tmp/sandbox106401185/main.go:9: mtx.Lock undefined (type myMutex has no field or method Lock) /tmp/sandbox106401185/main.go:10: mtx.Unlock undefined (type myMutex has no field or method Unlock)
如果你确实需要原有类型的方法，你可以定义一个新的struct类型，用匿名方式把原有类型嵌入其中。

Works:

package main

import "sync"

type myLocker struct {  
    sync.Mutex
}

func main() {  
    var lock myLocker
    lock.Lock() //ok
    lock.Unlock() //ok
}
interface类型的声明也会保留它们的方法集合。

Works:

package main

import "sync"

type myLocker sync.Locker

func main() {  
    var lock myLocker = new(sync.Mutex)
    lock.Lock() //ok
    lock.Unlock() //ok
}
从"for switch"和"for select"代码块中跳出
level: intermediate
没有标签的“break”声明只能从内部的switch/select代码块中跳出来。如果无法使用“return”声明的话，那就为外部循环定义一个标签是另一个好的选择。

package main

import "fmt"

func main() {  
    loop:
        for {
            switch {
            case true:
                fmt.Println("breaking out...")
                break loop
            }
        }

    fmt.Println("out!")
}
"goto"声明也可以完成这个功能。。。

"for"声明中的迭代变量和闭包
level: intermediate
这在Go中是个很常见的技巧。 for语句中的迭代变量在每次迭代时被重新使用。这就意味着你在 for循环中创建的闭包（即函数字面量）将会引用同一个变量（而在那些goroutine开始执行时就会得到那个变量的值）。

Incorrect:

package main

import (  
    "fmt"
    "time"
)

func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}
最简单的解决方法（不需要修改goroutine）是，在 for循环代码块内把当前迭代的变量值保存到一个局部变量中。

Works:

package main

import (  
    "fmt"
    "time"
)

func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        vcopy := v //
        go func() {
            fmt.Println(vcopy)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
另一个解决方法是把当前的迭代变量作为匿名goroutine的参数。

Works:

package main

import (  
    "fmt"
    "time"
)

func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
        go func(in string) {
            fmt.Println(in)
        }(v)
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
下面这个陷阱稍微复杂一些的版本。

Incorrect:

package main

import (  
    "fmt"
    "time"
)

type field struct {  
    name string
}

func (p *field) print() {  
    fmt.Println(p.name)
}

func main() {  
    data := []field{{"one"},{"two"},{"three"}}

    for _,v := range data {
        go v.print()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}
Works:

package main

import (  
    "fmt"
    "time"
)

type field struct {  
    name string
}

func (p *field) print() {  
    fmt.Println(p.name)
}

func main() {  
    data := []field{{"one"},{"two"},{"three"}}

    for _,v := range data {
        v := v
        go v.print()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: one, two, three
}
在运行这段代码时你认为会看到什么结果？（原因是什么？）

package main

import (  
    "fmt"
    "time"
)

type field struct {  
    name string
}

func (p *field) print() {  
    fmt.Println(p.name)
}

func main() {  
    data := []*field{{"one"},{"two"},{"three"}}

    for _,v := range data {
        go v.print()
    }

    time.Sleep(3 * time.Second)
}
Defer函数调用参数的求值
level: intermediate
被defer的函数的参数会在defer声明时求值（而不是在函数实际执行时）。 
Arguments for a deferred function call are evaluated when the defer statement is evaluated (not when the function is actually executing).

package main

import "fmt"

func main() {  
    var i int = 1

    defer fmt.Println("result =>",func() int { return i * 2 }())
    i++
    //prints: result => 2 (not ok if you expected 4)
}
被Defer的函数调用执行
level: intermediate
被defer的调用会在包含的函数的末尾执行，而不是包含代码块的末尾。对于Go新手而言，一个很常犯的错误就是无法区分被defer的代码执行规则和变量作用规则。如果你有一个长时运行的函数，而函数内有一个 for循环试图在每次迭代时都 defer资源清理调用，那就会出现问题。

package main

import (  
    "fmt"
    "os"
    "path/filepath"
)

func main() {  
    if len(os.Args) != 2 {
        os.Exit(-1)
    }

    start, err := os.Stat(os.Args[1])
    if err != nil || !start.IsDir(){
        os.Exit(-1)
    }

    var targets []string
    filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !fi.Mode().IsRegular() {
            return nil
        }

        targets = append(targets,fpath)
        return nil
    })

    for _,target := range targets {
        f, err := os.Open(target)
        if err != nil {
            fmt.Println("bad target:",target,"error:",err) //prints error: too many open files
            break
        }
        defer f.Close() //will not be closed at the end of this code block
        //do something with the file...
    }
}
解决这个问题的一个方法是把代码块写成一个函数。

package main

import (  
    "fmt"
    "os"
    "path/filepath"
)

func main() {  
    if len(os.Args) != 2 {
        os.Exit(-1)
    }

    start, err := os.Stat(os.Args[1])
    if err != nil || !start.IsDir(){
        os.Exit(-1)
    }

    var targets []string
    filepath.Walk(os.Args[1], func(fpath string, fi os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !fi.Mode().IsRegular() {
            return nil
        }

        targets = append(targets,fpath)
        return nil
    })

    for _,target := range targets {
        func() {
            f, err := os.Open(target)
            if err != nil {
                fmt.Println("bad target:",target,"error:",err)
                return
            }
            defer f.Close() //ok
            //do something with the file...
        }()
    }
}
另一个方法是去掉 defer语句 :-)

失败的类型断言
level: intermediate
失败的类型断言返回断言声明中使用的目标类型的“零值”。这在与隐藏变量混合时，会发生未知情况。

Incorrect:

package main

import "fmt"

func main() {  
    var data interface{} = "great"

    if data, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",data)
    } else {
        fmt.Println("[not an int] value =>",data)
        //prints: [not an int] value => 0 (not "great")
    }
}
Works:

package main

import "fmt"

func main() {  
    var data interface{} = "great"

    if res, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",res)
    } else {
        fmt.Println("[not an int] value =>",data)
        //prints: [not an int] value => great (as expected)
    }
}
阻塞的Goroutine和资源泄露
level: intermediate
Rob Pike在2012年的Google I/O大会上所做的“Go Concurrency Patterns”的演讲上，说道过几种基础的并发模式。从一组目标中获取第一个结果就是其中之一。

func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
这个函数在每次搜索重复时都会起一个goroutine。每个goroutine把它的搜索结果发送到结果的channel中。结果channel的第一个值被返回。

那其他goroutine的结果会怎样呢？还有那些goroutine自身呢？

在 First()函数中的结果channel是没缓存的。这意味着只有第一个goroutine返回。其他的goroutine会困在尝试发送结果的过程中。这意味着，如果你有不止一个的重复时，每个调用将会泄露资源。

为了避免泄露，你需要确保所有的goroutine退出。一个不错的方法是使用一个有足够保存所有缓存结果的channel。

func First(query string, replicas ...Search) Result {  
    c := make(chan Result,len(replicas))
    searchReplica := func(i int) { c <- replicas[i](query) }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
另一个不错的解决方法是使用一个有 default情况的 select语句和一个保存一个缓存结果的channel。 default情况保证了即使当结果channel无法收到消息的情况下，goroutine也不会堵塞。

func First(query string, replicas ...Search) Result {  
    c := make(chan Result,1)
    searchReplica := func(i int) {
        select {
        case c <- replicas[i](query):
        default:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }
    return <-c
}
你也可以使用特殊的取消channel来终止workers。

func First(query string, replicas ...Search) Result {  
    c := make(chan Result)
    done := make(chan struct{})
    defer close(done)
    searchReplica := func(i int) {
        select {
        case c <- replicas[i](query):
        case <- done:
        }
    }
    for i := range replicas {
        go searchReplica(i)
    }

    return <-c
}
为何在演讲中会包含这些bug？Rob Pike仅仅是不想把演示复杂化。这么作是合理的，但对于Go新手而言，可能会直接使用代码，而不去思考它可能有问题。

高级篇
使用指针接收方法的值的实例
level: advanced
只要值是可取址的，那在这个值上调用指针接收方法是没问题的。换句话说，在某些情况下，你不需要在有一个接收值的方法版本。

然而并不是所有的变量是可取址的。Map的元素就不是。通过interface引用的变量也不是。

package main

import "fmt"

type data struct {  
    name string
}

func (p *data) print() {  
    fmt.Println("name:",p.name)
}

type printer interface {  
    print()
}

func main() {  
    d1 := data{"one"}
    d1.print() //ok

    var in printer = data{"two"} //error
    in.print()

    m := map[string]data {"x":data{"three"}}
    m["x"].print() //error
}
Compile Errors:

/tmp/sandbox017696142/main.go:21: cannot use data literal (type data) as type printer in assignment: data does not implement printer (print method has pointer receiver)
/tmp/sandbox017696142/main.go:25: cannot call pointer method on m["x"] /tmp/sandbox017696142/main.go:25: cannot take the address of m["x"]
更新Map的值
level: advanced
如果你有一个struct值的map，你无法更新单个的struct值。

Fails:

package main

type data struct {  
    name string
}

func main() {  
    m := map[string]data {"x":{"one"}}
    m["x"].name = "two" //error
}
Compile Error:

/tmp/sandbox380452744/main.go:9: cannot assign to m["x"].name
这个操作无效是因为map元素是无法取址的。

而让Go新手更加困惑的是slice元素是可以取址的。

package main

import "fmt"

type data struct {  
    name string
}

func main() {  
    s := []data {{"one"}}
    s[0].name = "two" //ok
    fmt.Println(s)    //prints: [{two}]
}
注意在不久之前，使用编译器之一（gccgo）是可以更新map的元素值的，但这一行为很快就被修复了 :-)它也被认为是Go 1.3的潜在特性。在那时还不是要急需支持的，但依旧在todo list中。

第一个有效的方法是使用一个临时变量。

package main

import "fmt"

type data struct {  
    name string
}

func main() {  
    m := map[string]data {"x":{"one"}}
    r := m["x"]
    r.name = "two"
    m["x"] = r
    fmt.Printf("%v",m) //prints: map[x:{two}]
}
另一个有效的方法是使用指针的map。

package main

import "fmt"

type data struct {  
    name string
}

func main() {  
    m := map[string]*data {"x":{"one"}}
    m["x"].name = "two" //ok
    fmt.Println(m["x"]) //prints: &{two}
}
顺便说下，当你运行下面的代码时会发生什么？

package main

type data struct {  
    name string
}

func main() {  
    m := map[string]*data {"x":{"one"}}
    m["z"].name = "what?" //???
}
"nil" Interfaces和"nil" Interfaces的值
level: advanced
这在Go中是第二最常见的技巧，因为interface虽然看起来像指针，但并不是指针。interface变量仅在类型和值为“nil”时才为“nil”。

interface的类型和值会根据用于创建对应interface变量的类型和值的变化而变化。当你检查一个interface变量是否等于“nil”时，这就会导致未预期的行为。

package main

import "fmt"

func main() {  
    var data *byte
    var in interface{}

    fmt.Println(data,data == nil) //prints: <nil> true
    fmt.Println(in,in == nil)     //prints: <nil> true

    in = data
    fmt.Println(in,in == nil)     //prints: <nil> false
    //'data' is 'nil', but 'in' is not 'nil'
}
当你的函数返回interface时，小心这个陷阱。

Incorrect:

package main

import "fmt"

func main() {  
    doit := func(arg int) interface{} {
        var result *struct{} = nil

        if(arg > 0) {
            result = &struct{}{}
        }

        return result
    }

    if res := doit(-1); res != nil {
        fmt.Println("good result:",res) //prints: good result: <nil>
        //'res' is not 'nil', but its value is 'nil'
    }
}
Works:

package main

import "fmt"

func main() {  
    doit := func(arg int) interface{} {
        var result *struct{} = nil

        if(arg > 0) {
            result = &struct{}{}
        } else {
            return nil //return an explicit 'nil'
        }

        return result
    }

    if res := doit(-1); res != nil {
        fmt.Println("good result:",res)
    } else {
        fmt.Println("bad result (res is nil)") //here as expected
    }
}
栈和堆变量
level: advanced
你并不总是知道变量是分配到栈还是堆上。在C++中，使用 new创建的变量总是在堆上。在Go中，即使是使用 new()或者 make()函数来分配，变量的位置还是由编译器决定。编译器根据变量的大小和“泄露分析”的结果来决定其位置。这也意味着在局部变量上返回引用是没问题的，而这在C或者C++这样的语言中是不行的。

如果你想知道变量分配的位置，在“go build”或“go run”上传入“-m“ gc标志（即， go run -gcflags -m app.go）。

GOMAXPROCS, 并发, 和并行
level: advanced
默认情况下，Go仅使用一个执行上下文/OS线程（在当前的版本）。这个数量可以通过设置 GOMAXPROCS来提高。

一个常见的误解是， GOMAXPROCS表示了CPU的数量，Go将使用这个数量来运行goroutine。而runtime.GOMAXPROCS()函数的文档让人更加的迷茫。 GOMAXPROCS变量描述（https://golang.org/pkg/runtime/）所讨论OS线程的内容比较好。

你可以设置 GOMAXPROCS的数量大于CPU的数量。 GOMAXPROCS的最大值是256。

package main

import (  
    "fmt"
    "runtime"
)

func main() {  
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 1
    fmt.Println(runtime.NumCPU())       //prints: 1 (on play.golang.org)
    runtime.GOMAXPROCS(20)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 20
    runtime.GOMAXPROCS(300)
    fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 256
}
读写操作的重排顺序
level: advanced
Go可能会对某些操作进行重新排序，但它能保证在一个goroutine内的所有行为顺序是不变的。然而，它并不保证多goroutine的执行顺序。

package main

import (  
    "runtime"
    "time"
)

var _ = runtime.GOMAXPROCS(3)

var a, b int

func u1() {  
    a = 1
    b = 2
}

func u2() {  
    a = 3
    b = 4
}

func p() {  
    println(a)
    println(b)
}

func main() {  
    go u1()
    go u2()
    go p()
    time.Sleep(1 * time.Second)
}
如果你多运行几次上面的代码，你可能会发现 a和 b变量有多个不同的组合：

1
2

3
4

0
2

0
0

1
4
a和 b最有趣的组合式是 "02"。这表明 b在 a之前更新了。

如果你需要在多goroutine内放置读写顺序的变化，你将需要使用channel，或者使用"sync"包构建合适的结构体。

优先调度
level: advanced
有可能会出现这种情况，一个无耻的goroutine阻止其他goroutine运行。当你有一个不让调度器运行的 for循环时，这就会发生。

package main

import "fmt"

func main() {  
    done := false

    go func(){
        done = true
    }()

    for !done {
    }
    fmt.Println("done!")
}
for循环并不需要是空的。只要它包含了不会触发调度执行的代码，就会发生这种问题。

调度器会在GC、“go”声明、阻塞channel操作、阻塞系统调用和lock操作后运行。它也会在非内联函数调用后执行。

package main

import "fmt"

func main() {  
    done := false

    go func(){
        done = true
    }()

    for !done {
        fmt.Println("not done!") //not inlined
    }
    fmt.Println("done!")
}
要想知道你在 for循环中调用的函数是否是内联的，你可以在“go build”或“go run”时传入“-m” gc标志（如， go build -gcflags -m）。

另一个选择是显式的唤起调度器。你可以使用“runtime”包中的 Goshed()函数。

package main

import (  
    "fmt"
    "runtime"
)

func main() {  
    done := false

    go func(){
        done = true
    }()

    for !done {
        runtime.Gosched()
    }
    fmt.Println("done!")
}