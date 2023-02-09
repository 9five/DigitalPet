package router

import (
	"goDDD/config"
	_foodHandler "goDDD/food/delivery/handler"
	_foodRepo "goDDD/food/repository/postgresql"
	_foodUsecase "goDDD/food/usecase"
	_petHandler "goDDD/pet/delivery/handler"
	_petRepo "goDDD/pet/repository/postgresql"
	_petUsecase "goDDD/pet/usecase"

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
	_petHandler.NewPetHandler(router.Group("/api/pet"), petUsecase, foodUsecase)
}

func setupFood(router *gin.Engine) {
	foodRepo := _foodRepo.NewPostgresqlFoodRepository(config.DB)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepo)
	_foodHandler.NewFoodHandler(router.Group("api/food"), foodUsecase)
}
