package database

import (
	sharedLogger "grib/pkg/shared/logger"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host    string
	Name    string
	User    string
	Pass    string
	Port    string
	Charset string
}

func NewDB(config DBConfig, logger *logrus.Logger) (*gorm.DB, error) {
	var DB *gorm.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.Name, config.Charset)
	logger.Info(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: sharedLogger.NewGormLogger(logger),
	})
	if err != nil {
		return nil, err
	}
	return DB, Ping(DB)
}

func CloseDB(
	logger *logrus.Logger,
	db *gorm.DB,
) {
	myDB, err := db.DB()
	if err != nil {
		logger.Errorf("Error while returning *sql.DB: %v", err)
	}

	logger.Info("Closing the DB connection pool")
	if err := myDB.Close(); err != nil {
		logger.Errorf("Error while closing the master DB connection pool: %v", err)
	}
}

func Ping(db *gorm.DB) error {
	myDB, err := db.DB()
	if err != nil {
		return err
	}
	return myDB.Ping()
}
