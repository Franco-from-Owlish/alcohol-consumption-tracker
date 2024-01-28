package gorm

import (
	"alcohol-consumption-tracker/internal/patrons"
	"gorm.io/gorm"
)

var _ patrons.PatronService = (*PatronService)(nil)

type PatronService struct {
	DB *gorm.DB
}

func NewPatronService(db *gorm.DB) *PatronService {
	return &PatronService{
		DB: db,
	}
}

func (s *PatronService) CreatePatron(Patron *patrons.Patron) error {
	return s.DB.Create(&Patron).Error
}

func (s *PatronService) UpdatePatron(Patron *patrons.Patron) error {
	return s.DB.Save(&Patron).Error
}

// GetPatronByID implements PatronService
// Gets a Patron from their ID.
func (s *PatronService) GetPatronByID(id string) (*patrons.Patron, error) {
	Patron := patrons.Patron{}
	if result := s.DB.First(&Patron, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}
	return &Patron, nil
}

func (s *PatronService) UpdateConsumption(id uint) (*patrons.Patron, error) {
	//TODO implement me
	panic("implement me")
}
