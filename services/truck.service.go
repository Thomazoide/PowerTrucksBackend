package services

import (
	dtos "github.com/Thomazoide/Trucks-Monitor/DTOs"
	"github.com/Thomazoide/Trucks-Monitor/db"
	"github.com/Thomazoide/Trucks-Monitor/models"
	"gorm.io/gorm"
)

type TruckService struct {
	DBInstance *gorm.DB
}

func (s *TruckService) NewTruckService() *TruckService {
	return &TruckService{
		DBInstance: db.GetPGInstance(),
	}
}

func (s *TruckService) NewTruck(truck *dtos.TruckDTO) (*models.Truck, error) {
	newTruck := fromDTOtoModel(truck)
	if err := s.DBInstance.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(newTruck).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return newTruck, nil
}

func fromDTOtoModel(truckDTO *dtos.TruckDTO) *models.Truck {
	return &models.Truck{
		Plate:     truckDTO.Plate,
		CompanyID: &truckDTO.CompanyID,
	}
}
