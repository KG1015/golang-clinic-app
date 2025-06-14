package utils

import (
    "golang.org/x/crypto/bcrypt"
    "time"
    "github.com/dgrijalva/jwt-go"
)


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

var jwtKey = []byte("your_secret_key") 

type Claims struct {
    UserID uint
    Role   string
    jwt.StandardClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func JwtKey() []byte {
    return jwtKey
}