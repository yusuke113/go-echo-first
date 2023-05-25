package main

import (
	"asobi/db"
	"asobi/model"
	"fmt"
)

func main() {
	// データベース接続
	dbConn := db.NewDB()

	// マイグレーション終了後にDBを切断する
	defer fmt.Println("Successfully Up Migrated")
	defer db.CloseDB(dbConn)

	// マイグレーション
	dbConn.AutoMigrate(&model.User{}, &model.Post{})
}
