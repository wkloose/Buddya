package services

import (
	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func GetWisataByCityID(cityID uint) ([]models.Place, error) {
	var places []models.Place
	result := initializers.DB.Where("city_id = ? AND kategori = ?", cityID, "wisata").Find(&places)
	if result.Error != nil {
		return nil, result.Error
	}
	return places, nil
}
func GetKulinerByCityID(cityID uint) ([]models.Place, error) {
	var places []models.Place
	result := initializers.DB.Where("city_id = ? AND kategori = ?", cityID, "kuliner").Find(&places)
	if result.Error != nil {
		return nil, result.Error
	}
	return places, nil
}
