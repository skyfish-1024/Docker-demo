package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Golang! Hello Docker!")
	})
	err := r.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
