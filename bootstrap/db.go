package bootstrap

import (
	"time"

	"github.com/merkurtran/goblog/app/models/article"
	"github.com/merkurtran/goblog/app/models/user"
	"github.com/merkurtran/goblog/pkg/model"
	"gorm.io/gorm"
)

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
