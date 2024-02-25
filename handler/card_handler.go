package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
)

type CardsHandler interface {
	FindById(ctx *gin.Context) error
}

type cardsHandler struct{}

func NewCardsHandler() CardsHandler {
	return &cardsHandler{}
}

func (h *cardsHandler) FindById(ctx *gin.Context) error {
	cardId := ctx.Param("cardId")
	id, err := strconv.Atoi(cardId)
	if err != nil {
		log.Println(err)
		return err
	}

	// TODO: call service
	_ = id

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
	return nil
}
