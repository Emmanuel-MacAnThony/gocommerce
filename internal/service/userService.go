package service

import (
	"errors"
	"fmt"

	"github.com/Emmanuel-MacAnThony/gocommerce/internal/domain"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/dto"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})
	userInfo := fmt.Sprintf("%v, %v, %v", user.ID, user.Email, user.UserType)
	return userInfo, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("User does not exist for the provided email id")
	}
	// compare password and generate token
	return user.Email, nil //@temporary hack
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, code int) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId int) (interface{}, error) {
	return nil, nil
}
