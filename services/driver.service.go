package services

import (
	"errors"

	dtos "github.com/Thomazoide/Trucks-Monitor/DTOs"
	"github.com/Thomazoide/Trucks-Monitor/db"
	"github.com/Thomazoide/Trucks-Monitor/models"
	"gorm.io/gorm"
)

type DriverService struct {
	DBInstance *gorm.DB
}

func (*DriverService) NewDriverService() *DriverService {
	return &DriverService{
		DBInstance: db.GetPGInstance(),
	}
}

func (s *DriverService) NewDriver(driver *dtos.DriverDTO) (*models.Driver, error) {
	var DriverExists *models.Driver
	var newDriver *models.Driver
	if err := s.DBInstance.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&DriverExists).Where("rut = ?", driver.Rut).Error; err != nil {
			return err
		}
		if DriverExists != nil {
			return errors.New("Conductor ya existente en BBDD")
		}
		newDriver = &models.Driver{
			Rut:       driver.Rut,
			Name:      driver.Nombre,
			Lastname:  driver.Apellido,
			Email:     driver.Email,
			Cellphone: driver.Cellphone,
		}
		if err := tx.Save(newDriver).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return newDriver, nil
}

func (s *DriverService) UpdateDriver(driver *models.Driver) (*models.Driver, error) {
	var DriverExists *models.Driver
	s.DBInstance.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&DriverExists).Where("rut = ?", driver.Rut).Error; err != nil {
			return err
		}
		if DriverExists == nil || DriverExists.ID == 0 {
			return errors.New("Conductor no encontrado en BBDD")
		}
		if err := tx.Save(&driver).Error; err != nil {
			return err
		}
		return nil
	})
	return nil, nil
}

func fromDTOtoModel(model *models.Driver, dto *dtos.DriverDTO) *models.Driver {
	auxModel := model
	auxModel.Cellphone = dto.Cellphone
	auxModel.Name = dto.Nombre
	auxModel.Lastname = dto.Apellido
	auxModel.Rut = dto.Rut
	auxModel.Email = dto.Email
	return auxModel
}
