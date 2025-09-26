package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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

type EventResponse struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type AllResult struct {
	Wisata  []models.Place  `json:"wisata"`
	Kuliner []models.Place  `json:"kuliner"`
	Budaya  string          `json:"budaya"`
	Events  []EventResponse `json:"events"`
}

func GetEventsFromGemini(cityName string) ([]EventResponse, error) {
	tmpl, err := GetEventPromptTemplate()
	if err != nil {
		return nil, err
	}

	promptFinal := strings.Replace(tmpl.PromptText, "{kota}", cityName, -1)

	aiResp, err := CallGeminiAPI(promptFinal)
	if err != nil {
		return nil, err
	}

	events, err := ParseAIEventResponse(aiResp)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func ParseAIEventResponse(aiResponse string) ([]EventResponse, error) {
	s := strings.TrimSpace(aiResponse)

	if strings.HasPrefix(s, "```") {
		s = strings.TrimPrefix(s, "```json")
		s = strings.TrimPrefix(s, "```")
		s = strings.TrimSuffix(s, "```")
		s = strings.TrimSpace(s)
	}

	if idx := strings.IndexAny(s, "[{"); idx > 0 {
		s = s[idx:]
	}

	var events []EventResponse
	if err := json.Unmarshal([]byte(s), &events); err == nil && len(events) > 0 {
		return events, nil
	}

	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal([]byte(s), &wrapper); err == nil {
		if raw, ok := wrapper["events"]; ok {
			if err := json.Unmarshal(raw, &events); err == nil {
				return events, nil
			}
		}
	}

	lines := strings.Split(s, "\n")
	events = []EventResponse{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if idx := strings.Index(line, ". "); idx > 0 && idx < 4 {
			line = strings.TrimSpace(line[idx+2:])
		}

		parts := strings.Split(line, " - ")
		if len(parts) >= 3 {
			ev := EventResponse{
				Name:     strings.TrimSpace(parts[0]),
				Date:     strings.TrimSpace(parts[1]),
				Location: strings.TrimSpace(parts[2]),
				Link:     "",
			}
			if len(parts) >= 4 {
				ev.Link = strings.TrimSpace(parts[3])
			}
			events = append(events, ev)
		}
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("gagal parsing AI response menjadi events; raw: %s", aiResponse)
	}
	return events, nil
}

func GetAllByCity(cityName string) (*AllResult, error) {
	var city models.City
	if err := initializers.DB.Where("LOWER(nama_kota) = LOWER(?)", cityName).First(&city).Error; err != nil {
		return nil, err
	}

	var wisata []models.Place
	var kuliner []models.Place

	if err := initializers.DB.Where("city_id = ? AND kategori = ?", city.ID, "wisata").Find(&wisata).Error; err != nil {
		return nil, err
	}
	if err := initializers.DB.Where("city_id = ? AND kategori = ?", city.ID, "kuliner").Find(&kuliner).Error; err != nil {
		return nil, err
	}

	events, err := GetEventsFromGemini(city.NamaKota)
	if err != nil {
		log.Printf("warning: gagal memanggil Gemini untuk %s: %v", city.NamaKota, err)
		events = []EventResponse{}
	}

	return &AllResult{
		Wisata:  wisata,
		Kuliner: kuliner,
		Budaya:  city.DeskripsiBudaya,
		Events:  events,
	}, nil
}
