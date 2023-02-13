package handler

import (
	"digitalPet/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	FoodUsecase domain.FoodUsecase
}

func NewFoodHandler(router *gin.RouterGroup, foodUsecase domain.FoodUsecase) {
	handler := &foodHandler{
		FoodUsecase: foodUsecase,
	}

	router.POST("/newVarieties", handler.NewVarieties)
}

func (f *foodHandler) NewVarieties(ctx *gin.Context) {
	var input domain.InputFood
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	f.FoodUsecase.NewFood(ctx, &domain.Food{
		Name:     input.Name,
		Calories: input.Calories,
	})
}
