package service

import (
	"log"

	"github.com/Emmanuel-MacAnThony/gocommerce/internal/domain"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/dto"
)

type UserService struct {
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {
	log.Println("input", input)
	return "this is my token as on now", nil
}

func (s UserService) Login(input interface{}) (string, error) {
	return "", nil
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
