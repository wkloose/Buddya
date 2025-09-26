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
func GetPlacesByCityAndCategory(cityName, category string) ([]models.Place, error) {
    var places []models.Place
    err := initializers.DB.
        Joins("JOIN cities ON places.city_id = cities.id").
        Where("LOWER(cities.nama_kota) = LOWER(?) AND places.kategori = ?", cityName, category).
        Find(&places).Error

    return places, err
}
