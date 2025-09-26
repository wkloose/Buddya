package controllers

import (
	"strings"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/services"
	"github.com/wkloose/Buddya/models"
)

func GetBudayaByCity(c *gin.Context) {
    cityIDParam := c.Query("city_id")
    cityName := strings.TrimSpace(c.Query("city_name"))

    var city models.City

    if cityIDParam != "" {
        cityID, err := strconv.Atoi(cityIDParam)
        if err != nil {
            c.JSON(400, gin.H{"error": "city_id harus berupa angka"})
            return
        }
        if err := initializers.DB.First(&city, cityID).Error; err != nil {
            c.JSON(404, gin.H{"error": "city not found"})
            return
        }
    } else if cityName != "" {
        if err := initializers.DB.Where("LOWER(nama_kota) = LOWER(?)", cityName).First(&city).Error; err != nil {
            c.JSON(404, gin.H{"error": "city not found"})
            return
        }
    } else {
        c.JSON(400, gin.H{"error": "parameter city_id atau city_name wajib diisi"})
        return
    }

    budaya, err := services.GetBudayaByCityID(city.ID)
    if err != nil {
        c.JSON(500, gin.H{"error": "gagal mengambil data budaya"})
        return
    }

    c.JSON(200, gin.H{
        "city":   city.NamaKota,
        "budaya": budaya,
    })
}
