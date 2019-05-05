package services

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/repositories"
)

type UserService interface {
	IsLoggedIn(getUserID func() (string, error)) bool
	GetID(getUserID func() (string, error)) (string, error)
	SignUpAndLogin(nationality string, birth int, gender string, setUserID func(string) error) (string, bool)
	TryLogin(userID string, setUserID func(string) error) bool
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userService{userRepo}
}

func (us userService) IsLoggedIn(getUserID func() (string, error)) bool {
	userID, err := getUserID()

	return err == nil && us.userRepo.IsPresent(userID)
}

func (us userService) GetID(getUserID func() (string, error)) (string, error) {
	userID, err := getUserID()

	if err != nil || !us.userRepo.IsPresent(userID) {
		return "", fmt.Errorf("user is not logged in")
	}

	return userID, nil
}

func (us userService) SignUpAndLogin(nationality string, birth int, gender string, setUserID func(string) error) (string, bool) {
	userID := us.userRepo.Add(nationality, birth, gender)

	cookieErr := setUserID(userID)
	if cookieErr != nil {
		// ignore remove error because rare and it doesn't harm.
		_ = us.userRepo.Remove(userID)
		return "", false
	}

	return userID, true
}

func (us userService) TryLogin(userID string, setUserID func(string) error) bool {
	if us.userRepo.IsPresent(userID) {
		err := setUserID(userID)
		return err == nil
	}
	return false
}
