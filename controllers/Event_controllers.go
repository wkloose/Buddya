package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/services"
)

func GetEventByCity(c *gin.Context) {
	cityName := c.Query("city_name")
	if cityName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter city_name wajib diisi"})
		return
	}

	tmpl, err := services.GetEventPromptTemplate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	promptFinal := strings.Replace(tmpl.PromptText, "{kota}", cityName, -1)

	aiResponse, err := services.CallGeminiAPI(promptFinal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal memanggil Gemini API: " + err.Error()})
		return
	}

	events, err := services.ParseAIResponse(aiResponse)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"city_name":    cityName,
			"raw_response": aiResponse,
			"note":         "Parsing gagal, kirim raw text",
		})
		return
	}

	services.SavePromptLog("event budaya", promptFinal, aiResponse)

	c.JSON(http.StatusOK, gin.H{
		"city_name": cityName,
		"events":    events,
	})
}
