package app

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Photos    []Photo 	`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (u *User) GenerateJWT(secret string) (string, error) {
	claims := &Claims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJWT(tokenString string, secret string) (uint, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, err
	}
}
