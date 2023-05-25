package controller_test

import (
	"asobi/controller"
	"asobi/model"
	"asobi/usecase"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// userUseCaseのモック
type MockUserUseCase struct {
	usecase.IUserUseCase
	FakeStoreUser func(user model.User) (model.UserResponse, error)
}

func (m *MockUserUseCase) StoreUser(user model.User) (model.UserResponse, error) {
	return m.FakeStoreUser(user)
}

func TestUserController_Store(t *testing.T) {
	// userUseCaseのモックを作成
	MockUserUseCase := &MockUserUseCase{}

	// コントローラのインスタンスを作成
	controller := controller.NewUserController(MockUserUseCase)

	// Echoのインスタンスを作成
	e := echo.New()

	t.Run("ユーザー登録処理成功時のレスポンスが正しいこと", func(t *testing.T) {
		// モックの定義
		MockUserUseCase.FakeStoreUser = func(user model.User) (model.UserResponse, error) {
			resUser := model.UserResponse{
				ID:    1,
				Name:  "test",
				Email: "test@example.com",
			}
			return resUser, nil
		}

		// テスト用のユーザーデータ
		userData := model.User{
			Name:  "test",
			Email: "test@example.com",
		}
		// テスト用のリクエストを作成
		requestBody, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// コントローラのハンドラを実行
		err := controller.StoreUser(c)

		// エラーが発生していないか検証
		assert.NoError(t, err)

		// レスポンスを取得
		res := rec.Result()

		// レスポンスのステータスコードを検証
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("ユーザーのリクエストデータ型が不正な場合のレスポンスが正しいこと", func(t *testing.T) {
		// モックの定義
		MockUserUseCase.FakeStoreUser = func(user model.User) (model.UserResponse, error) {
			return model.UserResponse{}, nil
		}

		// テスト用のデータを構造体ではなくjsonで
		userData := map[string]interface{}{
			"name":  123,
			"email": "test@example.com",
		}

		// テスト用のリクエストを作成
		requestBody, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// コントローラのハンドラを実行
		err := controller.StoreUser(c)

		// エラーが発生していないか検証
		assert.NoError(t, err)

		// レスポンスを取得
		res := rec.Result()

		// レスポンスのステータスコードを検証
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("ユーザー登録処理でエラーが発生した場合のレスポンスが正しいこと", func(t *testing.T) {
		// モックの定義
		MockUserUseCase.FakeStoreUser = func(user model.User) (model.UserResponse, error) {
			return model.UserResponse{}, errors.New("usecase error")
		}

		// テスト用のユーザーデータ
		userData := model.User{
			Name:  "test",
			Email: "test@example.com",
		}
		// テスト用のリクエストを作成
		requestBody, _ := json.Marshal(userData)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// コントローラのハンドラを実行
		err := controller.StoreUser(c)

		// エラーが発生していないか検証
		assert.NoError(t, err)

		// レスポンスを取得
		res := rec.Result()

		// レスポンスのステータスコードを検証
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
