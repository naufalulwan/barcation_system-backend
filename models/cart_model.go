package models

import (
	"barcation_be/config"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int  `gorm:"not null" json:"quantity"`
	Total     int  `gorm:"not null" json:"total"`
	ProductID uint `gorm:"not null" json:"product_id"`
	UserID    uint `gorm:"not null" json:"user_id"`
	Product   Product
	User      User
}

func (c *Cart) DeleteCart(id uint) error {
	err := config.DB.Where("id = ?", id).Delete(&Cart{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Cart) GetCart() ([]Cart, error) {
	var cart []Cart
	err := config.DB.Preload("Product").Preload("User").Find(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *Cart) AddCart() error {
	if err := config.DB.First(&Product{}, c.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if err := config.DB.First(&User{}, c.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	var data Cart

	config.DB.Where("product_id = ? AND user_id = ?", c.ProductID, c.UserID).Find(&data)

	if data.ID != 0 {
		data.Quantity += c.Quantity
		data.Total += c.Total

		if err := config.DB.Save(&data).Error; err != nil {
			return err
		}

		return nil
	}

	if err := config.DB.Create(&c).Error; err != nil {
		return err
	}

	err := config.DB.Preload("Product").Preload("User").First(&c).Error
	if err != nil {
		return err
	}

	return nil
}
