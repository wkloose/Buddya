package models

import (
	"time"
	"gorm.io/gorm"
)
type PromptLog struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserInput   string         `gorm:"type:text" json:"user_input"`
	PromptFinal string         `gorm:"type:text" json:"prompt_final"`
	AIResponse  string         `gorm:"type:text" json:"ai_response"`
	Timestamp   time.Time      `gorm:"autoCreateTime" json:"timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}