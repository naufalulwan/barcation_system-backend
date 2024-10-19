package models

import (
	"barcation_be/config"
	"errors"
	"gorm.io/gorm"
)

type Inquiry struct {
	gorm.Model
	ProductID     uint    `gorm:"not null" json:"product_id"`
	UserID        uint    `gorm:"not null" json:"user_id"`
	TotalQuantity int     `gorm:"not null" json:"total_quantity"`
	TotalPrice    int     `gorm:"not null" json:"total_price"`
	Status        bool    `gorm:"not null" json:"status"`
	Product       Product `gorm:"foreignKey:ProductID;"`
	User          User    `gorm:"foreignKey:UserID;"`
}

func (i *Inquiry) GetInquiryById(ID uint) (*Inquiry, error) {
	err := config.DB.Model(Inquiry{}).Where("id = ?", ID).Take(&i).Error
	if err != nil {
		return nil, err
	}

	err = config.DB.Preload("Inquiry").Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Inquiry) GetInquiry() ([]Inquiry, error) {
	var inq []Inquiry

	if err := config.DB.Model(Inquiry{}).Find(&inq).Error; err != nil {
		return nil, err
	}

	err := config.DB.Preload("Product").Preload("User").Find(&inq).Error
	if err != nil {
		return nil, err
	}

	return inq, nil
}

func (i *Inquiry) CreateInquiry() error {
	if err := config.DB.First(&Product{}, i.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if err := config.DB.First(&User{}, i.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := config.DB.Create(&i).Error; err != nil {
		return err
	}

	err := config.DB.Preload("Product").Preload("User").First(&i, i.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *Inquiry) UpdateInquiry(Id uint) error {
	if err := config.DB.First(&Product{}, i.ProductID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if err := config.DB.First(&User{}, i.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := config.DB.Model(&Inquiry{}).Where("id = ?", Id).Updates(&i).Error; err != nil {
		return err
	}

	return nil
}

func (i *Inquiry) DeleteInquiry(Id uint) error {
	if err := config.DB.Where("id = ?", Id).Delete(&Inquiry{}).Error; err != nil {
		return err
	}

	return nil
}
