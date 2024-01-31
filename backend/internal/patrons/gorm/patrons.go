package gorm

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
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

func (s *PatronService) DeletePatron(patron *patrons.Patron) error {
	return s.DB.Delete(&patron).Error
}

func (s *PatronService) DeletePatronByID(id string) error {
	return s.DB.Delete(&patrons.Patron{}, "id = ?", id).Error
}

func (s *PatronService) UpdateConsumption(
	patron *patrons.Patron, drink *cocktail.Cocktail) error {
	patron.TotalAlcohol += drink.TotalAlcohol

	tx := s.DB.Begin()
	errAss := tx.Model(&patron).Association("Cocktails").Append(drink)
	if errAss != nil {
		tx.Rollback()
		return errAss
	}
	errUpt := tx.Save(&patron).Error
	if errUpt != nil {
		tx.Rollback()
		return errUpt
	}
	//tx.Preload("Cocktails").Find(&patron)
	return tx.Commit().Error
}

// GetAllPatrons implements PatronService
// Gets all Patrons.
func (s *PatronService) GetAllPatrons() (*[]patrons.Patron, error) {
	var patron []patrons.Patron
	if result := s.DB.Find(&patron); result.Error != nil {
		return nil, result.Error
	}
	return &patron, nil
}

// GetPatronByID implements PatronService
// Gets a Patron from their ID.
func (s *PatronService) GetPatronByID(id string) (*patrons.Patron, error) {
	patron := patrons.Patron{}
	if result := s.DB.Preload("Cocktails").
		First(&patron, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return &patron, nil
}
