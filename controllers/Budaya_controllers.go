package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/tempproject.git/services"
)

func GetBudayaByCity(c *gin.Context) {
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

	budaya, err := services.GetBudayaByCityID(uint(cityID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city_id": cityID,
		"budaya":  budaya,
	})
}
