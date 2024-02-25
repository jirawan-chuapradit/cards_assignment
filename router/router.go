package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/handler"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", handler.HealthCheckHandler)

	baseRouter := r.Group("/api")
	cardsRouter := baseRouter.Group("/cards")

	_ = cardsRouter

	return r
}
