package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PGInstance *gorm.DB

func InitPGInstance() error {
	var err error
	if PGInstance, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disabled timezone=America/Santiago", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASS")))); err != nil {
		return err
	}
	//generar los PGInstance.AutoMigrate(&models) con los modelos correspondientes
	return nil
}

func GetPGInstance() *gorm.DB {
	return PGInstance
}
