package database

import (
	"fmt"
	"github.com/RafatMeraz/ecom-micro/auth/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseInstance(cnf *configs.Config) *gorm.DB {
	dsnStr := "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Dhaka"
	dsn := fmt.Sprintf(dsnStr, cnf.DB.Host, cnf.DB.User, cnf.DB.Password, cnf.DB.Name, cnf.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database => " + err.Error())
	}
	return db
}
