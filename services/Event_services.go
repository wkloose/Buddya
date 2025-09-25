package services

import (
	"errors"
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func GetEventPromptTemplate() (models.PromptTemplate, error) {
	var tmpl models.PromptTemplate
	result := initializers.DB.Where("kategori = ?", "event").First(&tmpl)
	if result.Error != nil {
		return tmpl, errors.New("prompt template event tidak ditemukan")
	}
	return tmpl, nil
}

func SaveEvent(cityID uint, namaEvent string, tanggal time.Time, lokasi string, gmapsLink string) error {
	event := models.Event{
		NamaEvent:      namaEvent,
		Tanggal:        tanggal,
		Lokasi:         lokasi,
		GoogleMapsLink: gmapsLink,
		CityID:         cityID,
	}
	result := initializers.DB.Create(&event)
	return result.Error
}

func SavePromptLog(userInput, promptFinal, aiResponse string) {
	log := models.PromptLog{
		UserInput:   userInput,
		PromptFinal: promptFinal,
		AIResponse:  aiResponse,
		Timestamp:   time.Now(),
	}
	initializers.DB.Create(&log)
}
