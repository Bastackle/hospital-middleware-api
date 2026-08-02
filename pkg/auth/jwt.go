package auth

import (
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your-secret-key")

type JwtClaims struct {
	StaffID uint `json:"staff_id"`
	Username string `json:"username"`
	HospitalID uint `json:"hospital_id"`
	jwt.RegisteredClaims
}

// HashPassword แปลง Password เป็น Hash
func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPwd), err
}

// CheckPasswordHash ตรวจสอบ Password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken สร้าง JWT Token ที่ฝัง hospital_id ไว้ข้างใน
func GenerateToken(staffID uint, username string, hospitalID uint) (string, error) {
	claims := JwtClaims{
		StaffID: staffID,
		Username: username,
		HospitalID: hospitalID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // หมดอายุใน 24 ชม.
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}