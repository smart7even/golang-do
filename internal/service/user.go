package service

import (
	"fmt"

	"github.com/smart7even/golang-do/internal/domain"
)

type UserRepo interface {
	Create(token string) error
	ReadAll() ([]domain.User, error)
	ReadByToken(token string) (*domain.User, error)
	Update(user domain.User) error
	Delete(token string) error
}

type UserDoesNotExist struct {
	UserId string
}

func (e UserDoesNotExist) Error() string {
	return fmt.Sprintf("todo with id %v does not exist", e.UserId)
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(token string) error {
	return s.userRepo.Create(token)
}

func (s *UserService) ReadAll() ([]domain.User, error) {
	return s.userRepo.ReadAll()
}

func (s *UserService) Update(user domain.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) Delete(userId string) error {
	return s.userRepo.Delete(userId)
}
