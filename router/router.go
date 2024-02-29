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

	// repository
	userRepo := repository.NewUsersRepository(conn)
	cardsRepository := repository.NewCardsRepository(conn)
	cardsHistoryRepository := repository.NewCardsHistoryRepository(conn)
	// service
	authServ := auth.NewAuthService(server.RedisCli)
	userServ := service.NewUsersService(userRepo)
	cardsServ := service.NewCardsService(cardsRepository)
	cardsHistoryServ := service.NewCardsHistoryService(cardsHistoryRepository)
	commentServ := service.NewCommentService(cardsRepository)
	// handler
	authHandler := handler.NewAuthHandler(authServ, userServ)
	cardHandler := handler.NewCardsHandler(cardsServ)
	cardsHistoryHandler := handler.NewCardsHistoryHandler(cardsHistoryServ)
	commentHandler := handler.NewCommentHandler(commentServ)

	r.POST("/login", authHandler.Login)
	r.POST("/signup", authHandler.SignUp)

	authorized := r.Group("/")

	// middleware
	authMiddleware := middleware.NewAuthMiddleware(authServ)
	authorized.Use(authMiddleware.TokenAuthMiddleware)
	{
		authorized.GET("/test", middleware.Authorize("resource", "read", fileadapter), handler.TestAPI)
		authorized.POST("/logout", authHandler.Logout)

		// cards
		authorized.GET("/api/cards", middleware.Authorize("cards", "read", fileadapter), cardHandler.FindAll)
		authorized.GET("/api/cards/:cardId", middleware.Authorize("cards", "read", fileadapter), cardHandler.FindById)
		authorized.POST("/api/cards", middleware.Authorize("cards", "write", fileadapter), cardHandler.Create)
		authorized.PUT("/api/cards/:cardId", middleware.Authorize("cards", "write", fileadapter), cardHandler.Update)
		authorized.PUT("/api/cards/archive/:cardId", middleware.Authorize("cards", "write", fileadapter), cardHandler.Store)
		authorized.GET("/api/cards/history/:cardId", middleware.Authorize("cards", "read", fileadapter), cardsHistoryHandler.FindHistoryById)

		// comments
		authorized.PUT("/api/comments", middleware.Authorize("comments", "write", fileadapter), commentHandler.Create)
		authorized.PUT("/api/comments/:commentId", middleware.Authorize("comments", "write", fileadapter), commentHandler.Update)
		authorized.DELETE("/api/comments/:commentId", middleware.Authorize("comments", "delete", fileadapter), commentHandler.Delete)
	}

	return r
}
