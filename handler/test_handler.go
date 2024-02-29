package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
)

func TestAPI(c *gin.Context) {
	auth := auth.TokenManager{}
	metadata, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   metadata.UserId,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
