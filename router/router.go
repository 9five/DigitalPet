package router

import (
	"digitalPet/config"
	_foodHttp "digitalPet/food/delivery/http"
	_foodRepo "digitalPet/food/repository/postgresql"
	_foodUsecase "digitalPet/food/usecase"
	_petHttp "digitalPet/pet/delivery/http"
	_petRepo "digitalPet/pet/repository/postgresql"
	_petUsecase "digitalPet/pet/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	setupRouter(r)
	return r
}

func setupRouter(router *gin.Engine) {
	setupPet(router)
}

func setupPet(router *gin.Engine) {
	petRepo := _petRepo.NewPostgresqlPetRepository(config.DB)
	petUsecase := _petUsecase.NewPetUsecase(petRepo)
	foodRepo := _foodRepo.NewPostgresqlFoodRepository(config.DB)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepo)
	_petHttp.NewPetHandler(router.Group("/api/pet"), petUsecase, foodUsecase)
}

func setupFood(router *gin.Engine) {
	foodRepo := _foodRepo.NewPostgresqlFoodRepository(config.DB)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepo)
	_foodHttp.NewFoodHandler(router.Group("api/food"), foodUsecase)
}
