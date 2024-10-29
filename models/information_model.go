package models

import (
	"barcation_be/config"
	"gorm.io/gorm"
)

type Information struct {
	gorm.Model
	Type    string `gorm:"not null" json:"type"`
	Title   string `gorm:"not null" json:"title"`
	Message string `gorm:"not null" json:"message"`
	Image   string `gorm:"not null" json:"image"`
}

func (i *Information) GetInformation(Type string) ([]Information, error) {
	var inf []Information

	if err := config.DB.Model(Information{}).Where("type = ?", Type).Find(&inf).Error; err != nil {
		return nil, err
	}

	return inf, nil
}

func (i *Information) CreateInformation() error {
	if err := config.DB.Create(&i).Error; err != nil {
		return err
	}

	return nil
}

func (i *Information) UpdateInformation(ID uint) error {
	if err := config.DB.Model(&Information{}).Where("id = ?", ID).Updates(&i).Error; err != nil {
		return err
	}

	return nil
}

func (i *Information) DeleteInformation(ID uint) error {
	if err := config.DB.Where("id = ?", ID).Delete(&Information{}).Error; err != nil {
		return err
	}

	return nil
}
