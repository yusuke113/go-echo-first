package main

import (
	"asobi/controller"
	"asobi/db"
	"asobi/repository"
	"asobi/router"
	"asobi/usecase"
	"asobi/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)
	helloController := controller.NewHelloController()
	e := router.NewRouter(helloController, userController)
	e.Logger.Fatal(e.Start(":8888"))
}
