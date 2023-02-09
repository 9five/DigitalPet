package domain

import (
	"context"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}

type FoodRepository interface {
	New(ctx context.Context, food *Food) (*Food, error)
	Get(ctx context.Context, food *Food, optsWhere ...map[string]interface{}) (*Food, error)
	Gets(ctx context.Context, food *Food, optsWhere ...map[string]interface{}) ([]*Food, error)
}

type FoodUsecase interface {
	NewFood(ctx context.Context, food *Food) (*Food, error)
	GetFood(ctx context.Context, food *Food) (*Food, error)
}

type InputFood struct {
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}
