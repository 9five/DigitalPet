package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name         string    `json:"name"`
	Age          int       `json:"age"`
	Legs         int       `json:"legs"`
	Length       int       `json:"length"` //cm
	Satisfaction int       `json:"satisfaction"`
	LastVisit    time.Time `json:"last_visit"`
}

type PetRepository interface {
	New(ctx context.Context, pet *Pet) (*Pet, error)
	Get(ctx context.Context, pet *Pet, optsWhere ...map[string]interface{}) (*Pet, error)
	Gets(ctx context.Context, pet *Pet, optsWhere ...map[string]interface{}) ([]*Pet, error)
	Update(ctx context.Context, pet *Pet) error
}

type PetUsecase interface {
	NewPet(ctx context.Context, pet *Pet) (*Pet, error)
	GetPet(ctx context.Context, pet *Pet) (*Pet, error)
	UpdatePet(ctx context.Context, pet *Pet) error
	Meal(ctx context.Context, pet *Pet, food *Food) error
}

type InputPet struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Legs   int    `json:"legs"`
	Length int    `json:"length"` //cm
}
