package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/handler"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"github.com/jirawan-chuapradit/cards_assignment/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(conn *mongo.Client) *gin.Engine {

	r := gin.Default()

	r.GET("/ping", handler.HealthCheckHandler)

	baseRouter := r.Group("/api")
	cardsRouter := baseRouter.Group("/cards")

	// repository
	cardsRepository := repository.NewCardsRepository(conn)

	// service
	cardsServ := service.NewCardsService(cardsRepository)

	// handler
	cardHandler := handler.NewCardsHandler(cardsServ)
	cardsRouter.GET("/:cardId", cardHandler.FindById)
	return r
}
