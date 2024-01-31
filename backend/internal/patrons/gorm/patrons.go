package gorm

import (
	"alcohol-consumption-tracker/internal/patrons"
	"gorm.io/gorm"
)

var _ patrons.PatronService = (*PatronService)(nil)

type PatronService struct {
	DB *gorm.DB
	//IngredientService *cocktail
}

func NewPatronService(db *gorm.DB) *PatronService {
	return &PatronService{
		DB: db,
		//IngredientService:
	}
}

func (s *PatronService) CreatePatron(patron *patrons.Patron) error {
	return s.DB.Create(&patron).Error
}

func (s *PatronService) UpdatePatron(patron *patrons.Patron) error {
	return s.DB.Save(&patron).Error
}

func (s *PatronService) UpdateConsumption(patron *patrons.Patron) error {
	// TODO implement
	panic("not implemented")
}

// GetPatronByID implements PatronService
// Gets a Patron from their ID.
func (s *PatronService) GetPatronByID(id string) (*patrons.Patron, error) {
	patron := patrons.Patron{}
	if result := s.DB.First(&patron, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return &patron, nil
}
