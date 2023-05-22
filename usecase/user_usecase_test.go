package usecase_test

import (
	"asobi/model"
	"asobi/repository"
	"asobi/usecase"
	"asobi/validator"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// モックのUserRepository
type MockUserRepository struct {
	repository.IUserRepository
	FakeCreate func(user *model.User) error
}

func (m *MockUserRepository) Create(user *model.User) error {
	return m.FakeCreate(user)
}

// モックのUserValidator
type MockUserValidator struct {
	validator.IUserValidator
	FakeUserValidator func(user model.User) error
}

func (m *MockUserValidator) UserValidator(user model.User) error {
	return m.FakeUserValidator(user)
}

func TestTest(t *testing.T) {
	// モックオブジェクトの作成
	MockUserRepository := &MockUserRepository{}
	MockUserValidator := &MockUserValidator{}

	// UseCaseの作成
	uu := usecase.NewUserUseCase(MockUserRepository, MockUserValidator)

	t.Run("ユーザー情報を新規登録できること", func(t *testing.T) {
		// モックの定義
		MockUserRepository.FakeCreate = func(user *model.User) error {
			user.ID = 1
			user.Name = "test"
			user.Email = "test@example.com"
			return nil
		}
		MockUserValidator.FakeUserValidator = func(user model.User) error {
			return nil
		}

		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "test@example.com",
		}

		// ユーザーの作成
		resUser, err := uu.StoreUser(user)

		// エラーがないことを確認
		assert.NoError(t, err)

		// 期待されるユーザーデータ
		expectedUser := model.UserResponse{
			ID:    1,
			Name:  "test",
			Email: "test@example.com",
		}

		// 期待される結果の検証
		assert.Equal(t, expectedUser, resUser)
	})

	t.Run("ユーザー情報のバリデーションエラーが発生すること", func(t *testing.T) {
		// モックの定義
		MockUserRepository.FakeCreate = func(user *model.User) error {
			return nil
		}
		MockUserValidator.FakeUserValidator = func(user model.User) error {
			return errors.New("名前は必須です")
		}

		// テストデータの作成
		user := model.User{
			Name:  "",
			Email: "test@example.com",
		}

		// ユーザーの作成
		_, err := uu.StoreUser(user)

		// エラーが発生することを確認
		assert.Error(t, err)
	})

	t.Run("ユーザー情報の登録エラーが発生すること", func(t *testing.T) {
		// モックの定義
		MockUserRepository.FakeCreate = func(user *model.User) error {
			return errors.New("データベースエラー")
		}
		MockUserValidator.FakeUserValidator = func(user model.User) error {
			return nil
		}

		// テストデータの作成
		user := model.User{
			Name:  "test",
			Email: "test@example.com",
		}

		// ユーザーの作成
		_, err := uu.StoreUser(user)

		// エラーが発生することを確認
		assert.Error(t, err)
	})
}
