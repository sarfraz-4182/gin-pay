package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Amount       int64          `json:"amount"`
	Currency     string         `json:"currency"`
	CustomerID   string         `json:"customer_id"`
	Status       string         `json:"status"` // pending, succeeded, failed
	ExternalRef  string         `json:"external_ref"` // webhook or gateway ref
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
