package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
	"github.com/wkloose/Buddya/services"
)

func GetWisataByCity(c *gin.Context) {
	cityName := c.Query("city_name")
	cityIDParam := c.Query("city_id")

	var city models.City

	if cityIDParam != "" {
		cityID, err := strconv.Atoi(cityIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "city_id harus berupa angka"})
			return
		}
		if err := initializers.DB.First(&city, cityID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
			return
		}
	} else if cityName != "" {
		if err := initializers.DB.Where("LOWER(nama_kota) = LOWER(?)", cityName).First(&city).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter city_id atau city_name wajib diisi"})
		return
	}

	places, err := services.GetWisataByCityID(city.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data wisata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city":   city.NamaKota,
		"wisata": places,
	})
}

func GetKulinerByCity(c *gin.Context) {
    cityIDParam := c.Query("city_id")
    cityName := strings.TrimSpace(c.Query("city_name"))

    var city models.City

    if cityIDParam != "" {
        cityID, err := strconv.Atoi(cityIDParam)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "city_id harus berupa angka"})
            return
        }
        if err := initializers.DB.First(&city, cityID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
            return
        }
    } else if cityName != "" {
        if err := initializers.DB.Where("LOWER(nama_kota) = LOWER(?)", cityName).First(&city).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "city not found"})
            return
        }
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "parameter city_id atau city_name wajib diisi"})
        return
    }

    places, err := services.GetKulinerByCityID(city.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data kuliner"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "city":    city.NamaKota,
        "kuliner": places,
    })
}
func GetPlaces(c *gin.Context) {
	cityName := c.Query("city_name")
	category := c.Query("category")

	if cityName == "" || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city_name dan category wajib diisi"})
		return
	}

	places, err := services.GetPlacesByCityAndCategory(cityName, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city":     cityName,
		"category": category,
		"places":   places,
	})
}
