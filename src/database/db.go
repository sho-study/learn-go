package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"my-app/src/entities"
	"os"
)

var DB *gorm.DB

func Connect() {
  var err error

	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DBNAME")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("データベースの接続に失敗しました")
	}
}

func Migrate() {
	DB.AutoMigrate(
		&entities.User{},
	)
}
