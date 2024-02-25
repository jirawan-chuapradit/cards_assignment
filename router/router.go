package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/handler"
)

func Setup() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", handler.HealthCheckHandler)

	baseRouter := r.Group("/api")
	cardsRouter := baseRouter.Group("/cards")

	_ = cardsRouter
	cardHandler := handler.NewCardsHandler()
	cardsRouter.GET("/:cardId", cardHandler.FindById)
	return r
}
