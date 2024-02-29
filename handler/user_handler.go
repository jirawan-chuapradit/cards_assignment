package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
	"github.com/jirawan-chuapradit/cards_assignment/service"
)

type AuthHandler interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	tokenManager auth.TokenInterface
	auth         auth.AuthInterface
	userService  service.UsersService
}

func NewAuthHandler(authServ auth.AuthInterface, userServ service.UsersService) AuthHandler {
	return &authHandler{
		userService:  userServ,
		auth:         authServ,
		tokenManager: auth.NewTokenService(),
	}
}
func (h *authHandler) SignUp(c *gin.Context) {
	var u request.SignUp
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	if err := h.userService.CreateUser(c, u); err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Failed",
			Data:   "can not find history because internal server error",
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Successfully sign up",
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
func (h *authHandler) Login(c *gin.Context) {
	var u models.Login
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//find user with username
	user, _ := models.UserRepo.FindByID(2)

	//compare the user from the request, with the one we defined:
	if user.UserName != u.UserName || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	ts, err := h.tokenManager.CreateToken(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := h.auth.CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: map[string]interface{}{
			"tokens": tokens,
		},
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (h *authHandler) Logout(c *gin.Context) {
	//If metadata is passed and the tokens valid, delete them from the redis store
	metadata, _ := h.tokenManager.ExtractTokenMetadata(c.Request)
	if metadata != nil {
		deleteErr := h.auth.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
