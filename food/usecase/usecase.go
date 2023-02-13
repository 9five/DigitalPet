package usecase

import (
	"context"
	"digitalPet/domain"
)

type foodUsecase struct {
	foodRepo domain.FoodRepository
}

func NewFoodUsecase(foodRepository domain.FoodRepository) domain.FoodUsecase {
	return &foodUsecase{
		foodRepo: foodRepository,
	}
}

func (f *foodUsecase) NewFood(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	return f.foodRepo.New(ctx, food)
}

func (f *foodUsecase) GetFood(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	return f.foodRepo.Get(ctx, food)
}
