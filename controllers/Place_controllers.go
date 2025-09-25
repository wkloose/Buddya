package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/services"
)

func GetWisataByCity(c *gin.Context) {
	cityIDParam := c.Query("city_id")
	if cityIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter city_id wajib diisi"})
		return
	}

	cityID, err := strconv.Atoi(cityIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city_id harus berupa angka"})
		return
	}

	places, err := services.GetWisataByCityID(uint(cityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data wisata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city_id": cityID,
		"wisata":  places,
	})
}
func GetKulinerByCity(c *gin.Context) {
	cityIDParam := c.Query("city_id")
	if cityIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter city_id wajib diisi"})
		return
	}

	cityID, err := strconv.Atoi(cityIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city_id harus berupa angka"})
		return
	}

	places, err := services.GetKulinerByCityID(uint(cityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mengambil data kuliner"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city_id":  cityID,
		"kuliner":  places,
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