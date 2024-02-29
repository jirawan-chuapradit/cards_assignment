package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	tokenManager auth.TokenInterface
	authServ     auth.AuthInterface
}

func NewAuthHandler(authServ auth.AuthInterface) AuthHandler {
	return &authHandler{
		authServ:     authServ,
		tokenManager: auth.NewTokenService(),
	}
}

func (h *authHandler) Login(c *gin.Context) {
	var u models.User
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

	auth := auth.TokenManager{}
	ts, err := auth.CreateToken(user.ID, user.UserName)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := h.authServ.CreateAuth(user.ID, ts)
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
		deleteErr := h.authServ.DeleteTokens(metadata)
		if deleteErr != nil {
			c.JSON(http.StatusBadRequest, deleteErr.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
