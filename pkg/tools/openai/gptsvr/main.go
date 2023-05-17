// Package main
// Create  2023-03-25 00:54:41
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: 监听HTTP 响应.

type app struct {
	apiSvr *http.Server
	cancel context.CancelFunc
	close  chan error
}

// init 初始化svr.
func (a *app) init(ctx context.Context) {
	// 监听信号量: 及时退出服务.
}

// run 运行svr.
func (a *app) run(ctx context.Context) {

	// 监听指定信号：ctrl+c以及kill.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		select {
		case r := <-sigs:
			a.close <- fmt.Errorf("syscall:[%+v]\n", r)
		}
	}()
	if err := a.apiSvr.ListenAndServe(); err != nil {
		a.close <- err
	}
}

func (a *app) quit() {
	select {
	case err := <-a.close:
		fmt.Println("svr done because err:" + err.Error())
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	r := gin.New()
	r.GET("/openai", QueryOpenAI)
	app := &app{
		apiSvr: &http.Server{
			Addr:              ":8888",
			Handler:           r,
			ReadHeaderTimeout: 5 * time.Minute,
			WriteTimeout:      5 * time.Minute,
		},
		cancel: cancel,
		close:  make(chan error, 1), // 容量为1不阻塞.
	}

	go func() {
		app.init(ctx)
		app.run(ctx)
	}()
	app.quit()
}
