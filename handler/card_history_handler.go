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

type CardsHistoryHandler interface {
	FindHistoryById(ctx *gin.Context)
}

type cardsHistoryHandler struct {
	cardsHistoryService service.CardsHistoryService
}

func NewCardsHistoryHandler(cardsHistoryServ service.CardsHistoryService) CardsHistoryHandler {
	return &cardsHistoryHandler{
		cardsHistoryService: cardsHistoryServ,
	}
}

func (h *cardsHistoryHandler) FindHistoryById(ctx *gin.Context) {
	cardId := ctx.Param("cardId")
	log.Println(cardId)
	objID, err := primitive.ObjectIDFromHex(cardId)
	if err != nil { // TODO: handle
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "invalid card id",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	histories, err := h.cardsHistoryService.FindHistoryById(ctx, objID)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			webResponse := response.Response{
				Code:   http.StatusBadRequest,
				Status: "Failed",
				Data:   "history not found",
			}

			ctx.Header("Content-Type", "application/json")
			ctx.JSON(http.StatusBadRequest, webResponse)
			return
		}

		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not find history because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   histories,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
