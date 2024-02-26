package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
	"github.com/jirawan-chuapradit/cards_assignment/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardsHandler interface {
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
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

func (h *cardsHandler) Create(ctx *gin.Context) {
	createCardRequest := request.CreateCardRequestBody{}
	err := ctx.ShouldBindJSON(&createCardRequest)
	if err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "can not create cards because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	card, err := h.cardsService.Create(ctx, createCardRequest)
	if err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not create cards because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   card,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)

}

func (h *cardsHandler) Update(ctx *gin.Context) {
	updateCardRequest := request.UpdateCardRequestBody{}
	err := ctx.ShouldBindJSON(&updateCardRequest)
	if err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "can not update cards because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	cardId := ctx.Param("cardId")
	objID, err := primitive.ObjectIDFromHex(cardId)
	if err != nil { // TODO: handle
		log.Println(err)
		return
	}

	// validate
	if updateCardRequest.ID != objID {
		log.Printf("expected card %s, got %s", updateCardRequest.ID.Hex(), objID.Hex())
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Ok",
			Data:   "invalid request",
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	// service
	if err := h.cardsService.Update(ctx, updateCardRequest); err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not update card because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "update successfully",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
