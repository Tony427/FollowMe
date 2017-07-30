package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.LoadHTMLGlob("views/*")
	app.Static("/public", "./public")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"status": "true"})
	})

	ip := "127.0.0.1"
	port := "9000"

	if port == "" {
		port = "9000"
	}

	bind := fmt.Sprintf("%s:%s", ip, port)

	app.Run(bind)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", runtime.Version())
}
