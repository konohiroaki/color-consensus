package services

import (
	"fmt"
	"github.com/konohiroaki/color-consensus/backend/repositories"
	"sync"
)

type UserService interface {
	IsLoggedIn(getUserID func() (string, error)) bool
	GetID(getUserID func() (string, error)) (string, error)
	SignUpAndLogin(nationality string, birth int, gender string, setUserID func(string) error) (string, error)
	TryLogin(userID string, setUserID func(string) error) bool
}

type userService struct {
	userRepo   repositories.UserRepository
	nationRepo repositories.NationalityRepository
	genderRepo repositories.GenderRepository
}

var (
	userServiceInstance UserService
	userServiceOnce     sync.Once
)

func GetUserService(env string) UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = newUserService(repositories.GetUserRepository(env), repositories.GetNationalityRepository(), repositories.GetGenderRepository())
	})
	return userServiceInstance
}

func newUserService(userRepo repositories.UserRepository, nationRepo repositories.NationalityRepository, genderRepo repositories.GenderRepository) UserService {
	return userService{userRepo, nationRepo, genderRepo}
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

func (us userService) SignUpAndLogin(nationality string, birth int, gender string, setUserID func(string) error) (string, error) {
	if !us.genderRepo.IsPresent(gender) {
		return "", NewValidationError("gender format is not correct")
	}
	if !us.nationRepo.IsCodePresent(nationality) {
		return "", NewValidationError("nationality format is not correct")
	}

	userID := us.userRepo.Add(nationality, birth, gender)

	cookieErr := setUserID(userID)
	if cookieErr != nil {
		// ignore remove error because rare and it doesn't harm.
		_ = us.userRepo.Remove(userID)
		return "", NewInternalServerError("internal server error")
	}

	return userID, nil
}

func (us userService) TryLogin(userID string, setUserID func(string) error) bool {
	if us.userRepo.IsPresent(userID) {
		err := setUserID(userID)
		return err == nil
	}
	return false
}
