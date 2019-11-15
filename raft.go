/***
https://github.com/wongcony/grpcstudy
grpc use

protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
***/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"time"
)

const (
	FORLLOW = 1
	COMCOS  = 2
	LOADER  = 3
)

type LogEntry struct {
	Term  int
	Index int
	Data  interface{}
}

type Raft struct {
	team     int
	State    int
	Id       int
	Ip       string
	Caapfor  int
	Caapnum  int
	ConmitId int
	LastMsg  int
	SendIn   []int
	nc       []Node
	Log      []LogEntry

	Leader  chan int
	tiHello chan int
}

type Node struct {
	ipaddr string
	id     int
}

func NewNode(id int, ip string) Node {
	var p = Node{id: id, ipaddr: ip}
	return p

}

type HelloQuery struct {
	Log      []LogEntry
	ConmitId int
	LastMsg  int
}

type HelloReply struct {
	Team  int
	Index int
}

type VoteOut struct {
	Team int
	Id   int
}

type VoteIn struct {
	Team int
	Id   int
}

func (p *Raft) RequestVote(out VoteOut, in *VoteIn) error {
	fmt.Println("return vote")
	if out.Team >= p.team {
		p.team = out.Team
		p.State = FORLLOW
		p.tiHello <- 1
		in.Id = out.Id
		p.Caapfor = out.Id
		in.Team = out.Team
	} else {

		//		p.State = FORLLOW
	}
	return nil
}

func (p *Raft) getVote(ip string) {

	client, err := rpc.DialHTTP("tcp", ip)
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	defer client.Close()
	fmt.Println("qing qiu vote")
	out := VoteOut{}
	out.Id = p.Id
	out.Team = p.team
	in := VoteIn{}

	client.Call("Raft.RequestVote", out, &in)

	if in.Team < p.team {

		if p.Caapfor == in.Id {
			if p.State != LOADER {
				p.State = LOADER
				p.Leader <- 1
				fmt.Println("this node is lead node")
			}
		}
	} else {
		p.State = FORLLOW

		fmt.Println("this node is forllow node")
	}

}

func (p *Raft) sendhellomsg(ip string, q HelloQuery, r *HelloReply) {
	client, err := rpc.DialHTTP("tcp", ip)
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	defer client.Close()
	fmt.Println("sendhellomsg send")
	client.Call("Raft.Recvmsg", q, r)

	// ji lu dang qian jie dian commit index

}

func (p *Raft) Recvmsg(q HelloQuery, r *HelloReply) error {
	p.tiHello <- 1
	fmt.Println("sendhellomsg recv")
	//tong bu xiaoxi
	return nil

}

func (p *Raft) createNode(id int, ipaddr string) {
	n1 := NewNode(id, ipaddr)
	p.nc = append(p.nc, n1)

}

func (p *Raft) showNode() {
	for i, k := range p.nc {
		fmt.Println(i, k)
	}
}

func (p *Raft) start() {
	p.tiHello = make(chan int, 1)
	p.Leader = make(chan int, 1)
	p.State = FORLLOW
	rpc.Register(p)

	rpc.HandleHTTP()
	go func() {
		http.ListenAndServe(p.Ip, nil)
	}()
}

func (p *Raft) run() {
	for {
		if p.State == FORLLOW {
			//shoubao
			select {
			case <-p.tiHello:
				fmt.Println("FORLLOW state")

			case <-time.After(5 * time.Second):
				fmt.Println("sleep 5 miao")
				p.State = COMCOS
			}
		} else if p.State == COMCOS {

			p.team++
			for _, k := range p.nc {
				go p.getVote(k.ipaddr)
			}
			select {
			case <-time.After(time.Millisecond * 3000):
				p.State = FORLLOW
				fmt.Println("jiang")
			case <-p.Leader:
				p.State = LOADER
				go func() {
					for i := 0; i < 30; i++ {
						p.Log = append(p.Log, LogEntry{Term: p.team, Index: i})
						time.Sleep(time.Second * 3)
					}
				}()
			}

		} else if p.State == LOADER {

			select {
			case <-time.After(time.Millisecond * 3000):
				out := HelloQuery{}
				in := HelloReply{}
				for _, k := range p.nc {
					go p.sendhellomsg(k.ipaddr, out, &in)
				}

				if in.Index > p.Id {
					p.State = FORLLOW
				}
				fmt.Println("jiang")
			}

		}

	}
}

func main() {
	//s1 := rpc.NewServer()
	if len(os.Args) < 2 {
		return
	}
	nodeid, _ := strconv.Atoi(os.Args[1])

	ser := &Raft{}
	ser.start()
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(300))
	ser.Id = nodeid
	if nodeid == 1 {
		ser.Ip = ":3001"
		ser.createNode(2, ":3002")
	} else {
		ser.Ip = ":3002"
		ser.createNode(1, ":3001")
	}

	ser.showNode()
	ser.run()
	select {}

}
