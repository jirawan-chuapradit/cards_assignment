package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin"

	"github.com/casbin/casbin/persist"
	"github.com/gin-gonic/gin"
	"github.com/jirawan-chuapradit/cards_assignment/auth"
	"github.com/jirawan-chuapradit/cards_assignment/models/response"
)

type AuthMiddleware interface {
	TokenAuthMiddleware(c *gin.Context)
}

type authMiddleware struct {
	tokenManager auth.TokenInterface
	authServ     auth.AuthInterface
}

func NewAuthMiddleware(authService auth.AuthInterface) AuthMiddleware {
	return &authMiddleware{
		authServ:     authService,
		tokenManager: auth.NewTokenService(),
	}
}
func (m *authMiddleware) TokenAuthMiddleware(c *gin.Context) {
	if err := auth.TokenValid(c.Request); err != nil {
		webResponse := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Failed",
			Data:   "unauthorized",
		}

		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
	}

	metadata, err := m.tokenManager.ExtractTokenMetadata(c.Request)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Failed",
			Data:   "unauthorized",
		}

		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
	}
	if _, err := m.authServ.FetchAuth(metadata.TokenUuid); err != nil {
		webResponse := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Failed",
			Data:   "unauthorized",
		}

		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
	}
	fmt.Println("token valid")

	c.Next()
}

func Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.TokenValid(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
			c.Abort()
			return
		}

		auth := auth.TokenManager{}
		metadata, err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		// casbin enforces policy
		fmt.Println("metadata.UserName: ", metadata.UserName, obj, act, adapter)
		ok, err := enforce(metadata.UserName, obj, act, adapter)
		//ok, err := enforce(val.(string), obj, act, adapter)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, "error occurred when authorizing user")
			return
		}
		if !ok {
			fmt.Println("forbidden !!")
			c.AbortWithStatusJSON(403, "forbidden")
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	err := enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok := enforcer.Enforce(sub, obj, act)
	return ok, nil
}
