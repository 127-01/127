package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()package main

import (
	"context"
	"douyin/cache"
	"douyin/model"
	"douyin/router"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	//初始化
	model.InitDB()
	cache.InitCache()

	cache.InitCache()

	service.InitAllServiceMgr()

	//初始化router
	router.InitAllRouters(r)

	//r.Run()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
	
	gracefulExitWeb(server)
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer cache.OnSeverClose()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}
}


	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
