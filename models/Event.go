package models
import (
	"time"
	"gorm.io/gorm"
)
type Event struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	NamaEvent      string         `gorm:"size:150;not null" json:"nama_event"`
	Tanggal        time.Time      `json:"tanggal"`
	Lokasi         string         `gorm:"size:255" json:"lokasi"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	GoogleMapsLink string         `gorm:"size:255" json:"google_maps_link"`
	CityID         uint           `json:"city_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}