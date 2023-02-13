package handler

import (
	"digitalPet/domain"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PetHandler struct {
	PetUsecase  domain.PetUsecase
	FoodUsecase domain.FoodUsecase
}

func NewPetHandler(router *gin.RouterGroup, petUsecase domain.PetUsecase, foodUsecase domain.FoodUsecase) {
	handler := &PetHandler{
		PetUsecase:  petUsecase,
		FoodUsecase: foodUsecase,
	}

	router.POST("/adopt", handler.Adopt)
	router.PUT("/:petID/eat/:foodID", handler.Eat)
	router.PUT("/:petID/poo", handler.Eat)
	router.PUT("/:petID/sleep", handler.Eat)
	router.PUT("/:petID/play", handler.Eat)
	router.GET("/:petID/status", handler.Status)
}

func (p *PetHandler) Adopt(ctx *gin.Context) {
	var input domain.InputPet
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	p.PetUsecase.NewPet(ctx, &domain.Pet{
		Name:   input.Name,
		Age:    input.Age,
		Legs:   input.Legs,
		Length: input.Length,
	})
}

func (p *PetHandler) Eat(ctx *gin.Context) {
	petID := ctx.Param("petID")
	if petID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("which pet is eating?"),
		})
	}

	foodID := ctx.Param("foodID")
	if foodID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("what are pets eating?"),
		})
	}

	pet, err := getPetById(p, ctx, petID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	food, err := getFoodById(p, ctx, foodID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err := p.PetUsecase.Meal(ctx, pet, food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (p *PetHandler) Poo(ctx *gin.Context) {
	petID := ctx.Param("petID")
	if petID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("which pet is eating?"),
		})
	}

	pet, err := getPetById(p, ctx, petID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	pet.Satisfaction += 5
	if err := p.PetUsecase.UpdatePet(ctx, pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (p *PetHandler) Sleep(ctx *gin.Context) {
	petID := ctx.Param("petID")
	if petID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("which pet is eating?"),
		})
	}

	pet, err := getPetById(p, ctx, petID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	pet.Satisfaction += 60
	if err := p.PetUsecase.UpdatePet(ctx, pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (p *PetHandler) Play(ctx *gin.Context) {
	petID := ctx.Param("petID")
	if petID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("which pet is eating?"),
		})
	}

	pet, err := getPetById(p, ctx, petID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	pet.Satisfaction += 20
	if err := p.PetUsecase.UpdatePet(ctx, pet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (p *PetHandler) Status(ctx *gin.Context) {
	petID := ctx.Param("petID")
	if petID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("which pet is eating?"),
		})
	}

	pet, err := getPetById(p, ctx, petID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":         pet.Name,
		"age":          pet.Age,
		"legs":         pet.Legs,
		"length":       pet.Length,
		"satisfaction": pet.Satisfaction,
	})
}

func getPetById(p *PetHandler, ctx *gin.Context, petID string) (result *domain.Pet, err error) {
	u64PetID, err := strconv.ParseUint(petID, 10, 32)
	if err != nil {
		return
	}

	result, err = p.PetUsecase.GetPet(ctx, &domain.Pet{
		Model: gorm.Model{ID: uint(u64PetID)},
	})
	return
}

func getFoodById(p *PetHandler, ctx *gin.Context, foodID string) (result *domain.Food, err error) {
	u64FoodID, err := strconv.ParseUint(foodID, 10, 32)
	if err != nil {
		return
	}

	result, err = p.FoodUsecase.GetFood(ctx, &domain.Food{
		Model: gorm.Model{ID: uint(u64FoodID)},
	})
	return
}
