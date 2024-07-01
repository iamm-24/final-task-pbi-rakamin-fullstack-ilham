package helpers

import "rakamin/app"

var SecretKey = "secret_key"

func GenerateJWT(user app.User) (string, error) {
	return user.GenerateJWT(SecretKey)
}

func ValidateJWT(tokenString string) (uint, error) {
	return app.ParseJWT(tokenString, SecretKey)
}
