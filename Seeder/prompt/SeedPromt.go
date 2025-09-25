package Prompt

import (
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)

func SeedPrompt() {
	prompt := models.PromptTemplate{
		Kategori:   "event",
		PromptText: "Tolong carikan event budaya di {kota} dalam 1 bulan ke depan. Sertakan nama event, tanggal, alamat, dan link Google Maps.",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	initializers.DB.Create(&prompt)
}
