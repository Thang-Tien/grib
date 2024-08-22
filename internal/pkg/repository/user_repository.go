package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

type UserRepositoryInterface interface {
	GetByEmail(string) (interface{}, error)
}

type UserRepository struct {
	Logger *logrus.Logger
	DB *gorm.DB
}

func NewUserRepository(Logger *logrus.Logger, DB *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		Logger: Logger,
		DB: DB,
	}
}

func (ur *UserRepository) GetByEmail(email string) (interface{}, error) {
	var user interface{}
	user = ur.DB.Take(&user, email)
	ur.Logger.Info(user)
	return user, nil
}