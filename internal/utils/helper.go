package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"os"
	"strconv"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/gin-gonic/gin"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key);
	if value == ""{
		
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {

		return defaultValue
	}
	
	return intValue
}

func GenerateRandomString(lenght int) (string, error) {
	bytes := make([]byte, lenght)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenerateRandomInt(lenght int) (string, error) {
	digits := "0123456789"
	number := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		number[i] = digits[num.Int64()]
	}

	return string(number), nil
}

func GetUserLogged(ctx *gin.Context) (v1dto.UserPayload, error) {
	var user v1dto.UserPayload
	payload, _ := ctx.Get("data")
	bytes, _ := json.Marshal(payload)
	
	if err := json.Unmarshal(bytes, &user); err != nil {
		return v1dto.UserPayload{}, err
	}

	return user, nil
}