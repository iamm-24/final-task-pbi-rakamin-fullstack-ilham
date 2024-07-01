package models

import (
	"rakamin/app"
	"rakamin/database"
)

func FindPhotoByID(photoID uint) (app.Photo, error) {
	var photo app.Photo
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		return photo, result.Error
	}
	return photo, nil
}

func UpdatePhoto(photo app.Photo) error {
	result := database.DB.Save(&photo)
	return result.Error
}
