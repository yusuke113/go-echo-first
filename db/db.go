package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// 開発環境の場合は環境変数を読み込む
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("環境変数読み込み成功")

	// データベースに接続するためのURLを作成
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	// DB接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB接続成功")

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, _ := db.DB()
	err := dbSQL.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB切断成功")
}
