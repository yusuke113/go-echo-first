package repository_test

import (
	"asobi/model"
	"asobi/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserRepositoryCreate(t *testing.T) {
	// テスト用のデータベースを作成
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to connect to database:", err)
	}
	// マイグレーションを実行してテーブルを作成
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		t.Fatal("failed to perform database migration:", err)
	}

	repo := repository.NewUserRepository(db)

	t.Run("ユーザーを新規登録できること", func(t *testing.T) {
		user := &model.User{Name: "test1", Email: "test@example.com"}

		err := repo.Create(user)

		// エラーがないことを確認
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)                                        // IDが設定されていることを確認
		assert.Equal(t, "test1", user.Name)                                // Nameが設定されていることを確認
		assert.Equal(t, "test@example.com", user.Email)                   // Emailが設定されていることを確認
		assert.NotZero(t, user.CreatedAt)                                 // CreatedAtが設定されていることを確認
		assert.NotZero(t, user.UpdatedAt)                                 // UpdatedAtが設定されていることを確認
		assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second) // CreatedAtが現在時刻と近いことを確認
		assert.WithinDuration(t, time.Now(), user.UpdatedAt, time.Second) // UpdatedAtが現在時刻と近いことを確認
	})

	t.Run("同一メールアドレスでの登録時例外が発生すること", func(t *testing.T) {
		user2 := &model.User{Name: "test2", Email: "test@example.com"}
	
		err := repo.Create(user2)
	
		assert.Error(t, err)
		assert.EqualError(t, err, "UNIQUE constraint failed: users.email")
	})

	// テストが終わったらデータベースをクリーンアップ
	db.Migrator().DropTable(&model.User{})
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer dbSQL.Close()
}
