package validator_test

import (
	"asobi/model"
	"asobi/validator"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidator(t *testing.T) {
	// テスト対象のバリデータを作成
	userValidator := validator.NewUserValidator()

	t.Run("正常なユーザー情報の場合、エラーが発生しないこと", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "test@example.com",
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生しないことを確認
		assert.NoError(t, err)
	})

	t.Run("名前が未入力の場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "", // 空の名前
			Email: "test@example.com",
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "名前は必須です."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})

	t.Run("名前の文字数が範囲外の場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "test Doe Doe Doe Doe Doe", // 21文字の名前
			Email: "test@example.com",
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "名前は1文字以上20文字以下です."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})

	t.Run("名前が空白文字のみの場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  " ", // 空白文字のみの名前
			Email: "test@example.com",
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "名前は空白文字のみではない必要があります."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})

	t.Run("メールアドレスが未入力の場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "", // 空のメールアドレス
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "メールアドレスは必須です."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})

	t.Run("メールアドレスの文字数が範囲外の場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "test@example.com" + strings.Repeat("a", 250), // 256文字のメールアドレス
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "メールアドレスは1文字以上255文字以下です."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})

	t.Run("メールアドレスの形式が不正な場合、エラーが発生すること", func(t *testing.T) {
		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "test@example", // 不正な形式のメールアドレス
		}

		// バリデーションの実行
		err := userValidator.UserValidator(user)

		// エラーが発生することを確認
		expectedErrorMessage := "メールアドレスの形式が不正です."
		assert.Error(t, err)
		assert.Contains(t, err.Error(), expectedErrorMessage)
	})
}
