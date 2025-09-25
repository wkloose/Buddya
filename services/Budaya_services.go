package services

import (
	"errors"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func GetBudayaByCityID(cityID uint) (string, error) {
	var city models.City
	result := initializers.DB.First(&city, cityID)
	if result.Error != nil {
		return "", errors.New("kota tidak ditemukan")
	}

	if city.DeskripsiBudaya == "" {
		return "", errors.New("data budaya belum tersedia untuk kota ini")
	}

	return city.DeskripsiBudaya, nil
}
