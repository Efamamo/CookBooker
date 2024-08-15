package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Count struct {
	Count int
}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Static("/styles", "./styles")

	r.LoadHTMLGlob("views/*")
	count := Count{Count: 0}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", count)
	})

	r.Run()
}
