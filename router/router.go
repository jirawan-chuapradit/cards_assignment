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

	// repository
	cardsRepository := repository.NewCardsRepository(conn)
	cardsHistoryRepository := repository.NewCardsHistoryRepository(conn)

	// service
	cardsServ := service.NewCardsService(cardsRepository)
	cardsHistoryServ := service.NewCardsHistoryService(cardsHistoryRepository)
	commentServ := service.NewCommentService(cardsRepository)

	// handler
	cardHandler := handler.NewCardsHandler(cardsServ)
	cardsHistoryHandler := handler.NewCardsHistoryHandler(cardsHistoryServ)
	commentHandler := handler.NewCommentHandler(commentServ)

	cardRouter := baseRouter.Group("/cards")
	cardRouter.GET("", cardHandler.FindAll)
	cardRouter.GET("/:cardId", cardHandler.FindById)
	cardRouter.POST("", cardHandler.Create)
	cardRouter.PUT("/:cardId", cardHandler.Update)

	archiveRouter := cardRouter.Group("/archive")
	archiveRouter.POST("/:cardId", cardHandler.Store)

	cardsHistoryRouter := cardRouter.Group("/history")
	cardsHistoryRouter.GET("/:cardId", cardsHistoryHandler.FindHistoryById)

	commentRouter := baseRouter.Group("/comments")
	commentRouter.POST("", commentHandler.Create)              // TODO: create comment
	commentRouter.PUT("/:commentId", commentHandler.Update)    // TODO: edit comment
	commentRouter.DELETE("/:commentId", commentHandler.Delete) // TODO: delete comment

	return r
}
