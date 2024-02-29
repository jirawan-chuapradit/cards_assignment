package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	"github.com/jirawan-chuapradit/cards_assignment/models/request"
	"github.com/jirawan-chuapradit/cards_assignment/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	CreateUser(ctx context.Context, userReq request.SignUp) error
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

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
