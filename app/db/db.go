package db

import (
	"fmt"
	"goAPISample/app/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// DB初期化
func Init() {
	// 環境変数取得
	godotenv.Load(".env")
	godotenv.Load()

	// DB接続
	user := os.Getenv("MySQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	db, err = gorm.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	//migration
	autoMigration()
}

// DB取得
func GetDB() *gorm.DB {
	return db
}

// DB接続終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

// マイグレーション
func autoMigration() {
	db.AutoMigrate(&model.Post{})
}
