package validator

import (
	"asobi/model"
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidator(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidator(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Name,
			validation.Required.Error("名前は必須です"),
			validation.Length(1, 20).Error("名前は1文字以上20文字以下です"),
			validation.By(nameValidator),
		),
		validation.Field(
			&user.Email,
			validation.Required.Error("メールアドレスは必須です"),
			validation.Length(1, 255).Error("メールアドレスは1文字以上255文字以下です"),
			is.Email.Error("メールアドレスの形式が不正です"),
		),
	)
}

// カスタムバリデーションルール
func nameValidator(value interface{}) error {
	name, ok := value.(string)
	if !ok {
		return errors.New("名前は文字列である必要があります")
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("名前は空白文字のみではない必要があります")
	}
	return nil
}
