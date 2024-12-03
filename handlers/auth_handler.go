package handlers

import (
	"barcation_be/config"
	"barcation_be/models"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func AuthHandler(username string, password string, deviceId string, deviceToken string, isInfoSave bool) (string, string, models.User, error) {
	var err error
	res := models.User{}

	err = config.DB.Model(models.User{}).Where("username = ?", username).Take(&res).Error
	if err != nil {
		return "", "", res, fmt.Errorf("username atau password salah, silahkan cek kembali")
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", "", res, fmt.Errorf("username atau password salah, silahkan cek kembali")
	}

	if res.DeviceId == "" {
		err = config.DB.Model(models.User{}).Where("id = ?", res.ID).Update("device_id", deviceId).Error
		if err != nil {
			return "", "", res, fmt.Errorf("terjadi kesalahan saat menyimpan data id device")
		}
	} else if res.DeviceId != deviceId {
		return "", "", res, fmt.Errorf("akun anda sudah digunakan pada device lain, silahkan hubungi ke admin")
	}

	err = config.DB.Model(models.User{}).Where("id = ?", res.ID).Update("save_login", isInfoSave).Update("device_token", deviceToken).Update("last_login", time.Now()).Error
	if err != nil {
		return "", "", res, fmt.Errorf("terjadi kesalahan saat menyimpan data device ID, device token, dan last login")
	}

	accessToken, refreshToken, err := GenerateToken(res.ID, res.Username, res.DeviceId, res.Email, res.Level)
	if err != nil {
		return "", "", res, fmt.Errorf("terjadi kesalahan saat membuat token")
	}

	res.Password = ""

	return accessToken, refreshToken, res, nil
}
