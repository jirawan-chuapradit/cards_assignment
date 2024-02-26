package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
	"github.com/jirawan-chuapradit/cards_assignment/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsHandler interface {
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type cardsHandler struct {
	cardsService service.CardsService
}

func NewCardsHandler(cardsServ service.CardsService) CardsHandler {
	return &cardsHandler{
		cardsService: cardsServ,
	}
}

func (h *cardsHandler) FindById(ctx *gin.Context) {
	cardId := ctx.Param("cardId")
	objID, err := primitive.ObjectIDFromHex(cardId)
	if err != nil { // TODO: handle
		log.Println(err)
		return
	}
	cardDetails, err := h.cardsService.FindById(ctx, objID)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			webResponse := response.Response{
				Code:   http.StatusBadRequest,
				Status: "Failed",
				Data:   "card not found",
			}

			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusBadRequest, webResponse)
			return
		}

		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not find card because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   cardDetails,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *cardsHandler) FindAll(ctx *gin.Context) {
	cards, err := h.cardsService.FindAll(ctx)
	if err != nil {
		log.Println(err)

		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not find cards because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   cards,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
