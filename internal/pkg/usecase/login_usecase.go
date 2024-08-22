package usecase

import "grib/internal/pkg/repository"

type LoginUsecaseInterface interface {
	GetByEmail(email string) (interface{}, error)
}

type LoginUsecase struct {
	userRepo repository.UserRepositoryInterface
}

func NewLoginUsecase (userRepo repository.UserRepositoryInterface) LoginUsecaseInterface {
	return &LoginUsecase{
		userRepo: userRepo,
	}
}

func (lu *LoginUsecase) GetByEmail(email string) (interface{}, error) {
	return lu.userRepo.GetByEmail(email)
}