package main

import (
	"errors"
	"fmt"
)

type User struct {
	Login    string
	Password string
}

type UserStore interface {
	FindByLogin(login string)(User, error)
	AddUser(user User) error
}

func login(store UserStore, login, password string) bool{
	user, err := store.FindByLogin(login)
	if err != nil {
		return false
	}
	return password == user.Password
}

type FakeStore struct {
		users map[string]User // приватное поле
}

func NewFakeStore() *FakeStore {
	return &FakeStore{
		users: make(map[string]User),
	}
}

func (s *FakeStore) AddUser(user User) error {
	if _, exists := s.users[user.Login]; exists {
		return errors.New("User already exists")
	}
	s.users[user.Login] = user
	return nil
}

func (s *FakeStore) FindByLogin(login string) (User, error){
	user, exists := s.users[login]
	if exists {
		return user, nil
	}
	return User{}, errors.New("user no found")
}

func main() {
	fmt.Print("\033[H\033[2J")

	store := NewFakeStore()
	store.AddUser(User{
		Login: "admin",
		Password: "123",
	})

	fmt.Println(login(store, "admin", "123"))
	fmt.Println(login(store, "admin", "11223"))
	


}