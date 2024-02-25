package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
	"github.com/jirawan-chuapradit/cards_assignment/service"
)

type CardsHandler interface {
	FindById(ctx *gin.Context)
}

type cardsHandler struct {
	cardsService service.CardsService
}

func NewCardsHandler() CardsHandler {
	cardsServ := service.NewCardsService()
	return &cardsHandler{
		cardsService: cardsServ,
	}
}

func (h *cardsHandler) FindById(ctx *gin.Context) {
	cardId := ctx.Param("cardId")
	id, err := strconv.Atoi(cardId)
	if err != nil { // TODO: handle
		log.Println(err)
		return
	}

	cardDetails, err := h.cardsService.FindById(id)
	if err != nil {
		log.Println(err)
		return
	}
	_ = cardDetails

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: map[string]interface{}{
			"id":          1,
			"title":       "นัดสัมภาษณ์",
			"description": "mock",
			"created_at":  "",
			"status":      "TODO",
		},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
