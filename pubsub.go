/*
charu paixufa
redis fuwuqi
*/
package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

type RedisMap struct {
	sync.Mutex
	Data map[string]interface{}
}

func (r *RedisMap) Write(key string, v interface{}) {
	r.Lock()
	defer r.Unlock()
	r.Data[key] = v

}

func (r *RedisMap) Get(key string) interface{} {
	r.Lock()
	defer r.Unlock()
	return r.Data[key]

}

type HashMap struct {
	sync.Mutex
	Data map[interface{}]interface{}
}

func (r *HashMap) Write(key string, v interface{}) {
	r.Lock()
	defer r.Unlock()
	b1 := sha256.Sum256([]byte(key))
	r.Data[b1] = v

}

func (r *HashMap) Get(key string) interface{} {
	r.Lock()
	defer r.Unlock()
	b1 := sha256.Sum256([]byte(key))
	return r.Data[b1]

}

type Message struct {
	Id   int
	Data []byte
}

type PubSub interface {
	Done()
	Publish(msg Message) bool
	Subscribe(client QClient) error
	Unsubscribe(qClient QClient)
	SubscriberCount() int
}

type QClient interface {
	Id() string
	Notify(message Message) error
}

type Msglist struct {
	sync.Mutex
	Data    interface{}
	TimeOut int
}

type play int

type job struct {
	sum play
}

func (r *job) doing(key string) {
	fmt.Println("do some thing")
	time.Sleep(time.Second * 5)

}

var jobpool = make(chan job, 1)

var workerpool = make(chan chan job, 4)

type worker struct {
	joblist chan job
}

func (r *worker) ganhuo(key string) {
	fmt.Println("do some thing")
	time.Sleep(time.Second * 2)

}

func (r *worker) run(key string) {

	for {
		workerpool <- r.joblist
		select {
		case p := <-r.joblist:
			p.doing("jiang exe")

			fmt.Println("zhixing")
		}
	}
}

func createWorker() {
	for i := 0; i < 4; i++ {
		work := &worker{joblist: make(chan job, 1)}
		go func(work *worker) {
			work.run("ok")
		}(work)

	}
}

func createJob() {

	for i := 0; i < 10; i++ {
		p := job{sum: play(i)}
		jobpool <- p

	}

}

func dispach() {
	for {
		select {
		case work := <-jobpool:

			go func(job job) {
				p := <-workerpool
				p <- job

			}(work)

			fmt.Println("zhixing")
		}
	}
}

func main() {
	var s1 = &RedisMap{}
	s1.Data = make(map[string]interface{}, 100)
	s1.Write("jiang", "ok")
	s1.Write("bobo", []byte("ok"))
	createWorker()
	go dispach()
	createJob()
	time.Sleep(time.Second * 100)

}
