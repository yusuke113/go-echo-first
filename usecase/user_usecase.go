package usecase

import (
	"asobi/model"
	"asobi/repository"
	"asobi/validator"
)

type IUserUseCase interface {
	StoreUser(user model.User) (model.UserRespose, error)
}

type userUseCase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUseCase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUseCase {
	return &userUseCase{ur: ur, uv: uv}
}

func (uu *userUseCase) StoreUser(user model.User) (model.UserRespose, error) {
	// バリデーション
	if err := uu.uv.UserValidator(user); err != nil {
		return model.UserRespose{}, err
	}

	// User構造体を作成
	newUser := model.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := uu.ur.Create(&newUser); err != nil {
		return model.UserRespose{}, err
	}

	// レスポンス用のUser構造体を作成
	resUser := model.UserRespose{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	return resUser, nil
}
