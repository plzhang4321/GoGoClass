package main

import (
	"container/ring"
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	limitedFreqSecond int32 = 10
	CurCount          int32 = 0
	Bucket            int   = 6
	head              *ring.Ring
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9090") //获取一个tcpAddr
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr) //监听一个端口
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	// init head for store
	head := ring.New(Bucket)
	for i := 0; i < Bucket; i++ { //初始化为0 ring的默认初始化是nil
		head.Value = 0
		head = head.Next()

	}
	go func() {
		timer := time.NewTicker(time.Second * 1)
		for range timer.C {
			atomic.AddInt32(&CurCount, int32(-head.Value.(int)))
			head.Value = 0
			head = head.Next()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(&conn)
	}
}

func handle(conn *net.Conn) {
	defer (*conn).Close()
	n := atomic.AddInt32(&CurCount, 1)
	if n > limitedFreqSecond {
		(*conn).Write([]byte("too many request"))
	} else {
		mu := sync.Mutex{}
		mu.Lock()
		value := head.Value.(int)
		head.Value = value + 1
		mu.Unlock()
		// do something
		(*conn).Write([]byte("Congratulation"))
	}

}
