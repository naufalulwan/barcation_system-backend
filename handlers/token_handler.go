package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var tokenBlacklist = make(map[string]bool)

func IsTokenBlacklisted(c *gin.Context) bool {
	token, _ := ExtractToken(c)
	return tokenBlacklist[token]
}

func AddTokenToBlacklist(token string) {
	tokenBlacklist[token] = true
}

func GenerateToken(user_id uint, username, device_id, email, level string) (string, string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))

	if err != nil {
		return "", "", err
	}

	claims := jwt.MapClaims{}

	claims["user_id"] = user_id
	claims["username"] = username
	claims["email"] = email
	claims["level"] = level
	claims["device_id"] = device_id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenUse, err := accessToken.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{}

	refreshClaims["user_id"] = user_id
	refreshClaims["username"] = username
	refreshClaims["email"] = email
	refreshClaims["level"] = level
	refreshClaims["device_id"] = device_id
	refreshClaims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshTokenUse, err := refreshToken.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return "", "", err
	}

	return accessTokenUse, refreshTokenUse, nil
}

func ValidateToken(c *gin.Context) error {
	tokenString, refreshToken := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		refreshToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})

		if err != nil {
			return err
		}

		newAccessToken, _, err := GenerateToken(uint(refreshToken.Claims.(jwt.MapClaims)["user_id"].(float64)), refreshToken.Claims.(jwt.MapClaims)["username"].(string), refreshToken.Claims.(jwt.MapClaims)["device_id"].(string), refreshToken.Claims.(jwt.MapClaims)["email"].(string), refreshToken.Claims.(jwt.MapClaims)["level"].(string))

		if err != nil {
			return err
		}

		c.Writer.Header().Set("Authorization", newAccessToken)
		return nil
	}
	return nil
}

func ExtractToken(c *gin.Context) (string, string) {
	token := c.Query("token")
	refreshToken := c.Query("refresh-token")
	if token != "" {
		return token, refreshToken
	}
	bearerToken := c.Request.Header.Get("Authorization")
	refreshBearerToken := c.Request.Header.Get("Refresh-Token")
	if (len(strings.Split(refreshBearerToken, " ")) == 2) && (len(strings.Split(bearerToken, " ")) == 2) {
		return strings.Split(bearerToken, " ")[1], strings.Split(refreshBearerToken, " ")[1]
	}

	return "", ""
}

func ExtractTokenById(c *gin.Context) (uint, error) {
	tokenString, _ := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, nil
	}

	uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
	if err != nil {
		return 0, err
	}
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	currentTime := time.Now()
	if currentTime.After(expirationTime) {
		return 0, err
	}
	return uint(uid), nil
}
