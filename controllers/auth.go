package controllers

import (
    "net/http"
    "golang-clinic-app/config"
    "golang-clinic-app/models"
    "golang-clinic-app/utils"
    "github.com/gin-gonic/gin"
)

type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Check password hash
    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "role":  user.Role,
        "name":  user.Name,
    })
}

type RegisterInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"` // "doctor" or "receptionist"
}

func RegisterUser(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: hashedPassword,
        Role:     input.Role,
    }

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists or invalid data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}