package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/handler"
	"github.com/jirawan-chuapradit/cards_assignment/middleware"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"github.com/jirawan-chuapradit/cards_assignment/service"
)

func Setup(server models.Server) *gin.Engine {
	fileadapter := server.FileAdapter
	conn := server.DB
	r := gin.Default()
	r.GET("/ping", handler.HealthCheckHandler)

	// service
	authServ := auth.NewAuthService(server.RedisCli)
	// handler
	authHandler := handler.NewAuthHandler(authServ)

	r.POST("/login", authHandler.Login)

	authorized := r.Group("/")

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(authServ)
	authorized.Use(authMiddleware.TokenAuthMiddleware)
	{
		authorized.GET("/test", middleware.Authorize("resource", "read", fileadapter), handler.TestAPI)
		authorized.POST("/logout", authHandler.Logout)
	}
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
	archiveRouter.PUT("/:cardId", cardHandler.Store)

	cardsHistoryRouter := cardRouter.Group("/history")
	cardsHistoryRouter.GET("/:cardId", cardsHistoryHandler.FindHistoryById)

	commentRouter := baseRouter.Group("/comments")
	commentRouter.PUT("", commentHandler.Create)
	commentRouter.PUT("/:commentId", commentHandler.Update)
	commentRouter.DELETE("/:commentId", commentHandler.Delete)

	return r
}
