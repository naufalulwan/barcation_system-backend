package models

import (
	"barcation_be/config"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
	Icon string `gorm:"not null" json:"icon"`
}

func (c *Category) DeleteCategory(id uint) error {

	err := config.DB.Where("category_id = ?", id).Delete(&Product{}).Error

	if err != nil {
		return err
	}

	err = config.DB.Delete(&Category{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) UpdateCategory(id uint) error {
	err := config.DB.Model(&Category{}).Where("id = ?", id).Updates(&c).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Category) GetCategory() ([]Category, error) {
	var cat []Category

	if err := config.DB.Model(Category{}).Find(&cat).Error; err != nil {
		return nil, err
	}

	return cat, nil
}

func (c *Category) SaveCategory() error {
	if err := config.DB.Create(&c).Error; err != nil {
		return err
	}

	return nil
}
