package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Age      int
	Password string
	Email    string
}

type EmailValidationError struct {
	Reason string
}

func (e EmailValidationError) Error() string {
	return fmt.Sprintf("invalid email: %s", e.Reason)
}

type NameValidationError struct {
	Reason string
}

func (s NameValidationError) Error() string{
	return fmt.Sprintf("invalid name: %s", s.Reason)
}

type AgeValidationError struct {
	Age int
}

func (s AgeValidationError) Error() string{
	return fmt.Sprintf("invalid age: %d", s.Age)
}

type PasswordValidationError struct {
	Password	int
}

func (s PasswordValidationError) Error() string{
	return fmt.Sprintf("invalid password: %d", s.Password)
}


func ValidateUser(u User) error {
	if u.Name == "" {
		return errors.New(NameValidationError(u.Name))
	}
	if u.Age < 18 {
		return 
	}
	if len(u.Password) >= 6 {
		return 
	}
	if strings.Contains(u.Email, "@"){
		return 
	}
	return nil
}


func main() {
	fmt.Print("\033[H\033[2J")



}