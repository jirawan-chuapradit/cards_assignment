package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/twinj/uuid"
)

type TokenInterface interface {
	CreateToken(userId, userName string) (*models.TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*models.AccessDetails, error)
}

type TokenManager struct{}

func NewTokenService() *TokenManager {
	return &TokenManager{}
}

func (t *TokenManager) CreateToken(userId, userName string) (*models.TokenDetails, error) {
	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()
	td.TokenUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	claims := jwt.MapClaims{}
	claims["access_uuid"] = td.TokenUuid
	claims["user_id"] = userId
	claims["user_name"] = userName
	claims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	td.AccessToken, err = at.SignedString([]byte(config.AccessSecretKey))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userId

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["user_name"] = userName
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(config.RefreshSecretKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t *TokenManager) ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
func extract(token *jwt.Token) (*models.AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, userOk := claims["user_id"].(string)
		userName, userNameOk := claims["user_name"].(string)
		if !ok || !userOk || !userNameOk {
			return nil, errors.New("unauthorized")
		} else {
			return &models.AccessDetails{
				TokenUuid: accessUuid,
				UserId:    userId,
				UserName:  userName,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AccessSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	return nil
}
