package controllers

import (
	"net/http"
	"strconv"

	"rakamin/app"
	"rakamin/database"

	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	var photo app.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	photo.UserID = userID.(uint)
	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func GetPhotos(c *gin.Context) {
	var photos []app.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
	var photo app.Photo
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists || photo.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func DeletePhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid photo ID"})
		return
	}

	var photo app.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	userID, exists := c.Get("userID")
	if !exists || photo.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
