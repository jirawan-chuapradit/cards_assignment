package models

import (
	"errors"
	"log"
	"strconv"
)

var us = []Login{
	{
		ID:       "2",
		UserName: "users",
		Password: "pass",
	}, {
		ID:       "3",
		UserName: "username",
		Password: "password",
	},
}
var UserRepo = UserRepository{
	Users: us,
}

type UserRepository struct {
	Users []Login
}

func (r *UserRepository) FindAll() ([]Login, error) {
	return r.Users, nil
}

func (r *UserRepository) FindByID(id int) (Login, error) {

	for _, v := range r.Users {
		uid, err := strconv.Atoi(v.ID)
		if err != nil {
			return Login{}, err
		}
		if uid == int(id) {
			return v, nil
		}
	}

	return Login{}, errors.New("Not found")
}

func (r *UserRepository) Save(user Login) (Login, error) {
	r.Users = append(r.Users, user)

	return user, nil
}

func (r *UserRepository) Delete(user Login) {
	id := -1
	for i, v := range r.Users {
		if v.ID == user.ID {
			id = i
			break
		}
	}

	if id == -1 {
		log.Fatal("Not found user ")
		return
	}

	r.Users[id] = r.Users[len(r.Users)-1] // Copy last element to index i.
	r.Users[len(r.Users)-1] = Login{}     // Erase last element (write zero value).
	r.Users = r.Users[:len(r.Users)-1]    // Truncate slice.

	return
}
