package postgresql

import (
	"context"
	"goDDD/domain"

	"gorm.io/gorm"
	// "gorm.io/gorm/clause"
)

type postgresqlPetRepository struct {
	db *gorm.DB
}

func NewPostgresqlPetRepository(db *gorm.DB) domain.PetRepository {
	return &postgresqlPetRepository{db}
}

func (p *postgresqlPetRepository) New(ctx context.Context, pet *domain.Pet) (*domain.Pet, error) {
	result := p.db.Create(&pet)
	return pet, result.Error
}

func (p *postgresqlPetRepository) Get(ctx context.Context, pet *domain.Pet, optsWhere ...map[string]interface{}) (result *domain.Pet, err error) {
	query := p.db.Model(&domain.Pet{}).Where(pet)
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

func (p *postgresqlPetRepository) Gets(ctx context.Context, pet *domain.Pet, optsWhere ...map[string]interface{}) (result []*domain.Pet, err error) {
	query := p.db.Model(&domain.Pet{}).Where(pet)
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

func (p *postgresqlPetRepository) Update(ctx context.Context, pet *domain.Pet) error {
	return p.db.Save(&pet).Error
}
