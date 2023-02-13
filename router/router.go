package router

import (
	"goDDD/config"
	_foodHttp "goDDD/food/delivery/http"
	_foodRepo "goDDD/food/repository/postgresql"
	_foodUsecase "goDDD/food/usecase"
	_petHttp "goDDD/pet/delivery/http"
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
	_petHttp.NewPetHandler(router.Group("/api/pet"), petUsecase, foodUsecase)
}

func setupFood(router *gin.Engine) {
	foodRepo := _foodRepo.NewPostgresqlFoodRepository(config.DB)
	foodUsecase := _foodUsecase.NewFoodUsecase(foodRepo)
	_foodHttp.NewFoodHandler(router.Group("api/food"), foodUsecase)
}
