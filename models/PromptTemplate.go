package models
import (
	"time"
	"gorm.io/gorm"
)
type PromptTemplate struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Kategori   string         `gorm:"size:50;not null" json:"kategori"` // ex: event
	PromptText string         `gorm:"type:text;not null" json:"prompt_text"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}