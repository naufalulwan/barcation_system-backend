package models

import (
	"barcation_be/config"
	"errors"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PaymentType      string  `gorm:"not null" json:"payment_type"`
	PaymentDate      string  `gorm:"not null" json:"payment_date"`
	PaymentStatus    bool    `gorm:"not null" json:"payment_status"`
	PaymentReference string  `json:"payment_reference"`
	PaymentSignature string  `gorm:"not null;unique" json:"payment_signature"`
	UserID           uint    `gorm:"not null" json:"user_id"`
	InquiryID        uint    `gorm:"not null" json:"inquiry_id"`
	User             User    `gorm:"foreignKey:UserID;"`
	Inquiry          Inquiry `gorm:"foreignKey:InquiryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (i *Payment) GetPaymentById(ID uint) (*Payment, error) {
	err := config.DB.Model(Payment{}).Where("id = ?", ID).Take(&i).Error
	if err != nil {
		return nil, err
	}

	err = config.DB.Preload("User").Preload("Inquiry").Preload("Inquiry.Product").Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Payment) GetPayment() ([]Payment, error) {
	var pay []Payment

	if err := config.DB.Model(Payment{}).Find(&pay).Error; err != nil {
		return nil, err
	}

	err := config.DB.Preload("User").Preload("Inquiry").Preload("Inquiry.Product").Find(&pay).Error
	if err != nil {
		return nil, err
	}

	return pay, nil
}

func (i *Payment) CreatePayment() error {
	if err := config.DB.First(&User{}, i.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := config.DB.First(&Inquiry{}, i.Inquiry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("inquiry not found")
		}
		return err
	}

	if err := config.DB.Create(&i).Error; err != nil {
		return err
	}

	err := config.DB.Preload("User").Preload("Inquiry").Preload("Inquiry.Product").First(&i, i.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *Payment) UpdatePayment(ID uint) error {
	if err := config.DB.First(&User{}, i.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if err := config.DB.First(&Inquiry{}, i.InquiryID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("inquiry not found")
		}
		return err
	}

	if err := config.DB.Model(&Payment{}).Where("id = ?", ID).Updates(&i).Error; err != nil {
		return err
	}

	return nil
}

func (i *Payment) DeletePayment(ID uint) error {
	if err := config.DB.Where("id = ?", ID).Delete(&Payment{}).Error; err != nil {
		return err
	}

	return nil
}
