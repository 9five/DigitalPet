package usecase

import (
	"context"
	"goDDD/domain"
	"math"
	"math/rand"
	"time"
)

type petUsecase struct {
	petRepo domain.PetRepository
}

func NewPetUsecase(petRepository domain.PetRepository) domain.PetUsecase {
	return &petUsecase{
		petRepo: petRepository,
	}
}

func (p *petUsecase) NewPet(ctx context.Context, pet *domain.Pet) (*domain.Pet, error) {
	return p.petRepo.New(ctx, pet)
}

func (p *petUsecase) GetPet(ctx context.Context, pet *domain.Pet) (*domain.Pet, error) {
	result, err := p.petRepo.Get(ctx, pet)
	if err != nil {
		return result, err
	}

	elapsedTime := time.Now().Sub(pet.LastVisit).Minutes()
	decline := math.Floor(elapsedTime / 12)

	pet.Satisfaction -= int(decline)
	pet.LastVisit = time.Now()
	p.petRepo.Update(ctx, pet)

	return p.petRepo.Get(ctx, pet)
}

func (p *petUsecase) UpdatePet(ctx context.Context, pet *domain.Pet) error {
	if pet.Satisfaction > 100 {
		pet.Satisfaction = 100
	}

	if pet.Satisfaction < 0 {
		pet.Satisfaction = 0
	}
	return p.petRepo.Update(ctx, pet)
}

func (p *petUsecase) Meal(ctx context.Context, pet *domain.Pet, food *domain.Food) error {
	increase := float64(food.Calories) / 50

	freshness := rand.Intn(5)
	increase *= float64(freshness)

	pet.Satisfaction += int(increase)
	return p.petRepo.Update(ctx, pet)
}
