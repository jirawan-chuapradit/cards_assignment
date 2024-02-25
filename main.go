package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/handler"
)

func main() {
	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/ping", handler.HealthCheckHandler)
	r.Run()
}
