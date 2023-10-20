package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic(err)
	}

	/**
	** MIGRATIONS
	**/

	//db.AutoMigrate(&models.User{})

	return db
}
