package models

import (
	"barcation_be/config"
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(100);not null" json:"name"`
	Price       int      `gorm:"not null" json:"price"`
	Quantity    int      `gorm:"not null" json:"quantity"`
	Status      bool     `gorm:"not null" json:"status"`
	Image       string   `gorm:"type:text" json:"image"`
	Description string   `gorm:"type:text" json:"description"`
	CategoryID  uint     `gorm:"not null" json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Product) GetProductById(uid uint) (*Product, error) {

	err := config.DB.Model(Product{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return nil, err
	}

	err = config.DB.Preload("Category").Take(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) DeleteProduct(id uint) error {
	if err := config.DB.Where("id = ?", id).Delete(&Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(id uint) error {
	if p.CategoryID != 0 {
		if err := config.DB.First(&Category{}, p.CategoryID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("category not found")
			}
			return err
		}
	}

	if err := config.DB.Model(&Product{}).Where("id = ?", id).Updates(&p).Update("status", p.Status).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) GetProduct() ([]Product, error) {
	var pro []Product

	if err := config.DB.Model(Product{}).Find(&pro).Error; err != nil {
		return nil, err
	}

	err := config.DB.Preload("Category").Find(&pro).Error
	if err != nil {
		return nil, err
	}

	return pro, nil
}

func (p *Product) SaveProduct() error {

	if err := config.DB.First(&Category{}, p.CategoryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Category not found")
		}
		return err
	}

	if err := config.DB.Create(&p).Error; err != nil {
		return err
	}

	err := config.DB.Preload("Category").First(&p, p.ID).Error
	if err != nil {
		return err
	}

	return nil
}
