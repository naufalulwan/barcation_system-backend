package models

import (
	"barcation_be/config"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string    `gorm:"type:varchar(100);not null;unique" json:"username"`
	Password    string    `gorm:"type:varchar(100);not null" json:"password"`
	Email       string    `gorm:"type:varchar(100);not null" json:"email"`
	Level       string    `gorm:"type:varchar(100);not null" json:"level"`
	Address     string    `gorm:"type:varchar(100)" json:"address"`
	Phone       string    `gorm:"type:varchar(100)" json:"phone"`
	Status      bool      `gorm:"not null" json:"status"`
	Position    string    `gorm:"type:varchar(100)" json:"position"`
	DeviceId    string    `gorm:"type:varchar(100)" json:"device_id"`
	DeviceToken string    `gorm:"type:varchar(100)" json:"device_token"`
	SaveLogin   bool      `gorm:"not null" json:"save_login"`
	LastLogin   time.Time `json:"last_login"`
	Ssn         string    `gorm:"type:varchar(100)" json:"ssn"`
}

func (u *User) UpdateSaveInfoLogin() error {
	err := config.DB.Model(&User{}).Where("id = ?", u.ID).Update("save_login", u.SaveLogin).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdatePassword(password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	u.Password = string(hashedPassword)

	err := config.DB.Model(&User{}).Where("id = ?", u.ID).Updates(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) RecoveryUser(uid uint, status bool) error {
	err := config.DB.Unscoped().Model(&User{}).Where("id = ?", uid).Update("deleted_at", nil).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUser() error {

	tx := config.DB.Begin()

	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser() error {
	err := config.DB.Model(&User{}).Where("id = ?", u.ID).Updates(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByDeviceId(deviceId, username string) (*User, error) {
	err := config.DB.Model(User{}).Where("device_id = ?", deviceId).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (u *User) GetUserById(uid uint) (*User, error) {
	err := config.DB.Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (u *User) GetUserByDelete(uid uint) (*User, error) {
	err := config.DB.Unscoped().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (u *User) GetUser() ([]User, error) {
	var users []User

	err := config.DB.Model(User{}).Find(&users).Error
	if err != nil {
		return users, err
	}

	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (u *User) SaveUser() error {
	var err error

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	err = config.DB.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	err := config.DB.Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}
