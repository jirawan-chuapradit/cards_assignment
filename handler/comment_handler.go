package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
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
	tokenManager auth.TokenInterface
	commentServ  service.CommentService
}

func NewCommentHandler(commentServ service.CommentService) CommentHandler {
	return &commentHandler{
		tokenManager: auth.NewTokenService(),
		commentServ:  commentServ,
	}
}

func (h *commentHandler) Create(ctx *gin.Context) {
	createCommentRequest := request.CreateCommentBody{}
	if err := ctx.ShouldBindJSON(&createCommentRequest); err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "invalid request",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
	}
	if err := h.commentServ.Create(ctx, createCommentRequest); err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   "created comment successfully",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (h *commentHandler) Update(ctx *gin.Context) {
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

	updateCommentRequest := request.UpdateCommentBody{}
	if err := ctx.ShouldBindJSON(&updateCommentRequest); err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "invalid request",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	if updateCommentRequest.ID != objID {
		log.Printf("expected comment %s, got %s", updateCommentRequest.ID, objID)
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Failed",
			Data:   "can not update comment invalid comment id",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
	}

	metadata, err := h.tokenManager.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not update comment because internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}
	now := time.Now().In(config.Location)
	updateCommentRequest.UpdatedBy = metadata.UserName
	updateCommentRequest.UpdatedAt = &now

	if err := h.commentServ.Update(ctx, updateCommentRequest); err != nil {
		log.Println(err)
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not update comment because internal server error",
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

func (h *commentHandler) Delete(ctx *gin.Context) {
	commentId := ctx.Param("commentId")
	objID, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
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
	metadata, err := h.tokenManager.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "internal server error",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	if err := h.commentServ.Delete(ctx, objID, metadata.UserName); err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "internal server error",
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
