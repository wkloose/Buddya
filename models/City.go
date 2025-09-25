package models
import (
	"time"
	"gorm.io/gorm"
)
type City struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	NamaKota        string         `gorm:"size:100;not null" json:"nama_kota"`
	Provinsi        string         `gorm:"size:100" json:"provinsi"`
	Latitude        float64        `json:"latitude"`
	Longitude       float64        `json:"longitude"`
	DeskripsiBudaya string         `gorm:"type:text" json:"deskripsi_budaya"` // budaya khas kota
	Places          []Place        `gorm:"foreignKey:CityID" json:"places"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}