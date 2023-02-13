package postgresql

import (
	"context"
	"digitalPet/domain"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type postgresqlFoodRepository struct {
	db *gorm.DB
}

func NewPostgresqlFoodRepository(db *gorm.DB) domain.FoodRepository {
	return &postgresqlFoodRepository{db}
}

func (p *postgresqlFoodRepository) New(ctx context.Context, food *domain.Food) (*domain.Food, error) {
	result := p.db.Create(&food)
	return food, result.Error
}

func (p *postgresqlFoodRepository) Get(ctx context.Context, food *domain.Food, optsWhere ...map[string]interface{}) (result *domain.Food, err error) {
	query := p.db.Model(&domain.Food{}).Where(food)
	if len(optsWhere) != 0 {
		for _, optWhere := range optsWhere {
			for k, v := range optWhere {
				if v == nil {
					query.Where(k)
				} else {
					query.Where(k, v)
				}
			}
		}
	}
	err = query.First(&result).Error
	return
}

func (p *postgresqlFoodRepository) Gets(ctx context.Context, food *domain.Food, optsWhere ...map[string]interface{}) (result []*domain.Food, err error) {
	query := p.db.Model(&domain.Food{}).Where(food)
	if len(optsWhere) != 0 {
		for _, optWhere := range optsWhere {
			for k, v := range optWhere {
				if v == nil {
					query.Where(k)
				} else {
					query.Where(k, v)
				}
			}
		}
	}
	err = query.Find(&result).Error
	return
}
