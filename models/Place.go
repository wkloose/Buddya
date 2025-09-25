package models
import (
	"time"
	"gorm.io/gorm"
)
type Place struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	NamaTempat string         `gorm:"size:150;not null" json:"nama_tempat"`
	Kategori   string         `gorm:"type:varchar(20);not null" json:"kategori"` // wisata / kuliner
	Alamat     string         `gorm:"size:255" json:"alamat"`
	Latitude   float64        `json:"latitude"`
	Longitude  float64        `json:"longitude"`
	Deskripsi  string         `gorm:"type:text" json:"deskripsi"`
	ImageURL   string         `gorm:"size:255" json:"image_url"`
	CityID     uint           `json:"city_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}