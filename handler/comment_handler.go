package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
	"github.com/jirawan-chuapradit/cards_assignment/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type commentHandler struct {
	commentServ service.CommentService
}

func NewCommentHandler(commentServ service.CommentService) CommentHandler {
	return &commentHandler{
		commentServ: commentServ,
	}
}

func (h *commentHandler) Create(ctx *gin.Context) {
	return
}

func (h *commentHandler) Update(ctx *gin.Context) {
	return
}

func (h *commentHandler) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	objID, err := primitive.ObjectIDFromHex(commentId)
	if err != nil { // TODO: handle
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "invalid request",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}
	if err := h.commentServ.Delete(ctx, objID); err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not delete comment because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "delete comment successfully",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
