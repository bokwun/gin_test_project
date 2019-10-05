package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10*time.Second)
		c.String(200, "hello world!")

	})

	//handler选择gin的实例
	//实例化一个结构体
	srv := &http.Server{
		Addr:":8085",
		Handler:r,
	}
	//调用shutdown或者close之后，errserverclosed由listenandserver函数返回
	//把服务器放到goroutine里执行
	go func() {
													//手动调用shutdown时会产生这个错误
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	//退出信号捕获 ctrl+c == SIGINT  服务器在后台运行，使用kill $pid后会产生SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//channel阻塞，等待退出信号,读取值然后忽略
	<-quit
	log.Println("shutdown server....")

	//关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}

	log.Println("server exit")



}
