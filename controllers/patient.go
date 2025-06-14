package controllers

import (
    "net/http"
    "golang-clinic-app/config"
    "golang-clinic-app/models"
    "github.com/gin-gonic/gin"
)

// Create a new patient (Receptionist only)
func CreatePatient(c *gin.Context) {
    role, _ := c.Get("role")
    if role != "receptionist" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only receptionists can create patients"})
        return
    }
    var patient models.Patient
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Create(&patient).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, patient)
}

// Get all patients (Doctor and Receptionist)
func GetPatients(c *gin.Context) {
    role, _ := c.Get("role")
    if role != "doctor" && role != "receptionist" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
        return
    }
    var patients []models.Patient
    if err := config.DB.Find(&patients).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, patients)
}

// Get a single patient by ID (Doctor and Receptionist)
func GetPatient(c *gin.Context) {
    role, _ := c.Get("role")
    if role != "doctor" && role != "receptionist" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
        return
    }
    var patient models.Patient
    if err := config.DB.First(&patient, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }
    c.JSON(http.StatusOK, patient)
}

// Update a patient (Doctor and Receptionist)
func UpdatePatient(c *gin.Context) {
    role, _ := c.Get("role")
    if role != "doctor" && role != "receptionist" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
        return
    }
    var patient models.Patient
    if err := config.DB.First(&patient, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
        return
    }
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&patient)
    c.JSON(http.StatusOK, patient)
}

// Delete a patient (Receptionist only)
func DeletePatient(c *gin.Context) {
    role, _ := c.Get("role")
    if role != "receptionist" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only receptionists can delete patients"})
        return
    }
    if err := config.DB.Delete(&models.Patient{}, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}