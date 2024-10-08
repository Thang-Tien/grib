package migration

import (
	"grib/internal/pkg/domain/model/entity"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(entity.Customer{}, entity.Driver{}, entity.Vehicle{}, entity.Wallet{})
	return err
}
