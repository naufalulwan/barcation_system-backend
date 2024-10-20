package handlers

import (
	"barcation_be/config"
	"barcation_be/models"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(username, password, deviceId, ssn string, isInfoSave bool) (string, string, models.User, error) {
	var err error
	res := models.User{}

	err = config.DB.Model(models.User{}).Where("username = ?", username).Take(&res).Error
	if err != nil {
		return "", "", res, err
	}

	err = config.DB.Model(models.User{}).Where("id = ?", res.ID).Update("device_id", deviceId).Update("save_login", isInfoSave).Update("last_login", time.Now()).Update("ssn", ssn).Error
	if err != nil {
		return "", "", res, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", "", res, err
	}

	accessToken, refreshToken, err := GenerateToken(res.ID, res.Username, res.DeviceId, res.Email, res.Level)
	if err != nil {
		return "", "", res, err
	}

	res.Password = ""

	return accessToken, refreshToken, res, nil
}
