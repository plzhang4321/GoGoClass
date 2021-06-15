package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")

}
func StartHttp(ser *http.Server) error {
	http.HandleFunc("/hello", helloServer)
	fmt.Println("start server")
	err := ser.ListenAndServe()
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, err := errgroup.WithContext(ctx)
	ser := &http.Server{Addr: "8080"}
	g.Go(func() error {
		return StartHttp(ser)
	})
	//处理linux信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	//还没有上下文
	g.Go(func() error {
		<-err.Done()
		fmt.Println("server stop")
		return ser.Shutdown(err)
	})

	g.Go(func() error {
		for {
			select {
			case <-err.Done():
				return err.Err() //全部关闭
			case <-ch:
				cancel()
			}
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("error: ", err)
	}
}
