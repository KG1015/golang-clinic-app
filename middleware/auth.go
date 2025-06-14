package middleware

import (
    "net/http"
    "strings"
    "golang-clinic-app/utils"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // Expecting header format: "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
            c.Abort()
            return
        }

        tokenStr := parts[1]
        claims := &utils.Claims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return utils.JwtKey(), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        // Store user info in context for later use
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
        c.Next()
    }
}