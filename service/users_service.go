package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	CreateUser(ctx context.Context, userReq request.SignUp) error
	ValidateUser(ctx context.Context, userReq request.Login) (models.User, error)
}

type usersService struct {
	usersRepository repository.UsersRepository
}

func NewUsersService(usersRepository repository.UsersRepository) UsersService {
	return &usersService{
		usersRepository: usersRepository,
	}
}

func (s *usersService) CreateUser(ctx context.Context, userReq request.SignUp) error {
	hash, err := hashAndSalt([]byte(userReq.Password))
	if err != nil {
		return err
	}
	userReq.Password = hash
	rb, err := json.Marshal(userReq)
	if err != nil {
		return err
	}

	user := models.User{}
	if err := json.Unmarshal(rb, &user); err != nil {
		return err
	}
	now := time.Now().In(config.Location)
	user.CreatedAt = now
	user.UpdatedAt = now
	return s.usersRepository.Create(ctx, user)
}

func (s *usersService) ValidateUser(ctx context.Context, userReq request.Login) (models.User, error) {
	user, err := s.usersRepository.FindByEmail(ctx, userReq.Email)
	if err != nil {
		return user, err
	}
	if ok := comparePasswords(user.Password, userReq.Password); !ok {
		return user, errors.New("invalid user and password")
	}
	return user, nil
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
