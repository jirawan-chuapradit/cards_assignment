package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/jirawan-chuapradit/cards_assignment/models"
)

type AuthInterface interface {
	CreateAuth(string, *models.TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*models.AccessDetails) error
}

type AuthService struct {
	client *redis.Client
}

func NewAuthService(client *redis.Client) *AuthService {
	return &AuthService{client: client}
}

func (tk *AuthService) CreateAuth(userId string, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(td.TokenUuid, userId, at.Sub(now)).Result() // token exp redis key will removed
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(td.RefreshUuid, userId, rt.Sub(now)).Result() // token exp redis key will removed
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

func (tk *AuthService) FetchAuth(tokenUuid string) (string, error) {
	userid, err := tk.client.Get(tokenUuid).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func (tk *AuthService) DeleteTokens(authD *models.AccessDetails) error {
	//delete access token
	deletedAt, err := tk.client.Del(authD.TokenUuid).Result()
	if err != nil {
		return err
	}

	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UserId)

	//delete refresh token
	deletedRt, err := tk.client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}

	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *AuthService) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	deleted, err := tk.client.Del(refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
