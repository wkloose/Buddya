package Prompt

import (
	"time"

	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/models"
)
func SeedPrompt() {
	prompt := models.PromptTemplate{
		Kategori: "event",
		PromptText: `Tolong carikan maksimal 5 event budaya dalam 1 bulan ke depan di kota {kota}.
Jawab HANYA dalam format JSON array tanpa teks tambahan, dengan struktur:
[
  {
    "name": "Nama event",
    "date": "YYYY-MM-DD atau rentang tanggal",
    "location": "Alamat lengkap acara",
    "description": "Deskripsi singkat acara",
    "link": "Link Google Maps dalam format panjang (bukan shortlink), harus mengandung '/place/' atau koordinat lat-lng."
  }
]`,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	initializers.DB.Create(&prompt)
}

